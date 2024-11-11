package middleware

// Inspired by go-chi https://github.com/go-chi/chi/blob/master/middleware/wrap_writer.go

import "net/http"

// WrapResponseWriter кастомный ResponseWriter для переопределения методов
type WrapResponseWriter interface {
	http.ResponseWriter
	Status() int
	BytesWritten() int
}

type wrapResponseWriter struct {
	http.ResponseWriter
	statusCode int
	bytes      int
}

// NewWrapResponseWriter создает кастомный ResponseWriter
func NewWrapResponseWriter(w http.ResponseWriter) WrapResponseWriter {
	return &wrapResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		bytes:          0,
	}
}

// WriteHeader overrides WriteHeader to capture status code
func (w *wrapResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Write captures the number of bytes written to the response.
func (w *wrapResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

// Status returns the captured status code
func (w *wrapResponseWriter) Status() int {
	return w.statusCode
}

// BytesWritten returns the captured bytes
func (w *wrapResponseWriter) BytesWritten() int {
	return w.bytes
}
