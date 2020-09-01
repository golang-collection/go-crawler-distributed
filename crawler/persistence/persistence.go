package persistence

/**
* @Author: super
* @Date: 2020-08-16 09:01
* @Description:
**/
type ParseStorage func(interface{}) error

type FuncStorage struct {
	Name      string
	ParseFunc ParseStorage
}
