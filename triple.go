package gomarkov

import (
//	"rand"
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

/* String() functions */
func (t *triple) String() {
	fmt.Printf("%s %s -> %s", t.First, t.Second, t.Third)
}
