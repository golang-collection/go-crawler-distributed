package handler

import (
	"github.com/gin-gonic/gin"
	"go-crawler-distributed/pkg/errno"
	"net/http"
)

/**
* @Author: super
* @Date: 2020-09-03 20:56
* @Description:
**/

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}