package uniqsym

import (
	"regexp"
	"testing"
)

func TestRegexMatch(t *testing.T) {
	var sunnyTests = []struct {
		name string
		str  string
	}{
		{"single char", "a"},
		{"two chars", "ab"},
		{"single-char segments", "a.b"},
		{"two-char segments", "aa.bb"},
	}
	regex := regexp.MustCompile(`^` + regex + `$`)
	for _, test := range sunnyTests {
		t.Run(test.name, func(t *testing.T) {
			isMatched := regex.MatchString(test.str)
			if !isMatched {
				t.Error("unexpected mismatch")
			}
		})
	}
}

func TestRegexMismatch(t *testing.T) {
	var rainyTests = []struct {
		name string
		str  string
	}{
		{"empty string", ""},
		{"single sep", "."},
		{"two seps", ".."},
		{"char beetween seps", ".a."},
		{"empty sym", "a."},
		{"empty ns", ".a"},
	}
	regex := regexp.MustCompile(`^` + regex + `$`)
	for _, test := range rainyTests {
		t.Run(test.name, func(t *testing.T) {
			isMatched := regex.MatchString(test.str)
			if isMatched {
				t.Error("unexpected match")
			}
		})
	}
}
