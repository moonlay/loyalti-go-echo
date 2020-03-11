package repository
//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/jinzhu/gorm"
//	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
//	"github.com/stretchr/testify/mock"
//	"github.com/stretchr/testify/require"
//	"github.com/stretchr/testify/suite"
//	"testing"
//	"time"
//)
//
//type SpecialSuite struct {
//	suite.Suite
//	DB *gorm.DB
//	mock sqlmock.Sqlmock
//
//	special_repository SpecialRepository
//	special *model.SpecialProgram
//}
//
//func (s *SpecialSuite) SetupSuite(){
//	fmt.Println("test suite")
//	var (
//		db *sql.DB
//		err error
//	)
//	fmt.Println("test 2")
//	db, s.mock, err = sqlmock.New()
//	require.NoError(s.T(), err)
//	fmt.Println("test 3")
//
//	s.DB, err = gorm.Open("sqlserver", db)
//	require.NoError(s.T(), err)
//	fmt.Println("test 4")
//
//	s.DB.LogMode(true)
//	fmt.Println("test 5")
//
//	s.special_repository = CreateSpecialRepository(s.DB)
//	fmt.Println("test terlewati")
//}
//
//func (s *SpecialSuite) AfterTest(_, _ string) {
//	require.NoError(s.T(), s.mock.ExpectationsWereMet())
//}
//
//func TestInitSpecial(t *testing.T){
//	suite.Run(t, new(SpecialSuite))
//}
//
//type special_connection_mock struct {
//	mock.Mock
//}
//
//func (c special_connection_mock) ConnectionDB() (*gorm.DB) {
//	fmt.Println("connection mock")
//	return
//}
//
//func (s *SpecialSuite) Test_Create_Card(){
//	fmt.Println("test 1")
//	var (
//		special = model.SpecialProgram{
//			Created:            time.Now(),
//			CreatedBy:          "",
//			Modified:           time.Now(),
//			ModifiedBy:         "",
//			Active:             true,
//			IsDeleted:          false,
//			Deleted:            nil,
//			Deleted_by:         "",
//			ProgramName:        "",
//			ProgramImage:       "",
//			ProgramStartDate:   time.Time{},
//			ProgramEndDate:     time.Time{},
//			ProgramDescription: "",
//			Card:               "",
//			OutletID:           "",
//			MerchantId:         0,
//			MerchantName:       "",
//			CategoryId:         0,
//			Benefit:            nil,
//			TermsAndCondition:  nil,
//			Tier:               nil,
//			RedeemRules:        nil,
//			RewardTarget:       nil,
//			QRCodeId:           nil,
//		}
//	)
//	fmt.Println("test 2")
//	err := s.special_repository.CreateSpecial(&special)
//	fmt.Println("test 3")
//	require.NoError(s.T(), err)
//	fmt.Println("test 4")
//}
//
