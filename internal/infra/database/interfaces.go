package database

import "github.com/pedrohscramos/enube/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type SupplierInterface interface {
	FindAll(page, limit int, sort string) ([]entity.Supplier, error)
	FindByID(id int) (*entity.Supplier, error)
}
