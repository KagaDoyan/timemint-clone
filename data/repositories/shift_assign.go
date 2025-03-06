package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type shiftAssignRepository struct {
	db *gorm.DB
}

type ShiftAssignRepository interface {
	Create(shiftAssign entities.ShiftAssignment) (*entities.ShiftAssignment, error)
	BatchCreate(shiftAssigns []entities.ShiftAssignment) ([]entities.ShiftAssignment, error)
	Delete(id uint) error
	FindAll(page, limit int) ([]entities.ShiftAssignment, int64, error)
	FindById(id uint) (*entities.ShiftAssignment, error)
	CalendarShift(month, year int) ([]entities.ShiftAssignment, error)
	ShiftAssignmentReport(start, end string) ([]entities.ShiftAssignment, error)
}

func NewShiftAssignRepository(db *gorm.DB) ShiftAssignRepository {
	db.AutoMigrate(&entities.ShiftAssignment{})
	return &shiftAssignRepository{
		db: db,
	}
}

func (r shiftAssignRepository) Create(shiftAssign entities.ShiftAssignment) (*entities.ShiftAssignment, error) {
	err := r.db.Where("employee_id = ? AND shift_id = ? AND date = ?", shiftAssign.EmployeeID, shiftAssign.ShiftID, shiftAssign.Date).FirstOrCreate(&shiftAssign).Error
	if err != nil {
		return nil, err
	}
	var result entities.ShiftAssignment
	err = r.db.Preload(clause.Associations).First(&result, shiftAssign.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r shiftAssignRepository) BatchCreate(shiftAssigns []entities.ShiftAssignment) ([]entities.ShiftAssignment, error) {
	err := r.db.Create(&shiftAssigns).Error
	if err != nil {
		return nil, err
	}
	return shiftAssigns, nil
}

func (r shiftAssignRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.ShiftAssignment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r shiftAssignRepository) FindAll(page, limit int) ([]entities.ShiftAssignment, int64, error) {
	var shiftAssigns []entities.ShiftAssignment
	var total int64
	err := r.db.Model(&entities.ShiftAssignment{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Order("created_at desc").Preload(clause.Associations).Preload("Employee.Role").Find(&shiftAssigns).Error
	if err != nil {
		return nil, 0, err
	}
	return shiftAssigns, total, nil
}

func (r shiftAssignRepository) FindById(id uint) (*entities.ShiftAssignment, error) {
	var shiftAssign entities.ShiftAssignment
	err := r.db.Preload(clause.Associations).Preload("Employee.Role").First(&shiftAssign, id).Error
	if err != nil {
		return nil, err
	}
	return &shiftAssign, nil
}

func (r shiftAssignRepository) CalendarShift(month int, year int) ([]entities.ShiftAssignment, error) {
	var shiftAssigns []entities.ShiftAssignment
	err := r.db.Preload(clause.Associations).Preload("Employee.Role").Where("MONTH(date) = ? AND YEAR(date) = ?", month, year).Find(&shiftAssigns).Error
	if err != nil {
		return nil, err
	}
	return shiftAssigns, nil
}

func (r shiftAssignRepository) ShiftAssignmentReport(start, end string) ([]entities.ShiftAssignment, error) {
	var daterangewhere string
	query := r.db.Preload(clause.Associations).Preload("Employee.Role")
	if start != "" {
		if end == "" {
			end = start
		}
		daterangewhere = `date(date) BETWEEN ? AND ?`
		query = query.Where(daterangewhere, start, end)
	}
	var shiftAssigns []entities.ShiftAssignment
	err := query.Find(&shiftAssigns).Error
	if err != nil {
		return nil, err
	}
	return shiftAssigns, nil
}
