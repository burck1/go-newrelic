// Code generated by go-swagger; DO NOT EDIT.

package mobile_applications

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

// NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams creates a new V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams object
// with the default values initialized.
func NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams() *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	var ()
	return &V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParamsWithTimeout creates a new V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParamsWithTimeout(timeout time.Duration) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	var ()
	return &V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams{

		timeout: timeout,
	}
}

// NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParamsWithContext creates a new V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParamsWithContext(ctx context.Context) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	var ()
	return &V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams{

		Context: ctx,
	}
}

// NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParamsWithHTTPClient creates a new V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParamsWithHTTPClient(client *http.Client) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	var ()
	return &V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams{
		HTTPClient: client,
	}
}

/*V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams contains all the parameters to send to the API endpoint
for the v2 mobile applications metrics data Json by mobile application Id get operation typically these are written to a http.Request
*/
type V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams struct {

	/*From
	  Retrieve metrics after this time

	*/
	From *strfmt.DateTime
	/*MobileApplicationID
	  Mobile application ID

	*/
	MobileApplicationID int64
	/*Names
	  Retrieve specific metrics by name

	*/
	Names []string
	/*Period
	  Period of timeslices in seconds

	*/
	Period *int64
	/*Raw
	  Return unformatted raw values

	*/
	Raw *bool
	/*Summarize
	  Summarize the data

	*/
	Summarize *bool
	/*To
	  Retrieve metrics before this time

	*/
	To *strfmt.DateTime
	/*Values
	  Retrieve specific metric values

	*/
	Values []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithTimeout(timeout time.Duration) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithContext(ctx context.Context) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithHTTPClient(client *http.Client) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithFrom(from *strfmt.DateTime) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetFrom(from *strfmt.DateTime) {
	o.From = from
}

// WithMobileApplicationID adds the mobileApplicationID to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithMobileApplicationID(mobileApplicationID int64) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetMobileApplicationID(mobileApplicationID)
	return o
}

// SetMobileApplicationID adds the mobileApplicationId to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetMobileApplicationID(mobileApplicationID int64) {
	o.MobileApplicationID = mobileApplicationID
}

// WithNames adds the names to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithNames(names []string) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetNames(names []string) {
	o.Names = names
}

// WithPeriod adds the period to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithPeriod(period *int64) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetPeriod(period *int64) {
	o.Period = period
}

// WithRaw adds the raw to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithRaw(raw *bool) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetRaw(raw)
	return o
}

// SetRaw adds the raw to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetRaw(raw *bool) {
	o.Raw = raw
}

// WithSummarize adds the summarize to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithSummarize(summarize *bool) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetSummarize(summarize)
	return o
}

// SetSummarize adds the summarize to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetSummarize(summarize *bool) {
	o.Summarize = summarize
}

// WithTo adds the to to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithTo(to *strfmt.DateTime) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetTo(to *strfmt.DateTime) {
	o.To = to
}

// WithValues adds the values to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WithValues(values []string) *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams {
	o.SetValues(values)
	return o
}

// SetValues adds the values to the v2 mobile applications metrics data Json by mobile application Id get params
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) SetValues(values []string) {
	o.Values = values
}

// WriteToRequest writes these params to a swagger request
func (o *V2MobileApplicationsMetricsDataJSONByMobileApplicationIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom strfmt.DateTime
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom.String()
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	// path param mobile_application_id
	if err := r.SetPathParam("mobile_application_id", swag.FormatInt64(o.MobileApplicationID)); err != nil {
		return err
	}

	valuesNames := o.Names

	joinedNames := swag.JoinByFormat(valuesNames, "csv")
	// query array param names
	if err := r.SetQueryParam("names", joinedNames...); err != nil {
		return err
	}

	if o.Period != nil {

		// query param period
		var qrPeriod int64
		if o.Period != nil {
			qrPeriod = *o.Period
		}
		qPeriod := swag.FormatInt64(qrPeriod)
		if qPeriod != "" {
			if err := r.SetQueryParam("period", qPeriod); err != nil {
				return err
			}
		}

	}

	if o.Raw != nil {

		// query param raw
		var qrRaw bool
		if o.Raw != nil {
			qrRaw = *o.Raw
		}
		qRaw := swag.FormatBool(qrRaw)
		if qRaw != "" {
			if err := r.SetQueryParam("raw", qRaw); err != nil {
				return err
			}
		}

	}

	if o.Summarize != nil {

		// query param summarize
		var qrSummarize bool
		if o.Summarize != nil {
			qrSummarize = *o.Summarize
		}
		qSummarize := swag.FormatBool(qrSummarize)
		if qSummarize != "" {
			if err := r.SetQueryParam("summarize", qSummarize); err != nil {
				return err
			}
		}

	}

	if o.To != nil {

		// query param to
		var qrTo strfmt.DateTime
		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo.String()
		if qTo != "" {
			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}

	}

	valuesValues := o.Values

	joinedValues := swag.JoinByFormat(valuesValues, "csv")
	// query array param values
	if err := r.SetQueryParam("values", joinedValues...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
