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

// PortalSettingsViewModel portal settings view model
// swagger:model PortalSettingsViewModel
type PortalSettingsViewModel struct {

	// settings
	Settings []*PortalSetting `json:"settings"`

	// warnings
	Warnings []string `json:"warnings"`
}

// Validate validates this portal settings view model
func (m *PortalSettingsViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortalSettingsViewModel) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	for i := 0; i < len(m.Settings); i++ {
		if swag.IsZero(m.Settings[i]) { // not required
			continue
		}

		if m.Settings[i] != nil {
			if err := m.Settings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortalSettingsViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortalSettingsViewModel) UnmarshalBinary(b []byte) error {
	var res PortalSettingsViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
