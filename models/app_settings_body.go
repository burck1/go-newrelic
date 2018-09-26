// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AppSettingsBody AppSettingsBody
// swagger:model AppSettingsBody
type AppSettingsBody struct {

	// app apdex threshold
	AppApdexThreshold float64 `json:"app_apdex_threshold,omitempty"`

	// enable real user monitoring
	EnableRealUserMonitoring bool `json:"enable_real_user_monitoring,omitempty"`

	// end user apdex threshold
	EndUserApdexThreshold float64 `json:"end_user_apdex_threshold,omitempty"`
}

// Validate validates this app settings body
func (m *AppSettingsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppSettingsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppSettingsBody) UnmarshalBinary(b []byte) error {
	var res AppSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
