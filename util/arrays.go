package util

func Contains[T comparable](slice []T, element T) bool {
	if len(slice) == 0 {
		return false
	}

	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}
