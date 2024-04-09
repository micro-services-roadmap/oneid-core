package httpx

import (
	"github.com/micro-services-roadmap/oneid-core/model"
	"testing"
)

func TestGenerateJwt(t *testing.T) {

	jwt, err := GenerateJwt("localhost:8888", &model.JwtReq{Subject: "wpp-admin",
		KeyID:            "058d1d5f-a23f-47bc-8e9e-d31f412786ce",
		EncodedKeySecret: "972eff62-bd7c-4757-b332-b32b094a7aa5"})
	if err != nil || len(jwt.Msg) == 0 {
		return
	}
}
