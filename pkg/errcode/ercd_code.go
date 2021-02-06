package errcode

/**
* @Author: super
* @Date: 2021-02-06 19:41
* @Description:
**/

var (
	ErrorSaveFail   = NewError(30060001, "存储数据到ETCD失败")
)