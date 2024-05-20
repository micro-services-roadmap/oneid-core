package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/micro-services-roadmap/oneid-core/model"
	"strings"
)

func DecodeJwt(jwt string) (*model.CustomClaims, error) {

	parts := strings.Split(RemoveBearer(jwt), ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid JWT token")
	}

	return DecodePayload(parts[1])
}

func DecodePayload(payload string) (*model.CustomClaims, error) {
	data, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload[%v] due to %s", payload, err.Error())
	}

	var claims model.CustomClaims
	if err := json.Unmarshal(data, &claims); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload[%v] due to %s", payload, err.Error())
	}

	return &claims, nil
}
