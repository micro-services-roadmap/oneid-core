package util

import (
	"fmt"
	"testing"
)

func TestJWTParseToken(t *testing.T) {
	jwt, err := GetParser("../source/cert/public_key.pem")
	panic(err)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJBUElDb25zdW1lcklEIjoxLCJLZXlJRCI6InN0cmluZyIsIkF1ZGllbmNlIjoic3RyaW5nIiwiU3ViamVjdCI6IiIsIlR5cGUiOjEsIk5pY2tuYW1lIjoic3RyaW5nIiwiR2VuZGVyIjoic3RyaW5nIiwiQ292ZXIiOiJzdHJpbmciLCJQaG9uZSI6InN0cmluZyIsIkVtYWlsIjoic3RyaW5nIiwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6Ik9ORUlEX0FBQ1MiLCJhdWQiOlsic3RyaW5nIl0sImV4cCI6MTcxODQ0NzU3NSwibmJmIjoxNzEyMzk5NTc1fQ.B1Fy8Rs9P_2AePrBGH9S3rDkdMOou-pthkky8NioJVGn_QSs4rBu-_Ae1euHLtluywhqdlLQZGbJwJGxIu4veNHM9YSv5ORnkyz1-XhRbERxCTlD7Fb30jU5WY1ey7qVfLwoIoQ7P-xrUEWc_ScxuC02GQTtqWC2ppSvqL_y5CCIVRHqL3h279uBIzYIQlc8QWqeNNuJqEk1Uxn89R_q-6C_26nJiIEoEz-mtk6130Sw0INyEO13sdbN2NfIHUHmm4l83P-mJKVG5agUMgjVJY0_2_zsHEjlhvizjAxv2mKRuOyyTIlvyg0TP4G4iksq3iAOojONV9kPynH5LHjd-Q"

	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		return
	}

	fmt.Println(parseToken)

}
