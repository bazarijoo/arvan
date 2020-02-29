package Reository

import (
	"arvan/wallet/Entity"
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) Entity.UserRepository {
	return &repository{
		db:     db,
		logger: log.With(logger, "repository", "sqlite3"),
	}
}

func (repo *repository) GetBalance(ctx context.Context, phoneNumber string) (int, error) {
	var user Entity.UserEntity

	if phoneNumber == "" {
		return -1, RepoErr
	}

	if err := repo.db.Where("phone_number = ?", phoneNumber).Find(&user).Error; err != nil {
		fmt.Println("error: ", err.Error())
		return -1, err
	}

	return user.Balance, nil
}

func (repo *repository) UpdateBalance(ctx context.Context, phoneNumber string, amount int) (string, error) {

	var user Entity.UserEntity
	if err := repo.db.Where("phone_number = ?", phoneNumber).Find(&user).Error; err != nil {
		fmt.Println("error: ", err.Error())
		return "error", err
	}

	return "", nil
}
