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

type CardTypeSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	cardtype_repository CardTypeRepository
	tipe *model.CardType
}

func (s *CardTypeSuite) SetupSuite(){
	fmt.Println("test suite")
	var (
		db *sql.DB
		err error
	)
	fmt.Println("test 2")
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	fmt.Println("test 3")

	s.DB, err = gorm.Open("sqlserver", db)
	require.NoError(s.T(), err)
	fmt.Println("test 4")

	s.DB.LogMode(true)
	fmt.Println("test 5")

	s.cardtype_repository = CreateTypeRepository(s.DB)
	fmt.Println("test terlewati")
}

func (s *CardTypeSuite) AfterTest(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitCardType(t *testing.T){
	suite.Run(t, new(CardTypeSuite))
}

func (s *CardTypeSuite) Test_Create_Type_Card(){
	fmt.Println("test 1 aman")
	var (
		tipe = model.CardType{
			CardTypeName:"abc",
			Id:4,
			Active:true,
		}
	)
	fmt.Println("test 2 aman")
	err := s.cardtype_repository.CreateCard(&tipe)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}

func (s *CardTypeSuite) Test_Update_Type_Card(){
	fmt.Println("test 1 aman")
	var (
		tipe = model.CardType{
			CardTypeName:"adc",
			CreatedBy:"admin",
		}
	)
	fmt.Println("test 2 aman")
	err := s.cardtype_repository.UpdateCard(&tipe)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}

func (s *CardTypeSuite) Test_Delete_Type_Card(){
	fmt.Println("test 1 aman")
	var (
		tipe = model.CardType{
			CardTypeName:"adc",
			CreatedBy:"admin",
			Id: 1,
		}
	)
	fmt.Println("test 2 aman")
	err := s.cardtype_repository.DeleteCard(&tipe)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}





