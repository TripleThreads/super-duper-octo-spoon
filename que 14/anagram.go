package main

import (
	"strings"
)

func isAnagram(str1 string, str2 string) bool {

	if len(str1) == 0 && len(str2) == 0 {
		return true
	}
	if len(str1) != len(str2) {
		return false
	}
	// I could make this array's length to 26 but Idk maybe for special characters and whitespaces
	// not bad after all
	count := make([]int, 150)

	for i := 0; i < len(str1); i++ {
		count[str1[i]]++
		count[str2[i]]--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"Madam Curie!", "ABC", "madiur mace!"}
	word := "Radium came!"

	for _, w := range words {
		if isAnagram(strings.ToLower(w), strings.ToLower(word)) {
			println(w)
		}
	}
}
