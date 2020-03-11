package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetCategory(page *int, size *int, sort *int) []model.MerchantCategory {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var category []model.MerchantCategory
	db.Find(&category)

	if sort != nil {
		switch *sort {
		case 1 :
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"category_name desc"},
				}, &category)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"category_name asc"},
				}, &category)
			}
		}
	}
	db.Close()
	return category
}
