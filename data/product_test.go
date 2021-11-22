package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidProductDoesNotReturnError(t *testing.T) {
	p := &Product{
		Name:  "Milk",
		Price: 1,
		SKU:   "abs-fgb-kfk",
	}
	validation := NewValidation()
	errs := validation.Validate(p)
	assert.Len(t, errs, 0)
}

func TestProductMissingNameReturnsError(t *testing.T) {
	p := &Product{
		Price: 1,
		SKU:   "abs-fgb-kfk",
	}
	validation := NewValidation()
	errs := validation.Validate(p)

	assert.Len(t, errs, 1)

}
