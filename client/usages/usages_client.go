// Code generated by go-swagger; DO NOT EDIT.

package usages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new usages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for usages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
V2UsagesJSONByProductGet lists

List
*/
func (a *Client) V2UsagesJSONByProductGet(params *V2UsagesJSONByProductGetParams) (*V2UsagesJSONByProductGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2UsagesJSONByProductGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2UsagesJsonByProductGet",
		Method:             "GET",
		PathPattern:        "/v2/usages/{product}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V2UsagesJSONByProductGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2UsagesJSONByProductGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
