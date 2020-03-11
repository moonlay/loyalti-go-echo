package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/beevik/guid"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type CardSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	card_repository CardRepository
	card *model.Card
}

func (s *CardSuite) SetupSuite(){
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

	s.card_repository = CreateCardRepository(s.DB)
	fmt.Println("test terlewati")
}

func (s *CardSuite) AfterTest(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitCard(t *testing.T){
	suite.Run(t, new(CardSuite))
}

type card_connection_mock struct {
	mock.Mock
}

func (c card_connection_mock) ConnectionDB() (gorm *gorm.DB) {
	fmt.Println("connection mock")
	return
}

func (s *CardSuite) Test_Create_Card(){
	fmt.Println("test 1")
	var (
		card = model.Card{
			Id:                guid.NewString(),
			Created:           time.Now(),
			CreatedBy:         "",
			Modified:          time.Now(),
			ModifiedBy:        "",
			Active:            true,
			IsDeleted:         false,
			Deleted:           nil,
			DeletedBy:         "",
			Title:             "Kartue",
			Description:       "Deskripsi",
			FontColor:         "Black",
			TemplateColor:     "White",
			IconImage:         "",
			TermsAndCondition: "TnC",
			Benefit:           "1 gelas TeaJus",
			ValidUntil:        time.Time{},
			CurrentPoint:      0,
			IsValid:           true,
			ProgramId:         2,
			CardType:          "Chop",
			IconImageStamp:    "",
			MerchantId:        2,
			Tier:              "",
		}
	)
	fmt.Println("test 2")
	err := s.card_repository.CreateCard(&card)
	fmt.Println("test 3")
	require.NoError(s.T(), err)
	fmt.Println("test 4")
}

func (s *CardSuite) Test_Delete_Card(){
	fmt.Println("test 1")
	var(
		kartu = model.Card{
			Id:                "088c4c24-7ff6-4f4b-a059-50187e5941e1",
		}
	)
	fmt.Println("test 2")
	err := s.card_repository.DeleteCard(&kartu)
	fmt.Println("test 3")
	require.NoError(s.T(), err)
	fmt.Println("test 4")
}

func (s *CardSuite) Test_Update_Card(){
	fmt.Println("Test update card")
	var (
		kartu = model.Card{
			Id:                "f584d39f-ac6b-445b-a592-a98071d5d4e1",
			Active:            true,
			DeletedBy:         "",
			Title:             "KARTU POINT",
			Description:       "Deskripsi",
			FontColor:         "White",
			TemplateColor:     "Black",
			IconImage:         "Icon",
			TermsAndCondition: "TnC",
			Benefit:           "1 Kopi",
			ValidUntil:        time.Time{},
			CurrentPoint:      200,
			IsValid:           true,
			ProgramId:         1,
			CardType:          "Point",
			IconImageStamp:    "",
			MerchantId:        2,
			Tier:              "",
		}
	)
	fmt.Println("test update 1")
	err := s.card_repository.UpdateCard(&kartu)
	fmt.Println("test update 2")
	require.NoError(s.T(), err)
	fmt.Println("test update 3")
}