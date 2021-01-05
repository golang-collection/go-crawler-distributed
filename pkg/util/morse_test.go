package util

import (
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-24 10:04
* @Description:
**/

type Morse struct {
	str  string
	real string
	err  error
}

func TestGenerateMorse(t *testing.T) {
	var morses = []struct {
		str  string
		code string
		err  error
	}{
		{"aa11", ".-.-.----.----", nil},
		{"11aa", ".----.----.-.-", nil},
		{"", "", lengthError},
		{"111,as", "", unsupportedError},
		{"中文", "", unsupportedError},
		{"1a12 ", ".----.-.----..---", nil},
		{"   ", "", lengthError},
		{"asdj$%#, 441", "", unsupportedError},
		{"!@#$", "", unsupportedError},
	}

	for i, v := range morses {
		code, e := GenerateMorse(v.str)
		if code != v.code {
			t.Errorf("%d. %s morse code %s, wanted: %s, error= %v", i, v.str, code, v.code, e)
		} else if e != v.err {
			t.Errorf("%d. %s morse code %s, wanted: %s, error= %v", i, v.str, code, v.code, e)
		}
	}
}

func BenchmarkGenerateMorse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateMorse("asasd12454")
	}
}
