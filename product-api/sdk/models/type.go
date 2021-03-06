// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Type Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of type before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.
//
// Type values are comparable, such as with the == operator,
// so they can be used as map keys.
// Two Type values are equal if they represent identical types.
//
// swagger:model Type
type Type struct {

	// Align returns the alignment in bytes of a value of
	// this type when allocated in memory.
	Align int64 `json:"Align,omitempty"`

	// Bits returns the size of the type in bits.
	// It panics if the type's Kind is not one of the
	// sized or unsized Int, Uint, Float, or Complex kinds.
	Bits int64 `json:"Bits,omitempty"`

	// chan dir
	ChanDir ChanDir `json:"ChanDir,omitempty"`

	// Comparable reports whether values of this type are comparable.
	// Even if Comparable returns true, the comparison may still panic.
	// For example, values of interface type are comparable,
	// but the comparison will panic if their dynamic type is not comparable.
	Comparable bool `json:"Comparable,omitempty"`

	// elem
	Elem *Type `json:"Elem,omitempty"`

	// FieldAlign returns the alignment in bytes of a value of
	// this type when used as a field in a struct.
	FieldAlign int64 `json:"FieldAlign,omitempty"`

	// IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	// implicit actual type []T.
	//
	// For concreteness, if t represents func(x int, y ... float64), then
	//
	// t.NumIn() == 2
	// t.In(0) is the reflect.Type for "int"
	// t.In(1) is the reflect.Type for "[]float64"
	// t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic bool `json:"IsVariadic,omitempty"`

	// key
	Key *Type `json:"Key,omitempty"`

	// kind
	Kind Kind `json:"Kind,omitempty"`

	// Len returns an array type's length.
	// It panics if the type's Kind is not Array.
	Len int64 `json:"Len,omitempty"`

	// Name returns the type's name within its package for a defined type.
	// For other (non-defined) types it returns the empty string.
	Name string `json:"Name,omitempty"`

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField int64 `json:"NumField,omitempty"`

	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	NumIn int64 `json:"NumIn,omitempty"`

	// NumMethod returns the number of methods accessible using Method.
	//
	// Note that NumMethod counts unexported methods only for interface types.
	NumMethod int64 `json:"NumMethod,omitempty"`

	// NumOut returns a function type's output parameter count.
	// It panics if the type's Kind is not Func.
	NumOut int64 `json:"NumOut,omitempty"`

	// PkgPath returns a defined type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or not defined (*T, struct{},
	// []int, or A where A is an alias for a non-defined type), the package path
	// will be the empty string.
	PkgPath string `json:"PkgPath,omitempty"`

	// Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
	Size uint64 `json:"Size,omitempty"`

	// String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types. To test for type identity,
	// compare the Types directly.
	String string `json:"String,omitempty"`
}

// Validate validates this type
func (m *Type) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Type) validateChanDir(formats strfmt.Registry) error {
	if swag.IsZero(m.ChanDir) { // not required
		return nil
	}

	if err := m.ChanDir.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ChanDir")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ChanDir")
		}
		return err
	}

	return nil
}

func (m *Type) validateElem(formats strfmt.Registry) error {
	if swag.IsZero(m.Elem) { // not required
		return nil
	}

	if m.Elem != nil {
		if err := m.Elem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Elem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Elem")
			}
			return err
		}
	}

	return nil
}

func (m *Type) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if m.Key != nil {
		if err := m.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Key")
			}
			return err
		}
	}

	return nil
}

func (m *Type) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := m.Kind.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Kind")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Kind")
		}
		return err
	}

	return nil
}

// ContextValidate validate this type based on the context it is used
func (m *Type) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChanDir(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Type) contextValidateChanDir(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ChanDir.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ChanDir")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ChanDir")
		}
		return err
	}

	return nil
}

func (m *Type) contextValidateElem(ctx context.Context, formats strfmt.Registry) error {

	if m.Elem != nil {
		if err := m.Elem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Elem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Elem")
			}
			return err
		}
	}

	return nil
}

func (m *Type) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if m.Key != nil {
		if err := m.Key.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Key")
			}
			return err
		}
	}

	return nil
}

func (m *Type) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Kind.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Kind")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Kind")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Type) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Type) UnmarshalBinary(b []byte) error {
	var res Type
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
