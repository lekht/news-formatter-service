package format

import "strings"

type Formatter struct {
	words map[string]bool
}

var forbiddenWords = []string{
	"qwerty",
	"йцукен",
	"zxcvb",
}

func New() *Formatter {
	f := &Formatter{}

	wordBook := make(map[string]bool)

	for _, word := range forbiddenWords {
		wordBook[word] = true
	}

	f.words = wordBook

	return f
}

func (f *Formatter) CheckWord(s string) bool {

	words := strings.Fields(s)

	for _, word := range words {
		if _, ok := f.words[word]; ok {
			return true
		}
	}
	return false
}
