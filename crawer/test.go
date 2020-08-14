package main

import (
	"fmt"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-14 20:25
* @Description:
**/

func main() {
	toBeChange := "2003-8"
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, _ := time.ParseInLocation(timeLayout, toBeChange, loc)
	fmt.Println(theTime.Year())
	fmt.Println(theTime.Month())
	fmt.Println(theTime.Day())
}
