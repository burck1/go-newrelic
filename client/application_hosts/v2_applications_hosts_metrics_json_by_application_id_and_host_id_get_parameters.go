// Code generated by go-swagger; DO NOT EDIT.

package application_hosts

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

// NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams creates a new V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams object
// with the default values initialized.
func NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams() *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	var ()
	return &V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParamsWithTimeout creates a new V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParamsWithTimeout(timeout time.Duration) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	var ()
	return &V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams{

		timeout: timeout,
	}
}

// NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParamsWithContext creates a new V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParamsWithContext(ctx context.Context) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	var ()
	return &V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams{

		Context: ctx,
	}
}

// NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParamsWithHTTPClient creates a new V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParamsWithHTTPClient(client *http.Client) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	var ()
	return &V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams{
		HTTPClient: client,
	}
}

/*V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams contains all the parameters to send to the API endpoint
for the v2 applications hosts metrics Json by application Id and host Id get operation typically these are written to a http.Request
*/
type V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams struct {

	/*ApplicationID
	  Application ID

	*/
	ApplicationID int64
	/*Cursor
	  Cursor for next page (replacing page param)

	*/
	Cursor *string
	/*HostID
	  Application Host ID

	*/
	HostID int64
	/*Name
	  Filter metrics by name

	*/
	Name *string
	/*Page
	  Pagination index (will be deprecated)

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithTimeout(timeout time.Duration) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithContext(ctx context.Context) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithHTTPClient(client *http.Client) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithApplicationID(applicationID int64) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WithCursor adds the cursor to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithCursor(cursor *string) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithHostID adds the hostID to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithHostID(hostID int64) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetHostID(hostID int64) {
	o.HostID = hostID
}

// WithName adds the name to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithName(name *string) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WithPage(page *int64) *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the v2 applications hosts metrics Json by application Id and host Id get params
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *V2ApplicationsHostsMetricsJSONByApplicationIDAndHostIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application_id
	if err := r.SetPathParam("application_id", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	// path param host_id
	if err := r.SetPathParam("host_id", swag.FormatInt64(o.HostID)); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
