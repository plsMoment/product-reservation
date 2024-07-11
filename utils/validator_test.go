package utils

import (
	"errors"
	"product-storage/models"
	"testing"
)

var testCases = []struct {
	name     string
	input    []models.StorageProductAmount
	expected error
}{
	{"without err", []models.StorageProductAmount{{1, 1, 1, 1}}, nil},
	{"wrongProductId", []models.StorageProductAmount{{-1, 2, 2, 2}}, wrongProductIdValue},
	{"wrongStorageId", []models.StorageProductAmount{{1, 0, 2, 2}}, wrongStorageIdValue},
	{"wrongClientId", []models.StorageProductAmount{{1, 1, -2, 2}}, wrongClientIdValue},
	{"wrongAmount", []models.StorageProductAmount{{1, 1, 2, -2}}, wrongAmountValue},
	{"wrongProductId with other errors", []models.StorageProductAmount{{-1, -1, 2, 2}}, wrongProductIdValue},
	{"wrongStorageId with other errors", []models.StorageProductAmount{{1, -1, 0, -2}}, wrongStorageIdValue},
	{"wrongAmount in slice with len > 1", []models.StorageProductAmount{{1, 1, 1, 1}, {1, 2, 2, -2}}, wrongAmountValue},
}

func TestValidate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ans := Validate(tc.input)
			if !errors.Is(ans, tc.expected) {
				t.Errorf("Expected [%v] but got [%v]", tc.expected, ans)
			}
		})
	}
}
