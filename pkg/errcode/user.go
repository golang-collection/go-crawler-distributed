package errcode

/**
* @Author: super
* @Date: 2020-11-24 18:45
* @Description:
**/

var (
	ErrorUserSignInFail   = NewError(20060001, "用户登录失败")
	ErrorUserUpdateFail   = NewError(20060002, "用户信息更新失败")
	ErrorUserRegisterFail = NewError(20060003, "用户注册失败")
	ErrorUserCookieFail   = NewError(20060004, "用户cookie验证失败")
)
