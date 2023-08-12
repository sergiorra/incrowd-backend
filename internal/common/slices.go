package common

func Contains(values []string, search string) bool {
	for _, v := range values {
		if v == search {
			return true
		}
	}
	return false
}
