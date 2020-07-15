package task

import (
	"errors"
	"net"
	"time"

	"github.com/haormj/report-current-ip/option"
	"github.com/haormj/report-current-ip/provider"

	"github.com/haormj/log"
)

type ReportIPTask struct {
	opt option.TaskOpt
	ps  []provider.Provider
}

func NewReportIPTask(opt option.TaskOpt, ps []provider.Provider) (*ReportIPTask, error) {
	r := ReportIPTask{
		opt: opt,
		ps:  ps,
	}
	return &r, nil
}

func (r *ReportIPTask) Start() error {
	ps := r.ps
	for {
		ip, err := r.getCurrentIP(r.opt.ReportIPTaskOpt.IfName)
		if err != nil {
			log.Logger.Errorw("getCurrentIP error", "err", err)
		} else {
			ps = r.reportIPToProviders(ip, ps)
			if len(ps) == 0 {
				break
			}
		}
		time.Sleep(time.Second * time.Duration(r.opt.ReportIPTaskOpt.RetryInterval))
	}
	log.Logger.Infow("task finished")
	return nil
}

func (r *ReportIPTask) Stop() error {
	return nil
}

func (r *ReportIPTask) reportIPToProviders(ip string, ps []provider.Provider) []provider.Provider {
	reportFailedProviders := make([]provider.Provider, 0)
	for _, p := range ps {
		if err := p.ReportCurrentIP(ip); err != nil {
			reportFailedProviders = append(reportFailedProviders, p)
			log.Logger.Errorw("reportIPToProviders error", "err", err)
		}
	}
	return reportFailedProviders
}

func (r *ReportIPTask) getCurrentIP(name string) (string, error) {
	in, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}
	addrs, err := in.Addrs()
	if err != nil {
		return "", err
	}
	if len(addrs) > 0 {
		ip, _, err := net.ParseCIDR(addrs[0].String())
		if err != nil {
			return "", err
		}
		return ip.String(), nil
	}
	return "", errors.New("not exist")
}
