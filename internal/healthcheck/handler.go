// Package healthcheck handles the logic of simple endpoints to check the availability of the service.
package healthcheck

import "net/http"

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
