package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := `^[A-Za-z0-9]+$`
	testStrings := []string{
		"abc123",
		"ABC456",
		"AbC789",
		"123",
		"abcDEF",
		"abc 123", // invalid (contains space)
		"abc-123", // invalid (contains hyphen)
	}

	re := regexp.MustCompile(pattern)
	for _, str := range testStrings {
		matched := re.MatchString(str)
		fmt.Printf("Does \"%s\" match? %v\n", str, matched)
	}
}
