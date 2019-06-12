package authentication

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockHeadRequest(t *testing.T, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Errorf("Expected request method: %q, but got: %q", http.MethodHead, r.Method)
		}
		w.WriteHeader(statusCode)
	}))
}

func TestExtractAuthorizationHeaderFromContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), authKey, "Bearer 123")
	actualToken, _ := extractAuthorizationHeaderFromContext(ctx)
	if actualToken != "123" {
		t.Fatalf("Expected token: %q, but got: %q", "123", actualToken)
	}
}
