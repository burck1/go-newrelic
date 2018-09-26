// Code generated by go-swagger; DO NOT EDIT.

package alerts_channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/burck1/go-newrelic/models"
)

// V2AlertsChannelsJSONByChannelIDDeleteReader is a Reader for the V2AlertsChannelsJSONByChannelIDDelete structure.
type V2AlertsChannelsJSONByChannelIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2AlertsChannelsJSONByChannelIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2AlertsChannelsJSONByChannelIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewV2AlertsChannelsJSONByChannelIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewV2AlertsChannelsJSONByChannelIDDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewV2AlertsChannelsJSONByChannelIDDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewV2AlertsChannelsJSONByChannelIDDeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewV2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewV2AlertsChannelsJSONByChannelIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2AlertsChannelsJSONByChannelIDDeleteOK creates a V2AlertsChannelsJSONByChannelIDDeleteOK with default headers values
func NewV2AlertsChannelsJSONByChannelIDDeleteOK() *V2AlertsChannelsJSONByChannelIDDeleteOK {
	return &V2AlertsChannelsJSONByChannelIDDeleteOK{}
}

/*V2AlertsChannelsJSONByChannelIDDeleteOK handles this case with default header values.

OK
*/
type V2AlertsChannelsJSONByChannelIDDeleteOK struct {
	Payload *models.ChannelResponse
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_channels/{channel_id}.json][%d] v2AlertsChannelsJsonByChannelIdDeleteOK  %+v", 200, o.Payload)
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsChannelsJSONByChannelIDDeleteUnauthorized creates a V2AlertsChannelsJSONByChannelIDDeleteUnauthorized with default headers values
func NewV2AlertsChannelsJSONByChannelIDDeleteUnauthorized() *V2AlertsChannelsJSONByChannelIDDeleteUnauthorized {
	return &V2AlertsChannelsJSONByChannelIDDeleteUnauthorized{}
}

/*V2AlertsChannelsJSONByChannelIDDeleteUnauthorized handles this case with default header values.

Invalid API key
*/
type V2AlertsChannelsJSONByChannelIDDeleteUnauthorized struct {
	Payload interface{}
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_channels/{channel_id}.json][%d] v2AlertsChannelsJsonByChannelIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsChannelsJSONByChannelIDDeleteForbidden creates a V2AlertsChannelsJSONByChannelIDDeleteForbidden with default headers values
func NewV2AlertsChannelsJSONByChannelIDDeleteForbidden() *V2AlertsChannelsJSONByChannelIDDeleteForbidden {
	return &V2AlertsChannelsJSONByChannelIDDeleteForbidden{}
}

/*V2AlertsChannelsJSONByChannelIDDeleteForbidden handles this case with default header values.

New Relic API access has not been enabled
*/
type V2AlertsChannelsJSONByChannelIDDeleteForbidden struct {
	Payload interface{}
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_channels/{channel_id}.json][%d] v2AlertsChannelsJsonByChannelIdDeleteForbidden  %+v", 403, o.Payload)
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsChannelsJSONByChannelIDDeleteNotFound creates a V2AlertsChannelsJSONByChannelIDDeleteNotFound with default headers values
func NewV2AlertsChannelsJSONByChannelIDDeleteNotFound() *V2AlertsChannelsJSONByChannelIDDeleteNotFound {
	return &V2AlertsChannelsJSONByChannelIDDeleteNotFound{}
}

/*V2AlertsChannelsJSONByChannelIDDeleteNotFound handles this case with default header values.

No channel was found with the given ID
*/
type V2AlertsChannelsJSONByChannelIDDeleteNotFound struct {
	Payload interface{}
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_channels/{channel_id}.json][%d] v2AlertsChannelsJsonByChannelIdDeleteNotFound  %+v", 404, o.Payload)
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsChannelsJSONByChannelIDDeleteNotAcceptable creates a V2AlertsChannelsJSONByChannelIDDeleteNotAcceptable with default headers values
func NewV2AlertsChannelsJSONByChannelIDDeleteNotAcceptable() *V2AlertsChannelsJSONByChannelIDDeleteNotAcceptable {
	return &V2AlertsChannelsJSONByChannelIDDeleteNotAcceptable{}
}

/*V2AlertsChannelsJSONByChannelIDDeleteNotAcceptable handles this case with default header values.

Cannot delete USER channels
*/
type V2AlertsChannelsJSONByChannelIDDeleteNotAcceptable struct {
	Payload interface{}
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_channels/{channel_id}.json][%d] v2AlertsChannelsJsonByChannelIdDeleteNotAcceptable  %+v", 406, o.Payload)
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity creates a V2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity with default headers values
func NewV2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity() *V2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity {
	return &V2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity{}
}

/*V2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity handles this case with default header values.

Validation or internal error occurred
*/
type V2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity struct {
	Payload interface{}
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_channels/{channel_id}.json][%d] v2AlertsChannelsJsonByChannelIdDeleteUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2AlertsChannelsJSONByChannelIDDeleteDefault creates a V2AlertsChannelsJSONByChannelIDDeleteDefault with default headers values
func NewV2AlertsChannelsJSONByChannelIDDeleteDefault(code int) *V2AlertsChannelsJSONByChannelIDDeleteDefault {
	return &V2AlertsChannelsJSONByChannelIDDeleteDefault{
		_statusCode: code,
	}
}

/*V2AlertsChannelsJSONByChannelIDDeleteDefault handles this case with default header values.

A server error occurred, please contact <a href="http://support.newrelic.com">New Relic support</a>
*/
type V2AlertsChannelsJSONByChannelIDDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the v2 alerts channels Json by channel Id delete default response
func (o *V2AlertsChannelsJSONByChannelIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v2/alerts_channels/{channel_id}.json][%d] V2AlertsChannelsJsonByChannelIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *V2AlertsChannelsJSONByChannelIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
