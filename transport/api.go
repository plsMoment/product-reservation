package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"product-storage/models"
	"product-storage/service"
	"product-storage/utils"
	"strconv"
)

type Handler struct {
	srv service.StorageManager
}

func NewHandler(srv service.StorageManager) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) ReserveProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.StorageProductAmount
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		log.Println("decoding request json failed in ReserveProducts", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = utils.Validate(products); err != nil {
		log.Println("Validation products from request failed: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.srv.ReserveProductsOnStorages(context.TODO(), products)
	if err != nil {
		log.Println("ReserveProductsOnStorages failed: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("ReserveProducts successful")
}

func (h *Handler) ReleaseProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.StorageProductAmount
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		log.Println("decoding request json failed in ReleaseProducts", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = utils.Validate(products); err != nil {
		log.Println("Validation products from request failed: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.srv.ReleaseProductsOnStorages(context.TODO(), products)
	if err != nil {
		log.Println("ReleaseProductsOnStorages failed: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("ReleaseProducts successful")
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	rawStorageId := r.URL.Query().Get("storage_id")
	storageId, err := strconv.Atoi(rawStorageId)
	if err != nil {
		log.Println("conversion to storageId failed: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if storageId <= 0 {
		log.Println("storage id should be greater than zero")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.srv.GetProductsCountByStorage(context.TODO(), int32(storageId))
	if err != nil {
		log.Println("GetProductsCountByStorage failed: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Println("marshaling products in GetProducts failed: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(b); err != nil {
		log.Println("writing response data in GetProducts failed: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("GetProducts successful")
}
