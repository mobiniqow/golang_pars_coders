package repository

import (
	"mobiniqow/model"
	"time"

	"github.com/jinzhu/gorm"
)

type WalletRepositoryImp struct {
}

//test repository
func GetWalletRepository() *WalletRepositoryImp {
	instance := &WalletRepositoryImp{}
	return instance
}

func (c *WalletRepositoryImp) CreateWallet(db *gorm.DB, name string) (*model.Wallet, error) {
	//todo save new wallet to db
	var err error
	new_wallet := model.NewWallet(name)
	err = db.Debug().Create(&new_wallet).Error
	if err != nil {
		return &model.Wallet{}, err
	}
	return new_wallet, nil

}

func (c *WalletRepositoryImp) GetAllWallet(db *gorm.DB) ([]*model.Wallet, error) {
	var err error
	wallets := []*model.Wallet{}
	err = db.Debug().Model([]*model.Wallet{}).Find(&wallets).Error
	if err != nil {
		return []*model.Wallet{}, err
	}
	for idex, ele := range wallets {
		coin := []model.Coin{}
		cerr := db.Debug().Model([]model.Coin{}).Where("wallet_id = ?", ele.ID).Find(&coin).Error
		if cerr != nil {
			break
		}
		balance := 0.0
		for _, element := range coin {
			// index is the index where we are
			// element is the element from someSlice for where we are
			balance += element.Amount * element.Rate
		}
		wallets[idex].Balance = float64(balance)
		wallets[idex].Coins = coin
	}

	return wallets, err
}

func (c *WalletRepositoryImp) UpdateWallet(db *gorm.DB, old_name string, new_name string) (*model.Wallet, error) {
	//todo save new wallet to db
	wallet := &model.Wallet{}
	last_time := time.Now()
	err := db.Debug().Model(model.Wallet{}).Where("name = ?", old_name).First(wallet).Error
	if err != nil {
		return &model.Wallet{}, err
	}
	wallet.Name = new_name
	wallet.Last_update = last_time
	err_up := db.Save(&wallet).Error
	if err_up != nil {
		return &model.Wallet{}, err_up
	}
	return wallet, nil
}

func (c *WalletRepositoryImp) DeleteWallet(db *gorm.DB, name string) (*model.Wallet, error) {
	//todo save new wallet to db
	wallet := &model.Wallet{}
	err := db.Debug().Model(model.Wallet{}).Where("name = ?", name).First(wallet).Error
	if err != nil {
		return &model.Wallet{}, err
	}
	err_up := db.Delete(&wallet).Error
	if err_up != nil {
		return &model.Wallet{}, err_up
	}
	return wallet, nil
}
