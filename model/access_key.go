package model

type AccessKeyReq struct {
	AccessKeyUpdateReq

	//Type             *int32  `json:"type" form:"type" query:"type" binding:"required"`           //  0:  用户(admin), 1: 应用
	Subject          string `json:"subject"  form:"subject" query:"subject" binding:"required"` // consumer-name(appid) mapping to APIConsumerID
	KeyID            string `json:"key_id" form:"key_id" query:"key_id"`                        // use for Type=0
	EncodedKeySecret string `json:"key_secret" form:"key_secret" query:"key_secret"`            // use for Type=0
}

type AccessKeyUpdateReq struct {
	ID int64 `json:"id" form:"id" query:"id"`

	//APIConsumerID int64   ` json:"api_consumer_id" form:"api_consumer_id" query:"api_consumer_id" binding:"required"`
	Enabled  *bool   `json:"enabled" form:"enabled" query:"enabled" binding:"required"`
	Nickname *string `json:"nickname" form:"nickname" query:"nickname"`
	Gender   *string `json:"gender" form:"gender" query:"gender"`
	Cover    *string `json:"cover" form:"cover" query:"cover"`
	Phone    *string `json:"phone" form:"phone" query:"phone"`
	Email    *string `json:"email" form:"email" query:"email"`
}
