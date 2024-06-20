package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	results := make(map[string]int, len(s))
	
	for _, word := range words {
		count, ok := results[word]
		if ok { count = count + 1} else { count = 1 }
		results[word] = count
	}
	return results
}

func main() {
	wc.Test(WordCount)
}
