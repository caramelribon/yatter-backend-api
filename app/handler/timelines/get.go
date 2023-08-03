package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"

	"yatter-backend-go/app/domain/object"
)

// Handle request for `Get /v1/statuses/{id}`
func (h *handler) GetPublicTimelines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// クエリの取得
	queryIntParams := getQueryParams(r, "max_id", "since_id", "limit")
	onlyMediaStr := r.URL.Query().Get("only_media")
	onlyMedia, _ := strconv.ParseBool(onlyMediaStr)
	queryParams := &object.QueryParams{
		MaxId:     queryIntParams["max_id"],
		SinceId:   queryIntParams["since_id"],
		Limit:     queryIntParams["limit"],
		OnlyMedia: onlyMedia,
	}

	// idからstatusを取得
	timelines := new([]object.Status)
	timelines, err := h.sr.GetStatuses(ctx, queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timelines); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getQueryParams(r *http.Request, keys ...string) map[string]int64 {
	params := make(map[string]int64)
	for _, key := range keys {
		value := r.URL.Query().Get(key)
		// Convert to int64 if the value is not empty
		intValue, _ := toInt64(value)
		params[key] = intValue
	}
	return params
}

func toInt64(value string) (int64, error) {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
