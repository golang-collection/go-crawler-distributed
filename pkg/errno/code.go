package errno

/**
* @Author: super
* @Date: 2020-08-26 18:03
* @Description: code
         错误码设计原则 10002
         1  服务级错误 1为系统级错误 2为普通错误，通常是非法用户造成
         00 服务模块为2位数
         02 错误码为2位数
         code = 0说明正确
**/

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
)
