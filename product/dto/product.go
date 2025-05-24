package dto

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductResponse struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
