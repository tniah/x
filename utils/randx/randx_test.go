package randx_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tniah/x/utils/randx"
	"regexp"
	"testing"
)

func TestCharset(t *testing.T) {
	for k, v := range []struct {
		charset []rune
		pattern string
	}{
		{randx.AlphaNum, "^[a-zA-Z0-9]{62}$"},
		{randx.Alpha, "^[a-zA-Z]{52}$"},
		{randx.AlphaLowerNum, "^[a-z0-9]{36}$"},
		{randx.AlphaUpperNum, "^[A-Z0-9]{36}$"},
		{randx.AlphaLower, "^[a-z]{26}$"},
		{randx.AlphaUpperVowels, "^[AEIOUY]{6}$"},
		{randx.AlphaUpperNoVowels, "^[^AEIOUY]{20}$"},
		{randx.AlphaUpper, "^[A-Z]{26}$"},
		{randx.Numeric, "^[0-9]{10}$"},
		{randx.SecretCharset, "^[a-zA-Z0-9_.~-]{66}$"},
	} {
		re := regexp.MustCompile(v.pattern)
		valid := re.Match([]byte(string(v.charset)))
		assert.True(t, valid, "Case %d", k+1)
	}
}
