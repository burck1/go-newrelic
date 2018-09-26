// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2ComponentsJSONByIDGetReader is a Reader for the V2ComponentsJSONByIDGet structure.
type V2ComponentsJSONByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ComponentsJSONByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2ComponentsJSONByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2ComponentsJSONByIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2ComponentsJSONByIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewV2ComponentsJSONByIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2ComponentsJSONByIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2ComponentsJSONByIDGetOK creates a V2ComponentsJSONByIDGetOK with default headers values
func NewV2ComponentsJSONByIDGetOK() *V2ComponentsJSONByIDGetOK {
	return &V2ComponentsJSONByIDGetOK{}
}

/*V2ComponentsJSONByIDGetOK handles this case with default header values.

OK
*/
type V2ComponentsJSONByIDGetOK struct {
	Payload *models.ComponentResponse
}

func (o *V2ComponentsJSONByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/components/{component_id}.json][%d] v2ComponentsJsonByIdGetOK  %+v", 200, o.Payload)
}

func (o *V2ComponentsJSONByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ComponentsJSONByIDGetUnauthorized creates a V2ComponentsJSONByIDGetUnauthorized with default headers values
func NewV2ComponentsJSONByIDGetUnauthorized() *V2ComponentsJSONByIDGetUnauthorized {
	return &V2ComponentsJSONByIDGetUnauthorized{}
}

/*V2ComponentsJSONByIDGetUnauthorized handles this case with default header values.

Invalid API key
*/
type V2ComponentsJSONByIDGetUnauthorized struct {
	Payload interface{}
}

func (o *V2ComponentsJSONByIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/components/{component_id}.json][%d] v2ComponentsJsonByIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ComponentsJSONByIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ComponentsJSONByIDGetForbidden creates a V2ComponentsJSONByIDGetForbidden with default headers values
func NewV2ComponentsJSONByIDGetForbidden() *V2ComponentsJSONByIDGetForbidden {
	return &V2ComponentsJSONByIDGetForbidden{}
}

/*V2ComponentsJSONByIDGetForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2ComponentsJSONByIDGetForbidden struct {
	Payload interface{}
}

func (o *V2ComponentsJSONByIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/components/{component_id}.json][%d] v2ComponentsJsonByIdGetForbidden  %+v", 403, o.Payload)
}

func (o *V2ComponentsJSONByIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ComponentsJSONByIDGetNotFound creates a V2ComponentsJSONByIDGetNotFound with default headers values
func NewV2ComponentsJSONByIDGetNotFound() *V2ComponentsJSONByIDGetNotFound {
	return &V2ComponentsJSONByIDGetNotFound{}
}

/*V2ComponentsJSONByIDGetNotFound handles this case with default header values.

No Plugin found with given ID
*/
type V2ComponentsJSONByIDGetNotFound struct {
	Payload interface{}
}

func (o *V2ComponentsJSONByIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/components/{component_id}.json][%d] v2ComponentsJsonByIdGetNotFound  %+v", 404, o.Payload)
}

func (o *V2ComponentsJSONByIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ComponentsJSONByIDGetDefault creates a V2ComponentsJSONByIDGetDefault with default headers values
func NewV2ComponentsJSONByIDGetDefault(code int) *V2ComponentsJSONByIDGetDefault {
	return &V2ComponentsJSONByIDGetDefault{
		_statusCode: code,
	}
}

/*V2ComponentsJSONByIDGetDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2ComponentsJSONByIDGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 components Json by Id get default response
func (o *V2ComponentsJSONByIDGetDefault) Code() int {
	return o._statusCode
}

func (o *V2ComponentsJSONByIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /v2/components/{component_id}.json][%d] V2ComponentsJsonByIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *V2ComponentsJSONByIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
