package models

type StorageProductAmount struct {
	ProductId int32 `json:"product_id"`
	StorageId int32 `json:"storage_id"`
	ClientId  int32 `json:"client_id"`
	Amount    int32 `json:"amount"`
}

type ProductAmount struct {
	ProductId int32 `json:"product_id"`
	Amount    int32 `json:"amount"`
}

type ProductsByStorage struct {
	Products           []ProductAmount `json:"products"`
	IsStorageAvailable bool            `json:"is_storage_available"`
}
