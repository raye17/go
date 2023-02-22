package util

func CompareSlice(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, s1 := range str1 {
		found := false
		for _, s2 := range str2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
