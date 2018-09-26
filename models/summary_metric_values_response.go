// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SummaryMetricValuesResponse SummaryMetricValuesResponse
// swagger:model SummaryMetricValuesResponse
type SummaryMetricValuesResponse struct {

	// formatted
	Formatted string `json:"formatted,omitempty"`

	// raw
	Raw float64 `json:"raw,omitempty"`
}

// Validate validates this summary metric values response
func (m *SummaryMetricValuesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SummaryMetricValuesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryMetricValuesResponse) UnmarshalBinary(b []byte) error {
	var res SummaryMetricValuesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
