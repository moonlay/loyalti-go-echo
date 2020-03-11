package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

type CardTypeRepository interface {
	CreateCard(card *model.CardType) error
	UpdateCard(card *model.CardType) error
	DeleteCard(card *model.CardType) error
}

type cardtype_repo struct {
	DB *gorm.DB
}

func CreateTypeRepository(db *gorm.DB) CardTypeRepository {
	return &cardtype_repo{
		DB:db,
	}
}

func (p *cardtype_repo) CreateCard(card *model.CardType) error{
	db := database.ConnectionDB()
	//cardObj := *card
	err := db.Create(&card).Error
	if err !=  nil {
		fmt.Print("error : ", err.Error())
	}
	return err
}

func (p *cardtype_repo) UpdateCard(card *model.CardType) error {
	db := database.ConnectionDB()
	err := db.Model(&card).Where("id = ?", card.Id).Update(&card).Error
	return err
}

func (p *cardtype_repo) DeleteCard(card *model.CardType) error {
	db := database.ConnectionDB()
	err := db.Model(&card).Where("id = ?", card.Id).Update("active", false).Error
	return err
}

func CreateCard(card *model.CardType) string{
	db := database.ConnectionDB()
	cardObj := *card
	db.Create(&card)
	return cardObj.CardTypeName
}

func UpdateCard(card *model.CardType) string {
	db := database.ConnectionDB()
	db.Model(&card).Where("id = ?", card.Id).Update(&card)
	return card.CardTypeName
}

func DeleteCard(card *model.CardType) string {
	db := database.ConnectionDB()
	db.Model(&card).Where("id = ?", card.Id).Update("active", false)
	return "berhasil dihapus"
}

func GetCardType(page *int, size *int, sort *int) []model.CardType {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var card []model.CardType
	db.Find(&card)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"card_type_name desc"},
				}, &card)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"card_type_name asc"},
				}, &card)
			}
		}
	}
	db.Close()
	return card
}
