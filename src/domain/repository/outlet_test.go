package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type OutletSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	outlet_repository OutletRepository
	outlet *model.Outlet
}

func (s *OutletSuite) SetupSuite(){
	fmt.Println("test suite aman")
	var (
		db *sql.DB
		err error
	)
	fmt.Println("test 20 aman")
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	fmt.Println("test 30 aman")

	s.DB, err = gorm.Open("sqlserver", db)
	require.NoError(s.T(), err)
	fmt.Println("test40 aman")

	s.DB.LogMode(true)
	fmt.Println("test50 aman")

	s.outlet_repository = CreateOutletRepository(s.DB)
	fmt.Println("test ini aman")
}

func (s *OutletSuite) AfterTest(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitOutlet(t *testing.T){
	suite.Run(t, new(OutletSuite))
}

type outlet_connection_mock struct {
	mock.Mock
}

func (c outlet_connection_mock) ConnectionDB2() (gorm *gorm.DB){
	fmt.Println("masuk connection mock")
	return
}

func (s *OutletSuite) Test_Create_Outlet(){
	fmt.Println("test 1 aman")
	var (
		outlet = model.Outlet{
			Created:          time.Now(),
			CreatedBy:        "",
			Modified:         time.Now(),
			ModifiedBy:       "",
			Active:           true,
			IsDeleted:        false,
			Deleted:          nil,
			Deleted_by:       "",
			OutletName:       "P3",
			OutletAddress:    "Energy Tower",
			OutletPhone:      "081357867890",
			OutletCity:       "Jakarta",
			OutletProvince:   "DKI Jakarta",
			OutletPostalCode: "14078",
			OutletLongitude:  "",
			OutletLatitude:   "",
			OutletDay:        time.Time{},
			OutletHour:       time.Time{},
			MerchantId:       2,
		}
	)
	fmt.Println("test 2 aman")
	err := s.outlet_repository.CreateOutlet(&outlet)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman")
}

func (s *OutletSuite) Test_Update_Outlet(){
	fmt.Println("test 1 update outlet")
	var (
		outlet = model.Outlet{
			Created:          time.Time{},
			CreatedBy:        "",
			Modified:         time.Time{},
			ModifiedBy:       "",
			Active:           false,
			IsDeleted:        false,
			Deleted:          nil,
			Deleted_by:       "",
			OutletName:       "P3",
			OutletAddress:    "Energy Building",
			OutletPhone:      "081357867890",
			OutletCity:       "Jakarta",
			OutletProvince:   "DKI Jakarta -update",
			OutletPostalCode: "14078",
			OutletLongitude:  "",
			OutletLatitude:   "",
			OutletDay:        time.Time{},
			OutletHour:       time.Time{},
			MerchantId:       2,
		}
	)
	fmt.Println("test 2 update")
	err := s.outlet_repository.UpdateOutlet(&outlet)
	fmt.Println("test 3 update")
	require.NoError(s.T(), err)
	fmt.Println("test4 update")
}

func (s *OutletSuite) Test_Delete_Outlet(){
	fmt.Println("test 1")
	var (
		outlet = model.Outlet{
			Id:               1,
		}
	)
	fmt.Println("test 2")
	err := s.outlet_repository.DeleteOutlet(&outlet)
	fmt.Println("test 3")
	require.NoError(s.T(), err)
	fmt.Println("test 4")
}