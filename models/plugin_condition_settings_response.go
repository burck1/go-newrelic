// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PluginConditionSettingsResponse PluginConditionSettingsResponse
// swagger:model PluginConditionSettingsResponse
type PluginConditionSettingsResponse struct {

	// guid
	GUID string `json:"guid,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this plugin condition settings response
func (m *PluginConditionSettingsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConditionSettingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConditionSettingsResponse) UnmarshalBinary(b []byte) error {
	var res PluginConditionSettingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
