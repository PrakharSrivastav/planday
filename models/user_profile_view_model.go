// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserProfileViewModel user profile view model
// swagger:model UserProfileViewModel
type UserProfileViewModel struct {

	// employee
	Employee *UserProfileEmployeeViewModel `json:"employee,omitempty"`

	// personal details
	PersonalDetails *UserProfileDetailsViewModel `json:"personal_details,omitempty"`

	// portal
	Portal *UserProfilePortalViewModel `json:"portal,omitempty"`

	// preferences
	Preferences *UserProfilePreferencesViewModel `json:"preferences,omitempty"`

	// security
	Security *UserProfileSecurityViewModel `json:"security,omitempty"`

	// settings
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// Validate validates this user profile view model
func (m *UserProfileViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmployee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserProfileViewModel) validateEmployee(formats strfmt.Registry) error {

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

func (m *UserProfileViewModel) validatePersonalDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.PersonalDetails) { // not required
		return nil
	}

	if m.PersonalDetails != nil {
		if err := m.PersonalDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("personal_details")
			}
			return err
		}
	}

	return nil
}

func (m *UserProfileViewModel) validatePortal(formats strfmt.Registry) error {

	if swag.IsZero(m.Portal) { // not required
		return nil
	}

	if m.Portal != nil {
		if err := m.Portal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("portal")
			}
			return err
		}
	}

	return nil
}

func (m *UserProfileViewModel) validatePreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	if m.Preferences != nil {
		if err := m.Preferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preferences")
			}
			return err
		}
	}

	return nil
}

func (m *UserProfileViewModel) validateSecurity(formats strfmt.Registry) error {

	if swag.IsZero(m.Security) { // not required
		return nil
	}

	if m.Security != nil {
		if err := m.Security.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserProfileViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserProfileViewModel) UnmarshalBinary(b []byte) error {
	var res UserProfileViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
