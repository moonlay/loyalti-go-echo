package main

import (
	"github.com/beevik/guid"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func main() {
	db := database.ConnectionDB()
	db.AutoMigrate(&model.Reward{})
	reward := model.Reward{
		Id:                guid.NewString(),
		Created:           time.Now(),
		CreatedBy:         "Admin",
		Modified:          time.Now(),
		ModifiedBy:        "Admin",
		Active:            true,
		IsDeleted:         false,
		Deleted:           nil,
		DeletedBy:         "",
		RedeemPoints:      100,
		RewardName:        "Beli 1 dapat 1",
		RedeemRules:       "Harus belanja minimal 14 juta",
		TermsAndCondition: "Tidak bisa diuangkan, tidak bisa diwakilkan",
		ProgramId:         1,
		MerchantEmail:     "contact@nike.com",
	}
	db.Create(&reward)
}
