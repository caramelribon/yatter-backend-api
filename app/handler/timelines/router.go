package timelines

import (
	"net/http"

	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
	sr repository.Status
}

// Create Handler for `/v1/timelines`
func NewRouter(ar repository.Account, sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, sr}
	r.Get("/public", h.GetPublicTimelines)
	r.With(auth.Middleware(ar)).Get("/home", h.GetHomeTimelines)

	return r
}