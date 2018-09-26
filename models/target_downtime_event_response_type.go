// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TargetDowntimeEventResponseType TargetDowntimeEventResponseType
// swagger:model TargetDowntimeEventResponseType
type TargetDowntimeEventResponseType struct {

	// duration
	Duration int32 `json:"duration,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// ip address
	IPAddress string `json:"ip_address,omitempty"`

	// last failure
	// Format: date-time
	LastFailure strfmt.DateTime `json:"last_failure,omitempty"`

	// number of failures
	NumberOfFailures int32 `json:"number_of_failures,omitempty"`

	// opened at
	// Format: date-time
	OpenedAt strfmt.DateTime `json:"opened_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this target downtime event response type
func (m *TargetDowntimeEventResponseType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetDowntimeEventResponseType) validateLastFailure(formats strfmt.Registry) error {

	if swag.IsZero(m.LastFailure) { // not required
		return nil
	}

	if err := validate.FormatOf("last_failure", "body", "date-time", m.LastFailure.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TargetDowntimeEventResponseType) validateOpenedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.OpenedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("opened_at", "body", "date-time", m.OpenedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TargetDowntimeEventResponseType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetDowntimeEventResponseType) UnmarshalBinary(b []byte) error {
	var res TargetDowntimeEventResponseType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
