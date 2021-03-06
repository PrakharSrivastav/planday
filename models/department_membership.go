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

// DepartmentMembership department membership
// swagger:model DepartmentMembership
type DepartmentMembership struct {

	// created at
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	// Read Only: true
	CreatedBy int64 `json:"created_by,omitempty"`

	// date created
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"date_created,omitempty"`

	// date deleted
	// Read Only: true
	// Format: date-time
	DateDeleted strfmt.DateTime `json:"date_deleted,omitempty"`

	// date modified
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"date_modified,omitempty"`

	// deactivation date
	// Read Only: true
	DeactivationDate string `json:"deactivation_date,omitempty"`

	// delete origin enum
	DeleteOriginEnum int32 `json:"delete_origin_enum,omitempty"`

	// deleted at
	// Read Only: true
	DeletedAt string `json:"deleted_at,omitempty"`

	// deleted by
	// Read Only: true
	DeletedBy int64 `json:"deleted_by,omitempty"`

	// department
	Department *Department `json:"department,omitempty"`

	// department id
	DepartmentID int64 `json:"department_id,omitempty"`

	// dismissed date
	// Read Only: true
	// Format: date-time
	DismissedDate strfmt.DateTime `json:"dismissed_date,omitempty"`

	// employee
	Employee *Employee `json:"employee,omitempty"`

	// employee id
	EmployeeID int64 `json:"employee_id,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// is deleted
	// Read Only: true
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// is dismissed
	// Read Only: true
	IsDismissed *bool `json:"is_dismissed,omitempty"`

	// last modified
	// Read Only: true
	LastModified string `json:"last_modified,omitempty"`

	// modified by
	// Read Only: true
	ModifiedBy int64 `json:"modified_by,omitempty"`
}

// Validate validates this department membership
func (m *DepartmentMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDismissedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepartmentMembership) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("date_created", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepartmentMembership) validateDateDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateDeleted) { // not required
		return nil
	}

	if err := validate.FormatOf("date_deleted", "body", "date-time", m.DateDeleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepartmentMembership) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("date_modified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepartmentMembership) validateDepartment(formats strfmt.Registry) error {

	if swag.IsZero(m.Department) { // not required
		return nil
	}

	if m.Department != nil {
		if err := m.Department.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("department")
			}
			return err
		}
	}

	return nil
}

func (m *DepartmentMembership) validateDismissedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DismissedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dismissed_date", "body", "date-time", m.DismissedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepartmentMembership) validateEmployee(formats strfmt.Registry) error {

	if swag.IsZero(m.Employee) { // not required
		return nil
	}

	if m.Employee != nil {
		if err := m.Employee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employee")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DepartmentMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepartmentMembership) UnmarshalBinary(b []byte) error {
	var res DepartmentMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
