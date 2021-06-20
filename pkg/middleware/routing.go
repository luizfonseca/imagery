package middleware

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type HandlerFunction func(context RouterContext)
type RouterContext struct {
	Response http.ResponseWriter
	Request  *http.Request
	Logger   *zap.Logger
	Fetch    FetchFunc
}

type ContextOptions struct {
	Logger *zap.Logger
	Client *http.Client
}

type FetchOptions struct {
	Method  string
	Url     string
	Headers map[string]interface{}
}
type FetchFunc func(options FetchOptions) *http.Response

func NewContextOptions(logger *zap.Logger, client *http.Client) ContextOptions {
	ctxOptions := ContextOptions{
		Logger: logger,
		Client: client,
	}

	return ctxOptions
}

func NewFetch(ctxOptions ContextOptions) FetchFunc {
	return func(options FetchOptions) *http.Response {
		req, _ := http.NewRequest(options.Method, options.Url, nil)
		res, _ := ctxOptions.Client.Do(req)

		return res
	}
}

func NewRoute(handlerFn HandlerFunction, ctxOptions ContextOptions) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		context := RouterContext{
			Logger:   ctxOptions.Logger,
			Response: rw,
			Request:  r,
			Fetch:    NewFetch(ctxOptions),
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
