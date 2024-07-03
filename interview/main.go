package main

import "fmt"

// write to
//input of two string give bool output
//s1 = "aaabbc", "pqrrrpq"
//s2 = "pppqqr", "01www01"

// abcr
// rztf

func patterMatch(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1tos2 := make(map[rune]rune)
	s2tos1 := make(map[rune]rune)

	for i := 0; i < len(s1); i++ {
		if ch, ok := s1tos2[rune(s1[i])]; ok {
			if ch != rune(s2[i]) {
				return false
			}
		} else {
			s1tos2[rune(s1[i])] = rune(s2[i])
		}

		if ch, ok := s2tos1[rune(s2[i])]; ok {
			if ch != rune(s1[i]) {
				return false
			}
		} else {
			s2tos1[rune(s2[i])] = rune(s1[i])
		}
	}

	return true
}

func main() {
	s1 := "aaabbc"
	s2 := "pppqqr"

	fmt.Printf("Pattern match of s1 and s2: %v\n", patterMatch(s1, s2))
	s1 = "pqrrrpq"
	s2 = "01www01"
	fmt.Printf("Pattern match of s1 and s2: %v\n", patterMatch(s1, s2))

	s1 = "pqrrrpq"
	s2 = "01www02"
	fmt.Printf("Pattern match of s1 and s2: %v\n", patterMatch(s1, s2))

	s1 = "abcabc"
	s2 = "aaaaaa"
	fmt.Printf("Pattern match of s1 and s2: %v\n", patterMatch(s1, s2))
}
