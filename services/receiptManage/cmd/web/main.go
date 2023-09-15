package main

import (
	"log"
	"messanger/pkg/postgres"
	httpDelivery "messanger/services/receiptManage/internal/Delivery/http"
	"messanger/services/receiptManage/internal/Repository"
	"messanger/services/receiptManage/internal/Use_Case"
)

func main() {
	db, err := postgres.OpenDb("localhost", "5432", "postgres", "123456", "receipts")
	if err != nil {
		log.Fatal(err)
		return
	}
	app := httpDelivery.NewApp(Use_Case.NewReceiptUseCase(Repository.NewReceiptRepository(db)))

	err = app.Route().Run(":4000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
