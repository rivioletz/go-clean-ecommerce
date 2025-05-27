package usecase

import (
	"github.com/rivioletz/go-clean-ecommerce/product/dto"
)

type ProductUseCase interface {
	Create(req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	Update(id int64, req *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	Delete(id int64) error
	FindByID(id int64) (*dto.ProductResponse, error)
	FindAll() ([]*dto.ProductResponse, error)
}
