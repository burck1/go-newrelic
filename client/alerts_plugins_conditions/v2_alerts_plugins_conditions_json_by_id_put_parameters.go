// Code generated by go-swagger; DO NOT EDIT.

package alerts_plugins_conditions

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

// NewV2AlertsPluginsConditionsJSONByIDPutParams creates a new V2AlertsPluginsConditionsJSONByIDPutParams object
// with the default values initialized.
func NewV2AlertsPluginsConditionsJSONByIDPutParams() *V2AlertsPluginsConditionsJSONByIDPutParams {
	var ()
	return &V2AlertsPluginsConditionsJSONByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AlertsPluginsConditionsJSONByIDPutParamsWithTimeout creates a new V2AlertsPluginsConditionsJSONByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AlertsPluginsConditionsJSONByIDPutParamsWithTimeout(timeout time.Duration) *V2AlertsPluginsConditionsJSONByIDPutParams {
	var ()
	return &V2AlertsPluginsConditionsJSONByIDPutParams{

		timeout: timeout,
	}
}

// NewV2AlertsPluginsConditionsJSONByIDPutParamsWithContext creates a new V2AlertsPluginsConditionsJSONByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AlertsPluginsConditionsJSONByIDPutParamsWithContext(ctx context.Context) *V2AlertsPluginsConditionsJSONByIDPutParams {
	var ()
	return &V2AlertsPluginsConditionsJSONByIDPutParams{

		Context: ctx,
	}
}

// NewV2AlertsPluginsConditionsJSONByIDPutParamsWithHTTPClient creates a new V2AlertsPluginsConditionsJSONByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AlertsPluginsConditionsJSONByIDPutParamsWithHTTPClient(client *http.Client) *V2AlertsPluginsConditionsJSONByIDPutParams {
	var ()
	return &V2AlertsPluginsConditionsJSONByIDPutParams{
		HTTPClient: client,
	}
}

/*V2AlertsPluginsConditionsJSONByIDPutParams contains all the parameters to send to the API endpoint
for the v2 alerts plugins conditions Json by Id put operation typically these are written to a http.Request
*/
type V2AlertsPluginsConditionsJSONByIDPutParams struct {

	/*ConditionID
	  Alerts condition ID to update

	*/
	ConditionID int64
	/*PluginsCondition
	  Condition schema

	*/
	PluginsCondition *models.PluginsCondition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) WithTimeout(timeout time.Duration) *V2AlertsPluginsConditionsJSONByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) WithContext(ctx context.Context) *V2AlertsPluginsConditionsJSONByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) WithHTTPClient(client *http.Client) *V2AlertsPluginsConditionsJSONByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConditionID adds the conditionID to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) WithConditionID(conditionID int64) *V2AlertsPluginsConditionsJSONByIDPutParams {
	o.SetConditionID(conditionID)
	return o
}

// SetConditionID adds the conditionId to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) SetConditionID(conditionID int64) {
	o.ConditionID = conditionID
}

// WithPluginsCondition adds the pluginsCondition to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) WithPluginsCondition(pluginsCondition *models.PluginsCondition) *V2AlertsPluginsConditionsJSONByIDPutParams {
	o.SetPluginsCondition(pluginsCondition)
	return o
}

// SetPluginsCondition adds the pluginsCondition to the v2 alerts plugins conditions Json by Id put params
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) SetPluginsCondition(pluginsCondition *models.PluginsCondition) {
	o.PluginsCondition = pluginsCondition
}

// WriteToRequest writes these params to a swagger request
func (o *V2AlertsPluginsConditionsJSONByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param condition_id
	if err := r.SetPathParam("condition_id", swag.FormatInt64(o.ConditionID)); err != nil {
		return err
	}

	if o.PluginsCondition != nil {
		if err := r.SetBodyParam(o.PluginsCondition); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
