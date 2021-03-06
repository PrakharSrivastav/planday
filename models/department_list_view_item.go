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

// DepartmentListViewItem department list view item
// swagger:model DepartmentListViewItem
type DepartmentListViewItem struct {

	// city
	City string `json:"city,omitempty"`

	// country id
	CountryID int64 `json:"country_id,omitempty"`

	// county
	County string `json:"county,omitempty"`

	// employee count
	EmployeeCount int32 `json:"employee_count,omitempty"`

	// employee ids
	// Unique: true
	EmployeeIds []int64 `json:"employee_ids"`

	// employees
	// Read Only: true
	// Unique: true
	Employees []int64 `json:"employees"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// street
	Street string `json:"street,omitempty"`

	// street2
	Street2 string `json:"street2,omitempty"`

	// zip
	Zip string `json:"zip,omitempty"`
}

// Validate validates this department list view item
func (m *DepartmentListViewItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmployeeIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepartmentListViewItem) validateEmployeeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployeeIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("employee_ids", "body", m.EmployeeIds); err != nil {
		return err
	}

	return nil
}

func (m *DepartmentListViewItem) validateEmployees(formats strfmt.Registry) error {

	if swag.IsZero(m.Employees) { // not required
		return nil
	}

	if err := validate.UniqueItems("employees", "body", m.Employees); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DepartmentListViewItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepartmentListViewItem) UnmarshalBinary(b []byte) error {
	var res DepartmentListViewItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
