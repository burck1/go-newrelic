// Code generated by go-swagger; DO NOT EDIT.

package alerts_plugins_conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2AlertsPluginsConditionsJSONByConditionIDDeleteReader is a Reader for the V2AlertsPluginsConditionsJSONByConditionIDDelete structure.
type V2AlertsPluginsConditionsJSONByConditionIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2AlertsPluginsConditionsJSONByConditionIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewV2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2AlertsPluginsConditionsJSONByConditionIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2AlertsPluginsConditionsJSONByConditionIDDeleteOK creates a V2AlertsPluginsConditionsJSONByConditionIDDeleteOK with default headers values
func NewV2AlertsPluginsConditionsJSONByConditionIDDeleteOK() *V2AlertsPluginsConditionsJSONByConditionIDDeleteOK {
	return &V2AlertsPluginsConditionsJSONByConditionIDDeleteOK{}
}

/*V2AlertsPluginsConditionsJSONByConditionIDDeleteOK handles this case with default header values.

OK
*/
type V2AlertsPluginsConditionsJSONByConditionIDDeleteOK struct {
	Payload *models.PluginsConditionResponse
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_plugins_conditions/{condition_id}.json][%d] v2AlertsPluginsConditionsJsonByConditionIdDeleteOK  %+v", 200, o.Payload)
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PluginsConditionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized creates a V2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized with default headers values
func NewV2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized() *V2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized {
	return &V2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized{}
}

/*V2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized handles this case with default header values.

Invalid API key
*/
type V2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized struct {
	Payload interface{}
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_plugins_conditions/{condition_id}.json][%d] v2AlertsPluginsConditionsJsonByConditionIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden creates a V2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden with default headers values
func NewV2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden() *V2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden {
	return &V2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden{}
}

/*V2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden struct {
	Payload interface{}
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_plugins_conditions/{condition_id}.json][%d] v2AlertsPluginsConditionsJsonByConditionIdDeleteForbidden  %+v", 403, o.Payload)
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound creates a V2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound with default headers values
func NewV2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound() *V2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound {
	return &V2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound{}
}

/*V2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound handles this case with default header values.

No Alerts condition was found with the given ID
*/
type V2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound struct {
	Payload interface{}
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_plugins_conditions/{condition_id}.json][%d] v2AlertsPluginsConditionsJsonByConditionIdDeleteNotFound  %+v", 404, o.Payload)
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsPluginsConditionsJSONByConditionIDDeleteDefault creates a V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault with default headers values
func NewV2AlertsPluginsConditionsJSONByConditionIDDeleteDefault(code int) *V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault {
	return &V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault{
		_statusCode: code,
	}
}

/*V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 alerts plugins conditions Json by condition Id delete default response
func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_plugins_conditions/{condition_id}.json][%d] V2AlertsPluginsConditionsJsonByConditionIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *V2AlertsPluginsConditionsJSONByConditionIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
