package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	userTest := map[string]string{
		"email":    "test@example.com",
		"password": "test_password",
	}

	status, _, js := ts.post(t, "/api/v1/users", userTest)
	assert.Equal(t, http.StatusCreated, status)
	assert.Equal(t, float64(1), js["id"])
}
