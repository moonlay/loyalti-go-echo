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

type SpecialProgramSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	specialprogram_repository SpecialProgramRepository
	specialprogram *model.CardType
}

func (s *SpecialProgramSuite) SetupSuite(){
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

	s.specialprogram_repository = CreateSpecialRepository(s.DB)
	fmt.Println("test terlewati")
}

func (s *SpecialProgramSuite) AfterTest(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitSpecialProgram(t *testing.T){
	suite.Run(t, new(SpecialProgramSuite))
}

func (s *SpecialProgramSuite) Test_Create_Special_Program(){
	fmt.Println("test 1 aman")
	var (
		program = model.SpecialProgram{
				Id:2,
				ProgramName: "Program A",
				ProgramDescription:"Program Mantap",
				Active:true,
		}
	)
	fmt.Println("test 2 aman")
	err := s.specialprogram_repository.CreateSpecial(&program)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}

func (s *SpecialProgramSuite) Test_Update_Special_Program(){
	fmt.Println("test 1 aman")
	var (
		program = model.SpecialProgram{
			ProgramName: "Program A",
			ProgramDescription:"Program Mantap 2X",
		}
	)
	fmt.Println("test 2 aman")
	err := s.specialprogram_repository.UpdateSpecial(&program)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}

func (s *SpecialProgramSuite) Test_Delete_Special_Program(){
	fmt.Println("test 1 aman")
	var (
		program = model.SpecialProgram{
			ProgramName: "Program A",
			ProgramDescription:"Program Mantap 2X",
		}
	)
	fmt.Println("test 2 aman")
	err := s.specialprogram_repository.DeleteSpecial(&program)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}
