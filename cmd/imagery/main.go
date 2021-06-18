package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func withMiddleware(handlerFn http.HandlerFunc) http.HandlerFunc {
	if logger == nil {
		logger, _ = zap.NewProduction()
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		reqStart := time.Now()
		go func() {
			handlerFn(rw, r)
			wg.Done()
		}()
		reqEnd := time.Since(reqStart)
		logger.Info(
			fmt.Sprintf("%s %s", r.Method, r.URL.Path),
			zap.Duration("duration_ms", reqEnd*time.Millisecond),
			zap.Int64("content_length", r.ContentLength),
		)
		wg.Wait()
	}
}

// Handles: GET /
func baseHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(200)
	body, _ := json.Marshal(map[string]interface{}{"name": "Leoni"})
	rw.Write(body)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", withMiddleware(baseHandler))

	logger.Info("Server started on port :4000")
	http.ListenAndServe(":4000", router)
}
