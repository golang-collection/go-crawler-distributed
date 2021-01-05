package util

import "testing"

/**
* @Author: super
* @Date: 2020-09-24 20:08
* @Description:
**/

func TestEncodeBase64(t *testing.T) {
	s, err := EncodeBase64("hello world")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestDecodeBase64(t *testing.T) {
	s := DecodeBase64("H4sIAAAAAAAA/8pIzcnJVyjPL8pJAQAAAP//AQAA//+FEUoNCwAAAA==")
	t.Log(s)
}

func BenchmarkEncodeBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := EncodeBase64("helloworldasdafsdfasfsdgadfgadfweaweterteggdfsgdsbdfbvxvczxvfasdfasdfasdfsadfsadfsadfsd")
		if err != nil {
			b.Error(err)
		}
		b.Log(s)
	}
}
