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

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"voucher_name desc"},
				}, &voucher)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"voucher_name asc"},
				}, &voucher)
			}
		}
	}

	if merchant_id != nil {
		if page != nil && size != nil {
			pagination.Paging(&pagination.Param{
				DB:      db,
				Page:    *page,
				Limit:   *size,
				OrderBy: []string{"voucher_name asc"},
			}, &voucher)
			db.Order("merchant_id").Where("merchant_id = ?", merchant_id).Find(&voucher)

		} else {
			db.Order("merchant_id").Where("merchant_id = ?", merchant_id).Find(&voucher)
		}
	}
	db.Close()
	return voucher
}

