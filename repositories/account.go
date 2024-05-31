package repositories

import (
	"accounting_transaction_api/models"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func CreateNewAccountRepository(d *gorm.DB) *AccountRepo {
	return &AccountRepo{
		db: d,
	}
}

func (repo *AccountRepo) GetTransactions() ([]models.AccountingHeadTransaction, error) {
	var accountDetails []models.AccountingHeadTransaction
	if err := repo.db.Find(&accountDetails).Error; err != nil {
		return nil, err
	}
	return accountDetails, nil
}
