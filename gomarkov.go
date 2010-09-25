package gomarkov

import (
	"fmt"
	"rand"
	"regexp"
	"strings"
	"time"
)

const (
	WORDSIZE          = 64
	WORDS_IN_SENTANCE = 25
)

var mchain = make(map[string][]string)
var beginnings = make([]triple, 0, 4)

func ProcessString(data string) string {
	re, _ := regexp.Compile("\n+")
	re2, _ := regexp.Compile("[ ]+")

	cleanedText := re.ReplaceAllString(data, "\n")
	cleanedText = re2.ReplaceAllString(data, " ")
	var t triple
	for _, paragraph := range strings.Split(cleanedText, "\n", -1) {
		words := strings.Split(paragraph, " ", -1)
		if len(words) > 2 {
			addBeginning(words[0], words[1])
		}
		for _, word := range words {
			t.addWord(word)
			t.addToChain()
		}
	}

	return re.ReplaceAllString(data, "\n")
}

func Generate(count int) {
	rand.Seed(time.Nanoseconds())
	t := beginnings[rand.Intn(len(beginnings)-1)]
	first, second := t.First, t.Second

	fmt.Printf("%s %s ", first, second)

	var word string
	for i := 0; i < count; i++ {
		t := &triple{first, second, "", nil}
		word = t.getThird()
		first, second = second, word
		fmt.Printf("%s ", word)
		if word == "" {
			return
		}
		if char := word[len(word)-1]; char == '.' {
			if i+WORDS_IN_SENTANCE >= count {
				break
			}
		}
	}
	fmt.Println()
}

func addBeginning(first, second string) {
	var t triple
	t.First = first
	t.Second = second
	beginnings = append(beginnings, t)
}

func append(ar []triple, t triple) []triple {
	n := len(ar)
	if n+1 > cap(ar) { // reallocate
		a := make([]triple, n, 2*n+1)
		copy(a, ar)
		ar = a
	}
	ar = ar[0 : n+1]
	ar[n] = t
	return ar
}
