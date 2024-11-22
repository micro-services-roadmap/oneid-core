package utilo

import (
	"fmt"
	"testing"
)

func TestGenRsaCertPair(t *testing.T) {

	GenRsaCertPair()
}

func TestReadPubKeyOfAbPath(t *testing.T) {
	key, err := ReadPubKeyOfAbPath("/home/zack/kong/kong/v3.5.0/plugins/oneid/cert/public_key.pem")
	if err != nil {
		panic(err)
	}

	fmt.Println(key)
}
