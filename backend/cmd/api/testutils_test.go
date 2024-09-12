package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.org/tawanr/sensorly/internal/data/mocks"
)

func newTestConfig(t *testing.T) *config {
	cfg := &config{
		port: 8000,
		env:  "test",
	}
	return cfg
}

func newTestApplication(t *testing.T) *application {
	cfg := newTestConfig(t)
	models := mocks.NewMockModels()
	app := &application{logger: slog.New(slog.NewTextHandler(io.Discard, nil)), config: cfg, models: models}
	return app
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, path string) (int, http.Header, map[string]string) {
	res, err := http.Get(ts.URL + path)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	status := res.StatusCode
	headers := res.Header

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var js map[string]string
	json.Unmarshal(body, &js)
	return status, headers, js
}

func (ts *testServer) post(t *testing.T, path string, body interface{}) (int, http.Header, map[string]any) {
	payload, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post(ts.URL+path, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	status := res.StatusCode
	headers := res.Header

	response, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var js map[string]any
	json.Unmarshal(response, &js)
	return status, headers, js
}
