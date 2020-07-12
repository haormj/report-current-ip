package option

import (
	"fmt"

	"github.com/haormj/log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Options program config
type Option struct {
	v           *viper.Viper
	LogOpt      LogOpt      `mapstructure:"log_opt"`
	ProviderOpt ProviderOpt `mapstructure:"provider_opt"`
	TaskOpt     TaskOpt     `mapstructure:"task_opt"`
}

func New() (*Option, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	opt := Option{
		v: v,
	}
	if err := v.Unmarshal(&opt); err != nil {
		return nil, err
	}

	log.Logger = log.NewLog(
		log.Level(log.ParseLevel(opt.LogOpt.Level)),
		log.Filename(opt.LogOpt.Filename),
		log.Encoder(log.ParseEncoder(opt.LogOpt.Encoder)),
	)

	return &opt, nil
}

func (o *Option) String() string {
	c := o.v.AllSettings()
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("unable to marshal config to YAML: %v", err)
	}
	return string(b)
}
