package http

import (
	. "arvan/voucher/pkg/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SubmitVoucherCode    endpoint.Endpoint
	GetVoucherCodeStatus endpoint.Endpoint
}

func MakeEndpoints(s VoucherService) Endpoints {
	return Endpoints{
		SubmitVoucherCode:    makeSubmitVoucherCodeEndpoint(s),
		GetVoucherCodeStatus: makeGetVoucherCodeStatusEndpoint(s),
	}
}

func makeGetVoucherCodeStatusEndpoint(s VoucherService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetVoucherCodeStatusRequest)
		usedVoucherUsers, err := s.GetVoucherCodeStatus(ctx, req.VoucherCode)
		return GetVoucherCodeStatusResponse{UsersUsedVoucherCode: usedVoucherUsers}, err
	}
}

func makeSubmitVoucherCodeEndpoint(s VoucherService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SubmitVoucherCodeRequest)
		err := s.SubmitVoucherCode(ctx, req.PhoneNumber, req.Voucher_code)

		result := "success"
		if err != nil {
			result = "error"
		}
		return SubmitVoucherCodeResponse{
			Result: result,
		}, err
	}
}
