// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ContractRulesRuleViewModel contract rules rule view model
// swagger:model ContractRulesRuleViewModel
type ContractRulesRuleViewModel struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this contract rules rule view model
func (m *ContractRulesRuleViewModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContractRulesRuleViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractRulesRuleViewModel) UnmarshalBinary(b []byte) error {
	var res ContractRulesRuleViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}