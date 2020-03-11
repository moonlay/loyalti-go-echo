package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var SqlDB *gorm.DB

type Connection_Interface interface {
	ConnectionDB2() *gorm.DB
}

type Conn struct {
	Connection_Interface
}

type Connection struct {
}

//func ConnectionDB() *gorm.DB{
//	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
//	if err != nil {
//		panic("failed to connect database")
//	}
//	//defer db.Close()
//	SqlDB = db
//	return db
//}

func ConnectionDB() *gorm.DB {
	db, err := gorm.Open("mssql", "sqlserver://moonlay:Standar123.@loyaltiexpress.database.windows.net?database=loyalti.MerchantDB.Dev")
	//db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=Loyalti.MerchantDb.Dev")

	if err != nil {
		//panic("failed to connect database")
		fmt.Println("Error : ", err.Error)
	}
	//defer db.Close()
	SqlDB = db
	return db
}

//func (c Connection) ConnectionDB2() *gorm.DB{
//	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
//	if err != nil {
//		panic("failed to connect database")
//	}
//	//defer db.Close()
//
//	return db
//}

//func ConnectPostgre() *gorm.DB {
//	db, err := gorm.Open("postgres", "host=localhost port=5432 user=sa dbname=loyalti-postgre sslmode=disable password=Moonlay2019.")
//	//db, err := gorm.Open("postgres", "postgres:postgres@/loyalti")
//	//defer db.Close()
//
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	return db
//}
