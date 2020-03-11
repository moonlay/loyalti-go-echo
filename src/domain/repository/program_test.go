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

type ProgramSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	program_repository ProgramRepository
	program *model.Program
}

func (s *ProgramSuite) SetupSuite(){
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

	s.program_repository = CreateRepositoryProgram(s.DB)
	fmt.Println("test ini aman")
}

func (s *ProgramSuite) AfterTestProgram(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitProgram(t *testing.T){
	suite.Run(t, new(ProgramSuite))
}


func (s *ProgramSuite) Test_Create_Program(){
	fmt.Println("test 1 aman")
	var (
			program = model.Program{
				Id:6,
				ProgramName: "dsfgsdf",
				Active:true,
			}
	)
	fmt.Println("test 2 aman")
	err := s.program_repository.CreateProgram(&program)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}

func (s *ProgramSuite) Test_Update_Program(){
	fmt.Println("test update 1 aman")
	var (
		program = model.Program{
			ProgramName:"Diskon Kurang Mantap",
			MerchantId: 0,
		}
	)
	fmt.Println("test update 2 aman")
	err := s.program_repository.UpdateProgram(&program)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 update aman")
}

func (s *ProgramSuite) Test_Delete_Program(){
	fmt.Println("test delete 1 aman")
	var (
		program = model.Program{
			Id: 2,
		}
	)
	fmt.Println("test delete 2 aman")
	err := s.program_repository.DeleteProgram(&program)
	fmt.Println("test delete 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 delete aman")
}

func (s *ProgramSuite) Test_Get_Program() {

}