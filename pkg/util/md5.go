package util

import (
	"crypto/md5"
	"encoding/hex"
)

/**
* @Author: super
* @Date: 2020-09-23 18:59
* @Description:
**/

//字符串md5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
