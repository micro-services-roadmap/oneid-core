package utilo

func ToPrt[T any](i T) *T {
	return &i
}

// ToVl 返回指针i指向的值，如果i是nil，则返回T的零值。
func ToVl[T any](i *T) T {
	if i == nil {
		var zero T
		return zero
	}
	return *i
}
