package repository

import (
	"mobiniqow/model"

	"github.com/jinzhu/gorm"
)

// cuz some time multi repositories
type WalletRepository interface {
	CreateWallet(db *gorm.DB, name string) (*model.Wallet, error)
	GetAllWallet(db *gorm.DB) ([]*model.Wallet, error)
	UpdateWallet(db *gorm.DB, old_name string, new_name string) (*model.Wallet, error)
	DeleteWallet(db *gorm.DB, name string) (*model.Wallet, error)
}
