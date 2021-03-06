// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// QueryResultSorting query result sorting
// swagger:model QueryResultSorting
type QueryResultSorting struct {

	// order
	Order string `json:"order,omitempty"`

	// order by
	OrderBy string `json:"order_by,omitempty"`
}

// Validate validates this query result sorting
func (m *QueryResultSorting) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryResultSorting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResultSorting) UnmarshalBinary(b []byte) error {
	var res QueryResultSorting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
