package server_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/metrics"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestWithPort(t *testing.T) {
	testServer := server.New(metrics.NewNoOp())
	server.WithPort(4567)(testServer)

	assert.Equal(t, ":4567", testServer.Addr())
}

func TestWithReadTimeout(t *testing.T) {
	testServer := server.New(metrics.NewNoOp())
	server.WithReadTimeout(time.Hour)(testServer)

	assert.Equal(t, time.Hour, testServer.ReadTimeout())
}

func TestWithWriteTimeout(t *testing.T) {
	testServer := server.New(metrics.NewNoOp())
	server.WithWriteTimeout(time.Hour)(testServer)

	assert.Equal(t, time.Hour, testServer.WriteTimeout())
}
