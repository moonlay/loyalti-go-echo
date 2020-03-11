package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)

func GetSocialMedia(page *int, size *int, sort *int) []model.MerchantSocialMedia {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var sosmed []model.MerchantSocialMedia
	db.Find(&sosmed)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"social_media_name desc"},
				}, &sosmed)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"social_media_name asc"},
				}, &sosmed)
			}
		}
	}
	db.Close()
	return sosmed
}
