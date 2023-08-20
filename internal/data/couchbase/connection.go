package couchbase

import (
	"fmt"
	"time"

	config "github.com/b-sea/supply-run-api/configs"
	"github.com/couchbase/gocb/v2"
)

const bucket_name = "supply_run"

func NewConnection(cfg config.Config) (*gocb.Cluster, error) {
	opts := gocb.ClusterOptions{
		Username: cfg.Couchbase.User,
		Password: cfg.Couchbase.Pwd,
	}

	cluster, err := gocb.Connect(cfg.Couchbase.URL, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to setup couchbase connection: %w", err)
	}

	if err := cluster.WaitUntilReady(10*time.Second, nil); err != nil {
		return nil, fmt.Errorf("failed to connect to couchbase: %w", err)
	}
	return cluster, nil
}
