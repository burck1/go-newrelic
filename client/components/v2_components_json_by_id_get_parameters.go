// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewV2ComponentsJSONByIDGetParams creates a new V2ComponentsJSONByIDGetParams object
// with the default values initialized.
func NewV2ComponentsJSONByIDGetParams() *V2ComponentsJSONByIDGetParams {
	var ()
	return &V2ComponentsJSONByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2ComponentsJSONByIDGetParamsWithTimeout creates a new V2ComponentsJSONByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2ComponentsJSONByIDGetParamsWithTimeout(timeout time.Duration) *V2ComponentsJSONByIDGetParams {
	var ()
	return &V2ComponentsJSONByIDGetParams{

		timeout: timeout,
	}
}

// NewV2ComponentsJSONByIDGetParamsWithContext creates a new V2ComponentsJSONByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2ComponentsJSONByIDGetParamsWithContext(ctx context.Context) *V2ComponentsJSONByIDGetParams {
	var ()
	return &V2ComponentsJSONByIDGetParams{

		Context: ctx,
	}
}

// NewV2ComponentsJSONByIDGetParamsWithHTTPClient creates a new V2ComponentsJSONByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2ComponentsJSONByIDGetParamsWithHTTPClient(client *http.Client) *V2ComponentsJSONByIDGetParams {
	var ()
	return &V2ComponentsJSONByIDGetParams{
		HTTPClient: client,
	}
}

/*V2ComponentsJSONByIDGetParams contains all the parameters to send to the API endpoint
for the v2 components Json by Id get operation typically these are written to a http.Request
*/
type V2ComponentsJSONByIDGetParams struct {

	/*ComponentID
	  Plugin ID

	*/
	ComponentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) WithTimeout(timeout time.Duration) *V2ComponentsJSONByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) WithContext(ctx context.Context) *V2ComponentsJSONByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) WithHTTPClient(client *http.Client) *V2ComponentsJSONByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) WithComponentID(componentID int64) *V2ComponentsJSONByIDGetParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the v2 components Json by Id get params
func (o *V2ComponentsJSONByIDGetParams) SetComponentID(componentID int64) {
	o.ComponentID = componentID
}

// WriteToRequest writes these params to a swagger request
func (o *V2ComponentsJSONByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", swag.FormatInt64(o.ComponentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
