package repository

import (
	"fmt"
	"github.com/beevik/guid"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

//Create Voucher
func CreateVoucher(voucher *model.Voucher) string {
	db := database.ConnectionDB()
	newvoucher := model.Voucher{
		Id:                       guid.NewString(),
		Created:                  time.Now(),
		CreatedBy:                voucher.CreatedBy,
		Modified:                 time.Now(),
		ModifiedBy:               voucher.ModifiedBy,
		Active:                   voucher.Active,
		IsDeleted:                voucher.IsDeleted,
		Deleted:                  nil,
		DeletedBy:                "",
		VoucherName:              voucher.VoucherName,
		StartDate:                time.Time{},
		EndDate:                  time.Time{},
		VoucherDescription:       voucher.VoucherDescription,
		VoucherTermsAndCondition: voucher.VoucherTermsAndCondition,
		IsPushNotification:       voucher.IsPushNotification,
		IsGiveVoucher:            voucher.IsGiveVoucher,
		VoucherPeriod:            voucher.VoucherPeriod,
		RewardTermsAndCondition:  voucher.VoucherTermsAndCondition,
		BackgroundVoucherPattern: voucher.BackgroundVoucherPattern,
		BackgroundVoucherColour:  voucher.BackgroundVoucherColour,
		MerchantId:               voucher.MerchantId,
		OutletId:                 voucher.OutletId,
		ProgramId:                voucher.ProgramId,
	}
	db.Create(&newvoucher)
	return "voucher berhasil dibuat"
}

//Update Voucher using program id
func UpdateVoucher(voucher *model.Voucher) string {
	db := database.ConnectionDB()
	db.Model(&voucher).Where("program_id = ? ", voucher.ProgramId).Update(&voucher)
	return "Update Berhasil"
}

//Delete Vouocher using program id
func DeleteVoucher(voucher *model.Voucher) string {
	db := database.ConnectionDB()
	err := db.Model(&voucher).Where("program_id = ?", voucher.ProgramId).Update("active", false)
	if err != nil {
		fmt.Println("error : ", err.Error)
		db.Model(&voucher).Where("program_id = ?", voucher.ProgramId).Update("is_deleted", true)
	}
	return "berhasil dihapus"
}

//Get Voucher by Merchant_id and have sorting
func GetVoucher (page *int, size *int, sort *int, merchant_id *int) []model.Voucher {
	db := database.ConnectionDB()
	var voucher []model.Voucher
	db.Find(&voucher)
		if page == nil && size == nil && sort == nil && merchant_id == nil {
			db.Model(&voucher).Find(&voucher)
		}
		if page != nil && size != nil && sort == nil && merchant_id == nil {
			fmt.Println("masuk ke 2")
			db.Model(&voucher).Find(&voucher)
			pagination.Paging(&pagination.Param{
				DB:      db,
				Page:    *page,
				Limit:   *size,
				OrderBy: []string{"voucher_name asc"},
			}, &voucher)
		}
		if page != nil && size != nil && sort != nil && merchant_id == nil {
			fmt.Println("masuk ke 3")
			switch *sort {
			case 1 :
				db.Model(&voucher).Find(&voucher)
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"voucher_name asc"},
				}, &voucher)
			case 2:
				db.Model(&voucher).Find(&voucher)
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"voucher_name asc"},
				}, &voucher)
			}
		}
		if page != nil && size != nil && sort != nil && merchant_id != nil {
			fmt.Println("masuk ke 4")
			switch *sort {
			case 1 :
				db.Model(&voucher).Where("merchant_id = ?", merchant_id).Find(&voucher)
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"voucher_name asc"},
				}, &voucher)
			case 2:
				db.Model(&voucher).Where("merchant_id = ?", merchant_id).Find(&voucher)
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"voucher_name asc"},
				}, &voucher)
			}
		}

	db.Close()
	return voucher
}

