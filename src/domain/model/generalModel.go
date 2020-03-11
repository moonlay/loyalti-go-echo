package model

import (
	"time"
)

type GeneralModels struct {
	Id         int 			`gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created    time.Time 	`json:"created"`
	CreatedBy  string		`json:"created_by"`
	Modified   time.Time	`json:"modified"`
	ModifiedBy string		`json:"modified_by"`
	Active     bool			`json:"active"`
	IsDeleted  bool			`json:"is_deleted"`
	Deleted    *time.Time	`json:"deleted"`
	Deleted_by string		`json:"deleted_by"`
}