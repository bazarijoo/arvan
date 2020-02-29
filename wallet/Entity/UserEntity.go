package Entity

import "context"

type UserEntity struct {
	PhoneNumber string `gorm:"primary_key;unique;not null"`
	Balance     int    `gorm:"default:0"`
}

func (UserEntity) TableName() string {
	return "user_entity"
}

type UserRepository interface {
	GetBalance(ctx context.Context, phoneNumber string) (int, error)
	UpdateBalance(ctx context.Context, phoneNumber string, amount int) (string, error)
}
