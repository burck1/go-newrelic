// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2ServersMetricsDataJSONByServerIDGetReader is a Reader for the V2ServersMetricsDataJSONByServerIDGet structure.
type V2ServersMetricsDataJSONByServerIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ServersMetricsDataJSONByServerIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2ServersMetricsDataJSONByServerIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2ServersMetricsDataJSONByServerIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2ServersMetricsDataJSONByServerIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2ServersMetricsDataJSONByServerIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2ServersMetricsDataJSONByServerIDGetOK creates a V2ServersMetricsDataJSONByServerIDGetOK with default headers values
func NewV2ServersMetricsDataJSONByServerIDGetOK() *V2ServersMetricsDataJSONByServerIDGetOK {
	return &V2ServersMetricsDataJSONByServerIDGetOK{}
}

/*V2ServersMetricsDataJSONByServerIDGetOK handles this case with default header values.

OK
*/
type V2ServersMetricsDataJSONByServerIDGetOK struct {
	Payload *models.MetricDataResponse
}

func (o *V2ServersMetricsDataJSONByServerIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/servers/{server_id}/metrics/data.json][%d] v2ServersMetricsDataJsonByServerIdGetOK  %+v", 200, o.Payload)
}

func (o *V2ServersMetricsDataJSONByServerIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ServersMetricsDataJSONByServerIDGetUnauthorized creates a V2ServersMetricsDataJSONByServerIDGetUnauthorized with default headers values
func NewV2ServersMetricsDataJSONByServerIDGetUnauthorized() *V2ServersMetricsDataJSONByServerIDGetUnauthorized {
	return &V2ServersMetricsDataJSONByServerIDGetUnauthorized{}
}

/*V2ServersMetricsDataJSONByServerIDGetUnauthorized handles this case with default header values.

Invalid API key
*/
type V2ServersMetricsDataJSONByServerIDGetUnauthorized struct {
	Payload interface{}
}

func (o *V2ServersMetricsDataJSONByServerIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/servers/{server_id}/metrics/data.json][%d] v2ServersMetricsDataJsonByServerIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ServersMetricsDataJSONByServerIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ServersMetricsDataJSONByServerIDGetForbidden creates a V2ServersMetricsDataJSONByServerIDGetForbidden with default headers values
func NewV2ServersMetricsDataJSONByServerIDGetForbidden() *V2ServersMetricsDataJSONByServerIDGetForbidden {
	return &V2ServersMetricsDataJSONByServerIDGetForbidden{}
}

/*V2ServersMetricsDataJSONByServerIDGetForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2ServersMetricsDataJSONByServerIDGetForbidden struct {
	Payload interface{}
}

func (o *V2ServersMetricsDataJSONByServerIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/servers/{server_id}/metrics/data.json][%d] v2ServersMetricsDataJsonByServerIdGetForbidden  %+v", 403, o.Payload)
}

func (o *V2ServersMetricsDataJSONByServerIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ServersMetricsDataJSONByServerIDGetDefault creates a V2ServersMetricsDataJSONByServerIDGetDefault with default headers values
func NewV2ServersMetricsDataJSONByServerIDGetDefault(code int) *V2ServersMetricsDataJSONByServerIDGetDefault {
	return &V2ServersMetricsDataJSONByServerIDGetDefault{
		_statusCode: code,
	}
}

/*V2ServersMetricsDataJSONByServerIDGetDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2ServersMetricsDataJSONByServerIDGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 servers metrics data Json by server Id get default response
func (o *V2ServersMetricsDataJSONByServerIDGetDefault) Code() int {
	return o._statusCode
}

func (o *V2ServersMetricsDataJSONByServerIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /v2/servers/{server_id}/metrics/data.json][%d] V2ServersMetricsDataJsonByServerIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *V2ServersMetricsDataJSONByServerIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
