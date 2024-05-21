package model

type OneidConf struct {
	AuthenticationUrl string `json:"authentication-url" form:"authentication-url" query:"authentication-url"` // kong
	AccessKeyId       string `json:"access_key_id" form:"access_key_id" query:"access-key-id"`                // ak
	AccessKeySecret   string `json:"access-key-secret" form:"access-key-secret" query:"access-key-secret"`    // sk
	Subject           string `json:"subject" form:"subject" query:"subject"`                                  // consumer name
}
