package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GetVoucherCodeStatusRequest struct {
	VoucherCode string `json:"voucher_code"`
}

type GetVoucherCodeStatusResponse struct {
	Balance int `json:"balance"`
}

type SubmitVoucherCodeRequest struct {
	PhoneNumber  string `json:"phone_number"`
	Voucher_code string `json:"voucher_code"`
}
type SubmitVoucherCodeResponse struct {
	Result string `json:"result"`
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetVoucherCodeStatusReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetVoucherCodeStatusRequest
	vars := mux.Vars(r)

	req = GetVoucherCodeStatusRequest{
		VoucherCode: vars["voucher_code"],
	}
	return req, nil
}

func decodeSubmitVoucherCodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req SubmitVoucherCodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
