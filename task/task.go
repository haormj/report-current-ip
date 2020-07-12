package task

import (
	"github.com/haormj/report-current-ip/option"
	"github.com/haormj/report-current-ip/provider"
)

type Task struct {
	opt          option.TaskOpt
	ps           []provider.Provider
	ReportIPTask *ReportIPTask
}

func NewTask(opt option.TaskOpt, ps []provider.Provider) (*Task, error) {
	reportIPTask, err := NewReportIPTask(opt, ps)
	if err != nil {
		return nil, err
	}
	t := Task{
		opt:          opt,
		ps:           ps,
		ReportIPTask: reportIPTask,
	}
	return &t, nil
}

func (t *Task) Run() error {
	if err := t.ReportIPTask.Start(); err != nil {
		return err
	}
	return nil
}

func (t Task) Close() error {
	return nil
}
