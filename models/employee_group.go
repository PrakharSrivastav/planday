// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmployeeGroup employee group
// swagger:model EmployeeGroup
type EmployeeGroup struct {

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

	// default pay rate
	DefaultPayRate float64 `json:"default_pay_rate,omitempty"`

	// deleted at
	// Read Only: true
	DeletedAt string `json:"deleted_at,omitempty"`

	// deleted by
	// Read Only: true
	DeletedBy int64 `json:"deleted_by,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// disko web id
	DiskoWebID int32 `json:"disko_web_id,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// is deleted
	// Read Only: true
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// last modified
	// Read Only: true
	LastModified string `json:"last_modified,omitempty"`

	// members
	// Read Only: true
	Members []*EmployeeGroupMembership `json:"members"`

	// modified by
	// Read Only: true
	ModifiedBy int64 `json:"modified_by,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// portal id
	PortalID int64 `json:"portal_id,omitempty"`

	// raised events
	// Read Only: true
	RaisedEvents []*IdomainEvent `json:"raised_events"`

	// salary code
	// Required: true
	SalaryCode *string `json:"salary_code"`

	// salary code2
	SalaryCode2 string `json:"salary_code2,omitempty"`
}

// Validate validates this employee group
func (m *EmployeeGroup) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaisedEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalaryCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmployeeGroup) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("date_created", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EmployeeGroup) validateDateDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateDeleted) { // not required
		return nil
	}

	if err := validate.FormatOf("date_deleted", "body", "date-time", m.DateDeleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EmployeeGroup) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("date_modified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EmployeeGroup) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *EmployeeGroup) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmployeeGroup) validateRaisedEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.RaisedEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.RaisedEvents); i++ {
		if swag.IsZero(m.RaisedEvents[i]) { // not required
			continue
		}

		if m.RaisedEvents[i] != nil {
			if err := m.RaisedEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raised_events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmployeeGroup) validateSalaryCode(formats strfmt.Registry) error {

	if err := validate.Required("salary_code", "body", m.SalaryCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmployeeGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmployeeGroup) UnmarshalBinary(b []byte) error {
	var res EmployeeGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
