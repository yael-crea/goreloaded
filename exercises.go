package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func toDecimal(w string, base int) string {
	value, err := strconv.ParseInt(w, base, 64)
	if err != nil {
		return w
	}
	return strconv.FormatInt(value, 10)
}

func capitalize(s string) string {
	runes := []rune(strings.ToLower(s))
	if len(runes) == 0 {
		return ""
	}

	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

var (
	commaSpacing   = regexp.MustCompile(`\s*,\s*`)
	ellipsisBefore = regexp.MustCompile(`\s+(\.\.\.)`)
)

func normalizePunctuation(s string) string {
	s = commaSpacing.ReplaceAllString(s, ", ")
	return ellipsisBefore.ReplaceAllString(s, "$1")
}

func collerVirgules(s string) string {
	return normalizePunctuation(s)
}

func upLast(words []string, n int) {
	if n <= 0 {
		return
	}
	if n > len(words) {
		n = len(words)
	}

	for i := len(words) - n; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
}

