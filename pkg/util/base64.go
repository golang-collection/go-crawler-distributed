package util

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

/**
* @Author: super
* @Date: 2020-09-24 20:04
* @Description: base64编码与解码
**/

//将文本通过gzip压缩后通过base64编码
func EncodeBase64(value string) (string, error) {
	s := []byte(value)
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(s); err != nil {
		return "", err
	}
	if err := gz.Flush(); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	return str, nil
}

//将编码的base64字符串解码回原文本
func DecodeBase64(value string) string {
	data, _ := base64.StdEncoding.DecodeString(value)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return string(s)
}
