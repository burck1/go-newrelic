// Code generated by go-swagger; DO NOT EDIT.

package alerts_nrql_conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2AlertsNrqlConditionsJSONByConditionIDDeleteReader is a Reader for the V2AlertsNrqlConditionsJSONByConditionIDDelete structure.
type V2AlertsNrqlConditionsJSONByConditionIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2AlertsNrqlConditionsJSONByConditionIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewV2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2AlertsNrqlConditionsJSONByConditionIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2AlertsNrqlConditionsJSONByConditionIDDeleteOK creates a V2AlertsNrqlConditionsJSONByConditionIDDeleteOK with default headers values
func NewV2AlertsNrqlConditionsJSONByConditionIDDeleteOK() *V2AlertsNrqlConditionsJSONByConditionIDDeleteOK {
	return &V2AlertsNrqlConditionsJSONByConditionIDDeleteOK{}
}

/*V2AlertsNrqlConditionsJSONByConditionIDDeleteOK handles this case with default header values.

OK
*/
type V2AlertsNrqlConditionsJSONByConditionIDDeleteOK struct {
	Payload *models.NrqlConditionResponse
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_nrql_conditions/{condition_id}.json][%d] v2AlertsNrqlConditionsJsonByConditionIdDeleteOK  %+v", 200, o.Payload)
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NrqlConditionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized creates a V2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized with default headers values
func NewV2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized() *V2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized {
	return &V2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized{}
}

/*V2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized handles this case with default header values.

Invalid API key
*/
type V2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized struct {
	Payload interface{}
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_nrql_conditions/{condition_id}.json][%d] v2AlertsNrqlConditionsJsonByConditionIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden creates a V2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden with default headers values
func NewV2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden() *V2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden {
	return &V2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden{}
}

/*V2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden struct {
	Payload interface{}
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_nrql_conditions/{condition_id}.json][%d] v2AlertsNrqlConditionsJsonByConditionIdDeleteForbidden  %+v", 403, o.Payload)
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound creates a V2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound with default headers values
func NewV2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound() *V2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound {
	return &V2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound{}
}

/*V2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound handles this case with default header values.

No Alerts condition was found with the given ID
*/
type V2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound struct {
	Payload interface{}
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_nrql_conditions/{condition_id}.json][%d] v2AlertsNrqlConditionsJsonByConditionIdDeleteNotFound  %+v", 404, o.Payload)
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsNrqlConditionsJSONByConditionIDDeleteDefault creates a V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault with default headers values
func NewV2AlertsNrqlConditionsJSONByConditionIDDeleteDefault(code int) *V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault {
	return &V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault{
		_statusCode: code,
	}
}

/*V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 alerts nrql conditions Json by condition Id delete default response
func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_nrql_conditions/{condition_id}.json][%d] V2AlertsNrqlConditionsJsonByConditionIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *V2AlertsNrqlConditionsJSONByConditionIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
