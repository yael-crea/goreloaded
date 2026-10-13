package main

import (
	"reflect"
	"testing"
)

func TestToDecimal(t *testing.T) {
	tests := []struct {
		name string
		word string
		base int
		want string
	}{
		{name: "hexadécimal", word: "ff", base: 16, want: "255"},
		{name: "binaire", word: "1010", base: 2, want: "10"},
		{name: "déjà décimal", word: "42", base: 10, want: "42"},
		{name: "cas limite invalide", word: "zz", base: 16, want: "zz"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := toDecimal(test.word, test.base); got != test.want {
				t.Fatalf("toDecimal(%q, %d) = %q, attendu %q", test.word, test.base, got, test.want)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := map[string]string{
		"éLAN":    "Élan",
		"bonjour": "Bonjour",
		"A":       "A",
		"":        "",
	}

	for input, want := range tests {
		if got := capitalize(input); got != want {
			t.Errorf("capitalize(%q) = %q, attendu %q", input, got, want)
		}
	}
}

func TestNormalizePunctuation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "virgule collée à droite", input: "chat ,chien", want: "chat, chien"},
		{name: "espaces autour de la virgule", input: "un  ,  deux", want: "un, deux"},
		{name: "points de suspension", input: "oui ... non", want: "oui... non"},
		{name: "cas limite sans ponctuation", input: "bonjour", want: "bonjour"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := normalizePunctuation(test.input); got != test.want {
				t.Fatalf("normalizePunctuation(%q) = %q, attendu %q", test.input, got, test.want)
			}
		})
	}
}

func TestUpLast(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		n     int
		want  []string
	}{
		{name: "deux derniers", words: []string{"un", "deux", "trois"}, n: 2, want: []string{"un", "DEUX", "TROIS"}},
		{name: "n trop grand", words: []string{"un", "deux"}, n: 5, want: []string{"UN", "DEUX"}},
		{name: "zéro", words: []string{"un"}, n: 0, want: []string{"un"}},
		{name: "tranche vide", words: []string{}, n: 3, want: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			upLast(test.words, test.n)
			if !reflect.DeepEqual(test.words, test.want) {
				t.Fatalf("upLast donne %#v, attendu %#v", test.words, test.want)
			}
		})
	}
}

