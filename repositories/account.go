package repositories

import "gorm.io/gorm"

type AccountRepo struct {
	db *gorm.DB
}

func CreateNewAccountRepository(d *gorm.DB) *AccountRepo {
	return &AccountRepo{
		db: d,
	}
}
