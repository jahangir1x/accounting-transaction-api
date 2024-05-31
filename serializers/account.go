package serializers

import (
	"github.com/google/uuid"
	"time"
)

type TransactionSerializer struct {
	TransactionUUID      *uuid.UUID       `json:"transaction_uuid"`
	TransactionAt        *time.Time       `json:"transaction_at"`
	TransactionClass     int              `json:"transaction_class"`
	TransactionNumberStr string           `json:"transaction_number_string"`
	AccountingHeads      []AccountingHead `json:"accounting_heads"`
}

type AccountingHead struct {
	TransactionAt     *time.Time `json:"-"`
	TransactionClass  int        `json:"-"`
	TransactionNumber uint64     `json:"-"`
	TransactionPrefix string     `json:"-"`

	AccountingHeadId uuid.UUID `json:"accounting_head_id"`
	Debit            float64   `json:"debit"`
	Credit           float64   `json:"credit"`
}
