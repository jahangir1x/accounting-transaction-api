package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AccountingHeadTransaction struct {
	ID                                  uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	CreatedAt                           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                           time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                           gorm.DeletedAt
	TransactionAt                       *time.Time `json:"transaction_at"`
	PaymentModeID                       uint       `json:"payment_mode_id" gorm:"index:idx_transactions_payment_mode_id" binding:"required"`
	AccountingHeadId                    uuid.UUID  `json:"accounting_head_id" gorm:"type:uuid"`
	BusinessId                          uint       `json:"business_id"`
	UserID                              uint       `json:"user_id"`
	AccountID                           uint       `json:"account_id"`
	MainAccountId                       *uint      `json:"main_account_id"`
	BranchAccountId                     *uint      `json:"branch_account_id"`
	ContactID                           *uint      `json:"contact_id"`
	ContactUUID                         *uuid.UUID `json:"contact_uuid" gorm:"type:uuid"`
	ContactType                         string     `json:"contact_type"`
	Remarks                             *string    `json:"remarks"`
	Type                                int8       `json:"type"`
	Amount                              float64    `json:"amount" binding:"required,gte=0"`
	ActualAmount                        float64    `gorm:"->;-:migration"`
	Balance                             float64    `json:"balance" gorm:"->;-:migration"`
	IsVerified                          bool       `json:"is_verified,default=false" gorm:"default:false"`
	VerifiedBy                          *uint      `json:"verified_by"`
	TransactionNumber                   uint64     `json:"transaction_number"`
	TransactionPrefix                   string     `json:"transaction_prefix"`
	IsImported                          bool       `json:"is_imported,default=false" gorm:"default:false"`
	VerifyRemarks                       *string    `json:"verify_remarks"`
	TimeZone                            int        `json:"time_zone"`
	IsDeletable                         bool       `json:"is_deletable,default=true"`
	LastUpdatedBy                       uint       `json:"last_updated_by"`
	TransactionType                     int8       `json:"transaction_type"`
	ImportFileId                        *string    `json:"import_file_id"`
	TaxID                               *uuid.UUID `gorm:"type:uuid;" json:"tax_id"`
	TaxType                             string     `json:"tax_type"`
	TaxAmount                           float64    `json:"tax_amount"`
	TransactionID                       *uint64    `json:"transaction_id"`
	OppositeAccountingHeadTransactionID *uuid.UUID `json:"opposite_accounting_head_transaction_id" gorm:"type:uuid"`
	TransactionUUID                     *uuid.UUID `json:"transaction_uuid" gorm:"type:uuid"`
	TransactionLineItemID               *uuid.UUID `json:"transaction_line_item_id" gorm:"type:uuid"`
	TransactionClass                    int        `json:"transaction_class"`
	FromBranchAccountId                 *uint      `json:"from_branch_account_id"`
	FromSubAccountId                    *uint      `json:"from_sub_account_id"`
	Order                               int        `json:"order"`
}
