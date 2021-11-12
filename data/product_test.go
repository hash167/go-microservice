package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Milk",
		Price: 1,
		SKU:   "abs-fgb-kfk",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
