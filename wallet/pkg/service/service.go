package service

import (
	. "arvan/wallet/Entity"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type WalletService interface {
	GetBalance(ctx context.Context, phoneNumber string) (int, error)
	UpdateBalance(ctx context.Context, phoneNumber string, credit int) (string, error)
}

type service struct {
	repository UserRepository
	logger     log.Logger
}

func NewService(rep UserRepository, logger log.Logger) WalletService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) GetBalance(ctx context.Context, phoneNumber string) (int, error) {
	logger := log.With(s.logger, "method", "GetBalance")

	balance, err := s.repository.GetBalance(ctx, phoneNumber)
	if err != nil {
		_ = level.Error(logger).Log("err", err)
		return -1, err
	}
	_ = logger.Log("Got balance for "+phoneNumber+" :", balance)
	return balance, nil

}

func (s service) UpdateBalance(ctx context.Context, phoneNumber string, credit int) (string, error) {
	logger := log.With(s.logger, "method", "UpdateBalance")
	msg, err := s.repository.UpdateBalance(ctx, phoneNumber, credit)

	if err != nil {
		_ = level.Error(logger).Log("err", err)
		return msg, err
	}
	return msg, nil
}
