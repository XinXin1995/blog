package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yam:"addr"`
	DbType        string `mpastructure:"db-type" json:"dbType" yaml:"dbType"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"useMultipoint"`
}
