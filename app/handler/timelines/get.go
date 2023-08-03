package timelines

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"

	"yatter-backend-go/app/domain/object"
)

type getResponse struct {
	ID       int64           `json:"id"`
	Account  *object.Account `json:"account"`
	Content  string          `json:"content"`
	CreateAt time.Time       `json:"create_at"`
}

// Handle request for `Get /v1/statuses/{id}`
func (h *handler) GetPublicTimelines(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
