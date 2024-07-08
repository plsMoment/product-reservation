package utils

import (
	"errors"
	"product-storage/models"
)

var (
	wrongProductIdValue = errors.New("wrong product id value: product_id must be greater than zero")
	wrongStorageIdValue = errors.New("wrong storage_id value: storage_id must be greater than zero")
	wrongClientIdValue  = errors.New("wrong client_id value: client_id must be greater than zero")
	wrongAmountValue    = errors.New("wrong amount value: amount must be greater than zero")
)

func Validate(products []models.StorageProductAmount) error {
	for _, item := range products {
		if item.ProductId <= 0 {
			return wrongProductIdValue
		} else if item.StorageId <= 0 {
			return wrongStorageIdValue
		} else if item.ClientId <= 0 {
			return wrongClientIdValue
		} else if item.Amount <= 0 {
			return wrongAmountValue
		}
	}
	return nil
}
