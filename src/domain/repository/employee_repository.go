package repository

import (
	"fmt"
	"github.com/beevik/guid"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

type EmployeeRepository interface {
	CreateEmployee (newemployee *model.Employee) error
	UpdateEmployee (newemployee *model.Employee) error
	DeleteEmployee (newemployee *model.Employee) error
}

type employee_repo struct {
	DB *gorm.DB
}

func (p employee_repo) CreateEmployee (newemployee *model.Employee) error {
	fmt.Println("masuk fungsi")
	employee := model.Employee{
		Id:            guid.NewString(),
		Created:       time.Now(),
		CreatedBy:     "",
		Modified:      time.Now(),
		ModifiedBy:    "",
		Active:        false,
		IsDeleted:     false,
		Deleted:       nil,
		Deleted_by:    "",
		EmployeeName:  newemployee.EmployeeName,
		EmployeeEmail: newemployee.EmployeeEmail,
		EmployeePin:   newemployee.EmployeePin,
		EmployeeRole:  newemployee.EmployeeRole,
		OutletId:      newemployee.OutletId,
		OutletName:    newemployee.OutletName,
	}

	db := database.ConnectionDB()
	err := db.Create(&employee).Error
	if err != nil {
		fmt.Println("Tak ada error")
	}
	return err
}

func CreateEmployeeRepository (db *gorm.DB) EmployeeRepository {
	return &employee_repo{
		DB:db,
	}
}

func (p *employee_repo) UpdateEmployee(updateemployee *model.Employee) error{
	db := database.ConnectionDB()
	employee := model.Employee{
		Id:            guid.NewString(),
		Created:       time.Now(),
		CreatedBy:     "",
		Modified:      time.Now(),
		ModifiedBy:    "",
		Active:        true,
		IsDeleted:     false,
		Deleted:       nil,
		Deleted_by:    "",
		EmployeeName:  updateemployee.EmployeeName,
		EmployeeEmail: updateemployee.EmployeeEmail,
		EmployeePin:   updateemployee.EmployeePin,
		EmployeeRole:  updateemployee.EmployeeRole,
		OutletId:      updateemployee.OutletId,
		OutletName:    updateemployee.OutletName,
	}
	err := db.Model(&employee).Where("outlet_id = ?", employee.OutletId).Update(&employee).Error
	return err
}

func (p *employee_repo) DeleteEmployee(deleteemployee *model.Employee) error {
	db := database.ConnectionDB()

	err := db.Model(&deleteemployee).Where("id = ?", deleteemployee.Id).Update("active", true).Error
	if err == nil {
		fmt.Println("tidak ada error")
	}
	return err
}

func CreateEmployee(employee *model.Employee) string{
	db := database.ConnectionDB()
	employee.Id = guid.NewString()
	db.Create(&employee)
	return employee.EmployeeEmail
}


func UpdateEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	db.Model(&employee).Where("employee_email = ?", employee.EmployeeEmail).Update(&employee)
	return employee.EmployeeEmail
}


func DeleteEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	db.Model(&employee).Where("outlet_id = ?", employee.OutletId).Delete(&employee)
	//db.Model(&employee).Where("employee_email = ?",employee.EmployeeEmail).Update("active", false)
	return "berhasil dihapus"
}

func GetEmployee(page *int, size *int, sort *int) []model.Employee {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var employee []model.Employee
	db.Find(&employee)



	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"employee_name desc"},
				}, &employee)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"employee_name asc"},
				}, &employee)
			}
		}
	}
	db.Close()
	return employee
}
