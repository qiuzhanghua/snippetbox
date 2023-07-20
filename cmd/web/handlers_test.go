package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"taiji.dev/snippetbox/internal/assert"
	"testing"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ping", nil)
	ping(rr, r)
	rs := rr.Result()
	assert.Equal(t, rs.StatusCode, http.StatusOK)
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	assert.Equal(t, err, nil)
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
