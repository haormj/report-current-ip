package main

import (
	"flag"
	"fmt"

	"github.com/haormj/report-current-ip/option"
	"github.com/haormj/report-current-ip/provider"
	"github.com/haormj/report-current-ip/task"

	"github.com/haormj/log"
	"github.com/haormj/version"
)

func main() {
	v := flag.Bool("v", false, "show version")
	c := flag.Bool("c", false, "show config")
	flag.Parse()

	if *v {
		fmt.Println(version.FullVersion())
		return
	}

	if *c {
		opt, err := option.New()
		if err != nil {
			fmt.Printf("[ERR] option.NewOption error %v\n", err)
			return
		}
		fmt.Println(opt)
		return
	}
	opt, err := option.New()
	if err != nil {
		log.Logger.Errorw("option.New error", "err", err)
		return
	}
	providers, err := provider.NewProviders(opt.ProviderOpt)
	if err != nil {
		log.Logger.Errorw("NewProviders error", "err", err)
		return
	}

	t, err := task.NewTask(opt.TaskOpt, providers)
	if err != nil {
		log.Logger.Errorw("NewTask error", "err", err)
		return
	}
	if err := t.Run(); err != nil {
		log.Logger.Errorw("task Run error", "err", err)
	}
	defer t.Close()
}
