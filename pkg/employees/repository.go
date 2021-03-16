package employees

import (
	"fmt"
	"gorm.io/gorm"
)

type Repo interface {
	Save(employee Employee) error
	GetById(id uint) (Employee, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repo {
	return &repoImpl{
		db: db,
	}
}

func (e *repoImpl) Save(employee Employee) error {
	err := e.db.Create(&employee).Error
	if err != nil {
		return fmt.Errorf("repo:Save %w", err)
	}
	return nil
}

func (e *repoImpl) GetById(id uint) (Employee, error) {
	var employee Employee
	err := e.db.Where("Id = ?", id).First(&employee).Error
	if err == gorm.ErrRecordNotFound {
		return Employee{}, fmt.Errorf("repo:GetById Employee with id %d not found", id)
	}
	return employee, nil
}
