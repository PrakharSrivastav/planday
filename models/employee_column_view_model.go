// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EmployeeColumnViewModel employee column view model
// swagger:model EmployeeColumnViewModel
type EmployeeColumnViewModel struct {

	// custom columns
	CustomColumns []*StandardColumn `json:"custom_columns"`

	// standard columns
	StandardColumns []*StandardColumn `json:"standard_columns"`
}

// Validate validates this employee column view model
func (m *EmployeeColumnViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardColumns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmployeeColumnViewModel) validateCustomColumns(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomColumns) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomColumns); i++ {
		if swag.IsZero(m.CustomColumns[i]) { // not required
			continue
		}

		if m.CustomColumns[i] != nil {
			if err := m.CustomColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmployeeColumnViewModel) validateStandardColumns(formats strfmt.Registry) error {

	if swag.IsZero(m.StandardColumns) { // not required
		return nil
	}

	for i := 0; i < len(m.StandardColumns); i++ {
		if swag.IsZero(m.StandardColumns[i]) { // not required
			continue
		}

		if m.StandardColumns[i] != nil {
			if err := m.StandardColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standard_columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmployeeColumnViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmployeeColumnViewModel) UnmarshalBinary(b []byte) error {
	var res EmployeeColumnViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
