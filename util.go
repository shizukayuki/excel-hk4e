package excel

import (
	"regexp"
	"slices"
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

func SlugLower(s string) string {
	return strings.ToLower(Slug(s))
}

func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	var r S
	for i := range s {
		if f(s[i]) {
			r = append(r, s[i])
		}
	}
	return r
}

func Find[S ~[]E, E any](s S, f func(E) bool) E {
	if i := slices.IndexFunc(s, f); i != -1 {
		return s[i]
	}
	return *new(E)
}
