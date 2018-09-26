// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2DashboardsJSONPostReader is a Reader for the V2DashboardsJSONPost structure.
type V2DashboardsJSONPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2DashboardsJSONPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2DashboardsJSONPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewV2DashboardsJSONPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewV2DashboardsJSONPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2DashboardsJSONPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewV2DashboardsJSONPostServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2DashboardsJSONPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2DashboardsJSONPostOK creates a V2DashboardsJSONPostOK with default headers values
func NewV2DashboardsJSONPostOK() *V2DashboardsJSONPostOK {
	return &V2DashboardsJSONPostOK{}
}

/*V2DashboardsJSONPostOK handles this case with default header values.

OK
*/
type V2DashboardsJSONPostOK struct {
	Payload *models.DashboardResponse
}

func (o *V2DashboardsJSONPostOK) Error() string {
	return fmt.Sprintf("[POST /v2/dashboards.json][%d] v2DashboardsJsonPostOK  %+v", 200, o.Payload)
}

func (o *V2DashboardsJSONPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DashboardsJSONPostBadRequest creates a V2DashboardsJSONPostBadRequest with default headers values
func NewV2DashboardsJSONPostBadRequest() *V2DashboardsJSONPostBadRequest {
	return &V2DashboardsJSONPostBadRequest{}
}

/*V2DashboardsJSONPostBadRequest handles this case with default header values.

Dashboard payload failed validation
*/
type V2DashboardsJSONPostBadRequest struct {
	Payload interface{}
}

func (o *V2DashboardsJSONPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/dashboards.json][%d] v2DashboardsJsonPostBadRequest  %+v", 400, o.Payload)
}

func (o *V2DashboardsJSONPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DashboardsJSONPostUnauthorized creates a V2DashboardsJSONPostUnauthorized with default headers values
func NewV2DashboardsJSONPostUnauthorized() *V2DashboardsJSONPostUnauthorized {
	return &V2DashboardsJSONPostUnauthorized{}
}

/*V2DashboardsJSONPostUnauthorized handles this case with default header values.

Invalid API key
*/
type V2DashboardsJSONPostUnauthorized struct {
	Payload interface{}
}

func (o *V2DashboardsJSONPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/dashboards.json][%d] v2DashboardsJsonPostUnauthorized  %+v", 401, o.Payload)
}

func (o *V2DashboardsJSONPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DashboardsJSONPostForbidden creates a V2DashboardsJSONPostForbidden with default headers values
func NewV2DashboardsJSONPostForbidden() *V2DashboardsJSONPostForbidden {
	return &V2DashboardsJSONPostForbidden{}
}

/*V2DashboardsJSONPostForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2DashboardsJSONPostForbidden struct {
	Payload interface{}
}

func (o *V2DashboardsJSONPostForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/dashboards.json][%d] v2DashboardsJsonPostForbidden  %+v", 403, o.Payload)
}

func (o *V2DashboardsJSONPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DashboardsJSONPostServiceUnavailable creates a V2DashboardsJSONPostServiceUnavailable with default headers values
func NewV2DashboardsJSONPostServiceUnavailable() *V2DashboardsJSONPostServiceUnavailable {
	return &V2DashboardsJSONPostServiceUnavailable{}
}

/*V2DashboardsJSONPostServiceUnavailable handles this case with default header values.

Service unavailable
*/
type V2DashboardsJSONPostServiceUnavailable struct {
	Payload interface{}
}

func (o *V2DashboardsJSONPostServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/dashboards.json][%d] v2DashboardsJsonPostServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2DashboardsJSONPostServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2DashboardsJSONPostDefault creates a V2DashboardsJSONPostDefault with default headers values
func NewV2DashboardsJSONPostDefault(code int) *V2DashboardsJSONPostDefault {
	return &V2DashboardsJSONPostDefault{
		_statusCode: code,
	}
}

/*V2DashboardsJSONPostDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2DashboardsJSONPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 dashboards Json post default response
func (o *V2DashboardsJSONPostDefault) Code() int {
	return o._statusCode
}

func (o *V2DashboardsJSONPostDefault) Error() string {
	return fmt.Sprintf("[POST /v2/dashboards.json][%d] V2DashboardsJsonPost default  %+v", o._statusCode, o.Payload)
}

func (o *V2DashboardsJSONPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
