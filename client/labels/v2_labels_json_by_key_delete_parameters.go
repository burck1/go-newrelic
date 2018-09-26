// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewV2LabelsJSONByKeyDeleteParams creates a new V2LabelsJSONByKeyDeleteParams object
// with the default values initialized.
func NewV2LabelsJSONByKeyDeleteParams() *V2LabelsJSONByKeyDeleteParams {
	var ()
	return &V2LabelsJSONByKeyDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2LabelsJSONByKeyDeleteParamsWithTimeout creates a new V2LabelsJSONByKeyDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2LabelsJSONByKeyDeleteParamsWithTimeout(timeout time.Duration) *V2LabelsJSONByKeyDeleteParams {
	var ()
	return &V2LabelsJSONByKeyDeleteParams{

		timeout: timeout,
	}
}

// NewV2LabelsJSONByKeyDeleteParamsWithContext creates a new V2LabelsJSONByKeyDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2LabelsJSONByKeyDeleteParamsWithContext(ctx context.Context) *V2LabelsJSONByKeyDeleteParams {
	var ()
	return &V2LabelsJSONByKeyDeleteParams{

		Context: ctx,
	}
}

// NewV2LabelsJSONByKeyDeleteParamsWithHTTPClient creates a new V2LabelsJSONByKeyDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2LabelsJSONByKeyDeleteParamsWithHTTPClient(client *http.Client) *V2LabelsJSONByKeyDeleteParams {
	var ()
	return &V2LabelsJSONByKeyDeleteParams{
		HTTPClient: client,
	}
}

/*V2LabelsJSONByKeyDeleteParams contains all the parameters to send to the API endpoint
for the v2 labels Json by key delete operation typically these are written to a http.Request
*/
type V2LabelsJSONByKeyDeleteParams struct {

	/*Key
	  Label key. Example: 'Language:Java'

	*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) WithTimeout(timeout time.Duration) *V2LabelsJSONByKeyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) WithContext(ctx context.Context) *V2LabelsJSONByKeyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) WithHTTPClient(client *http.Client) *V2LabelsJSONByKeyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) WithKey(key string) *V2LabelsJSONByKeyDeleteParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the v2 labels Json by key delete params
func (o *V2LabelsJSONByKeyDeleteParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *V2LabelsJSONByKeyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
