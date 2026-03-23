package handler

import (
	"encoding/json"
	"net"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ip := getClientIP(r)

	response := map[string]interface{}{
		"method":     r.Method,
		"url":        r.URL.String(),
		"proto":      r.Proto,
		"remoteAddr": r.RemoteAddr,
		"clientIP":   ip,
		"headers":    r.Header,
		"host":       r.Host,
		"userAgent":  r.UserAgent(),
		"referer":    r.Referer(),
		"query":      r.URL.Query(),
	}

	json.NewEncoder(w).Encode(response)
}

// Extract real client IP (handles proxies/load balancers)
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For (can contain multiple IPs)
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	// Check X-Real-IP
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// Fallback to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}