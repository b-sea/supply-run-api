package mock

import (
	"encoding/json"
	"net/http"
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

func (r *Recorder) Handler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(writer).Encode("hello, I am meant to display a metrics page")
	})
}

func (r *Recorder) ObserveRequestDuration(method string, path string, code string, duration time.Duration) {
	logrus.Infof("Request Duration: %s %s [%s] | %s", method, path, code, duration)
}

func (r *Recorder) ObserveResponseSize(method string, path string, code string, bytes int64) {
	logrus.Infof("Response Size: %s %s [%s] | %v", method, path, code, bytes)
}

func (r *Recorder) RequestAuthorized(username string) {
	logrus.Infof("Request Authorized: %s", username)
}

func (r *Recorder) ObserveResolverDuration(object string, field string, status string, duration time.Duration) {
	logrus.Infof("Resolver Duration: %s.%s, [%s] %s", object, field, status, duration)
}

func (r *Recorder) ObserveResolverError(object string, field string, code string) {
	logrus.Infof("Resolver Error: %s.%s | %s", object, field, code)
}
