package tea

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/go-chi/chi/middleware"
)

// Body parses a JSON request body and validates it's contents
func Body(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return Validate.StructCtx(r.Context(), v)
}

type key int

const loggerKey key = 1

// NewLogger creates a new logger from an *http.Request. It is used
// to create a new log.Interface for middleware and when an existing
// logger does not exist. It is exposed so that it can be overridden
//
// The default NewLogger creates a new logger with the request.id
// from middleware.GetReqID, and the client.ip from r.RemoteAddr
var NewLogger = func(r *http.Request) log.Interface {
	return log.WithFields(log.Fields{
		"request.id": middleware.GetReqID(r.Context()),
		"client.ip":  r.RemoteAddr,
	})
}

// Logger pulls a logger from request context or creates a new one
// with NewLogger if one does not exist
func Logger(r *http.Request) log.Interface {
	ctx := r.Context()
	if ll, ok := ctx.Value(loggerKey).(log.Interface); ok {
		return ll
	}
	return NewLogger(r)
}

// LoggerCtx uses Logger to get a log.Interface and appends it to
// the request context, returning the new context with the logger value
// to be retrieved within Handlers
func LoggerCtx(r *http.Request) (log.Interface, context.Context) {
	l := Logger(r)
	return l, context.WithValue(r.Context(), loggerKey, l)
}

// LoggerMiddleware appends a logger with request information using
// LoggerCtx to every request
func LoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, ctx := LoggerCtx(r)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
