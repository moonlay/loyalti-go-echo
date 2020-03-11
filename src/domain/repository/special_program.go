package repository

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/sirupsen/logrus"
)

type SpecialProgramRepository interface {
	CreateSpecial(special *model.SpecialProgram) error
	UpdateSpecial(special *model.SpecialProgram) error
	DeleteSpecial(special *model.SpecialProgram) error
}

type special_repo struct {
	DB *gorm.DB
}

func CreateSpecialRepository(db *gorm.DB) SpecialProgramRepository {
	return &special_repo{
		DB: db,
	}
}

func (p *special_repo) CreateSpecial(special *model.SpecialProgram) error {
	db := database.ConnectionDB()
	specialObj := *special
	err := db.Create(&specialObj).Error
	return err
}

func (p *special_repo) UpdateSpecial(special *model.SpecialProgram) error {
	db := database.ConnectionDB()
	err := db.Model(&special).Where("program_name = ?", special.ProgramName).Update(&special).Error
	return err
}

func (p *special_repo) DeleteSpecial(special *model.SpecialProgram) error {
	db := database.ConnectionDB()
	err := db.Model(&special).Where("program_name = ?", special.ProgramName).Update("active", false).Error
	return err
}

//type SpecialRepository interface {
//	CreateSpecial (newspecial *model.SpecialProgram) error
//	UpdateSpecial (newspecial *model.SpecialProgram) error
//	DeleteSpecial (newspecial *model.SpecialProgram) error
//}
//
//type special_repo struct {
//	DB *gorm.DB
//}
//
//func CreateSpecialRepository (db *gorm.DB) SpecialRepository {
//	return  &special_repo{
//		DB:db,
//	}
//}
//
//func (p *special_repo) CreateSpecial (newspecial *model.SpecialProgram) error {
//	db:= database.ConnectionDB()
//	err := db.Create(&newspecial).Error
//	if err == nil {
//		fmt.Println("Tidak ada Error!")
//	}
//	return err
//}
//
//func (p *special_repo) UpdateSpecial (newspecial *model.SpecialProgram) error {
//	db := database.ConnectionDB()
//
//}

func CreateSpecial(special *model.SpecialProgram) string {
	db := database.ConnectionDB()
	specialObj := *special
	db.Create(&specialObj)
	return specialObj.ProgramName
}

func UpdateSpecial(special *model.SpecialProgram) string {
	db := database.ConnectionDB()
	db.Model(&special).Where("program_name = ?", special.ProgramName).Update(&special)
	return special.ProgramName
}

func DeleteSpecial(special *model.SpecialProgram) string {
	db := database.ConnectionDB()
	db.Model(&special).Where("program_name = ?", special.ProgramName).Update("active", false)
	return "berhasil dihapus"
}

func GetSpecialProgram(page *int, size *int, sort *int, category *int, id *int) []model.SpecialProgram {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var program []model.SpecialProgram
	var rows *sql.Rows
	var err error
	var total int

	if sort != nil {
		switch *sort {
		case 1:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
				fmt.Println("test")
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
				fmt.Println("apakah masuk")
				if err != nil {
					panic(err)
				}
			}
		case 2:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 3:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 4:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
				rows.Close()
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		}
	} else {
		if page != nil && size != nil {
			rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("masuk ga")
			rows, err = db.Find(&program).Rows()
			if err != nil {
				panic(err)
			}
		}
	}
	if id != nil {
		rows, err = db.Where("id = ?", id).First(&program).Rows()
		if err != nil {
			panic(err)
		}
	}

	result := make([]model.SpecialProgram, 0)

	for rows.Next() {
		t := new(model.SpecialProgram)
		fmt.Println(t)
		err = rows.Scan(
			&t.Id,
			&t.Created,
			&t.CreatedBy,
			&t.Modified,
			&t.ModifiedBy,
			&t.Active,
			&t.IsDeleted,
			&t.Deleted,
			&t.Deleted_by,
			&t.ProgramName,
			&t.ProgramImage,
			&t.ProgramStartDate,
			&t.ProgramEndDate,
			&t.ProgramDescription,
			&t.Card,
			&t.OutletID,
			&t.MerchantId,
			&t.CategoryId,
			&t.Benefit,
			&t.TermsAndCondition,
			&t.Tier,
			&t.RedeemRules,
			&t.RewardTarget,
			&t.QRCodeId,
		)
		merchant := new(model.Merchant)

		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("id = ?", t.MerchantId).
			First(&merchant)
		//t.MerchantName = merchant.MerchantName
		if err != nil {
			logrus.Error(err)
			return nil
		}
		result = append(result, *t)
	}
	db.Close()
	return result
}
