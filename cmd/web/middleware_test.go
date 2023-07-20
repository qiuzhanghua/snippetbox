package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"taiji.dev/snippetbox/internal/assert"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	secureHeaders(next).ServeHTTP(rr, r)
	rs := rr.Result()
	expectedValue := "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expectedValue)
	expectedValue = "deny"
	assert.Equal(t, rs.Header.Get("X-Frame-Options"), expectedValue)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	assert.Equal(t, err, nil)
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
