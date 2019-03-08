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

// SecurityAction security action
// swagger:model SecurityAction
type SecurityAction struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// date created
	// Format: date-time
	DateCreated strfmt.DateTime `json:"date_created,omitempty"`

	// date modified
	// Format: date-time
	DateModified strfmt.DateTime `json:"date_modified,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// last modified
	LastModified string `json:"last_modified,omitempty"`
}

// Validate validates this security action
func (m *SecurityAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAction) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("date_created", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAction) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("date_modified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAction) UnmarshalBinary(b []byte) error {
	var res SecurityAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}