package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// CloseableResponseWriter : used when browser does not support gzip compression
type CloseableResponseWriter interface {
	http.ResponseWriter
	Close()
}

// gzipResponseWriter : used when browser does support gzip compression
type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (w gzipResponseWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

func (w gzipResponseWriter) Close() {
	w.Writer.Close()
}

func (w gzipResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

type closeableResponseWriter struct {
	http.ResponseWriter
}

func (w closeableResponseWriter) Close() {}

// GetResponseWriter : determine if gzip is supported and return proper writer accordingly
func GetResponseWriter(w http.ResponseWriter, req *http.Request) CloseableResponseWriter {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		return closeableResponseWriter{ResponseWriter: w}
	}

	w.Header().Set("Content-Encoding", "gzip")
	gRW := gzipResponseWriter{
		ResponseWriter: w,
		Writer:         gzip.NewWriter(w),
	}

	return gRW
}

// GzipHandler : handler for gzip compression
type GzipHandler struct{}

func (h *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseWriter := GetResponseWriter(w, r)
	defer responseWriter.Close()

	http.DefaultServeMux.ServeHTTP(responseWriter, r)
}
