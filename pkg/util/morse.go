package util

import (
	"errors"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-24 09:33
* @Description:
**/

var lengthError = errors.New("length must > 1")
var unsupportedError = errors.New("unsupported byte")

var morseMap = map[byte]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",
}

func GenerateMorse(str string) (string, error) {
	str = strings.TrimSpace(str)
	length := len(str)
	if length == 0 {
		return "", lengthError
	}
	var builder strings.Builder
	bytes := []byte(str)
	for _, v := range bytes {
		if value, ok := morseMap[v]; ok {
			builder.WriteString(value)
		} else {
			return "", unsupportedError
		}
	}
	return builder.String(), nil
}
