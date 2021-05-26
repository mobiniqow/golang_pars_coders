package model

import (
	"time"
)

type Wallet struct {
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	ID          int       `gorm:"primary_key;auto_increment" json:"id"`
	Balance     float64   `gorm:"type:decimal(7,6);" json:"balance"`
	Coins       []Coin    `gorm:"ForeignKey:WalletID" json:"coin"`
	Last_update time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_update"`
}

// Wallet constructor
func NewWallet(name string) *Wallet {
	// last_update:= time.Now().Format("2020-02-01 13:05:05")
	return &Wallet{Name: name, Last_update: time.Now(), Coins: []Coin{}, Balance: 0.0}
}
func (c *Wallet) IsValid() bool {
	return c.ID != -1
}
