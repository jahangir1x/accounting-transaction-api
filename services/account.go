package services

import (
	"accounting_transaction_api/repositories"
	"accounting_transaction_api/serializers"
	"github.com/google/uuid"
	"strconv"
)

type AccountService struct {
	accountRepo repositories.AccountRepo
}

func CreateNewAccountService(accountRepo repositories.AccountRepo) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

func (service *AccountService) GetTransactions() ([]serializers.TransactionSerializer, error) {

	transactions, err := service.accountRepo.GetTransactions()
	if err != nil {
		return nil, err
	}

	transactionMap := make(map[uuid.UUID][]serializers.AccountingHead)

	for _, transaction := range transactions {
		debit := 0.0
		credit := 0.0
		if transaction.Amount > 0 {
			debit = transaction.Amount
		} else {
			credit = transaction.Amount * -1.0
		}

		accountingHead := serializers.AccountingHead{
			TransactionAt:     transaction.TransactionAt,
			TransactionClass:  transaction.TransactionClass,
			TransactionNumber: transaction.TransactionNumber,
			TransactionPrefix: transaction.TransactionPrefix,
			AccountingHeadId:  transaction.AccountingHeadId,
			Debit:             debit,
			Credit:            credit,
		}

		transactionMap[*transaction.TransactionUUID] = append(transactionMap[*transaction.TransactionUUID], accountingHead)
	}

	var transactionSerializers []serializers.TransactionSerializer
	for key, value := range transactionMap {
		serializedTransaction := serializers.TransactionSerializer{
			TransactionUUID:      &key,
			TransactionAt:        value[0].TransactionAt,
			TransactionClass:     value[0].TransactionClass,
			TransactionNumberStr: value[0].TransactionPrefix + strconv.FormatUint(value[0].TransactionNumber, 10),
			AccountingHeads:      value,
		}
		transactionSerializers = append(transactionSerializers, serializedTransaction)
	}

	return transactionSerializers, nil
}
