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

type EmployeeSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	employee_repository EmployeeRepository
	employee *model.Employee
}

func (s *EmployeeSuite) SetupSuite(){
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

	s.employee_repository = CreateEmployeeRepository(s.DB)
	fmt.Println("test ini aman")
}

func (s *EmployeeSuite) AfterTest(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitEmployee(t *testing.T){
	suite.Run(t, new(EmployeeSuite))
}

type employee_connection_mock struct {
	mock.Mock
}

func (c employee_connection_mock) ConnectionDB() (gorm *gorm.DB) {
	fmt.Println("masuk connection mock")
	return
}

func (s *EmployeeSuite) Test_Create_Employee(){
	fmt.Println("test 1 aman")
	var (
		employee = model.Employee{
			Id:            guid.NewString(),
			Created:       time.Now(),
			CreatedBy:     "Admin",
			Modified:      time.Now(),
			ModifiedBy:    "Admin",
			Active:        true,
			IsDeleted:     false,
			Deleted:       nil,
			Deleted_by:    "",
			EmployeeName:  "Felix",
			EmployeeEmail: "felix@email.com",
			EmployeePin:   "123456",
			EmployeeRole:  "Cashier",
			OutletId:      "2",
			OutletName:    "Bu Bejo",
		}
	)
	fmt.Println("test 2 aman")
	err := s.employee_repository.CreateEmployee(&employee)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman")
}

func (s *EmployeeSuite) Test_Update_Employee(){
	fmt.Println("Test update employee")
	var (
		employe = model.Employee{
			Created:       time.Time{},
			CreatedBy:     "",
			Modified:      time.Time{},
			ModifiedBy:    "",
			Active:        false,
			IsDeleted:     false,
			Deleted:       nil,
			Deleted_by:    "",
			EmployeeName:  "David",
			EmployeeEmail: "david@email.com -update",
			EmployeePin:   "969696",
			EmployeeRole:  "Cashier -update",
			OutletId:      "1",
			OutletName:    "Bu Nanik -update",
		}
	)
	fmt.Println("test update 1")
	err := s.employee_repository.UpdateEmployee(&employe)
	fmt.Println("test update 2")
	require.NoError(s.T(), err)
	fmt.Println("test update 3")
}

func (s *EmployeeSuite) Test_Delete_Employee(){
	fmt.Println("test 1")
	var (
		employ = model.Employee{
			Id:"78cadd73-b45b-41e6-938f-b18b89e9866a",
		}
	)
	fmt.Println("test 2")
	err := s.employee_repository.DeleteEmployee(&employ)
	fmt.Println("test 3")
	require.NoError(s.T(), err)
	fmt.Println("test 4")
}