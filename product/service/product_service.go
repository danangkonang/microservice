package service

import (
	"context"
	"errors"
	"time"

	"github.com/danangkonang/product/config"
	"github.com/danangkonang/product/model"
)

type ServiceProduct interface {
	Findproduct() ([]*model.ProductResponse, error)
	CreateProduct(prd *model.ProductRequest) error
}

func NewServiceProduct(Con *config.DB) ServiceProduct {
	return &Database{
		Postgresql: Con.Postgresql,
	}
}

func (r *Database) Findproduct() ([]*model.ProductResponse, error) {
	prd := make([]*model.ProductResponse, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row, err := r.Postgresql.QueryContext(ctx, "SELECT product_id, product_name, qty, price FROM products")
	if err != nil {
		return nil, errors.New("internal server error")
	}

	for row.Next() {
		pd := new(model.ProductResponse)
		err := row.Scan(&pd.ProductId, &pd.ProductName, &pd.Qty, &pd.Price)
		if err != nil {
			return nil, err
		}
		prd = append(prd, pd)
	}
	return prd, nil
}

func (r *Database) CreateProduct(prd *model.ProductRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := "INSERT INTO products (product_id, product_name, qty, price) VALUES ($1, $2, $3, $4)"
	_, err := r.Postgresql.ExecContext(ctx, query, prd.ProductId, prd.ProductName, prd.Qty, prd.Price)
	if err != nil {
		return err
	}
	return nil
}
