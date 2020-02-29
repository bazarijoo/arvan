package Entity

import "context"

type VoucherEntity struct {
	Code        string              `gorm:"primary_key;unique;not null"`
	IsActive    bool                `gorm:"default:false"`
	CountUsed   int                 `gorm:"default:0"`
	Limit       int                 `gorm:"default:10"`
	Amount      int                 `gorm:"default:0"`
	Credit      int                 `gorm:"default:0"`
	VoucherUsed []VoucherUserEntity `gorm:"foreignkey:phone_number;association_foreignkey:voucher_code"`
}

func (VoucherEntity) TableName() string {
	return "voucher_entity"
}

type VoucherRepository interface {
	GetVoucherCodeStatus(ctx context.Context, voucherCode string) error
	SubmitVoucherCode(ctx context.Context, phoneNumber string, voucherCode string) error
}
