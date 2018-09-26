// Code generated by go-swagger; DO NOT EDIT.

package alerts_policies

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

	models "github.com/burck1/go-newrelic/models"
)

// NewV2AlertsPoliciesJSONByIDPutParams creates a new V2AlertsPoliciesJSONByIDPutParams object
// with the default values initialized.
func NewV2AlertsPoliciesJSONByIDPutParams() *V2AlertsPoliciesJSONByIDPutParams {
	var ()
	return &V2AlertsPoliciesJSONByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AlertsPoliciesJSONByIDPutParamsWithTimeout creates a new V2AlertsPoliciesJSONByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AlertsPoliciesJSONByIDPutParamsWithTimeout(timeout time.Duration) *V2AlertsPoliciesJSONByIDPutParams {
	var ()
	return &V2AlertsPoliciesJSONByIDPutParams{

		timeout: timeout,
	}
}

// NewV2AlertsPoliciesJSONByIDPutParamsWithContext creates a new V2AlertsPoliciesJSONByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AlertsPoliciesJSONByIDPutParamsWithContext(ctx context.Context) *V2AlertsPoliciesJSONByIDPutParams {
	var ()
	return &V2AlertsPoliciesJSONByIDPutParams{

		Context: ctx,
	}
}

// NewV2AlertsPoliciesJSONByIDPutParamsWithHTTPClient creates a new V2AlertsPoliciesJSONByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AlertsPoliciesJSONByIDPutParamsWithHTTPClient(client *http.Client) *V2AlertsPoliciesJSONByIDPutParams {
	var ()
	return &V2AlertsPoliciesJSONByIDPutParams{
		HTTPClient: client,
	}
}

/*V2AlertsPoliciesJSONByIDPutParams contains all the parameters to send to the API endpoint
for the v2 alerts policies Json by Id put operation typically these are written to a http.Request
*/
type V2AlertsPoliciesJSONByIDPutParams struct {

	/*Policy
	  Policy Schema

	*/
	Policy *models.Policy
	/*PolicyID
	  Policy ID

	*/
	PolicyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) WithTimeout(timeout time.Duration) *V2AlertsPoliciesJSONByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) WithContext(ctx context.Context) *V2AlertsPoliciesJSONByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) WithHTTPClient(client *http.Client) *V2AlertsPoliciesJSONByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) WithPolicy(policy *models.Policy) *V2AlertsPoliciesJSONByIDPutParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) SetPolicy(policy *models.Policy) {
	o.Policy = policy
}

// WithPolicyID adds the policyID to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) WithPolicyID(policyID int64) *V2AlertsPoliciesJSONByIDPutParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the v2 alerts policies Json by Id put params
func (o *V2AlertsPoliciesJSONByIDPutParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *V2AlertsPoliciesJSONByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	// path param policy_id
	if err := r.SetPathParam("policy_id", swag.FormatInt64(o.PolicyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
