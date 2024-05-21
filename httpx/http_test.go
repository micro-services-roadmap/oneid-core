package httpx

import (
	"fmt"
	"github.com/micro-services-roadmap/oneid-core/model"
	"strconv"
	"strings"
	"testing"
)

func TestGenerateJwt(t *testing.T) {

	jwt, err := GenerateJwt("", &model.JwtReq{Subject: "wpp-admin",
		KeyID:            "058d1d5f-a23f-47bc-8e9e-d31f412786ce",
		EncodedKeySecret: "972eff62-bd7c-4757-b332-b32b094a7aa5"})
	if err != nil || len(jwt.Msg) == 0 {
		return
	}
}

func TestRegister(t *testing.T) {

	s := "wpp-admin"
	var zero int32 = 0
	KeyID := "zack"
	EncodedKeySecret := "se"
	m := &model.AccessKeyReq{
		Subject:          &s,
		Type:             &zero,
		KeyID:            &KeyID,
		EncodedKeySecret: &EncodedKeySecret,
	}
	jwt, err := Register("localhost:8888", m)
	if err != nil || len(jwt.Msg) == 0 {
		return
	}
}

func TestGetCaptcha(t *testing.T) {
	jwt, err := GetCaptcha("localhost:8888")
	if err != nil || len(jwt.Msg) == 0 {
		return
	}
}

func TestReplace(t *testing.T) {
	UpdateUrl = strings.Replace(UpdateUrl, ":id", strconv.Itoa(9999), 1)
	fmt.Println(UpdateUrl)
}

func TestDoReq(t *testing.T) {
	// _, _ = DoReq("", "http://localhost", []byte{})
	// _, _ = DoReq("", "https://localhost", []byte{})
	_, _ = DoReq("", "localhost", []byte{})
}
