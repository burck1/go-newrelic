// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2ApplicationsJSONGetReader is a Reader for the V2ApplicationsJSONGet structure.
type V2ApplicationsJSONGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ApplicationsJSONGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2ApplicationsJSONGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2ApplicationsJSONGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2ApplicationsJSONGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2ApplicationsJSONGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2ApplicationsJSONGetOK creates a V2ApplicationsJSONGetOK with default headers values
func NewV2ApplicationsJSONGetOK() *V2ApplicationsJSONGetOK {
	return &V2ApplicationsJSONGetOK{}
}

/*V2ApplicationsJSONGetOK handles this case with default header values.

OK
*/
type V2ApplicationsJSONGetOK struct {
	Payload *models.ApplicationResponse
}

func (o *V2ApplicationsJSONGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/applications.json][%d] v2ApplicationsJsonGetOK  %+v", 200, o.Payload)
}

func (o *V2ApplicationsJSONGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsJSONGetUnauthorized creates a V2ApplicationsJSONGetUnauthorized with default headers values
func NewV2ApplicationsJSONGetUnauthorized() *V2ApplicationsJSONGetUnauthorized {
	return &V2ApplicationsJSONGetUnauthorized{}
}

/*V2ApplicationsJSONGetUnauthorized handles this case with default header values.

Invalid API key
*/
type V2ApplicationsJSONGetUnauthorized struct {
	Payload interface{}
}

func (o *V2ApplicationsJSONGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/applications.json][%d] v2ApplicationsJsonGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ApplicationsJSONGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsJSONGetForbidden creates a V2ApplicationsJSONGetForbidden with default headers values
func NewV2ApplicationsJSONGetForbidden() *V2ApplicationsJSONGetForbidden {
	return &V2ApplicationsJSONGetForbidden{}
}

/*V2ApplicationsJSONGetForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2ApplicationsJSONGetForbidden struct {
	Payload interface{}
}

func (o *V2ApplicationsJSONGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/applications.json][%d] v2ApplicationsJsonGetForbidden  %+v", 403, o.Payload)
}

func (o *V2ApplicationsJSONGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsJSONGetDefault creates a V2ApplicationsJSONGetDefault with default headers values
func NewV2ApplicationsJSONGetDefault(code int) *V2ApplicationsJSONGetDefault {
	return &V2ApplicationsJSONGetDefault{
		_statusCode: code,
	}
}

/*V2ApplicationsJSONGetDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2ApplicationsJSONGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 applications Json get default response
func (o *V2ApplicationsJSONGetDefault) Code() int {
	return o._statusCode
}

func (o *V2ApplicationsJSONGetDefault) Error() string {
	return fmt.Sprintf("[GET /v2/applications.json][%d] V2ApplicationsJsonGet default  %+v", o._statusCode, o.Payload)
}

func (o *V2ApplicationsJSONGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
