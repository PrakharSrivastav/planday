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

// PayRateViewModelExtended pay rate view model extended
// swagger:model PayRateViewModelExtended
type PayRateViewModelExtended struct {

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// employee group id
	EmployeeGroupID int64 `json:"employee_group_id,omitempty"`

	// employee id
	EmployeeID int64 `json:"employee_id,omitempty"`

	// employee name
	EmployeeName string `json:"employee_name,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// modified by
	ModifiedBy string `json:"modified_by,omitempty"`

	// pay rate
	PayRate float64 `json:"pay_rate,omitempty"`

	// role id
	RoleID int64 `json:"role_id,omitempty"`

	// valid from
	// Format: date-time
	ValidFrom strfmt.DateTime `json:"valid_from,omitempty"`

	// valid to
	// Format: date-time
	ValidTo strfmt.DateTime `json:"valid_to,omitempty"`
}

// Validate validates this pay rate view model extended
func (m *PayRateViewModelExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayRateViewModelExtended) validateValidFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("valid_from", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PayRateViewModelExtended) validateValidTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("valid_to", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PayRateViewModelExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayRateViewModelExtended) UnmarshalBinary(b []byte) error {
	var res PayRateViewModelExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}