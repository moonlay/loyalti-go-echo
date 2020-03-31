package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func CreateReward(reward *model.Reward) string {
	db := database.ConnectionDB()
	rewardobj := *reward
	db.Create(&rewardobj)

	return "reward berhasil dibuat"
}

func UpdateReward(reward *model.Reward) string {
	db := database.ConnectionDB()
	db.Model(&reward).Where("id = ?", reward.Id).Update(&reward)
	return "Update Berhasil"
}

func DeleteReward(reward *model.Reward) string {
	db := database.ConnectionDB()
	err := db.Model(&reward).Where("id = ?", reward.Id).Update("active", false)
	if err != nil {
		db.Model(&reward).Where("id = ?", reward.Id).Update("is_deleted", true)
	}
	return "Berhasil dihapus"
}

func GetReward(page *int, size *int, sort *int, merchant_email *string) []model.Reward {
	db := database.ConnectionDB()
	var reward []model.Reward
	db.Where("merchant_email = ? ", merchant_email).Find(&reward)
	fmt.Println(reward)
	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				db.Model(&reward).Where("merchant_email = ? ", merchant_email).Find(&reward)
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"reward_name desc"},
				}, &reward)
			}
		case 2:
			if size != nil && page != nil {
				db.Model(&reward).Where("merchant_email = ? ", merchant_email).Find(&reward)
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"reward_name asc"},
				}, &reward)
			}
		}
	}
	return reward
}
