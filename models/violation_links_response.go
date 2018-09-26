// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ViolationLinksResponse ViolationLinksResponse
// swagger:model ViolationLinksResponse
type ViolationLinksResponse struct {

	// condition id
	ConditionID int32 `json:"condition_id,omitempty"`

	// incident id
	IncidentID int32 `json:"incident_id,omitempty"`

	// policy id
	PolicyID int32 `json:"policy_id,omitempty"`
}

// Validate validates this violation links response
func (m *ViolationLinksResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViolationLinksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViolationLinksResponse) UnmarshalBinary(b []byte) error {
	var res ViolationLinksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
