package CodeWars

func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}

	length := len(strng)
	for i := 0; i < length; i++ {
		rotation := strng[i:] + strng[:i]
		found := false
		for _, word := range arr {
			if word == rotation {
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
