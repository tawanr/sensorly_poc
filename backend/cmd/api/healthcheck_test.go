package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/api/v1/health")
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, body["status"], "available")
	assert.Equal(t, body["environment"], "test")
	assert.Equal(t, body["version"], version)
}
