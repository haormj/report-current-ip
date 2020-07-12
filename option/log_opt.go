package option

// LogOpt opt
type LogOpt struct {
	Level    string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	Encoder  string `mapstructure:"encoder"`
}
