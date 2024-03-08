package excel

import (
	"regexp"
	"strings"
)

var reSlug = regexp.MustCompile(`[^a-zA-Z0-9 ]`)

func Slug(s string) string {
	s = reSlug.ReplaceAllString(s, "")
	words := strings.Split(s, " ")
	for i, w := range words {
		if len(w) == 0 {
			continue
		}
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, "")
}
