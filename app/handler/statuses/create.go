package statuses

import (
	"encoding/json"
	"net/http"
	"time"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Status string
}

type AddResponse struct {
	ID       int64           `json:"id"`
	Account  *object.Account `json:"account"`
	Content  string          `json:"content"`
	CreateAt time.Time       `json:"create_at"`
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

	// ステータスの作成
	if err := h.sr.CreateStatus(r.Context(), status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスの作成
	res := &AddResponse{
		ID:       status.ID,
		Account:  account,	
		Content:  status.Content,
		CreateAt: status.CreateAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
