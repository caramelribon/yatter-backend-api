package statuses

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
func (h *handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramId := chi.URLParam(r, "id")
	if paramId == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// idの型変換
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// idからstatusを取得
	status, err := h.sr.FindById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// status.account_idからaccountを取得
	account, err := h.ar.FindById(ctx, status.AccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスの作成
	res := &getResponse{
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
