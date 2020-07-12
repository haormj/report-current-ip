package provider

import (
	"github.com/haormj/report-current-ip/option"
)

type Provider interface {
	ReportCurrentIP(ip string) error
}

func NewProviders(opt option.ProviderOpt) ([]Provider, error) {
	ali, err := NewAliyunProvider(opt)
	if err != nil {
		return nil, err
	}
	providers := []Provider{ali}
	return providers, nil
}
