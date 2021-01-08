package errcode

/**
* @Author: super
* @Date: 2020-09-18 07:53
* @Description: 统一错误代码
**/

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFound      = NewError(10000002, "找不到")
)
