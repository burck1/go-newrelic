// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// NewV2ApplicationsJSONByIDDeleteParams creates a new V2ApplicationsJSONByIDDeleteParams object
// with the default values initialized.
func NewV2ApplicationsJSONByIDDeleteParams() *V2ApplicationsJSONByIDDeleteParams {
	var ()
	return &V2ApplicationsJSONByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2ApplicationsJSONByIDDeleteParamsWithTimeout creates a new V2ApplicationsJSONByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2ApplicationsJSONByIDDeleteParamsWithTimeout(timeout time.Duration) *V2ApplicationsJSONByIDDeleteParams {
	var ()
	return &V2ApplicationsJSONByIDDeleteParams{

		timeout: timeout,
	}
}

// NewV2ApplicationsJSONByIDDeleteParamsWithContext creates a new V2ApplicationsJSONByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2ApplicationsJSONByIDDeleteParamsWithContext(ctx context.Context) *V2ApplicationsJSONByIDDeleteParams {
	var ()
	return &V2ApplicationsJSONByIDDeleteParams{

		Context: ctx,
	}
}

// NewV2ApplicationsJSONByIDDeleteParamsWithHTTPClient creates a new V2ApplicationsJSONByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2ApplicationsJSONByIDDeleteParamsWithHTTPClient(client *http.Client) *V2ApplicationsJSONByIDDeleteParams {
	var ()
	return &V2ApplicationsJSONByIDDeleteParams{
		HTTPClient: client,
	}
}

/*V2ApplicationsJSONByIDDeleteParams contains all the parameters to send to the API endpoint
for the v2 applications Json by Id delete operation typically these are written to a http.Request
*/
type V2ApplicationsJSONByIDDeleteParams struct {

	/*ApplicationID
	  Application ID

	*/
	ApplicationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) WithTimeout(timeout time.Duration) *V2ApplicationsJSONByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) WithContext(ctx context.Context) *V2ApplicationsJSONByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) WithHTTPClient(client *http.Client) *V2ApplicationsJSONByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) WithApplicationID(applicationID int64) *V2ApplicationsJSONByIDDeleteParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the v2 applications Json by Id delete params
func (o *V2ApplicationsJSONByIDDeleteParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WriteToRequest writes these params to a swagger request
func (o *V2ApplicationsJSONByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application_id
	if err := r.SetPathParam("application_id", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
