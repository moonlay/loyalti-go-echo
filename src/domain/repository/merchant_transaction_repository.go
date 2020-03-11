package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

type TransactionRepository interface {
	CreateTransaction(transaction *model.TransactionMerchant) error
}

type repoTransaction struct {
	DB *gorm.DB
}

func CreateRepositoryTransaction(db *gorm.DB) TransactionRepository {
	return &repoTransaction{
		DB:db,
	}
}

func (p *repoTransaction) CreateTransaction(transaction *model.TransactionMerchant) error {
	db := database.ConnectionDB()
	transactionObj := *transaction
	err := db.Create(&transactionObj).Error
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	return err
}

func CreateTransaction(transaction *model.TransactionMerchant) string {
	db := database.ConnectionDB()
	transactionObj := *transaction
	db.Create(&transactionObj)
	return transactionObj.BillNumber
}

func UpdateTransaction(transaction *model.TransactionMerchant) string {
	db := database.ConnectionDB()
	db.Model(&transaction).Where("outlet_id = ?", transaction.OutletId).Update(&transaction)
	return transaction.OutletId
}

func DeleteTransaction(transaction *model.TransactionMerchant) string {
	db := database.ConnectionDB()
	db.Model(&transaction).Where("outlet_id = ?",transaction.OutletId).Update("active", false)
	return "berhasil dihapus"
}

func GetTransaction(page *int, size *int, sort *int, outletid *string) []model.TransactionMerchant {
	db := database.ConnectionDB()
	var transaction []model.TransactionMerchant
	db.Find(&transaction)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"created desc"},
				}, &transaction)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"created asc"},
				}, &transaction)
			}
		}
	}

	if outletid != nil {
		if page != nil && size != nil {
			pagination.Paging(&pagination.Param{
				DB:      db,
				Page:    *page,
				Limit:   *size,
				OrderBy: []string{"created desc"},
			}, &transaction)
			db.Order("created").Where("outlet_id = ?", outletid).Find(&transaction)
		}else {
			db.Order("created").Where("outlet_id = ?", outletid).Find(&transaction)
		}
	}

	db.Close()
	return transaction
}
