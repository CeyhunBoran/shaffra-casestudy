package logging

import (
	"log"
	"net/http"
	"time"
)

// This uses a channel to signal when the request has completed,
// allowing the logging to happen asynchronously without blocking the handler execution.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		ch := make(chan bool)
		go func() {
			next.ServeHTTP(w, r)
			ch <- true
		}()
		<-ch
		duration := time.Since(startTime)
		log.Printf("%s %s took %v\n", r.Method, r.URL.String(), duration)
	})
}
