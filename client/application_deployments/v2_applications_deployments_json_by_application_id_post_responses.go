// Code generated by go-swagger; DO NOT EDIT.

package application_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2ApplicationsDeploymentsJSONByApplicationIDPostReader is a Reader for the V2ApplicationsDeploymentsJSONByApplicationIDPost structure.
type V2ApplicationsDeploymentsJSONByApplicationIDPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2ApplicationsDeploymentsJSONByApplicationIDPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewV2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewV2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2ApplicationsDeploymentsJSONByApplicationIDPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewV2ApplicationsDeploymentsJSONByApplicationIDPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2ApplicationsDeploymentsJSONByApplicationIDPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2ApplicationsDeploymentsJSONByApplicationIDPostOK creates a V2ApplicationsDeploymentsJSONByApplicationIDPostOK with default headers values
func NewV2ApplicationsDeploymentsJSONByApplicationIDPostOK() *V2ApplicationsDeploymentsJSONByApplicationIDPostOK {
	return &V2ApplicationsDeploymentsJSONByApplicationIDPostOK{}
}

/*V2ApplicationsDeploymentsJSONByApplicationIDPostOK handles this case with default header values.

OK
*/
type V2ApplicationsDeploymentsJSONByApplicationIDPostOK struct {
	Payload *models.DeploymentResponse
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostOK) Error() string {
	return fmt.Sprintf("[POST /v2/applications/{application_id}/deployments.json][%d] v2ApplicationsDeploymentsJsonByApplicationIdPostOK  %+v", 200, o.Payload)
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest creates a V2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest with default headers values
func NewV2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest() *V2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest {
	return &V2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest{}
}

/*V2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest handles this case with default header values.

Revision parameter required
*/
type V2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest struct {
	Payload interface{}
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/applications/{application_id}/deployments.json][%d] v2ApplicationsDeploymentsJsonByApplicationIdPostBadRequest  %+v", 400, o.Payload)
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized creates a V2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized with default headers values
func NewV2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized() *V2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized {
	return &V2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized{}
}

/*V2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized handles this case with default header values.

Invalid API key
*/
type V2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized struct {
	Payload interface{}
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/applications/{application_id}/deployments.json][%d] v2ApplicationsDeploymentsJsonByApplicationIdPostUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsDeploymentsJSONByApplicationIDPostForbidden creates a V2ApplicationsDeploymentsJSONByApplicationIDPostForbidden with default headers values
func NewV2ApplicationsDeploymentsJSONByApplicationIDPostForbidden() *V2ApplicationsDeploymentsJSONByApplicationIDPostForbidden {
	return &V2ApplicationsDeploymentsJSONByApplicationIDPostForbidden{}
}

/*V2ApplicationsDeploymentsJSONByApplicationIDPostForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2ApplicationsDeploymentsJSONByApplicationIDPostForbidden struct {
	Payload interface{}
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/applications/{application_id}/deployments.json][%d] v2ApplicationsDeploymentsJsonByApplicationIdPostForbidden  %+v", 403, o.Payload)
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsDeploymentsJSONByApplicationIDPostNotFound creates a V2ApplicationsDeploymentsJSONByApplicationIDPostNotFound with default headers values
func NewV2ApplicationsDeploymentsJSONByApplicationIDPostNotFound() *V2ApplicationsDeploymentsJSONByApplicationIDPostNotFound {
	return &V2ApplicationsDeploymentsJSONByApplicationIDPostNotFound{}
}

/*V2ApplicationsDeploymentsJSONByApplicationIDPostNotFound handles this case with default header values.

No Application found with the given ID
*/
type V2ApplicationsDeploymentsJSONByApplicationIDPostNotFound struct {
	Payload interface{}
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/applications/{application_id}/deployments.json][%d] v2ApplicationsDeploymentsJsonByApplicationIdPostNotFound  %+v", 404, o.Payload)
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ApplicationsDeploymentsJSONByApplicationIDPostDefault creates a V2ApplicationsDeploymentsJSONByApplicationIDPostDefault with default headers values
func NewV2ApplicationsDeploymentsJSONByApplicationIDPostDefault(code int) *V2ApplicationsDeploymentsJSONByApplicationIDPostDefault {
	return &V2ApplicationsDeploymentsJSONByApplicationIDPostDefault{
		_statusCode: code,
	}
}

/*V2ApplicationsDeploymentsJSONByApplicationIDPostDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2ApplicationsDeploymentsJSONByApplicationIDPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 applications deployments Json by application Id post default response
func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostDefault) Code() int {
	return o._statusCode
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostDefault) Error() string {
	return fmt.Sprintf("[POST /v2/applications/{application_id}/deployments.json][%d] V2ApplicationsDeploymentsJsonByApplicationIdPost default  %+v", o._statusCode, o.Payload)
}

func (o *V2ApplicationsDeploymentsJSONByApplicationIDPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
