package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/dddong3/Bid_Backend/logger"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetRealIP(r *http.Request) string {
	cfConnectingIP := r.Header.Get("Cf-Connecting-Ip")
	if cfConnectingIP != "" {
		return cfConnectingIP
	}

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
		var graphqlPayload map[string]interface{}
		if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
			body, err := io.ReadAll(r.Body)
			if err == nil {
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				json.Unmarshal(body, &graphqlPayload)
			}
		}
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

		if graphqlPayload != nil {
			query, _ := graphqlPayload["query"].(string)
			variables, _ := graphqlPayload["variables"].(map[string]interface{})
			logger.Logger.Info("GraphQL Payload",
				" | query:", query,
				" | variables:", variables,
			)
		}
	})
}
