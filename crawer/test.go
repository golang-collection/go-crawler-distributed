package main

import (
	"fmt"
	"regexp"
)

/**
* @Author: super
* @Date: 2020-08-14 20:25
* @Description:
**/
var DateRe = regexp.MustCompile(`([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8])))`)
var priceRe = regexp.MustCompile(`[0-9]+[.]?[0-9]*`)

func main() {
	v := "2036-2-8"
	dateMatch := DateRe.FindAllSubmatch([]byte(v), -1)
	if len(dateMatch) == 0 {
		v = "2006-01-02"
	}
	fmt.Println(v)
}
