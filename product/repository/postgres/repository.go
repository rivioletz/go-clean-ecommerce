package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rivioletz/go-clean-ecommerce/product/entity"
	"github.com/rivioletz/go-clean-ecommerce/product/repository"
)

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) repository.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(product *entity.Product) (*entity.Product, error) {
	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(context.Background(), query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(product *entity.Product) (*entity.Product, error) {
	query := `UPDATE products SET name = $1, price = $2 WHERE id = $3 RETURNING id`
	cmdTag, err := r.db.Exec(context.Background(), query, product.Name, product.Price, product.ID)
	if err != nil {
		return nil, err
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, errors.New("No product updated, product may not exist")
	}
	return product, nil
}

func (r *productRepository) Delete(id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	cmdTag, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("No product deleted, product may not exist")
	}
	return nil
}

func (r *productRepository) FindByID(id int64) (*entity.Product, error) {
	query := `SELECT id, name, price FROM products WHERE id = $1`
	product := &entity.Product{}
	err := r.db.QueryRow(context.Background(), query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil // Product not found
		}
		return nil, err
	}
	return product, nil
}

func (r *productRepository) FindAll() ([]*entity.Product, error) {
	query := `SELECT id, name, price FROM products`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		product := &entity.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
