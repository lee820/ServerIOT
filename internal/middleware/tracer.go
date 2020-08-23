package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lee820/ServerIOT/global"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

//Tracing jaegerTracer中间件
func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
				global.Tracer, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
				global.Tracer, c.Request.URL.Path)
		}
		defer span.Finish()

		//记录spanID和traceID
		var SpanID string
		var traceID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.(jaeger.SpanContext).TraceID().String()
			SpanID = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", SpanID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
