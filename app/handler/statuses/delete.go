package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"yatter-backend-go/app/handler/auth"
)

// Handle request for `DELETE /v1/statuses/{id}`
func (h *handler) DeleteStatus(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(id)

	// アカウント情報の取得
	account := auth.AccountOf(r)

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
	if status != nil && status.Account.ID != account.ID {
		http.Error(w, "このstatusは削除できません", http.StatusNotFound)
		return
	}

	// ステータスの削除
	if err := h.sr.DeleteStatus(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("削除できました"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
