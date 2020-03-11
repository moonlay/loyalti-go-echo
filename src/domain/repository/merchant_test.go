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
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	merchant *model.NewMerchantCommand
}

func (s *Suite) SetupSuite(){
	fmt.Println("test 10 aman")
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
	fmt.Println("test 40 aman")

	s.DB.LogMode(true)
	fmt.Println("test 50 aman")

	s.repository = CreateRepository(s.DB)
	fmt.Println("test ini aman")
}

func (s *Suite) AfterTest(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T){
	suite.Run(t, new(Suite))
}

type connection_mock struct {
	mock.Mock
}

func (c connection_mock) ConnectionDB2() (gorm *gorm.DB){
	fmt.Println("masuk mockingan connection")
	return
}

func (s *Suite) Test_Create_Merchant(){
	fmt.Println("test 1 aman")
	var (
		merchant = model.NewMerchantCommand{
			Id:            19,
			MerchantEmail: "abc",
			Active:true,
		}
	)
	fmt.Println("test 2 aman")
	 err := s.repository.CreateMerchant(&merchant)
	fmt.Println("test 3 aman")
	 require.NoError(s.T(), err)
	 fmt.Println("test 4 aman", err)
}

func (s *Suite) Test_Update_Merchant(){
	fmt.Println("test 1 update aman")
	var (
		merchant = model.NewMerchantCommand{
			MerchantEmail: "abcd",
			MerchantAddress: "jalan 1",
		}
	)
	fmt.Println("test 2 update aman")
	err := s.repository.UpdateMerchant(&merchant)
	fmt.Println("test 3 update aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 update aman", err)
}

func (s *Suite) Test_Delete_Merchant(){
	fmt.Println("test 1 delete aman")
	var (
		merchant = model.NewMerchantCommand{
			MerchantEmail: "abcd",
		}
	)
	fmt.Println("test 2 delete aman")
	err := s.repository.DeleteMerchant(&merchant)
	fmt.Println("test 3 delete aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 delete aman")
}

