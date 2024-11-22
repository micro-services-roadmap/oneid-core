package utilo

// RemoveBearer remove token Bearer prefix if necessary
func RemoveBearer(token string) string {

	if len(token) == 0 || len(token) < 7 {
		return token
	}

	if token[:7] == "Bearer " || token[:7] == "bearer " {
		token = token[7:]
	}

	return token
}
