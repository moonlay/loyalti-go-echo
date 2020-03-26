package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func main() {
	db := database.ConnectionDB()

	outlet := model.Program{
		Id:                    0,
		Created:               time.Now(),
		CreatedBy:             "Admin",
		Modified:              time.Now(),
		ModifiedBy:            "Admin",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		ProgramName:           "Adidas VS Corona",
		ProgramImage:          "aaaa",
		ProgramStartDate:      time.Date(2020,time.March,31,23,59,59, 59, time.UTC),
		ProgramEndDate:        time.Date(2020,time.April,30,23,59,59, 59, time.UTC),
		ProgramDescription:    "Diskon 50% untuk yang berulang tahun bulan april dan free ongkir",
		Card:                  "Member",
		//OutletID:              "2",
		MerchantId:            5,
		CategoryId:            2,
		Benefit:               nil,
		TermsAndCondition:     nil,
		Tier:                  nil,
		RedeemRules:           nil,
		RewardTarget:          nil,
		QRCodeId:              nil,
		ProgramPoint:          nil,
		MinPayment:            nil,
		IsReqBillNumber:       true,
		IsReqTotalTransaction: true,
		IsPushNotification:    true,
		IsLendCard:            true,
		IsGiveCard:            true,
		IsWelcomeBonus:        true,
	}
	db.Create(&outlet)
}


