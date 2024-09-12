package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/")
	assert.Equal(t, http.StatusNotFound, code)
	assert.Equal(t, body["error"], "The requested resource could not be found.")
}
