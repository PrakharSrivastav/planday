// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EmployeePriorityViewModel employee priority view model
// swagger:model EmployeePriorityViewModel
type EmployeePriorityViewModel struct {

	// employee group id
	EmployeeGroupID int64 `json:"employee_group_id,omitempty"`

	// employee id
	EmployeeID int64 `json:"employee_id,omitempty"`

	// end date
	EndDate string `json:"end_date,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// portal id
	PortalID int64 `json:"portal_id,omitempty"`

	// start date
	StartDate string `json:"start_date,omitempty"`

	// value
	Value int32 `json:"value,omitempty"`
}

// Validate validates this employee priority view model
func (m *EmployeePriorityViewModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmployeePriorityViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmployeePriorityViewModel) UnmarshalBinary(b []byte) error {
	var res EmployeePriorityViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
