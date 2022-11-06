package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/danangkonang/product/config"
	"github.com/danangkonang/product/model"
)

type ServiceCart interface {
	FindMyCard(userId string) ([]*model.CardResponse, error)
	CreateCart(prd *model.CardRequest) error
}

func NewServiceCart(Con *config.DB) ServiceCart {
	return &Database{
		Postgresql: Con.Postgresql,
	}
}

func (r *Database) FindMyCard(userId string) ([]*model.CardResponse, error) {
	prd := make([]*model.CardResponse, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
		SELECT c.cart_id, c.product_id, p.product_name, c.qty, p.price, p.price * c.qty AS sub_total
		FROM carts c
		LEFT JOIN products p
		ON p.product_id = c.product_id
		WHERE c.user_id=$1 AND c.is_checkout=$2
	`

	row, err := r.Postgresql.QueryContext(ctx, query, userId, false)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("internal server error")
	}

	for row.Next() {
		pd := new(model.CardResponse)
		err := row.Scan(&pd.CartId, &pd.ProductId, &pd.ProductName, &pd.Qty, &pd.Price, &pd.SubTotal)
		if err != nil {
			return nil, err
		}
		prd = append(prd, pd)
	}
	return prd, nil
}

func (r *Database) CreateCart(prd *model.CardRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.Postgresql.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var qt int64
	findProdQuery := "SELECT qty FROM products WHERE product_id=$1"
	errFind := tx.QueryRowContext(ctx, findProdQuery, prd.ProductId).Scan(&qt)
	if err != nil {
		tx.Rollback()
		return errFind
	}

	var cid string
	var qtc int64
	isExistsQuery := "SELECT cart_id, qty FROM carts WHERE product_id=$1 AND user_id=$2"
	errExists := r.Postgresql.QueryRowContext(ctx, isExistsQuery, prd.ProductId, prd.UserId).Scan(&cid, &qtc)
	if errExists != nil {
		insQtyQuery := "INSERT INTO carts (cart_id, user_id, product_id, qty, is_checkout) VALUES ($1, $2, $3, $4, $5)"
		_, errIns := tx.ExecContext(ctx, insQtyQuery, prd.CartId, prd.UserId, prd.ProductId, prd.Qty, prd.IsCheckout)
		if errIns != nil {
			tx.Rollback()
			return errIns
		}
	} else {
		newQty := prd.Qty + qtc
		updQtyQuery := "UPDATE carts SET qty=$1 WHERE cart_id=$2"
		_, errUpt := tx.ExecContext(ctx, updQtyQuery, newQty, cid)
		if errUpt != nil {
			tx.Rollback()
			return errUpt
		}
	}

	lastQty := qt - prd.Qty
	updLatesQuery := "UPDATE products SET qty=$1 WHERE product_id=$2"
	_, errUpd := r.Postgresql.ExecContext(ctx, updLatesQuery, lastQty, prd.ProductId)
	if errUpd != nil {
		tx.Rollback()
		return errUpd
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
