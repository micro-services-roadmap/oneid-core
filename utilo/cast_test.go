package utilo

import (
	"fmt"
	"testing"
)

func TestToPrt(t *testing.T) {
	fmt.Println(ToPrt(1))
	fmt.Println(ToPrt("1"))
	fmt.Println(ToPrt(true))
	fmt.Println(ToPrt([]int{1, 2, 3}))
	fmt.Println(ToPrt(map[string]int{"a": 1, "b": 2}))
	fmt.Println(ToPrt(func() {}))
	fmt.Println(ToPrt(struct{}{}))
}

func TestToVl(t *testing.T) {
	var a *int64
	fmt.Println(ToVl(a))
	var b *string
	fmt.Println(ToVl(b))
	fmt.Println(ToVl(ToPrt(1)))
	fmt.Println(ToVl(ToPrt(true)))
	fmt.Println(ToVl(ToPrt([]int{1, 2, 3})))
	fmt.Println(ToVl(ToPrt(map[string]int{"a": 1, "b": 2})))
	//fmt.Println(ToVl(ToPrt(func() {})))
	fmt.Println(ToVl(ToPrt(struct{}{})))
}
