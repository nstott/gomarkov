package gomarkov

import (
	"rand"
	"regexp"
	"strings"
	"time"
)

const (
	WORDSIZE          = 64
	WORDS_IN_SENTANCE = 25
)


type Mchain struct {
	Mchain map[string][]string
	Beginnings []triple
}

func ProcessString(data string) *Mchain {
	var m = new(Mchain)
	m.init()

	re, _ := regexp.Compile("\n+")
	re2, _ := regexp.Compile("[ ]+")

	cleanedText := re.ReplaceAllString(data, "\n")
	cleanedText = re2.ReplaceAllString(data, " ")
	var t triple
	for _, paragraph := range strings.Split(cleanedText, "\n", -1) {
		words := strings.Split(paragraph, " ", -1)
		if len(words) > 2 {
			m.addBeginning(words[0], words[1])
		}
		for _, word := range words {
			t.addWord(word)
			m.addToChain(t)
		}
	}
	
	return m
}

func Generate(m *Mchain, count int) string {
	rand.Seed(time.Nanoseconds())
	t := m.Beginnings[rand.Intn(len(m.Beginnings)-1)]
	first, second := t.First, t.Second

	ret := make([]string, count, count)
	ret[0] = t.First; ret[1] = t.Second

	var word string
	for i := 2; i < count; i++ {
		t := &triple{first, second, "", nil}
		word = m.getThird(t)
		first, second = second, word
		ret[i] = word
		if word == "" {
			return strings.Join(ret, " ")
		}
		if char := word[len(word)-1]; char == '.' {
			if i+WORDS_IN_SENTANCE >= count {
				break
			}
		}
	}
	return strings.Join(ret, " ")
}


func (m *Mchain) addBeginning(first, second string) {
	var t triple
	t.First = first
	t.Second = second
	m.Beginnings = append(m.Beginnings, t)
}


func (m *Mchain) addToChain(t triple) {
	if t.First == "" || t.Second == "" || t.Third == "" {
		return
	}
	var newvec []string
	vec, present := m.Mchain[t.First+" "+t.Second]
	if !present || len(vec) == 0 {
		newvec = make([]string, 1, 1)
		newvec[0] = t.Third
	} else {
		newvec = make([]string, len(vec)+1, len(vec)+1)
		copy(newvec, vec)
		newvec[len(vec)-1] = t.Third
	}
	m.Mchain[t.First+" "+t.Second] = newvec
}


func (m *Mchain) getThird(t *triple) string {
	if t.First == "" || t.Second == "" {
		return ""
	}
	vec := m.Mchain[t.First+" "+t.Second]
	if len(vec) == 0 {
		return ""
	}
	return vec[rand.Intn(len(vec)-1)]
}


func (m *Mchain) init() {
	m.Beginnings = make([]triple, 0, 4)
	m.Mchain = make(map[string][]string)
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


