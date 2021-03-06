// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PhoneViewModel phone view model
// swagger:model PhoneViewModel
type PhoneViewModel struct {

	// id
	ID int64 `json:"id,omitempty"`

	// phone prefix
	PhonePrefix string `json:"phone_prefix,omitempty"`

	// short name
	ShortName string `json:"short_name,omitempty"`
}

// Validate validates this phone view model
func (m *PhoneViewModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhoneViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhoneViewModel) UnmarshalBinary(b []byte) error {
	var res PhoneViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
