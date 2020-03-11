package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

type OutletRepository interface {
	CreateOutlet (newoutlet *model.Outlet) error
	UpdateOutlet (newoutlet *model.Outlet) error
	DeleteOutlet (newoutlet *model.Outlet) error
}

type outlet_repo struct {
	DB *gorm.DB
	//database.Connection_Interface
}

func (p *outlet_repo) CreateOutlet (newoutlet *model.Outlet) error {
	fmt.Println("masuk fungsi")
	outlet := model.Outlet{
		Created:          time.Now(),
		CreatedBy:        "",
		Modified:         time.Now(),
		ModifiedBy:       "",
		Active:           true,
		IsDeleted:        false,
		Deleted:          nil,
		Deleted_by:       "",
		OutletName:       newoutlet.OutletName,
		OutletAddress:    newoutlet.OutletAddress,
		OutletPhone:      newoutlet.OutletPhone,
		OutletCity:       newoutlet.OutletCity,
		OutletProvince:   newoutlet.OutletProvince,
		OutletPostalCode: newoutlet.OutletPostalCode,
		OutletLongitude:  newoutlet.OutletLongitude,
		OutletLatitude:   newoutlet.OutletLatitude,
		OutletDay:        time.Time{},
		OutletHour:       time.Time{},
		MerchantId:       1,
	}
	db := database.ConnectionDB()
	err := db.Create(&outlet).Error
	if err != nil {
		fmt.Println("Tak ada error")
	}
	return err
}

func CreateOutletRepository (db *gorm.DB) OutletRepository {
	return &outlet_repo{
		DB:db,
	}
}

func (p *outlet_repo) UpdateOutlet(newoutlet *model.Outlet) error {
	db := database.ConnectionDB()
	outlet := model.Outlet{
		Created:          time.Time{},
		CreatedBy:        "",
		Modified:         time.Time{},
		ModifiedBy:       "",
		Active:           false,
		IsDeleted:        false,
		Deleted:          nil,
		Deleted_by:       "",
		OutletName:       newoutlet.OutletName,
		OutletAddress:    newoutlet.OutletAddress,
		OutletPhone:      newoutlet.OutletPhone,
		OutletCity:       newoutlet.OutletCity,
		OutletProvince:   newoutlet.OutletProvince,
		OutletPostalCode: newoutlet.OutletPostalCode,
		OutletLongitude:  newoutlet.OutletLongitude,
		OutletLatitude:   newoutlet.OutletLatitude,
		OutletDay:        time.Time{},
		OutletHour:       time.Time{},
		MerchantId:       newoutlet.MerchantId,
	}
	err := db.Model(&outlet).Where("merchant_id = ?", outlet.MerchantId).Update(&outlet).Error
	return err
}

func (p *outlet_repo) DeleteOutlet(newoutlet *model.Outlet) error {
	db := database.ConnectionDB()

	err := db.Model(&newoutlet).Where("id = ?", newoutlet.Id).Update("active", false).Error
	if err == nil {
		fmt.Println("tidak ada error")
	}
	return err
}

func CreateOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	outlet.Created = time.Now()
	outlet.Modified = time.Now()
	outletObj := *outlet

	db.Create(&outletObj)

	return outletObj.OutletName
}

func UpdateOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Where("id = ?", outlet.Id).Update(&outlet)
	return outlet.OutletName
}

func DeleteOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Where("id= ?", outlet.Id).Update("active", false)
	return "berhasil dihapus"
}

func GetOutlet(page *int, size *int, id *int, email *string) []model.Outlet {

	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var outlet []model.Outlet
	var rows *sql.Rows
	var err error
	var total int
	var merchantID int
	merchant_email := model.Merchant{}
	db.Model(&merchant_email).Where("merchant_email = ?", email).Find(&merchant_email)
	if email != nil {
		merchantID = merchant_email.Id
	}
	fmt.Println("email", email)

	if id != nil {

		if page != nil && size != nil  {
			fmt.Println("2")
			rows, err = db.Where("merchant_id = ?", id).Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
			fmt.Println("Error : ", err.Error())
      }
		}
	} else if  merchantID != 0 {

		if page != nil && size != nil {
			//fmt.Println("2")
			rows, err = db.Where("merchant_id = ?", merchantID).Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				fmt.Println("Error : ", err.Error())
			}
		}
	} else {
		if page != nil && size != nil {
			rows, err = db.Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		} else {
			rows, err = db.Find(&outlet).Rows()
			if err != nil {
				panic(err)
			}
		}
	}
	if id != nil {
		if page == nil && size == nil {
			db.Model(&outlet).Where("merchant_id = ? ", id).Find(&outlet)
		}
	}


	result := make([]model.Outlet, 0)

	for rows.Next() {
		o := new(model.Outlet)
		err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.Deleted_by,
			&o.OutletName,
			&o.OutletAddress,
			&o.OutletPhone,
			&o.OutletCity,
			&o.OutletProvince,
			&o.OutletPostalCode,
			&o.OutletLongitude,
			&o.OutletLatitude,
			&o.OutletDay,
			&o.OutletHour,
			&o.MerchantId,
		)

		//merchantEmail := new(model.Merchant)
		//db.Table("merchants").
		//	Select("merchants.merchant_email").
		//	Where("id = ?", o.MerchantId).
		//	First(&merchantEmail)
		//o.MerchantEmail = merchantEmail.MerchantEmail
		//if err != nil {
		//	logrus.Error(err)
		//	return nil
		//}

		result = append(result, *o)
	}

	db.Close()
	return result
}
