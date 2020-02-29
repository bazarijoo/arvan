package Reository

import (
	"arvan/voucher/Entity"
	"arvan/voucher/Model"
	"context"
	"errors"
	"fmt"
	"github.com/captaincodeman/couponcode"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) Entity.VoucherRepository {
	return &repository{
		db:     db,
		logger: log.With(logger, "repository", "sqlite3"),
	}
}

func (repo *repository) GetVoucherCodeStatus(ctx context.Context, phoneNumber string) error {
	//var user Entity.UserEntity
	//
	//if phoneNumber == "" {
	//	return -1, RepoErr
	//}
	//
	//if err := repo.db.Where("phone_number = ?", phoneNumber).Find(&user).Error; err != nil {
	//	fmt.Println("error: ", err.Error())
	//	return -1, err
	//}

	return nil
}

func (repo *repository) SubmitVoucherCode(ctx context.Context, phoneNumber string, voucherCode string) error {

	if err := repo.db.Exec("PRAGMA serializable = true").Error; err != nil { /// or serializable
		return err
	}

	//Transactions in SQLite are SERIALIZABLE. and changes are not visible  to ther db connections prior commit
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	validatedCode, err := couponcode.Validate(voucherCode)
	if err != nil {
		return err
	}
	var voucher Entity.VoucherEntity
	if err := repo.db.Where("code = ?", validatedCode).Find(&voucher).Error; err != nil {
		return err
	}
	if voucher.IsActive == false {
		return errors.New("Voucher Code not activated")
	}

	if voucher.CountUsed >= voucher.Limit {
		return errors.New("limit reached, not valid any more")
	}

	var voucherUser Entity.VoucherUserEntity
	if err := repo.db.Where("phone_number = ? AND voucher_code = ? ", phoneNumber, voucherCode).Find(&voucherUser).Error; err != nil {
		fmt.Println("error: ", err.Error())
		return err
	}

	if voucherUser.IsUsed == true {
		return errors.New("gift code is used before")
	}

	newIsActive := true
	if voucher.CountUsed+1 >= voucher.Limit {
		newIsActive = false
	}

	if err := repo.db.Model(&voucher).Updates(Entity.VoucherEntity{
		Code:      voucher.Code,
		IsActive:  newIsActive,
		CountUsed: voucher.CountUsed + 1,
		Amount:    voucher.Amount,
		Credit:    voucher.Credit,
	}).Error; err != nil {
		fmt.Println("error: ", err.Error())
		return err
	}

	if err := repo.db.Model(&voucherUser).Updates(Entity.VoucherUserEntity{
		VoucherCode: voucher.Code,
		PhoneNumber: voucherUser.PhoneNumber,
		IsUsed:      true,
	}).Error; err != nil {
		fmt.Println("error: ", err.Error())
		return err
	}

	//API Call to 8080
	err = Model.UpdateBalanceAPI(voucherUser.PhoneNumber, voucher.Credit)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return err
	}

	return tx.Commit().Error
}
