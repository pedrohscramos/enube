package database

import (
	"github.com/pedrohscramos/enube/internal/entity"
	"gorm.io/gorm"
)

type Supplier struct {
	DB *gorm.DB
}

func NewSupplier(db *gorm.DB) *Supplier {
	return &Supplier{DB: db}
}

func (s *Supplier) FindAll(page, limit int, sort string) ([]entity.Supplier, error) {
	var suppliers []entity.Supplier
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = s.DB.Limit(limit).Offset((page - 1) * limit).Order("id " + sort).Find(&suppliers).Error
	} else {
		err = s.DB.Order("id " + sort).Find(&suppliers).Error
	}
	return suppliers, err
}

func (s *Supplier) FindByID(id int) (*entity.Supplier, error) {
	var supplier entity.Supplier
	err := s.DB.First(&supplier, "id = ?", id).Error
	return &supplier, err
}
