package modelo

type OneidConf struct {
	AuthenticationUrl string `mapstructure:"authentication-url"  yaml:"authentication-url" json:"authentication-url" form:"authentication-url" query:"authentication-url"` // kong
	AccessKeyId       string `mapstructure:"access-key-id" yaml:"access-key-id" json:"access-key-id" form:"access-key-id" query:"access-key-id"`                           // ak
	AccessKeySecret   string `mapstructure:"access-key-secret" yaml:"access-key-secret" json:"access-key-secret" form:"access-key-secret" query:"access-key-secret"`       // sk
	Subject           string `mapstructure:"subject" yaml:"subject" json:"subject" form:"subject" query:"subject"`                                                         // consumer name
}
