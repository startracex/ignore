package ignore

import (
	"regexp"
	"strings"
)

// Match returns matching results
func Match(origin, target string) bool {
	if len(origin) < 1 || len(target) < 1 {
		return false
	}

	if strings.HasSuffix(origin, "/") {
		origin = origin + "*"
	}

	// if top level
	if strings.HasPrefix(origin, "/") || strings.HasPrefix(origin, "!/") {
		return MatchTop(origin, target)
	}

	regstring := QuoteNoBrackets(origin)

	regstring = GlobToRegexp(regstring)

	reg, err := regexp.Match(regstring, []byte(target))
	if err != nil {
		return false
	}
	return reg
}

// MatchTop receive "/" and "!/"
func MatchTop(origin, target string) bool {
	// "!/"
	if origin[0] == '!' {
		return !Match(origin[2:], target)
	}
	// "/"
	origin = origin[1:]
	regstring := QuoteNoBrackets(origin)

	regstring = GlobToRegexp(regstring)
	if !ContainsQuite(regstring, "*") {
		regstring = "^" + regstring
	}
	m, err := regexp.Match(regstring, []byte(target))
	if err != nil {
		return false
	}
	return m
}

// QuoteNoBrackets quote string without brackets
func QuoteNoBrackets(src string) string {
	src = regexp.QuoteMeta(src)
	src = ReplaceQuote(src, "[", "[", -1)
	return ReplaceQuote(src, "]", "]", -1)
}

// ReplaceQuote replace quoted old to new
func ReplaceQuote(s, old, new string, n int) string {
	return strings.Replace(s, regexp.QuoteMeta(old), new, n)
}

func ContainsQuite(s string, substr string) bool {
	return strings.Contains(s, regexp.QuoteMeta(substr))
}

// GlobToRegexp
func GlobToRegexp(glob string) string {
	glob = ReplaceQuote(glob, "**", ".*", -1)
	glob = ReplaceQuote(glob, "*", "[^/]*", -1)
	glob = ReplaceQuote(glob, "?", ".", -1)
	for i := 0; i < len(glob); i++ {
		if glob[i] == '!' {
			if i+1 >= len(glob) {
				return "^"
			} else if glob[i+1] == '[' {
				return "^"
			}
			return "[^" + glob[i+1:] + "]"
		}
	}

	return glob
}
