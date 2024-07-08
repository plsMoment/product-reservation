package storage

import (
	"context"
	"product-storage/models"
)

type Storage interface {
	ReserveProducts(ctx context.Context, products []models.StorageProductAmount) error
	ReleaseProducts(ctx context.Context, products []models.StorageProductAmount) error
	GetProductsCount(ctx context.Context, storageId int32) (models.ProductsByStorage, error)
}
