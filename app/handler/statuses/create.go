package statuses

import (
	"encoding/json"
	"net/http"
	
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Status string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)
	status.Content = req.Status

	// アカウント情報の取得
	account := auth.AccountOf(r)
	status.AccountID = account.ID
	status.Account = account

	// ステータスの作成
	if err := h.sr.CreateStatus(r.Context(), status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
