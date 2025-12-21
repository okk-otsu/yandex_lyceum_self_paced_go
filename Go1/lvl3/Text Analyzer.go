package main

import (
	"fmt"
	"sort"
	"strings"
)

func someFromatingText(text string) string {
	text = strings.ReplaceAll(text, "!", " ")
	text = strings.ReplaceAll(text, "?", " ")
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, "     ", " ")
	text = strings.ReplaceAll(text, "     ", " ")
	text = strings.ReplaceAll(text, "    ", " ")
	text = strings.ReplaceAll(text, "   ", " ")
	text = strings.ReplaceAll(text, "  ", " ")
	text = strings.TrimSpace(text)

	text = strings.ToLower(text)
	// fmt.Printf("\n\n\n\n\n\n")
	// fmt.Printf(text)
	// fmt.Printf("\n\n\n\n\n\n")
	return text
}

func MakeMap(words []string) map[string]int {
	res := map[string]int{}
	for _, word := range words {
		if word != "" {
			res[word]++
		}
	}
	return res
}

func TheBestPopularWord(mp map[string]int) string {
	var resWord string = ""
	var resCount int = 0
	for word, count := range mp {
		if resCount < count {
			resCount = count
			resWord = word
		}
	}

	return resWord
}

func getTopWords(wordMap map[string]int, n int) []string {

	slice := make([]int, 0, len(wordMap))
	intMap := map[int]string{}
	for key, val := range wordMap {
		slice = append(slice, -val)
		intMap[val] = key
	}

	sort.Ints(slice)
	slice = slice[:n]
	res := make([]string, 0, n)
	for _, val := range slice {
		res = append(res, intMap[-val])
	}
	return res
}

func AnalyzeText(text string) {

	text = someFromatingText(text)
	words := strings.Split(text, " ")

	mp := MakeMap(words)

	fmt.Printf("Количество слов: %d\n", len(words))
	fmt.Printf("Количество уникальных слов: %d\n", len(mp))
	ThePopularWord := TheBestPopularWord(mp)
	fmt.Printf(`Самое часто встречающееся слово: "%s" (встречается %d раз)`, ThePopularWord, mp[ThePopularWord])

	fmt.Printf("\nТоп-5 самых часто встречающихся слов:\n")

	topWords := getTopWords(mp, 5)
	for _, word := range topWords {
		fmt.Printf(`"%s": %d раз`, word, mp[word])
		fmt.Println()
	}

}
