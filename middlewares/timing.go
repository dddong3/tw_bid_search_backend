package middlewares

import (
	"time"
	"net/http"
	"strings"
	"github.com/dddong3/Bid_Backend/logger"
)

func GetRealIP(r *http.Request) string {
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	return r.RemoteAddr
}

func TimingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		remoteAddr := GetRealIP(r)
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