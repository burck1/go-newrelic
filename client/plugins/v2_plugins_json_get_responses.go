// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2PluginsJSONGetReader is a Reader for the V2PluginsJSONGet structure.
type V2PluginsJSONGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2PluginsJSONGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2PluginsJSONGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2PluginsJSONGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2PluginsJSONGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2PluginsJSONGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2PluginsJSONGetOK creates a V2PluginsJSONGetOK with default headers values
func NewV2PluginsJSONGetOK() *V2PluginsJSONGetOK {
	return &V2PluginsJSONGetOK{}
}

/*V2PluginsJSONGetOK handles this case with default header values.

OK
*/
type V2PluginsJSONGetOK struct {
	Payload *models.PluginResponse
}

func (o *V2PluginsJSONGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/plugins.json][%d] v2PluginsJsonGetOK  %+v", 200, o.Payload)
}

func (o *V2PluginsJSONGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PluginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2PluginsJSONGetUnauthorized creates a V2PluginsJSONGetUnauthorized with default headers values
func NewV2PluginsJSONGetUnauthorized() *V2PluginsJSONGetUnauthorized {
	return &V2PluginsJSONGetUnauthorized{}
}

/*V2PluginsJSONGetUnauthorized handles this case with default header values.

Invalid API key
*/
type V2PluginsJSONGetUnauthorized struct {
	Payload interface{}
}

func (o *V2PluginsJSONGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/plugins.json][%d] v2PluginsJsonGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V2PluginsJSONGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2PluginsJSONGetForbidden creates a V2PluginsJSONGetForbidden with default headers values
func NewV2PluginsJSONGetForbidden() *V2PluginsJSONGetForbidden {
	return &V2PluginsJSONGetForbidden{}
}

/*V2PluginsJSONGetForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2PluginsJSONGetForbidden struct {
	Payload interface{}
}

func (o *V2PluginsJSONGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/plugins.json][%d] v2PluginsJsonGetForbidden  %+v", 403, o.Payload)
}

func (o *V2PluginsJSONGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2PluginsJSONGetDefault creates a V2PluginsJSONGetDefault with default headers values
func NewV2PluginsJSONGetDefault(code int) *V2PluginsJSONGetDefault {
	return &V2PluginsJSONGetDefault{
		_statusCode: code,
	}
}

/*V2PluginsJSONGetDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2PluginsJSONGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 plugins Json get default response
func (o *V2PluginsJSONGetDefault) Code() int {
	return o._statusCode
}

func (o *V2PluginsJSONGetDefault) Error() string {
	return fmt.Sprintf("[GET /v2/plugins.json][%d] V2PluginsJsonGet default  %+v", o._statusCode, o.Payload)
}

func (o *V2PluginsJSONGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
