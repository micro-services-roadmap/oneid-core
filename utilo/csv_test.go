package utilo

import (
	"fmt"
	"testing"
)

func TestReadAsNestedMap(t *testing.T) {
	nestedMap, _ := ReadAsNestedMap("wosai_order_20240415.csv", ',', Convert2Float32, 0)

	fmt.Println(nestedMap)
}
