package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`

	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE" json:"transaction_details,omitempty"`
}

func (Transaction) TableName() string {
	return "transactions"
}
