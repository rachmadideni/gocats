package repository

import (
	"gocats/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error
	CreateTransactionDetail(tx *gorm.DB, detail *models.TransactionDetail) error
	UpdateProductStock(tx *gorm.DB, productID uint, quantity int) error
	FindByID(id uint) (*models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	return tx.Create(transaction).Error
}

func (r *transactionRepository) CreateTransactionDetail(tx *gorm.DB, detail *models.TransactionDetail) error {
	return tx.Create(detail).Error
}

func (r *transactionRepository) UpdateProductStock(tx *gorm.DB, productID uint, quantity int) error {
	return tx.Model(&models.Product{}).Where("id = ?", productID).UpdateColumn("stock", gorm.Expr("stock - ?", quantity)).Error
}

func (r *transactionRepository) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("TransactionDetails.Product").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
