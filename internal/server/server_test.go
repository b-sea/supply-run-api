package server_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestServerStartStop(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	assert.NoError(t, err)

	listener, err := net.ListenTCP("tcp", addr)
	assert.NoError(t, err)

	port := listener.Addr().(*net.TCPAddr).Port
	assert.NoError(t, listener.Close())

	testServer := server.New(
		zerolog.Nop(),
		metrics.NewNoOp(),
		server.WithPort(port),
	)

	timer := time.NewTimer(500 * time.Millisecond)

	go func() {
		assert.NoError(t, testServer.Start())
	}()

	<-timer.C

	assert.NoError(t, testServer.Stop())
}

func TestServerMetrics(t *testing.T) {
	testServer := httptest.NewServer(
		server.New(
			zerolog.Nop(),
			metrics.NewNoOp(),
		),
	)

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
	testServer := httptest.NewServer(
		server.New(
			zerolog.Nop(),
			metrics.NewNoOp(),
		),
	)

	request, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/graphql", testServer.URL),
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
