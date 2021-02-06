package middleware

import (
	"context"
	"go-crawler-distributed/global"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

/**
* @Author: super
* @Date: 2020-09-24 08:14
* @Description: 调用链追踪中间件，调用结果可在jaeger ui看到
**/

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}
		defer span.Finish()

		var traceID string
		var spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
