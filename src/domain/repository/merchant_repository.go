package repository

import (
	"bytes"
	"crypto/tls"
	"github.com/jinzhu/gorm"

	"encoding/json"
	"fmt"

	"github.com/biezhi/gorm-paginator/pagination"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	//"github.com/stretchr/testify/require"

	"net/http"
	"os"
	"time"
)

type Repository interface {
	CreateMerchant (newmerchant *model.NewMerchantCommand) error
	UpdateMerchant(newmerchant *model.NewMerchantCommand) error
	DeleteMerchant(newmerchant *model.NewMerchantCommand) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) CreateMerchant (newmerchant *model.NewMerchantCommand) error {
	merchant := model.Merchant{
		Created:               time.Now(),
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	db := database.ConnectionDB()
	err := db.Create(&merchant).Error
	if err == nil {
		fmt.Println("tak ada error")
	}
	return err
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB:db,
	}
}

func CreateMerchantWSO2(newmerchant *model.NewMerchantCommand) (*http.Response, error) {
	user := model.AccountMerchant{
		Username: newmerchant.MerchantEmail,
		Password: newmerchant.MerchantPassword,
	}
	data, _:= json.Marshal(user)
	fmt.Println("Ini datanya : ",string(data))

	req, err := http.NewRequest("POST", "https://11.11.5.146:9443/scim2/Users", bytes.NewReader(data))
	fmt.Println("ini isi bytes reader : ",)
	fmt.Println(bytes.NewReader(data))
	//os.Exit(1)
	req.Header.Set("Authorization", "Basic YWRtaW5AZ21haWwuY29tOmFkbWlu")
	req.Header.Set("Content-Type","application/json")
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
	fmt.Println("ini response : ", resp)

	return resp, err
}

func  CreateMerchant2 (newmerchant *model.NewMerchantCommand) string {
	merchant := model.Merchant{
		Created:               time.Now(),
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	db := database.ConnectionDB()
	db.Create(&merchant)
	return merchant.MerchantEmail
}

func (p *repo) UpdateMerchant(newmerchant *model.NewMerchantCommand) error {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	err := db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update(&merchant).Error
	return err
}

func UpdateMerchant2(newmerchant *model.NewMerchantCommand) string {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update(&merchant)
	return merchant.MerchantEmail
}

func (p *repo) DeleteMerchant(newmerchant *model.NewMerchantCommand) error {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	err := db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update("active", false).Error
	if err == nil {
		fmt.Println("Tidak ada error")
	}
	return err
}

func DeleteMerchant2(newmerchant *model.NewMerchantCommand) string {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update("active", false)
	return "berhasil dihapus"
}

func GetMerchant(page *int, size *int, sort *int, email *string, id *int) []model.Merchant {

	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var merchant []model.Merchant
	db.Find(&merchant)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name desc"},
				}, &merchant)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name asc"},
				}, &merchant)
			}
		}
	}

	if email != nil {
		if page != nil && size != nil {
			pagination.Paging(&pagination.Param{
				DB:      db,
				Page:    *page,
				Limit:   *size,
				OrderBy: []string{"merchant_name asc"},
			}, &merchant)
			db.Order("id").Where("merchant_email = ?", email).Find(&merchant)

		} else {
			db.Order("id").Where("merchant_email = ?", email).Find(&merchant)
		}
	}

	if id != nil {
		if page != nil && size != nil {
			pagination.Paging(&pagination.Param{
				DB:      db,
				Page:    *page,
				Limit:   *size,
				OrderBy: []string{"merchant_name asc"},
			}, &merchant)
			db.Order("id").Where("id = ?", id).Find(&merchant)

		} else {
			db.Order("id").Where("id = ?", id).Find(&merchant)
		}
	}

	db.Close()
	return merchant
}
