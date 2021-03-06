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

// EnabledDepartmentForSecurityGroupMembership enabled department for security group membership
// swagger:model EnabledDepartmentForSecurityGroupMembership
type EnabledDepartmentForSecurityGroupMembership struct {

	// created by employee id
	// Read Only: true
	CreatedByEmployeeID int64 `json:"created_by_employee_id,omitempty"`

	// date created
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"date_created,omitempty"`

	// department id
	// Read Only: true
	DepartmentID int64 `json:"department_id,omitempty"`

	// only employees from primary department
	// Read Only: true
	OnlyEmployeesFromPrimaryDepartment *bool `json:"only_employees_from_primary_department,omitempty"`

	// security group id
	// Read Only: true
	SecurityGroupID int64 `json:"security_group_id,omitempty"`

	// security group membership
	// Read Only: true
	SecurityGroupMembership *SecurityGroupMembership `json:"security_group_membership,omitempty"`

	// user id
	// Read Only: true
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this enabled department for security group membership
func (m *EnabledDepartmentForSecurityGroupMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupMembership(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnabledDepartmentForSecurityGroupMembership) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("date_created", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EnabledDepartmentForSecurityGroupMembership) validateSecurityGroupMembership(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroupMembership) { // not required
		return nil
	}

	if m.SecurityGroupMembership != nil {
		if err := m.SecurityGroupMembership.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_group_membership")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnabledDepartmentForSecurityGroupMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnabledDepartmentForSecurityGroupMembership) UnmarshalBinary(b []byte) error {
	var res EnabledDepartmentForSecurityGroupMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
