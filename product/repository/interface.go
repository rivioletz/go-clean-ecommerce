package repository

import (
	"github.com/rivioletz/go-clean-ecommerce/product/entity"
)

type ProductRepository interface {
	Create(product *entity.Product) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id int64) error
	FindByID(id int64) (*entity.Product, error)
	FindAll() ([]*entity.Product, error)
}
