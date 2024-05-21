package model

type OneidConf struct {
	AuthenticationUrl string `json:"authentication_url" form:"authentication_url" query:"authentication_url"` // kong
	AccessKeyId       string `json:"access_key_id" form:"access_key_id" query:"access_key_id"`                // ak
	AccessKeySecret   string `json:"access_key_secret" form:"access_key_secret" query:"access_key_secret"`    // sk
	Subject           string `json:"subject" form:"subject" query:"subject"`                                  // consumer name
}
