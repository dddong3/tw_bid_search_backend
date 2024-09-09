package middlewares

import (
	"time"
	"net/http"
	"github.com/dddong3/Bid_Backend/logger"
)

func TimingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()
		durationMs := float64(duration.Milliseconds())
		method := r.Method
		path := r.URL.Path

		logger.Logger.Info("Request processed",
			" | remoteAddr:", remoteAddr,
			" | method:", method,
			" | path:", path,
			" | durationMs:", durationMs,
			" | userAgent:", userAgent,
		)

	})
}