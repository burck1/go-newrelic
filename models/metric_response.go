// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MetricResponse MetricResponse
// swagger:model MetricResponse
type MetricResponse struct {

	// name
	Name string `json:"name,omitempty"`

	// timeslices
	Timeslices []*TimesliceResponse `json:"timeslices"`
}

// Validate validates this metric response
func (m *MetricResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeslices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricResponse) validateTimeslices(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeslices) { // not required
		return nil
	}

	for i := 0; i < len(m.Timeslices); i++ {
		if swag.IsZero(m.Timeslices[i]) { // not required
			continue
		}

		if m.Timeslices[i] != nil {
			if err := m.Timeslices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeslices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricResponse) UnmarshalBinary(b []byte) error {
	var res MetricResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
