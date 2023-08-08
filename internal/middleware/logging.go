package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	constants "github.com/deanrtaylor1/backend-go/internal/constants"
)

type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func colorLog(method string, path string, statusCode int, duration time.Duration) string {
	ms := float64(duration) / float64(time.Millisecond)
	switch {
	case statusCode >= 500:
		return constants.BgRed + method + constants.Reset + " " + path + " " + constants.BgRed + fmt.Sprintf("%d", statusCode) + constants.Reset + " " + fmt.Sprintf("%.3fms", ms)
	case statusCode >= 400:
		return constants.BgYellow + method + constants.Reset + " " + path + " " + constants.BgYellow + fmt.Sprintf("%d", statusCode) + constants.Reset + " " + fmt.Sprintf("%.3fms", ms)
	case statusCode >= 300:
		return constants.BgCyan + method + constants.Reset + " " + path + " " + constants.BgCyan + fmt.Sprintf("%d", statusCode) + constants.Reset + " " + fmt.Sprintf("%.3fms", ms)
	default:
		return constants.BgGreen + method + constants.Reset + " " + path + " " + constants.BgGreen + fmt.Sprintf("%d", statusCode) + constants.Reset + " " + fmt.Sprintf("%.3fms", ms)
	}
}

func ColorLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		srw := &statusResponseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(srw, r)

		duration := time.Since(start)

		method := r.Method
		path := r.RequestURI
		statusCode := srw.status

		log.Println(colorLog(method, path, statusCode, duration))
	})
}
