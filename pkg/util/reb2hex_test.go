package util

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-11-30 20:21
* @Description:
**/

func TestRgbToHex(t *testing.T) {
	fmt.Println(RgbToHex(`{
								"red": 12,
								"green": 255,
								"blue": 255
								}`))
}

func BenchmarkRgbToHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = RgbToHex(`{
								"red": 255,
								"green": 255,
								"blue": 255
								}`)
	}
}
