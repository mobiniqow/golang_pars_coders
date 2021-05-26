package repository

import (
	"mobiniqow/model"

	"github.com/jinzhu/gorm"
)

// cuz some time multi repositories
type CoinRepository interface {
	CreateCoin(db *gorm.DB, wname string, name string, symbol string, amount float64, rate float64) (*model.Coin, error)
	GetAllCoin(db *gorm.DB, wname string) (*model.Wallet, error)
	UpdateCoin(db *gorm.DB, old_name string, new_name string, name string, symbol string, amount float64, rate float64) (*model.Coin, error)
	DeleteCoin(db *gorm.DB, wname string, symbol string) (*model.Coin, error)
}
