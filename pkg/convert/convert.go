package convert

import "strconv"

/**
* @Author: super
* @Date: 2020-09-18 10:10
* @Description: 用于处理数据转换
**/

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

func (s StrTo) MustInt64() int64 {
	v := s.MustInt()
	return int64(v)
}
