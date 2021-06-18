package middleware

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type RouterContext struct {
	Response http.ResponseWriter
	Request  *http.Request
	Logger   *zap.Logger
}

type HandlerFunction func(context RouterContext)

type ContextOptions struct {
	Logger *zap.Logger
}

func NewContextOptions(logger *zap.Logger) ContextOptions {
	ctxOptions := ContextOptions{
		Logger: logger,
	}

	return ctxOptions
}

func NewRoute(handlerFn HandlerFunction, ctxOptions ContextOptions) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		context := RouterContext{
			Logger:   ctxOptions.Logger,
			Response: rw,
			Request:  r,
		}

		reqStart := time.Now()
		defer func() {
			reqEnd := time.Since(reqStart)
			context.Logger.Info(
				fmt.Sprintf("%s %s", r.Method, r.URL.Path),
				zap.Duration("duration_ms", reqEnd*time.Millisecond),
				zap.Int64("content_length", r.ContentLength),
			)
		}()

		handlerFn(context)
		r.Close = true
	}
}
