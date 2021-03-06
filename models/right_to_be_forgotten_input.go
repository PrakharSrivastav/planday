// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RightToBeForgottenInput right to be forgotten input
// swagger:model RightToBeForgottenInput
type RightToBeForgottenInput struct {

	// delete activated
	DeleteActivated bool `json:"delete_activated,omitempty"`

	// num of years
	NumOfYears int32 `json:"num_of_years,omitempty"`
}

// Validate validates this right to be forgotten input
func (m *RightToBeForgottenInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RightToBeForgottenInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RightToBeForgottenInput) UnmarshalBinary(b []byte) error {
	var res RightToBeForgottenInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
