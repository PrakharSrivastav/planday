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

// RoleInputModelExtended role input model extended
// swagger:model RoleInputModelExtended
type RoleInputModelExtended struct {

	// employee group
	EmployeeGroup *EmployeeGroupInputModel `json:"employee_group,omitempty"`

	// pay rates
	PayRates []*PayRateInputModel `json:"pay_rates"`

	// role
	Role *EmployeeGroupInputModel `json:"role,omitempty"`

	// update employee group memberships
	UpdateEmployeeGroupMemberships bool `json:"update_employee_group_memberships,omitempty"`
}

// Validate validates this role input model extended
func (m *RoleInputModelExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmployeeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleInputModelExtended) validateEmployeeGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployeeGroup) { // not required
		return nil
	}

	if m.EmployeeGroup != nil {
		if err := m.EmployeeGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employee_group")
			}
			return err
		}
	}

	return nil
}

func (m *RoleInputModelExtended) validatePayRates(formats strfmt.Registry) error {

	if swag.IsZero(m.PayRates) { // not required
		return nil
	}

	for i := 0; i < len(m.PayRates); i++ {
		if swag.IsZero(m.PayRates[i]) { // not required
			continue
		}

		if m.PayRates[i] != nil {
			if err := m.PayRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pay_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RoleInputModelExtended) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleInputModelExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleInputModelExtended) UnmarshalBinary(b []byte) error {
	var res RoleInputModelExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}