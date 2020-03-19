package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

type OutletRepository interface {
	CreateOutlet(newoutlet *model.Outlet) error
	UpdateOutlet(newoutlet *model.Outlet) error
	DeleteOutlet(newoutlet *model.Outlet) error
}

type outlet_repo struct {
	DB *gorm.DB
	//database.Connection_Interface
}

func (p *outlet_repo) CreateOutlet(newoutlet *model.Outlet) error {
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

func CreateOutletRepository(db *gorm.DB) OutletRepository {
	return &outlet_repo{
		DB: db,
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
	db.Find(&outlet)

	if id == nil && size == nil && page == nil && email == nil {
		fmt.Println("1")
		fmt.Println(&outlet)
		db.Find(&outlet)
	}

	if id == nil && size != nil && page != nil && email != nil {
		fmt.Println("2")
		db.Model(&outlet).Where("merchant_email = ?", email).Limit(*size).Offset(*page).Find(&outlet)
		db.Model(&outlet).Find(&outlet)
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"outlet_name desc"},
		}, &outlet)
	}

	if id == nil && size != nil && page != nil && email == nil {
		fmt.Println("3")
		db.Model(&outlet).Find(&outlet)
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"outlet_name asc"},
		}, &outlet)
	}

	if id != nil && size != nil && page != nil && email == nil {
		fmt.Println("4")
		db.Model(&outlet).Where("merchant_id =  ?", id).Find(&outlet)
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"outlet_name asc"},
		}, &outlet)
	}
	return outlet
}
