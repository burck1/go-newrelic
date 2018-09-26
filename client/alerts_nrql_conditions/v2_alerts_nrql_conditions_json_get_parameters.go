// Code generated by go-swagger; DO NOT EDIT.

package alerts_nrql_conditions

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

// NewV2AlertsNrqlConditionsJSONGetParams creates a new V2AlertsNrqlConditionsJSONGetParams object
// with the default values initialized.
func NewV2AlertsNrqlConditionsJSONGetParams() *V2AlertsNrqlConditionsJSONGetParams {
	var ()
	return &V2AlertsNrqlConditionsJSONGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AlertsNrqlConditionsJSONGetParamsWithTimeout creates a new V2AlertsNrqlConditionsJSONGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AlertsNrqlConditionsJSONGetParamsWithTimeout(timeout time.Duration) *V2AlertsNrqlConditionsJSONGetParams {
	var ()
	return &V2AlertsNrqlConditionsJSONGetParams{

		timeout: timeout,
	}
}

// NewV2AlertsNrqlConditionsJSONGetParamsWithContext creates a new V2AlertsNrqlConditionsJSONGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AlertsNrqlConditionsJSONGetParamsWithContext(ctx context.Context) *V2AlertsNrqlConditionsJSONGetParams {
	var ()
	return &V2AlertsNrqlConditionsJSONGetParams{

		Context: ctx,
	}
}

// NewV2AlertsNrqlConditionsJSONGetParamsWithHTTPClient creates a new V2AlertsNrqlConditionsJSONGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AlertsNrqlConditionsJSONGetParamsWithHTTPClient(client *http.Client) *V2AlertsNrqlConditionsJSONGetParams {
	var ()
	return &V2AlertsNrqlConditionsJSONGetParams{
		HTTPClient: client,
	}
}

/*V2AlertsNrqlConditionsJSONGetParams contains all the parameters to send to the API endpoint
for the v2 alerts nrql conditions Json get operation typically these are written to a http.Request
*/
type V2AlertsNrqlConditionsJSONGetParams struct {

	/*Page
	  Pagination index

	*/
	Page *int64
	/*PolicyID
	  Alerts policy ID

	*/
	PolicyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) WithTimeout(timeout time.Duration) *V2AlertsNrqlConditionsJSONGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) WithContext(ctx context.Context) *V2AlertsNrqlConditionsJSONGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) WithHTTPClient(client *http.Client) *V2AlertsNrqlConditionsJSONGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) WithPage(page *int64) *V2AlertsNrqlConditionsJSONGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) SetPage(page *int64) {
	o.Page = page
}

// WithPolicyID adds the policyID to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) WithPolicyID(policyID int64) *V2AlertsNrqlConditionsJSONGetParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the v2 alerts nrql conditions Json get params
func (o *V2AlertsNrqlConditionsJSONGetParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *V2AlertsNrqlConditionsJSONGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	// query param policy_id
	qrPolicyID := o.PolicyID
	qPolicyID := swag.FormatInt64(qrPolicyID)
	if qPolicyID != "" {
		if err := r.SetQueryParam("policy_id", qPolicyID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
