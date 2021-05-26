package repository

import (
	"fmt"
	"mobiniqow/model"
	"time"

	"github.com/jinzhu/gorm"
)

type CoinRepositoryImp struct {
}

//test repository
func GetCoinRepository() *CoinRepositoryImp {
	instance := &CoinRepositoryImp{}
	return instance
}

func (c *CoinRepositoryImp) CreateCoin(db *gorm.DB, wname string, name string, symbol string, amount float64, rate float64) (*model.Coin, error) {
	var err error
	wallet := &model.Wallet{}
	last_time := time.Now()
	werr := db.Debug().Model(model.Wallet{}).Where("name = ?", wname).First(wallet).Error
	if werr != nil {
		return &model.Coin{}, werr
	}
	fmt.Print(name, amount, symbol, rate, wallet.ID)
	new_coin := model.NewCoin(name, amount, symbol, rate, wallet.ID)
	wallet.Last_update = last_time
	err = db.Debug().Create(&new_coin).Error
	if err != nil {
		return &model.Coin{}, err
	}
	db.Save(wallet)
	return new_coin, nil
}

func (c *CoinRepositoryImp) GetAllCoin(db *gorm.DB, wname string) (*model.Wallet, error) {
	var err error
	var balance = 0.0
	wallets := model.Wallet{}
	coin := []model.Coin{}
	err = db.Debug().Model(model.Wallet{}).Where("name = ?", wname).First(&wallets).Error
	if err != nil {
		return &model.Wallet{}, err
	}
	cerr := db.Debug().Model([]model.Coin{}).Where("wallet_id = ?", wallets.ID).Find(&coin).Error
	if cerr != nil {
		return &model.Wallet{}, cerr
	}
	for _, element := range coin {
		// index is the index where we are
		// element is the element from someSlice for where we are
		balance += element.Amount * element.Rate
	}
	wallets.Coins = coin
	wallets.Balance = float64(balance)
	return &wallets, err
}

func (c *CoinRepositoryImp) UpdateCoin(db *gorm.DB, wname string, wsymbol string, name string, symbol string, amount float64, rate float64) (*model.Coin, error) {

	wallet := &model.Wallet{}
	coin := &model.Coin{}
	last_time := time.Now()
	werr := db.Debug().Model(model.Wallet{}).Where("name = ?", wname).First(wallet).Error
	if werr != nil {
		fmt.Print("werr")

		return &model.Coin{}, werr
	}
	cerr := db.Debug().Model(model.Coin{}).Where("symbol = ?", wsymbol).Where("wallet_id = ?", wallet.ID).First(coin).Error
	if cerr != nil {
		fmt.Print(cerr)
		return &model.Coin{}, cerr
	}
	coin.Name = name
	coin.Symbol = symbol
	coin.Amount = amount
	coin.Rate = rate
	wallet.Last_update = last_time
	db.Save(wallet)
	db.Save(coin)
	return coin, nil
}

func (c *CoinRepositoryImp) DeleteCoin(db *gorm.DB, wname string, symbol string) (*model.Coin, error) {
	wallet := &model.Wallet{}
	coin := &model.Coin{}
	last_time := time.Now()
	werr := db.Debug().Model(model.Wallet{}).Where("name = ?", wname).First(wallet).Error
	if werr != nil {
		return &model.Coin{}, werr
	}
	cerr := db.Debug().Model(model.Coin{}).Where("symbol = ?", symbol).Where("wallet_id = ?", wallet.ID).First(coin).Error
	if cerr != nil {
		return &model.Coin{}, cerr
	}
	wallet.Last_update = last_time
	db.Delete(coin)
	db.Save(wallet)
	return coin, nil
}
