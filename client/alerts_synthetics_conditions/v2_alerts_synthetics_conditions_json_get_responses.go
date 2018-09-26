// Code generated by go-swagger; DO NOT EDIT.

package alerts_synthetics_conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2AlertsSyntheticsConditionsJSONGetReader is a Reader for the V2AlertsSyntheticsConditionsJSONGet structure.
type V2AlertsSyntheticsConditionsJSONGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2AlertsSyntheticsConditionsJSONGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2AlertsSyntheticsConditionsJSONGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2AlertsSyntheticsConditionsJSONGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2AlertsSyntheticsConditionsJSONGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewV2AlertsSyntheticsConditionsJSONGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2AlertsSyntheticsConditionsJSONGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2AlertsSyntheticsConditionsJSONGetOK creates a V2AlertsSyntheticsConditionsJSONGetOK with default headers values
func NewV2AlertsSyntheticsConditionsJSONGetOK() *V2AlertsSyntheticsConditionsJSONGetOK {
	return &V2AlertsSyntheticsConditionsJSONGetOK{}
}

/*V2AlertsSyntheticsConditionsJSONGetOK handles this case with default header values.

OK
*/
type V2AlertsSyntheticsConditionsJSONGetOK struct {
	Payload *models.SyntheticsConditionResponse
}

func (o *V2AlertsSyntheticsConditionsJSONGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/alerts_synthetics_conditions.json][%d] v2AlertsSyntheticsConditionsJsonGetOK  %+v", 200, o.Payload)
}

func (o *V2AlertsSyntheticsConditionsJSONGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyntheticsConditionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsSyntheticsConditionsJSONGetUnauthorized creates a V2AlertsSyntheticsConditionsJSONGetUnauthorized with default headers values
func NewV2AlertsSyntheticsConditionsJSONGetUnauthorized() *V2AlertsSyntheticsConditionsJSONGetUnauthorized {
	return &V2AlertsSyntheticsConditionsJSONGetUnauthorized{}
}

/*V2AlertsSyntheticsConditionsJSONGetUnauthorized handles this case with default header values.

Invalid API key
*/
type V2AlertsSyntheticsConditionsJSONGetUnauthorized struct {
	Payload interface{}
}

func (o *V2AlertsSyntheticsConditionsJSONGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/alerts_synthetics_conditions.json][%d] v2AlertsSyntheticsConditionsJsonGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V2AlertsSyntheticsConditionsJSONGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsSyntheticsConditionsJSONGetForbidden creates a V2AlertsSyntheticsConditionsJSONGetForbidden with default headers values
func NewV2AlertsSyntheticsConditionsJSONGetForbidden() *V2AlertsSyntheticsConditionsJSONGetForbidden {
	return &V2AlertsSyntheticsConditionsJSONGetForbidden{}
}

/*V2AlertsSyntheticsConditionsJSONGetForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2AlertsSyntheticsConditionsJSONGetForbidden struct {
	Payload interface{}
}

func (o *V2AlertsSyntheticsConditionsJSONGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/alerts_synthetics_conditions.json][%d] v2AlertsSyntheticsConditionsJsonGetForbidden  %+v", 403, o.Payload)
}

func (o *V2AlertsSyntheticsConditionsJSONGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsSyntheticsConditionsJSONGetNotFound creates a V2AlertsSyntheticsConditionsJSONGetNotFound with default headers values
func NewV2AlertsSyntheticsConditionsJSONGetNotFound() *V2AlertsSyntheticsConditionsJSONGetNotFound {
	return &V2AlertsSyntheticsConditionsJSONGetNotFound{}
}

/*V2AlertsSyntheticsConditionsJSONGetNotFound handles this case with default header values.

No Alerts policy was found for the given ID
*/
type V2AlertsSyntheticsConditionsJSONGetNotFound struct {
	Payload interface{}
}

func (o *V2AlertsSyntheticsConditionsJSONGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/alerts_synthetics_conditions.json][%d] v2AlertsSyntheticsConditionsJsonGetNotFound  %+v", 404, o.Payload)
}

func (o *V2AlertsSyntheticsConditionsJSONGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsSyntheticsConditionsJSONGetDefault creates a V2AlertsSyntheticsConditionsJSONGetDefault with default headers values
func NewV2AlertsSyntheticsConditionsJSONGetDefault(code int) *V2AlertsSyntheticsConditionsJSONGetDefault {
	return &V2AlertsSyntheticsConditionsJSONGetDefault{
		_statusCode: code,
	}
}

/*V2AlertsSyntheticsConditionsJSONGetDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2AlertsSyntheticsConditionsJSONGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 alerts synthetics conditions Json get default response
func (o *V2AlertsSyntheticsConditionsJSONGetDefault) Code() int {
	return o._statusCode
}

func (o *V2AlertsSyntheticsConditionsJSONGetDefault) Error() string {
	return fmt.Sprintf("[GET /v2/alerts_synthetics_conditions.json][%d] V2AlertsSyntheticsConditionsJsonGet default  %+v", o._statusCode, o.Payload)
}

func (o *V2AlertsSyntheticsConditionsJSONGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
