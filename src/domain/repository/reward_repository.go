package repository

import (
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func CreateReward (reward *model.Reward) string {
	db := database.ConnectionDB()
	rewardobj := *reward
	db.Create(&rewardobj)

	return "reward berhasil dibuat"
}

func UpdateReward (reward *model.Reward) string {
	db := database.ConnectionDB()
	db.Model(&reward).Where("id = ?", reward.Id).Update(&reward)
	return "Update Berhasil"
}

func DeleteReward (reward *model.Reward) string {
	db := database.ConnectionDB()
	err := db.Model(&reward).Where("id = ?",reward.Id).Update("active", false)
	if err != nil {
		db.Model(&reward).Where("id = ?",reward.Id).Update("is_deleted", true)
	}

}

func GetReward (page *int, size *int , sort *int, merchant_id *int ){

}
