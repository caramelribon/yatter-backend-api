package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"yatter-backend-go/app/handler/auth"
)

// Handle request for `DELETE /v1/statuses/{id}`
func (h *handler) DeleteStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get id from url
	paramId := chi.URLParam(r, "id")
	if paramId == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// id conversion from string to int64
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get account
	account := auth.AccountOf(r)

	// get status from id
	status, err := h.sr.FindById(ctx, id); 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == nil || (status != nil && status.Account.ID != account.ID) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// delete status
	if err := h.sr.DeleteStatus(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("Complete delete status"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
