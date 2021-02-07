package errcode

/**
* @Author: super
* @Date: 2021-02-06 19:41
* @Description:
**/

var (
	ErrorSaveFail   = NewError(30060001, "存储数据到ETCD失败")
	ErrorDeleteFail = NewError(30060002, "ETCD删除数据失败")
	ErrorListFail   = NewError(30060003, "ETCD获取数据列表失败")
)
