package main

import (
	"fmt"
	"strings"
)

type TestParam struct {
	Input    string
	Expected map[string]int
}

var tests []TestParam = []TestParam{
	{
		"Hello My name is Talha. And I am Talha.",
		map[string]int{
			"hello": 1,
			"talha": 2,
		},
	},
	{
		"Hello My name is Talha Cheese. I like chick-n-cheese.",
		map[string]int{
			"hello":  1,
			"talha":  1,
			"cheese": 2,
		},
	},
	{
		"We came, we saw, we conquered...then we ate Bill's (Mille-Feuille) cake.",
		map[string]int{
			"we":   4,
			"Bill": 1,
			"cake": 1,
		},
	},
}

func main() {
	runWordCloudTest(tests)
}

func runWordCloudTest(tests []TestParam) {
	for _, t := range tests {
		cloud := wordCloud(t.Input)
		for w, cnt := range t.Expected {
			if cnt != cloud[w] {
				fmt.Printf("*** FAIL *** Expected count for '%s' is %d. Got %d. \n", w, cnt, cloud[w])
			}
		}
	}
}

func wordCloud(para string) map[string]int {

	// para clean
	para = strings.Replace(para, "-", " ", -1)

	// need to break the para into words

	words := strings.Split(para, " ")

	var cloud map[string]int = make(map[string]int)
	for _, word := range words {
		// clean the word
		word = cleanWord(word)
		cloud[word] = cloud[word] + 1
	}

	return cloud
}

func cleanWord(word string) string {
	unwanted := []string{",", ":", ".", "'s"}
	for _, r := range unwanted {
		word = strings.Replace(word, r, "", -1)
	}

	return word
}
