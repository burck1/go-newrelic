// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EndUserSummaryDataResponse EndUserSummaryDataResponse
// swagger:model EndUserSummaryDataResponse
type EndUserSummaryDataResponse struct {

	// apdex score
	ApdexScore float64 `json:"apdex_score,omitempty"`

	// response time
	ResponseTime float64 `json:"response_time,omitempty"`

	// throughput
	Throughput float64 `json:"throughput,omitempty"`
}

// Validate validates this end user summary data response
func (m *EndUserSummaryDataResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndUserSummaryDataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndUserSummaryDataResponse) UnmarshalBinary(b []byte) error {
	var res EndUserSummaryDataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
