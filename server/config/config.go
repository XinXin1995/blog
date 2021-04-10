package config

type Server struct {
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Log     Log     `mapstructure:"log" json:"log" yaml:"log"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey"`
}


type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile"`
	Stdout  string `mapstructure:"stdout" json:"stdout"`
	File    string `mapstructure:"file" json:"file"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight"`
}
