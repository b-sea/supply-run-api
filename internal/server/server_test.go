package server_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestServerStartStop(t *testing.T) {
	testServer := server.New(metrics.NewBasicLogger())

	timer := time.NewTimer(500 * time.Millisecond)

	go func() {
		assert.NoError(t, testServer.Start())
	}()

	<-timer.C

	assert.NoError(t, testServer.Stop())
}

func TestWithPort(t *testing.T) {
	testServer := server.New(metrics.NewBasicLogger())
	server.WithPort(4567)(testServer)

	assert.Equal(t, ":4567", testServer.Addr())
}

func TestWithReadTimeout(t *testing.T) {
	testServer := server.New(metrics.NewBasicLogger())
	server.WithReadTimeout(time.Hour)(testServer)

	assert.Equal(t, time.Hour, testServer.ReadTimeout())
}

func TestWithWriteTimeout(t *testing.T) {
	testServer := server.New(metrics.NewBasicLogger())
	server.WithWriteTimeout(time.Hour)(testServer)

	assert.Equal(t, time.Hour, testServer.WriteTimeout())
}

func TestServerMetrics(t *testing.T) {
	testServer := httptest.NewServer(server.New(metrics.NewBasicLogger()))

	request, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/metrics", testServer.URL),
		nil,
	)

	request.Close = true

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	assert.NoError(t, response.Body.Close())

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "", response.Header.Get("Content-Type"))
	assert.Equal(t, ``, string(body))

	testServer.Close()
}

func TestServerAPIGraphql(t *testing.T) {
	testServer := httptest.NewServer(server.New(metrics.NewBasicLogger()))

	request, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/api/graphql", testServer.URL),
		strings.NewReader(`{}`),
	)

	request.Close = true
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	assert.NoError(t, response.Body.Close())

	assert.Equal(t, http.StatusUnprocessableEntity, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(
		t,
		`{"errors":[{"message":"no operation provided","extensions":{"code":"GRAPHQL_VALIDATION_FAILED"}}],"data":null}`,
		string(body),
	)

	testServer.Close()
}
