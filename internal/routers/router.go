package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/middleware"
	"go-crawler-distributed/pkg/limiter"
	//_ "go-crawler-distributed/docs"

	"time"
)

/**
* @Author: super
* @Date: 2021-01-05 14:33
* @Description:
**/

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.AppInfo())
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	//放到需要token的请求中
	//r.Use(middleware.JWT())
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//获取token
	//authManager := dao.NewAuthManager("auth", global.DBEngine)
	//authService := service.NewAuthService(authManager)
	//authController := api.NewAuthController(authService)
	//r.GET("/auth", authController.GetAuth)

	return r
}
