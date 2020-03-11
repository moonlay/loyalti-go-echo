package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	transaction_repository TransactionRepository
	program *model.TransactionMerchant
}

func (s *TransactionSuite) SetupSuite(){
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

	s.transaction_repository = CreateRepositoryTransaction(s.DB)
	fmt.Println("test ini aman")
}

func (s *TransactionSuite) AfterTestTransaction(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitTransaction(t *testing.T){
	suite.Run(t, new(TransactionSuite))
}

func (s *TransactionSuite) Test_Create_Transaction(){
	var (
		transaction = model.TransactionMerchant{
			Id:1,
			TotalTransaction:2000,
		}
	)
	err := s.transaction_repository.CreateTransaction(&transaction)
	if err != nil {
		fmt.Println("Error  : ", err.Error())
	}
	require.NoError(s.T(), err)
}