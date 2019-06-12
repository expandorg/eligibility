package authentication

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type ContextKey string

var authKey ContextKey = "auth"

// AuthMiddleware provides a middleware that retrieves the authentication header
// from the request and inject into the context
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authorization header was not provided"))
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authorization format must be <Bearer 123>"))
			return
		}

		ctx := context.WithValue(r.Context(), authKey, authHeader)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetAuthFromContext retrieves the Authorization from context
func GetAuthFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(authKey)
	switch t := val.(type) {
	case string:
		return t, nil
	default:
		return "", errors.New("missing auth from context")
	}
}
