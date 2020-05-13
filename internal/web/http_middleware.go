package web

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/rs/xid"
)

// CtxKey ...
type CtxKey string

const (
	traceIDKey CtxKey = "traceID"
)

func applyMiddlewares(r http.Handler) http.Handler {
	r = logging(newLogger(), r)
	r = tracing(generateID, r)
	return r
}

// TODO
func rateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check
		next.ServeHTTP(w, r)
	})
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "http: ", log.LstdFlags)
}

func logging(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			traceID, ok := r.Context().Value(traceIDKey).(string)
			if !ok {
				traceID = "unknown"
			}
			logger.Println(traceID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		}()
		next.ServeHTTP(w, r)
	})
}

func generateID() string {
	guid := xid.New()
	return guid.String()
}

func tracing(nextTraceID func() string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get("X-Request-Id")
		if traceID == "" {
			traceID = nextTraceID()
		}
		ctx := context.WithValue(r.Context(), traceIDKey, traceID)
		w.Header().Set("X-Request-Id", traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
