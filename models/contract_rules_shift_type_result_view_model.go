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

// ContractRulesShiftTypeResultViewModel contract rules shift type result view model
// swagger:model ContractRulesShiftTypeResultViewModel
type ContractRulesShiftTypeResultViewModel struct {

	// all selected
	AllSelected bool `json:"all_selected,omitempty"`

	// contract rule id
	ContractRuleID int64 `json:"contract_rule_id,omitempty"`

	// shift types
	ShiftTypes []*ContractRulesShiftTypeViewModel `json:"shift_types"`
}

// Validate validates this contract rules shift type result view model
func (m *ContractRulesShiftTypeResultViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShiftTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContractRulesShiftTypeResultViewModel) validateShiftTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ShiftTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ShiftTypes); i++ {
		if swag.IsZero(m.ShiftTypes[i]) { // not required
			continue
		}

		if m.ShiftTypes[i] != nil {
			if err := m.ShiftTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shift_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContractRulesShiftTypeResultViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractRulesShiftTypeResultViewModel) UnmarshalBinary(b []byte) error {
	var res ContractRulesShiftTypeResultViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
