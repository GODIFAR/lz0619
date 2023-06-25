package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	// words strings.Fields  返回空格分隔字符串的字符串切片
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		//strings.Trim 返回去掉`.,"-`的字符串
		word = strings.Trim(word, `.,"-`)
		frequency[word]++
	}
	return frequency
}

func main() {
	text := `As far far .far .far as eye coule reach he saw nothing but the stems of the great plants
	about him receding in the violet shade, and far overhead the multiole
	transparency of huge leaves filtering the sunshine to the solemn splendor of,
	conet--except.`

	frequency := countWords(text)
	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}
}
