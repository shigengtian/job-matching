package config

type Captcha struct {
	KeyLong            int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`
	ImgWidth           int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`
	ImgHeight          int `mapstructure:"img-height" json:"img-height" yaml:"img-height"`
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"`
}
