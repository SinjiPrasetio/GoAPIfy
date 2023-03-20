package stringable

import "strings"

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	first := strings.ToUpper(string(s[0]))
	return first + s[1:]
}

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func LowerCase(s string) string {
	return strings.ToLower(s)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RemoveVowels(s string) string {
	vowels := "aeiouAEIOU"
	result := ""
	for _, char := range s {
		if !strings.ContainsRune(vowels, char) {
			result += string(char)
		}
	}
	return result
}

func RemoveConsonants(s string) string {
	consonants := "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
	result := ""
	for _, char := range s {
		if !strings.ContainsRune(consonants, char) {
			result += string(char)
		}
	}
	return result
}
