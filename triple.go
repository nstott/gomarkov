package gomarkov

import (
	"rand"
	"fmt"
)

type triple struct {
	First       string
	Second      string
	Third       string
	SetOfThirds []string
}

/* 
 * a group of three words,
 * the the First two words form a key for a map
 * the value of the map is a list of all possible SetOfThirds words
 */
func (t *triple) addWord(s string) {
	if t.First == "" {
		t.First = s
	} else if t.Second == "" {
		t.Second = s
	} else if t.Third == "" {
		t.Third = s
	} else {
		t.First = t.Second
		t.Second = t.Third
		t.Third = s
	}
}

func (t *triple) addToChain() {
	if t.First == "" || t.Second == "" || t.Third == "" {
		return
	}
	var newvec []string
	vec, present := mchain[t.First+" "+t.Second]
	if !present || len(vec) == 0 {
		newvec = make([]string, 1, 1)
		newvec[0] = t.Third
	} else {
		newvec = make([]string, len(vec)+1, len(vec)+1)
		copy(newvec, vec)
		newvec[len(vec)-1] = t.Third
	}
	mchain[t.First+" "+t.Second] = newvec
}

func (t *triple) getThird() string {
	if t.First == "" || t.Second == "" {
		return ""
	}
	vec := mchain[t.First+" "+t.Second]
	if len(vec) == 0 {
		return ""
	}
	return vec[rand.Intn(len(vec)-1)]
}


/* String() functions */
func (t *triple) String() {
	fmt.Printf("%s %s -> %s", t.First, t.Second, t.Third)
}
