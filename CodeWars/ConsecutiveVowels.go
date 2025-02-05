package CodeWars

func GetTheVowels(word string) int {
	vowels := "aeiou"
	index := 0
	count := 0

	for _, char := range word {
		if char == rune(vowels[index]) {
			count++
			index = (index + 1) % len(vowels)
		}
	}

	return count
}
