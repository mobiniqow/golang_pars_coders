package service

import (
	"mobiniqow/model"
	"mobiniqow/repository"
)

type CoinService struct {
	repository repository.CoinRepositoryImp
	Db         *repository.Server
}

// initialization instance
func GetCoinService() *CoinService {
	new_db := repository.Server{}
	rep := repository.GetCoinRepository()
	instance := &CoinService{Db: new_db.Connect(), repository: *rep}
	return instance
}

func (c *CoinService) CreateCoin(wname string, name string, symbol string, amount float64, rate float64) (*model.Coin, error) {
	return c.repository.CreateCoin(c.Db.DB, wname, name, symbol, amount, rate)
}

func (c *CoinService) GetAllCoin(wname string) (*model.Wallet, error) {
	return c.repository.GetAllCoin(c.Db.DB, wname)
}

func (c *CoinService) UpdateCoin(wname string, old_symbol string, name string, symbol string, amount float64, rate float64) (*model.Coin, error) {
	return c.repository.UpdateCoin(c.Db.DB, wname, old_symbol, name, symbol, amount, rate)
}

func (c *CoinService) DeleteCoin(wname string, symbol string) (*model.Coin, error) {
	return c.repository.DeleteCoin(c.Db.DB, wname, symbol)
}
