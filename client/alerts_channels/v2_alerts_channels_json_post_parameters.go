// Code generated by go-swagger; DO NOT EDIT.

package alerts_channels

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

// NewV2AlertsChannelsJSONPostParams creates a new V2AlertsChannelsJSONPostParams object
// with the default values initialized.
func NewV2AlertsChannelsJSONPostParams() *V2AlertsChannelsJSONPostParams {
	var ()
	return &V2AlertsChannelsJSONPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AlertsChannelsJSONPostParamsWithTimeout creates a new V2AlertsChannelsJSONPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AlertsChannelsJSONPostParamsWithTimeout(timeout time.Duration) *V2AlertsChannelsJSONPostParams {
	var ()
	return &V2AlertsChannelsJSONPostParams{

		timeout: timeout,
	}
}

// NewV2AlertsChannelsJSONPostParamsWithContext creates a new V2AlertsChannelsJSONPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AlertsChannelsJSONPostParamsWithContext(ctx context.Context) *V2AlertsChannelsJSONPostParams {
	var ()
	return &V2AlertsChannelsJSONPostParams{

		Context: ctx,
	}
}

// NewV2AlertsChannelsJSONPostParamsWithHTTPClient creates a new V2AlertsChannelsJSONPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AlertsChannelsJSONPostParamsWithHTTPClient(client *http.Client) *V2AlertsChannelsJSONPostParams {
	var ()
	return &V2AlertsChannelsJSONPostParams{
		HTTPClient: client,
	}
}

/*V2AlertsChannelsJSONPostParams contains all the parameters to send to the API endpoint
for the v2 alerts channels Json post operation typically these are written to a http.Request
*/
type V2AlertsChannelsJSONPostParams struct {

	/*Channel
	  channel schema

	*/
	Channel *models.Channel
	/*PolicyIds
	  Policy IDs to associate with channel

	*/
	PolicyIds []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) WithTimeout(timeout time.Duration) *V2AlertsChannelsJSONPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) WithContext(ctx context.Context) *V2AlertsChannelsJSONPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) WithHTTPClient(client *http.Client) *V2AlertsChannelsJSONPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannel adds the channel to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) WithChannel(channel *models.Channel) *V2AlertsChannelsJSONPostParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) SetChannel(channel *models.Channel) {
	o.Channel = channel
}

// WithPolicyIds adds the policyIds to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) WithPolicyIds(policyIds []int64) *V2AlertsChannelsJSONPostParams {
	o.SetPolicyIds(policyIds)
	return o
}

// SetPolicyIds adds the policyIds to the v2 alerts channels Json post params
func (o *V2AlertsChannelsJSONPostParams) SetPolicyIds(policyIds []int64) {
	o.PolicyIds = policyIds
}

// WriteToRequest writes these params to a swagger request
func (o *V2AlertsChannelsJSONPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Channel != nil {
		if err := r.SetBodyParam(o.Channel); err != nil {
			return err
		}
	}

	var valuesPolicyIds []string
	for _, v := range o.PolicyIds {
		valuesPolicyIds = append(valuesPolicyIds, swag.FormatInt64(v))
	}

	joinedPolicyIds := swag.JoinByFormat(valuesPolicyIds, "csv")
	// query array param policy_ids
	if err := r.SetQueryParam("policy_ids", joinedPolicyIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
