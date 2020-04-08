package repository

import (
	"bytes"
	"github.com/beevik/guid"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"

	"encoding/json"
	"fmt"

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
		Id: guid.NewString(),
		Username: newmerchant.MerchantEmail,
		Password: newmerchant.MerchantPassword,
		Email: newmerchant.MerchantEmail,
	}
	data, _:= json.Marshal(user)
	fmt.Println("Ini datanya : ",string(data))

	req, err := http.NewRequest("POST", "https://identityserver-loyalti.azurewebsites.net/connect/register", bytes.NewReader(data))
	fmt.Println("ini isi bytes reader : ",)
	fmt.Println(bytes.NewReader(data))
	//os.Exit(1)
	//req.Header.Set("Authorization", "Basic YWRtaW5AZ21haWwuY29tOmFkbWlu")
	req.Header.Set("Content-Type","application/json")
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
	//}

	//client := &http.Client{Transport: tr}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
	fmt.Println("ini response : ", resp)
	//os.Exit(1)
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

func GetMerchant(page *int, size *int, sort *int, email *string) []model.Merchant {
	fmt.Println("masuk ke Fungsi Get")
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var merchant []model.Merchant
	db.Find(&merchant)

	if page == nil && size == nil && sort == nil && email == nil {
		fmt.Println("masuk 1")
		db.Model(&merchant).Find(&merchant)
	}

	if page != nil && size != nil && sort != nil && email == nil {
		fmt.Println("masuk 2", email)
		db.Find(&merchant)
		switch *sort {
		case 1 :
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name desc"},
				}, &merchant)

		case 2 :
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name asc"},
				}, &merchant)
		}
	}

	if page != nil && size != nil && sort != nil && email != nil {
		fmt.Println("masuk 3")
		db.Model(&merchant).Where("merchant_email = ? ", email).Find(&merchant)
		switch *sort {
		case 1 :
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name desc"},
				}, &merchant)
		case 2 :
			pagination.Paging(&pagination.Param{
				DB:      db,
				Page:    *page,
				Limit:   *size,
				OrderBy: []string{"merchant_name asc"},
			}, &merchant)
		}
	}

	if page == nil && size == nil && sort == nil && email != nil {
		fmt.Println("masuk 4")
		db.Model(&merchant).Where("merchant_email =  ?", email).Find(&merchant)
	}



	return merchant
}
