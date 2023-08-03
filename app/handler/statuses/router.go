package statuses

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

// Create Handler for `/v1/statuses/`
func NewRouter(ar repository.Account, sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, sr}
	r.With(auth.Middleware(ar)).Post("/", h.Create)

	return r
}