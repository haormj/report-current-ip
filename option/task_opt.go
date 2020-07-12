package option

type TaskOpt struct {
	ReportIPTaskOpt ReportIPTaskOpt `mapstructure:"report_ip_task_opt"`
}

type ReportIPTaskOpt struct {
	IfName        string `mapstructure:"if_name"`
	RetryInterval int64  `mapstructure:"retry_interval"`
}
