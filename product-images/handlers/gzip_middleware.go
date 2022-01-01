package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type WrappedResponseWriter struct {
	http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappendResponseWriter(rw http.ResponseWriter, r *http.Request) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WrappedResponseWriter{gw: gw}
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			wrw := NewWrappendResponseWriter(rw, r)
			wrw.Header().Set("Content-Encoding", "gzip")
			next.ServeHTTP(wrw, r)
			defer wrw.Flush()
			return
		}
		next.ServeHTTP(rw, r)
	})
}
