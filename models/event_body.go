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

// EventBody EventBody
// swagger:model EventBody
type EventBody struct {

	// closed at
	// Format: date-time
	ClosedAt strfmt.DateTime `json:"closed_at,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// level
	Level string `json:"level,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// payload
	Payload map[string]interface{} `json:"payload,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this event body
func (m *EventBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClosedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventBody) validateClosedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ClosedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("closed_at", "body", "date-time", m.ClosedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventBody) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventBody) UnmarshalBinary(b []byte) error {
	var res EventBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
