package helper

import (
	"log"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (wr *ResponseWriter) WriteHeader(statusCode int) {
	wr.StatusCode = statusCode
	wr.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &ResponseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("%s  %d  %s  %s", r.Method, rw.StatusCode, r.URL.Path, time.Since(start))
	}
}
