package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Handle request for `Get /v1/timelines/pubic`
func (h *handler) GetPublicTimelines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get query params
	queryParams := getQueryParams(r)

	// get public statuses
	timelines, err := h.sr.GetPublicStatuses(ctx, queryParams)
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

func (h *handler) GetHomeTimelines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get account
	account := auth.AccountOf(r)

	// get query params
	queryParams := getQueryParams(r)

	// get home statuses
	timelines, err := h.sr.GetHomeStatuses(ctx, queryParams, account.ID)
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


// getQueryParams returns query params as *object.QueryParams
func getQueryParams(r *http.Request) *object.QueryParams {
	queryIntParams := getQueryIntParams(r, "max_id", "since_id", "limit")
	onlyMediaStr := r.URL.Query().Get("only_media")
	onlyMedia, _ := strconv.ParseBool(onlyMediaStr)
	queryParams := &object.QueryParams{
		MaxId:     queryIntParams["max_id"],
		SinceId:   queryIntParams["since_id"],
		Limit:     queryIntParams["limit"],
		OnlyMedia: onlyMedia,
	}
	return queryParams
}


// getQueryIntParams returns query params as map[string]int64
func getQueryIntParams(r *http.Request, keys ...string) map[string]int64 {
	params := make(map[string]int64)
	for _, key := range keys {
		value := r.URL.Query().Get(key)
		// Convert to int64 if the value is not empty
		intValue, _ := toInt64(value)
		params[key] = intValue
	}
	return params
}

// toInt64 converts string to int64
func toInt64(value string) (int64, error) {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
