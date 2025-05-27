package usecase

import (
	"errors"

	"github.com/rivioletz/go-clean-ecommerce/product/dto"
	"github.com/rivioletz/go-clean-ecommerce/product/entity"
	"github.com/rivioletz/go-clean-ecommerce/product/repository"
)

type productUseCase struct {
	repo repository.ProductRepository
}

// NewProductUseCase creates a new instance of ProductUseCase
func NewProductUseCase(r repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: r,
	}
}

// Create creates a new product
func (uc *productUseCase) Create(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := &entity.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	createdProduct, err := uc.repo.Create(product)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:    createdProduct.ID,
		Name:  createdProduct.Name,
		Price: createdProduct.Price,
	}, nil
}

// Update updates an existing product
func (uc *productUseCase) Update(id int64, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	product.Name = req.Name
	product.Price = req.Price

	updatedProduct, err := uc.repo.Update(product)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:    updatedProduct.ID,
		Name:  updatedProduct.Name,
		Price: updatedProduct.Price,
	}, nil
}

// Delete deletes a product by ID
func (uc *productUseCase) Delete(id int64) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

// FindByID retrieves a product by ID
func (uc *productUseCase) FindByID(id int64) (*dto.ProductResponse, error) {
	product, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	return &dto.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}

// FindAll retrieves all products
func (uc *productUseCase) FindAll() ([]*dto.ProductResponse, error) {
	products, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []*dto.ProductResponse
	for _, product := range products {
		result = append(result, &dto.ProductResponse{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return result, nil
}
