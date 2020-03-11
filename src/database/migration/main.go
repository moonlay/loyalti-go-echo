package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func main() {
	db := database.ConnectionDB()
	voucher := model.Voucher{}
	db.AutoMigrate(&voucher)
}


