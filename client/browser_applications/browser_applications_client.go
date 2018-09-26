// Code generated by go-swagger; DO NOT EDIT.

package browser_applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new browser applications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for browser applications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
V2BrowserApplicationsJSONGet lists

List
*/
func (a *Client) V2BrowserApplicationsJSONGet(params *V2BrowserApplicationsJSONGetParams) (*V2BrowserApplicationsJSONGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2BrowserApplicationsJSONGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2BrowserApplicationsJsonGet",
		Method:             "GET",
		PathPattern:        "/v2/browser_applications.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V2BrowserApplicationsJSONGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2BrowserApplicationsJSONGetOK), nil

}

/*
V2BrowserApplicationsJSONPost creates

Create
*/
func (a *Client) V2BrowserApplicationsJSONPost(params *V2BrowserApplicationsJSONPostParams) (*V2BrowserApplicationsJSONPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2BrowserApplicationsJSONPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2BrowserApplicationsJsonPost",
		Method:             "POST",
		PathPattern:        "/v2/browser_applications.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V2BrowserApplicationsJSONPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2BrowserApplicationsJSONPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
