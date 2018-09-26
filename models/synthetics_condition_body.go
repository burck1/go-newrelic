// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SyntheticsConditionBody SyntheticsConditionBody
// swagger:model SyntheticsConditionBody
type SyntheticsConditionBody struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// monitor id
	MonitorID string `json:"monitor_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// runbook url
	RunbookURL string `json:"runbook_url,omitempty"`
}

// Validate validates this synthetics condition body
func (m *SyntheticsConditionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SyntheticsConditionBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyntheticsConditionBody) UnmarshalBinary(b []byte) error {
	var res SyntheticsConditionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
