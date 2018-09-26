// Code generated by go-swagger; DO NOT EDIT.

package browser_applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2BrowserApplicationsJSONPostReader is a Reader for the V2BrowserApplicationsJSONPost structure.
type V2BrowserApplicationsJSONPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2BrowserApplicationsJSONPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2BrowserApplicationsJSONPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2BrowserApplicationsJSONPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2BrowserApplicationsJSONPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2BrowserApplicationsJSONPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2BrowserApplicationsJSONPostOK creates a V2BrowserApplicationsJSONPostOK with default headers values
func NewV2BrowserApplicationsJSONPostOK() *V2BrowserApplicationsJSONPostOK {
	return &V2BrowserApplicationsJSONPostOK{}
}

/*V2BrowserApplicationsJSONPostOK handles this case with default header values.

OK
*/
type V2BrowserApplicationsJSONPostOK struct {
	Payload *models.BrowserApplicationResponse
}

func (o *V2BrowserApplicationsJSONPostOK) Error() string {
	return fmt.Sprintf("[POST /v2/browser_applications.json][%d] v2BrowserApplicationsJsonPostOK  %+v", 200, o.Payload)
}

func (o *V2BrowserApplicationsJSONPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BrowserApplicationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2BrowserApplicationsJSONPostUnauthorized creates a V2BrowserApplicationsJSONPostUnauthorized with default headers values
func NewV2BrowserApplicationsJSONPostUnauthorized() *V2BrowserApplicationsJSONPostUnauthorized {
	return &V2BrowserApplicationsJSONPostUnauthorized{}
}

/*V2BrowserApplicationsJSONPostUnauthorized handles this case with default header values.

Invalid API key
*/
type V2BrowserApplicationsJSONPostUnauthorized struct {
	Payload interface{}
}

func (o *V2BrowserApplicationsJSONPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/browser_applications.json][%d] v2BrowserApplicationsJsonPostUnauthorized  %+v", 401, o.Payload)
}

func (o *V2BrowserApplicationsJSONPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2BrowserApplicationsJSONPostForbidden creates a V2BrowserApplicationsJSONPostForbidden with default headers values
func NewV2BrowserApplicationsJSONPostForbidden() *V2BrowserApplicationsJSONPostForbidden {
	return &V2BrowserApplicationsJSONPostForbidden{}
}

/*V2BrowserApplicationsJSONPostForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2BrowserApplicationsJSONPostForbidden struct {
	Payload interface{}
}

func (o *V2BrowserApplicationsJSONPostForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/browser_applications.json][%d] v2BrowserApplicationsJsonPostForbidden  %+v", 403, o.Payload)
}

func (o *V2BrowserApplicationsJSONPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2BrowserApplicationsJSONPostDefault creates a V2BrowserApplicationsJSONPostDefault with default headers values
func NewV2BrowserApplicationsJSONPostDefault(code int) *V2BrowserApplicationsJSONPostDefault {
	return &V2BrowserApplicationsJSONPostDefault{
		_statusCode: code,
	}
}

/*V2BrowserApplicationsJSONPostDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2BrowserApplicationsJSONPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 browser applications Json post default response
func (o *V2BrowserApplicationsJSONPostDefault) Code() int {
	return o._statusCode
}

func (o *V2BrowserApplicationsJSONPostDefault) Error() string {
	return fmt.Sprintf("[POST /v2/browser_applications.json][%d] V2BrowserApplicationsJsonPost default  %+v", o._statusCode, o.Payload)
}

func (o *V2BrowserApplicationsJSONPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
