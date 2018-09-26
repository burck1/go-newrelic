// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2UsersJSONByIDGetReader is a Reader for the V2UsersJSONByIDGet structure.
type V2UsersJSONByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UsersJSONByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2UsersJSONByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2UsersJSONByIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2UsersJSONByIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewV2UsersJSONByIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2UsersJSONByIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2UsersJSONByIDGetOK creates a V2UsersJSONByIDGetOK with default headers values
func NewV2UsersJSONByIDGetOK() *V2UsersJSONByIDGetOK {
	return &V2UsersJSONByIDGetOK{}
}

/*V2UsersJSONByIDGetOK handles this case with default header values.

OK
*/
type V2UsersJSONByIDGetOK struct {
	Payload *models.UserResponse
}

func (o *V2UsersJSONByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/users/{user_id}.json][%d] v2UsersJsonByIdGetOK  %+v", 200, o.Payload)
}

func (o *V2UsersJSONByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UsersJSONByIDGetUnauthorized creates a V2UsersJSONByIDGetUnauthorized with default headers values
func NewV2UsersJSONByIDGetUnauthorized() *V2UsersJSONByIDGetUnauthorized {
	return &V2UsersJSONByIDGetUnauthorized{}
}

/*V2UsersJSONByIDGetUnauthorized handles this case with default header values.

Invalid API key
*/
type V2UsersJSONByIDGetUnauthorized struct {
	Payload interface{}
}

func (o *V2UsersJSONByIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/users/{user_id}.json][%d] v2UsersJsonByIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UsersJSONByIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UsersJSONByIDGetForbidden creates a V2UsersJSONByIDGetForbidden with default headers values
func NewV2UsersJSONByIDGetForbidden() *V2UsersJSONByIDGetForbidden {
	return &V2UsersJSONByIDGetForbidden{}
}

/*V2UsersJSONByIDGetForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2UsersJSONByIDGetForbidden struct {
	Payload interface{}
}

func (o *V2UsersJSONByIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/users/{user_id}.json][%d] v2UsersJsonByIdGetForbidden  %+v", 403, o.Payload)
}

func (o *V2UsersJSONByIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UsersJSONByIDGetNotFound creates a V2UsersJSONByIDGetNotFound with default headers values
func NewV2UsersJSONByIDGetNotFound() *V2UsersJSONByIDGetNotFound {
	return &V2UsersJSONByIDGetNotFound{}
}

/*V2UsersJSONByIDGetNotFound handles this case with default header values.

No user found with the given ID
*/
type V2UsersJSONByIDGetNotFound struct {
	Payload interface{}
}

func (o *V2UsersJSONByIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/users/{user_id}.json][%d] v2UsersJsonByIdGetNotFound  %+v", 404, o.Payload)
}

func (o *V2UsersJSONByIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UsersJSONByIDGetDefault creates a V2UsersJSONByIDGetDefault with default headers values
func NewV2UsersJSONByIDGetDefault(code int) *V2UsersJSONByIDGetDefault {
	return &V2UsersJSONByIDGetDefault{
		_statusCode: code,
	}
}

/*V2UsersJSONByIDGetDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2UsersJSONByIDGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 users Json by Id get default response
func (o *V2UsersJSONByIDGetDefault) Code() int {
	return o._statusCode
}

func (o *V2UsersJSONByIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /v2/users/{user_id}.json][%d] V2UsersJsonByIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *V2UsersJSONByIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
