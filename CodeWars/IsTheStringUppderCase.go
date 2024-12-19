package CodeWars

import "unicode"

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
