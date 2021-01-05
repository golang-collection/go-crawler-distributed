package global

import "github.com/opentracing/opentracing-go"

/**
* @Author: super
* @Date: 2020-09-24 08:10
* @Description: 配置全局统一的调用链追踪
**/

var (
	Tracer opentracing.Tracer
)
