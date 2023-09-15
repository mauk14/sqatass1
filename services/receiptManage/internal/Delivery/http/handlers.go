package httpDelivery

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"messanger/services/receiptManage/internal/Domain"
	"net/http"
	"strconv"
)

func (a *App) createReceipt(c *gin.Context) {
	var receipt Domain.Receipt

	if err := c.BindJSON(&receipt); err != nil {
		log.Print(err)
		return
	}
	err := a.receiptUseCase.Create(context.Background(), &receipt)

	if err != nil {
		log.Print(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, receipt)
}

func (a *App) deleteReceipt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Print(err)
		return
	}

	err = a.receiptUseCase.Delete(context.Background(), int64(id))
	if err != nil {
		log.Print(err)
		return
	}
	c.IndentedJSON(http.StatusOK, id)

}

func (a *App) getReceipt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Print(err)
		return
	}

	receipt, err := a.receiptUseCase.Get(context.Background(), int64(id))
	if err != nil {
		log.Print(err)
	}
	c.IndentedJSON(http.StatusOK, receipt)

}

func (a *App) getAllReceipt(c *gin.Context) {
	receipts, err := a.receiptUseCase.GetAll(context.Background())
	if err != nil {
		log.Print(err)
	}
	c.IndentedJSON(http.StatusOK, receipts)

}

func (a *App) updateReceipt(c *gin.Context) {
	var receipt Domain.Receipt

	if err := c.BindJSON(&receipt); err != nil {
		log.Print(err)
		return
	}
	err := a.receiptUseCase.Update(context.Background(), &receipt)

	if err != nil {
		log.Print(err)
	}

	c.IndentedJSON(http.StatusOK, receipt)
}
