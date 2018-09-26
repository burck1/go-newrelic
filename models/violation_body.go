// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ViolationBody ViolationBody
// swagger:model ViolationBody
type ViolationBody struct {

	// closed at
	ClosedAt int32 `json:"closed_at,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// opened at
	OpenedAt int32 `json:"opened_at,omitempty"`
}

// Validate validates this violation body
func (m *ViolationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViolationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViolationBody) UnmarshalBinary(b []byte) error {
	var res ViolationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
