package handlers

import (
	"fmt"
	"mobiniqow/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCoin(ctx *gin.Context) {
	wname := ctx.Param("wname")
	name := ctx.PostForm("name")
	symbol := ctx.PostForm("symbol")
	rate := ctx.PostForm("rate")
	amount := ctx.PostForm("amount")
	rate_pars, _ := strconv.ParseFloat(rate, 8)
	amount_pars, _ := strconv.ParseFloat(amount, 8)
	result, err := service.GetCoinService().CreateCoin(wname, name, symbol, amount_pars, rate_pars)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed create new wallet! name is duplicated",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    result.Name,
			"symbol":  result.Symbol,
			"amount":  result.Amount,
			"rate":    result.Rate,
			"code":    200,
			"message": "Coin added successfully!",
		})
	}
}

func GetCoin(ctx *gin.Context) {
	wname := ctx.Param("wname")
	result, err := service.GetCoinService().GetAllCoin(wname)
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
			"message":      "All coin reseved successfully!",
		})
	}
}

func UpdateCoin(ctx *gin.Context) {
	wname := ctx.Param("wname")
	old_symbol := ctx.Param("symbol")
	name := ctx.PostForm("name")
	symbol := ctx.PostForm("symbol")
	rate := ctx.PostForm("rate")
	amount := ctx.PostForm("amount")
	rate_pars, _ := strconv.ParseFloat(rate, 8)
	amount_pars, _ := strconv.ParseFloat(amount, 8)
	result, err := service.GetCoinService().UpdateCoin(wname, old_symbol, name, symbol, amount_pars, rate_pars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    result.Name,
			"symbol":  result.Symbol,
			"amount":  result.Amount,
			"rate":    result.Rate,
			"code":    200,
			"message": "Coin updated successfully!",
		})
	}
}

func DeleteCoin(ctx *gin.Context) {
	wname := ctx.Param("wname")
	symbol := ctx.Param("symbol")
	result, err := service.GetCoinService().DeleteCoin(wname, symbol)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    result.Name,
			"amount":  result.Amount,
			"symbol":  result.Symbol,
			"rate":    result.Rate,
			"code":    200,
			"message": "Coin deleted successfully!",
		})
	}
}
