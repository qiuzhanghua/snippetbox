package main

import (
	"bytes"
	"io"
	"log"
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

func TestPingE2E(t *testing.T) {
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/ping")
	assert.Equal(t, err, nil)
	assert.Equal(t, rs.StatusCode, http.StatusOK)
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	assert.Equal(t, err, nil)
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
