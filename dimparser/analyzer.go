package dimparser

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

var misleadingWords = []string{"座面高さ"}

// analyze given string and returns cleaned string
func analyze(s string) string {
	s = string(norm.NFKC.Bytes([]byte(s)))
	for _, w := range misleadingWords {
		s = strings.ReplaceAll(s, w, "#")
	}
	return s
}
