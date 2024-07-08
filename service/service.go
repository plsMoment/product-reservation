package service

import (
	"context"
	"product-storage/models"
	"product-storage/storage"
)

type StorageManager interface {
	ReserveProductsOnStorages(ctx context.Context, products []models.StorageProductAmount) error
	ReleaseProductsOnStorages(ctx context.Context, products []models.StorageProductAmount) error
	GetProductsCountByStorage(ctx context.Context, storageId int32) (models.ProductsByStorage, error)
}

type Service struct {
	s storage.Storage
}

func New(s storage.Storage) *Service {
	return &Service{s}
}

func (srv *Service) ReserveProductsOnStorages(ctx context.Context, products []models.StorageProductAmount) error {
	return srv.s.ReserveProducts(ctx, products)
}

func (srv *Service) ReleaseProductsOnStorages(ctx context.Context, products []models.StorageProductAmount) error {
	return srv.s.ReleaseProducts(ctx, products)
}

func (srv *Service) GetProductsCountByStorage(ctx context.Context, storageId int32) (models.ProductsByStorage, error) {
	return srv.s.GetProductsCount(ctx, storageId)
}
