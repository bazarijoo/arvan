package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GetBalanceRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type GetBalanceResponse struct {
	Balance int `json:"balance"`
}

type UpdateBalanceRequest struct {
	PhoneNumber string `json:"phone_number"`
	Credit      int    `json:"credit"`
}
type UpdateBalanceResponse struct {
	Ok string `json:"ok"`
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetBalanceReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetBalanceRequest
	vars := mux.Vars(r)

	req = GetBalanceRequest{
		PhoneNumber: vars["phone_number"],
	}
	return req, nil
}

func decodeUpdateBalanceReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateBalanceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
