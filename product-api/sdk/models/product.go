// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Product product
//
// swagger:model Product
type Product struct {

	// description
	Description string `json:"description,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// price
	Price float32 `json:"price,omitempty"`

	// s k u
	SKU string `json:"sku,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this product based on context it is used
func (m *Product) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}