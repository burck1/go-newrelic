// Code generated by go-swagger; DO NOT EDIT.

package alerts_external_service_conditions

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

// NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams creates a new V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams object
// with the default values initialized.
func NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams() *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	var ()
	return &V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParamsWithTimeout creates a new V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParamsWithTimeout(timeout time.Duration) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	var ()
	return &V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams{

		timeout: timeout,
	}
}

// NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParamsWithContext creates a new V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParamsWithContext(ctx context.Context) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	var ()
	return &V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams{

		Context: ctx,
	}
}

// NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParamsWithHTTPClient creates a new V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParamsWithHTTPClient(client *http.Client) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	var ()
	return &V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams{
		HTTPClient: client,
	}
}

/*V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams contains all the parameters to send to the API endpoint
for the v2 alerts external service conditions policies Json by policy Id post operation typically these are written to a http.Request
*/
type V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams struct {

	/*ExternalServiceCondition
	  Condition schema

	*/
	ExternalServiceCondition *models.ExternalServiceCondition
	/*PolicyID
	  Alerts policy ID

	*/
	PolicyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) WithTimeout(timeout time.Duration) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) WithContext(ctx context.Context) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) WithHTTPClient(client *http.Client) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalServiceCondition adds the externalServiceCondition to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) WithExternalServiceCondition(externalServiceCondition *models.ExternalServiceCondition) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	o.SetExternalServiceCondition(externalServiceCondition)
	return o
}

// SetExternalServiceCondition adds the externalServiceCondition to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) SetExternalServiceCondition(externalServiceCondition *models.ExternalServiceCondition) {
	o.ExternalServiceCondition = externalServiceCondition
}

// WithPolicyID adds the policyID to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) WithPolicyID(policyID int64) *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the v2 alerts external service conditions policies Json by policy Id post params
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *V2AlertsExternalServiceConditionsPoliciesJSONByPolicyIDPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExternalServiceCondition != nil {
		if err := r.SetBodyParam(o.ExternalServiceCondition); err != nil {
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
