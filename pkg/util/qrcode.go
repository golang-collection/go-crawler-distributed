package util

/**
* @Author: super
* @Date: 2020-08-21 22:08
* @Description:
**/

import "github.com/skip2/go-qrcode"

func GenerateQRCodeByte(str string) ([]byte, error) {
	return qrcode.Encode(str, qrcode.Highest, 256)
}
