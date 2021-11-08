package conf

type SystemConfig struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}