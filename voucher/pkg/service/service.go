package service

import (
	. "arvan/voucher/Entity"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type VoucherService interface {
	GetVoucherCodeStatus(ctx context.Context, voucherCode string) error
	SubmitVoucherCode(ctx context.Context, phoneNumber string, voucherCode string) error
}

type service struct {
	repository VoucherRepository
	logger     log.Logger
}

func NewService(rep VoucherRepository, logger log.Logger) VoucherService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) GetVoucherCodeStatus(ctx context.Context, phoneNumber string) error {
	//logger := log.With(s.logger, "method", "GetBalance")

	//balance, err := s.repository.GetBalance(ctx, phoneNumber)
	//if err != nil {
	//	_ = level.Error(logger).Log("err", err)
	//	return -1, err
	//}
	//_ = logger.Log("Got balance for "+phoneNumber+" :", balance)
	return nil

}

func (s service) SubmitVoucherCode(ctx context.Context, phoneNumber string, voucherCode string) error {
	logger := log.With(s.logger, "method", "SubmitVoucherCode")
	err := s.repository.SubmitVoucherCode(ctx, phoneNumber, voucherCode)

	if err != nil {
		_ = level.Error(logger).Log("err", err)
		return err
	}
	return nil
}
