// Code generated by go-swagger; DO NOT EDIT.

package alerts_policy_channels

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

// NewV2AlertsPolicyChannelsJSONDeleteParams creates a new V2AlertsPolicyChannelsJSONDeleteParams object
// with the default values initialized.
func NewV2AlertsPolicyChannelsJSONDeleteParams() *V2AlertsPolicyChannelsJSONDeleteParams {
	var ()
	return &V2AlertsPolicyChannelsJSONDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AlertsPolicyChannelsJSONDeleteParamsWithTimeout creates a new V2AlertsPolicyChannelsJSONDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AlertsPolicyChannelsJSONDeleteParamsWithTimeout(timeout time.Duration) *V2AlertsPolicyChannelsJSONDeleteParams {
	var ()
	return &V2AlertsPolicyChannelsJSONDeleteParams{

		timeout: timeout,
	}
}

// NewV2AlertsPolicyChannelsJSONDeleteParamsWithContext creates a new V2AlertsPolicyChannelsJSONDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AlertsPolicyChannelsJSONDeleteParamsWithContext(ctx context.Context) *V2AlertsPolicyChannelsJSONDeleteParams {
	var ()
	return &V2AlertsPolicyChannelsJSONDeleteParams{

		Context: ctx,
	}
}

// NewV2AlertsPolicyChannelsJSONDeleteParamsWithHTTPClient creates a new V2AlertsPolicyChannelsJSONDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AlertsPolicyChannelsJSONDeleteParamsWithHTTPClient(client *http.Client) *V2AlertsPolicyChannelsJSONDeleteParams {
	var ()
	return &V2AlertsPolicyChannelsJSONDeleteParams{
		HTTPClient: client,
	}
}

/*V2AlertsPolicyChannelsJSONDeleteParams contains all the parameters to send to the API endpoint
for the v2 alerts policy channels Json delete operation typically these are written to a http.Request
*/
type V2AlertsPolicyChannelsJSONDeleteParams struct {

	/*ChannelID
	  Channel ID

	*/
	ChannelID int64
	/*PolicyID
	  Policy ID

	*/
	PolicyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) WithTimeout(timeout time.Duration) *V2AlertsPolicyChannelsJSONDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) WithContext(ctx context.Context) *V2AlertsPolicyChannelsJSONDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) WithHTTPClient(client *http.Client) *V2AlertsPolicyChannelsJSONDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) WithChannelID(channelID int64) *V2AlertsPolicyChannelsJSONDeleteParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) SetChannelID(channelID int64) {
	o.ChannelID = channelID
}

// WithPolicyID adds the policyID to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) WithPolicyID(policyID int64) *V2AlertsPolicyChannelsJSONDeleteParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the v2 alerts policy channels Json delete params
func (o *V2AlertsPolicyChannelsJSONDeleteParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *V2AlertsPolicyChannelsJSONDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param channel_id
	qrChannelID := o.ChannelID
	qChannelID := swag.FormatInt64(qrChannelID)
	if qChannelID != "" {
		if err := r.SetQueryParam("channel_id", qChannelID); err != nil {
			return err
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
