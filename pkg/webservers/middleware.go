package webservers

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of HTTP requests and the time taken to handle them
func loggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            defer func() {
                logger.Printf("Handled request: %s %s, from: %s, took: %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
            }()
            next.ServeHTTP(w, r)
        })
    }
}

func corsMiddleware() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Set headers
            w.Header().Set("Access-Control-Allow-Origin", "*") // or specify your domain
            w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

            // If it's a preflight OPTIONS request, handle it
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}


// RecoverMiddleware recovers from any panics in the HTTP handlers and logs an error message
func recoverMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            defer func() {
                if err := recover(); err != nil {
                    logger.Printf("Recovered from a panic: %v", err)
                    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                }
            }()
            next.ServeHTTP(w, r)
        })
    }
}

// BasicAuthMiddleware provides basic authentication layer to protect sensitive routes
func basicAuthMiddleware(username, password string, logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            user, pass, ok := r.BasicAuth()
            if !ok || user != username || pass != password {
                logger.Printf("Unauthorized access attempt: %s", r.RemoteAddr)
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
