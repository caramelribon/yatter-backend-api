package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
	sr repository.Status
}

// Create Handler for `/v1/timelines/public`
func NewRouter(ar repository.Account, sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, sr}
	r.Get("/public", h.GetPublicTimelines)

	return r
}