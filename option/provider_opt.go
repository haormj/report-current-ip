package option

type ProviderOpt struct {
	DomainRecordOpt   DomainRecordOpt   `mapstructure:"domain_record_opt"`
	AliyunProviderOpt AliyunProviderOpt `mapstructure:"aliyun_provider_opt"`
}

type DomainRecordOpt struct {
	DomainName string `mapstructure:"domain_name"`
	RR         string `mapstructure:"rr"`
	Type       string `mapstructure:"type"`
}

type AliyunProviderOpt struct {
	RegionID        string `mapstructure:"region_id"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
}
