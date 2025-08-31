package mock

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/server"
	"github.com/sirupsen/logrus"
)

var (
	_ server.Recorder = (*Recorder)(nil)
	_ auth.Recorder   = (*Recorder)(nil)
)

type Recorder struct{}

func (r *Recorder) ObserveRequestDuration(method string, path string, code string, duration time.Duration) {
	logrus.Infof("Request Duration: %s, %s, %s, %s", method, path, code, duration)
}

func (r *Recorder) ObserveResponseSize(method string, path string, code string, bytes int64) {
	logrus.Infof("Response Size: %s %s %s %v", method, path, code, bytes)
}

func (r *Recorder) RequestAuthorized(username string) {
	logrus.Infof("Request Authorized: %s", username)
}
