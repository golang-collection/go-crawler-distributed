package middleware

import (
	"fmt"
	"time"

	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/app"
	"go-crawler-distributed/pkg/email"
	"go-crawler-distributed/pkg/errcode"

	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-09-23 20:45
* @Description: 自定义recovery，主要用于记录异常发生的时间以及错误信息
**/

func Recovery() gin.HandlerFunc {
	mailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)

				err := mailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
