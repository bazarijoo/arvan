package Entity

type VoucherUserEntity struct {
	PhoneNumber string `gorm:"primary_key;auto_increment:false"`
	VoucherCode string `gorm:"primary_key;auto_increment:false"`
	IsUsed      bool   `gorm:"default:false"`
}

func (VoucherUserEntity) TableName() string {
	return "voucher_user_entity"
}
