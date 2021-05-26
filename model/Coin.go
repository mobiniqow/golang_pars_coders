package model

type Coin struct {
	ID       uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Name     string  `gorm:"size:255;not null;unique" json:"name"`
	Symbol   string  `gorm:"size:255;not null;unique" json:"symbol"`
	Amount   float64 `gorm:"type:decimal(7,6);" json:"amount"`
	Rate     float64 `gorm:"type:decimal(7,6);" json:"rate"`
	WalletID int     `gorm:"column:wallet_id" json:"-"`
	// Wallet   Wallet
}

func NewCoin(name string, amount float64, symbol string, rate float64, walletId int) *Coin {
	// last_update:= time.Now().Format("2020-02-01 13:05:05")
	return &Coin{Name: name, Amount: amount, Symbol: symbol, Rate: rate, WalletID: walletId}
}
