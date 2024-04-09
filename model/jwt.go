package model

import "github.com/golang-jwt/jwt/v4"

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}

type JwtReq struct {
	KeyID            string  `json:"access_key" form:"access_key" query:"access_key" binding:"required"`
	EncodedKeySecret string  `json:"security_key" form:"security_key" query:"security_key" binding:"required"`
	Subject          string  `json:"subject"  form:"subject" query:"subject" binding:"required"` // appid
	Value            *string `json:"value"  form:"value" query:"value"`                          //需要被签名的内容
	Captcha          string  `json:"captcha"`                                                    // 验证码
	CaptchaId        string  `json:"captchaId"`                                                  // 验证码ID
}

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	//APIConsumerID int64
	ID       int64 // access key record is for kong plugin update
	KeyID    string
	Audience []string // routers join by ,
	Subject  string   // appid
	Value    *string  //需要被签名的内容
	// EncodedKeySecret string

	Type     int32
	Nickname string
	Gender   string
	Cover    string
	Phone    string
	Email    string
}
