package service

import (
	"mobiniqow/model"
	"mobiniqow/repository"
)

type WalletService struct {
	repository repository.WalletRepositoryImp
	Db         *repository.Server
}

// initialization instance
func GetWalletService() *WalletService {
	new_db := repository.Server{}
	rep := repository.GetWalletRepository()
	instance := &WalletService{Db: new_db.Connect(), repository: *rep}
	return instance
}

//query for find instance
func (c *WalletService) CreateNewWallet(name string) (*model.Wallet, error) {
	wallet, err := c.repository.CreateWallet(c.Db.DB, name)
	return wallet, err
}
func (c *WalletService) GetAllWallet() ([]*model.Wallet, error) {
	return c.repository.GetAllWallet(c.Db.DB)
}
func (c *WalletService) UpdateWallet(old_name string, new_name string) (*model.Wallet, error) {
	return c.repository.UpdateWallet(c.Db.DB, old_name, new_name)
}
func (c *WalletService) DeleteWallet(wname string) (*model.Wallet, error) {
	return c.repository.DeleteWallet(c.Db.DB, wname)
}
