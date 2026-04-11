package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func ReverseWords(str string) string {
	// words := strings.Split(str1, " ")
	words := strings.Fields(str)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

func ReverseLetters(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SplitString(s string, del string) []string {
	return strings.Split(s, del)
}

func main() {
	fmt.Println("Strings cheat-sheet")

	fmt.Println("---Reverse---")
	// Reverse the words in a string
	fmt.Println(ReverseWords("It's a beautiful day"))
	// Reverse the letters in a string
	fmt.Println(ReverseLetters("Hello World!"))

	fmt.Println("---Split Strings---")
	// Split String
	fmt.Println(SplitString("Hello World!", " "))
	fmt.Println(SplitString("apples, oranges, grapes, guava", ","))

	fmt.Println("---Prefix and Suffix---")
	// Prefix and Suffix check
	fmt.Println(strings.HasPrefix("Hello World!", "He"))
	fmt.Println(strings.HasPrefix("Hello World!", "The"))
	fmt.Println(strings.HasSuffix("Hello World!", "ld!"))
	fmt.Println(strings.HasSuffix("Hello World!", "old!"))

	fmt.Println("---Contains---")
	// Contains
	fmt.Println(strings.Contains("Hello World!", "ll"))
	fmt.Println(strings.Contains("Hello World!", "art"))
	// Contains Any
	fmt.Println(strings.ContainsAny("Hello World!", "old"))
	fmt.Println(strings.ContainsAny("Hello World!", "xyz"))

	fmt.Println("---Index---")
	// Index
	fmt.Println(strings.Index("Hello World!", "ll"))
	fmt.Println(strings.LastIndex("Hello World!", "l"))

	fmt.Println("---Count---")
	// Count (counts non-overlapping instances)
	fmt.Println(strings.Index("Hello World!", "l"))

	fmt.Println("---Compare---")
	// Compare
	fmt.Println(strings.Compare("apple", "oranges"))

	fmt.Println("---Substring---")
	// Substring
	fmt.Println("Hello World!"[:5])
	fmt.Println("Hello World!"[7:])

	fmt.Println("---Sort---")
	// Sorting
	strs := []string{"apple", "guava", "oranges", "grapes", "kiwi", "papaya", "berries"}
	copy := strs
	slices.Sort(copy)
	fmt.Println(copy)

	slices.SortFunc(copy, func(a, b string) int {
		// compare length
		if len(a) != len(b) {
			return cmp.Compare(len(a), len(b))
		}
		// then alphabetically
		return cmp.Compare(a, b)
	})
	fmt.Println(copy)

}
