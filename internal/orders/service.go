package orders

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	repo "github.com/Didul-arch/ecom-go-belajar/internal/adapters/mysql/sqlc"
)

type svc struct {
	repo *repo.Queries
	db   *sql.DB
}

func NewService(repo *repo.Queries, db *sql.DB) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has no stock")
)

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	// validate payload
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("Customer ID is required")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("Atleast one item is required")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback()
	qtx := s.repo.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}
	orderID, err := order.LastInsertId()
	if err != nil {
		return repo.Order{}, err
	}

	// look for the product if exists
	for _, item := range tempOrder.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNoStock
		}

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    orderID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCenters,
		})
		if err != nil {
			return repo.Order{}, err
		}

		// Challenge: Update the product stock quantity
	}
	tx.Commit()

	return repo.Order{
		ID:         orderID,
		CustomerID: tempOrder.CustomerID,
		CreatedAt: sql.NullTime{
			Time:  time.Now(), // Isi pake waktu sekarang
			Valid: true,       // Kasih tahu Go kalau ini valid (gak null)
		},
	}, nil
}
