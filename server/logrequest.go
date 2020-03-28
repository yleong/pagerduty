package server

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

//logRequest records metrics for each client request/response
func logRequest(w *responseWithStatus, r *http.Request, p *prometheus.CounterVec) {
	path := r.URL.Path
	status := w.status
	p.WithLabelValues(path, fmt.Sprintf("%v", status)).Inc()
}

//responseWithStatus is needed because there is no way to extract StatusCode from http.ResponseWriter
type responseWithStatus struct {
	http.ResponseWriter
	status int
}

//WriteHeader is overridden to also store the status code for later retrieval
func (r *responseWithStatus) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

//makeHandler converts our custom handlers (that uses responseWithStatus) into standard http handlers.
func makeHandler(fn func(*responseWithStatus, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w2 := &responseWithStatus{w, 0}
		defer logRequest(w2, r, clientRequests)
		fn(w2, r)
	}
}
