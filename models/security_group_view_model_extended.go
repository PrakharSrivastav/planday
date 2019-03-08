// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SecurityGroupViewModelExtended security group view model extended
// swagger:model SecurityGroupViewModelExtended
type SecurityGroupViewModelExtended struct {

	// item
	Item *SecurityGroupViewModel `json:"item,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`
}

// Validate validates this security group view model extended
func (m *SecurityGroupViewModelExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupViewModelExtended) validateItem(formats strfmt.Registry) error {

	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {
		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupViewModelExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupViewModelExtended) UnmarshalBinary(b []byte) error {
	var res SecurityGroupViewModelExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
