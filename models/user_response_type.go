// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserResponseType UserResponseType
// swagger:model UserResponseType
type UserResponseType struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// role
	Role string `json:"role,omitempty"`
}

// Validate validates this user response type
func (m *UserResponseType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserResponseType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserResponseType) UnmarshalBinary(b []byte) error {
	var res UserResponseType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
