package postgres

import (
	"context"
	"fmt"
	"product-storage/models"

	"github.com/jackc/pgx/v5"
)

const (
	reserveProductStmt = "INSERT INTO reservations (storage_id, product_id, client_id, amount) VALUES ($1, $2, $3, $4) " +
		"ON CONFLICT (storage_id, product_id, client_id) " +
		"DO UPDATE SET amount = reservations.amount + EXCLUDED.amount " +
		"WHERE reservations.storage_id = EXCLUDED.storage_id AND reservations.product_id = EXCLUDED.product_id AND reservations.client_id = EXCLUDED.client_id"
	decreaseProductStmt        = "UPDATE products_to_storages SET amount = amount - $3 WHERE storage_id = $1 AND product_id = $2"
	releaseProductStmt         = "UPDATE reservations SET amount = amount - $4 WHERE storage_id = $1 AND product_id = $2 AND client_id = $3 RETURNING amount"
	deleteReservation          = "DELETE FROM reservations WHERE storage_id = $1 AND product_id = $2 AND client_id = $3"
	increaseProductStmt        = "UPDATE products_to_storages SET amount = amount + $3 WHERE storage_id = $1 AND product_id = $2"
	getStorageAvailabilityStmt = "SELECT is_available FROM storages WHERE id = $1"
	getStorageProducts         = "SELECT product_id, amount FROM products_to_storages WHERE storage_id = $1"
)

func (s *Storage) ReserveProducts(ctx context.Context, products []models.StorageProductAmount) error {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, product := range products {
		_, err = tx.Exec(ctx, reserveProductStmt, product.StorageId, product.ProductId, product.ClientId, product.Amount)
		if err != nil {
			return fmt.Errorf("reserve product failed: %w", err)
		}
		_, err = tx.Exec(ctx, decreaseProductStmt, product.StorageId, product.ProductId, product.Amount)
		if err != nil {
			return fmt.Errorf("decrease product failed: %w", err)
		}
	}

	err = tx.Commit(ctx)
	return err
}

func (s *Storage) ReleaseProducts(ctx context.Context, products []models.StorageProductAmount) error {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var resAmount int32
	for _, product := range products {
		row := tx.QueryRow(ctx, releaseProductStmt, product.StorageId, product.ProductId, product.ClientId, product.Amount)
		if err = row.Scan(&resAmount); err != nil {
			return fmt.Errorf("scan product amount failed: %w", err)
		}

		if resAmount == 0 {
			_, err = tx.Exec(ctx, deleteReservation, product.StorageId, product.ProductId, product.ClientId)
			if err != nil {
				return fmt.Errorf("delete reservation failed: %w", err)
			}
		}

		_, err = tx.Exec(ctx, increaseProductStmt, product.StorageId, product.ProductId, product.Amount)
		if err != nil {
			return fmt.Errorf("increase product failed: %w", err)
		}
	}

	err = tx.Commit(ctx)
	return err
}

func (s *Storage) GetProductsCount(ctx context.Context, storageId int32) (models.ProductsByStorage, error) {

	var (
		res   models.ProductsByStorage
		batch = &pgx.Batch{}
	)

	batch.Queue(getStorageAvailabilityStmt, storageId)
	batch.Queue(getStorageProducts, storageId)
	resultsFromDb := s.pool.SendBatch(ctx, batch)
	defer resultsFromDb.Close()

	if err := resultsFromDb.QueryRow().Scan(&res.IsStorageAvailable); err != nil {
		return res, fmt.Errorf("scan storage avaliability failed: %w", err)
	}

	rows, err := resultsFromDb.Query()
	if err != nil {
		return res, err
	}

	res.Products, err = pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProductAmount])

	return res, err
}
