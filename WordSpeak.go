package main

import "fmt"
import "strings"
import "strconv"

func main() {
	fmt.Println("go faith")
	definitions := map[string]int{
		"print": 1,
		"add":   2,
	}
	pn := "print 1 print add 1 2 print 4"
	words := strings.Split(pn, " ")
	wordIndex := 2
	fmt.Println(phraseLength(words, wordIndex, definitions))
}

func phraseLength(words []string, wordIndex int, definitions map[string]int) int {
	word := words[wordIndex]
	length := 1
	_, err := strconv.ParseInt(word, 10, 64)
	if err == nil {
		return 1
	}
	argumentLength := definitions[word]
	for argumentIndex := 1; argumentIndex <= argumentLength; argumentIndex++ {
		length += phraseLength(words, wordIndex+length, definitions)
	}
	return length
}
