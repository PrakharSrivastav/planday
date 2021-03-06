// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GdprStatusResult gdpr status result
// swagger:model GdprStatusResult
type GdprStatusResult struct {

	// is gdpr enabled
	IsGdprEnabled bool `json:"is_gdpr_enabled,omitempty"`
}

// Validate validates this gdpr status result
func (m *GdprStatusResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GdprStatusResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GdprStatusResult) UnmarshalBinary(b []byte) error {
	var res GdprStatusResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
