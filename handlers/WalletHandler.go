package handlers

import (
	"fmt"
	"mobiniqow/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWallet(ctx *gin.Context) {
	name := ctx.PostForm("name")
	new_wallet, err := service.GetWalletService().CreateNewWallet(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed create new wallet! name is duplicated",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name":         new_wallet.Name,
			"balance":      new_wallet.Balance,
			"coin":         new_wallet.Coins,
			"last_updated": new_wallet.Last_update,
			"message":      "Food added successfully!",
		})
	}
}

func GetWallet(ctx *gin.Context) {
	allWallets, err := service.GetWalletService().GetAllWallet()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"size":    len(allWallets),
			"wallets": allWallets,
		})
	}

}

func UpdateWallet(ctx *gin.Context) {
	old_name := ctx.Param("wname")
	new_name := ctx.PostForm("name")
	result, err := service.GetWalletService().UpdateWallet(old_name, new_name)
	fmt.Print(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		t := result.Last_update
		formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		ctx.JSON(http.StatusOK, gin.H{
			"name":         result.Name,
			"balance":      result.Balance,
			"coin":         result.Coins,
			"last_updated": formatted,
			"code":         200,
			"message":      "Wallet name changed successfully!",
		})
	}
}

func DeleteWallet(ctx *gin.Context) {
	wname := ctx.Param("wname")
	result, err := service.GetWalletService().DeleteWallet(wname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})

	} else {
		t := result.Last_update
		formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		ctx.JSON(http.StatusOK, gin.H{
			"name":         result.Name,
			"balance":      result.Balance,
			"coin":         result.Coins,
			"last_updated": formatted,
			"code":         200,
			"message":      "Wallet deleted successfully!",
		})
	}
}
