// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmployeeGroupListViewItem employee group list view item
// swagger:model EmployeeGroupListViewItem
type EmployeeGroupListViewItem struct {

	// employee count
	EmployeeCount int32 `json:"employee_count,omitempty"`

	// employee ids
	EmployeeIds []int64 `json:"employee_ids"`

	// employees
	// Unique: true
	Employees []int64 `json:"employees"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pay rate
	PayRate float64 `json:"pay_rate,omitempty"`

	// pay rates
	PayRates float64 `json:"pay_rates,omitempty"`
}

// Validate validates this employee group list view item
func (m *EmployeeGroupListViewItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmployees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmployeeGroupListViewItem) validateEmployees(formats strfmt.Registry) error {

	if swag.IsZero(m.Employees) { // not required
		return nil
	}

	if err := validate.UniqueItems("employees", "body", m.Employees); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmployeeGroupListViewItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmployeeGroupListViewItem) UnmarshalBinary(b []byte) error {
	var res EmployeeGroupListViewItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
