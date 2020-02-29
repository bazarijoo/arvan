package http

import (
	. "arvan/wallet/pkg/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	UpdateBalance endpoint.Endpoint
	GetBalance    endpoint.Endpoint
}

func MakeEndpoints(s WalletService) Endpoints {
	return Endpoints{
		GetBalance:    makeGetBalanceEndpoint(s),
		UpdateBalance: makeUpdateBalanceEndpoint(s),
	}
}

func makeGetBalanceEndpoint(s WalletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBalanceRequest)
		balance, err := s.GetBalance(ctx, req.PhoneNumber)
		return GetBalanceResponse{Balance: balance}, err
	}
}

func makeUpdateBalanceEndpoint(s WalletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBalanceRequest)
		ok, err := s.UpdateBalance(ctx, req.PhoneNumber, req.Amount)

		return UpdateBalanceResponse{
			Ok: ok,
		}, err
	}
}
