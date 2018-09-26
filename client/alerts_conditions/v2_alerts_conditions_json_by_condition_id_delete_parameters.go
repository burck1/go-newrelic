// Code generated by go-swagger; DO NOT EDIT.

package alerts_conditions

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

// NewV2AlertsConditionsJSONByConditionIDDeleteParams creates a new V2AlertsConditionsJSONByConditionIDDeleteParams object
// with the default values initialized.
func NewV2AlertsConditionsJSONByConditionIDDeleteParams() *V2AlertsConditionsJSONByConditionIDDeleteParams {
	var ()
	return &V2AlertsConditionsJSONByConditionIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AlertsConditionsJSONByConditionIDDeleteParamsWithTimeout creates a new V2AlertsConditionsJSONByConditionIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AlertsConditionsJSONByConditionIDDeleteParamsWithTimeout(timeout time.Duration) *V2AlertsConditionsJSONByConditionIDDeleteParams {
	var ()
	return &V2AlertsConditionsJSONByConditionIDDeleteParams{

		timeout: timeout,
	}
}

// NewV2AlertsConditionsJSONByConditionIDDeleteParamsWithContext creates a new V2AlertsConditionsJSONByConditionIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AlertsConditionsJSONByConditionIDDeleteParamsWithContext(ctx context.Context) *V2AlertsConditionsJSONByConditionIDDeleteParams {
	var ()
	return &V2AlertsConditionsJSONByConditionIDDeleteParams{

		Context: ctx,
	}
}

// NewV2AlertsConditionsJSONByConditionIDDeleteParamsWithHTTPClient creates a new V2AlertsConditionsJSONByConditionIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AlertsConditionsJSONByConditionIDDeleteParamsWithHTTPClient(client *http.Client) *V2AlertsConditionsJSONByConditionIDDeleteParams {
	var ()
	return &V2AlertsConditionsJSONByConditionIDDeleteParams{
		HTTPClient: client,
	}
}

/*V2AlertsConditionsJSONByConditionIDDeleteParams contains all the parameters to send to the API endpoint
for the v2 alerts conditions Json by condition Id delete operation typically these are written to a http.Request
*/
type V2AlertsConditionsJSONByConditionIDDeleteParams struct {

	/*ConditionID
	  Alerts condition ID

	*/
	ConditionID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) WithTimeout(timeout time.Duration) *V2AlertsConditionsJSONByConditionIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) WithContext(ctx context.Context) *V2AlertsConditionsJSONByConditionIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) WithHTTPClient(client *http.Client) *V2AlertsConditionsJSONByConditionIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConditionID adds the conditionID to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) WithConditionID(conditionID int64) *V2AlertsConditionsJSONByConditionIDDeleteParams {
	o.SetConditionID(conditionID)
	return o
}

// SetConditionID adds the conditionId to the v2 alerts conditions Json by condition Id delete params
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) SetConditionID(conditionID int64) {
	o.ConditionID = conditionID
}

// WriteToRequest writes these params to a swagger request
func (o *V2AlertsConditionsJSONByConditionIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param condition_id
	if err := r.SetPathParam("condition_id", swag.FormatInt64(o.ConditionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
