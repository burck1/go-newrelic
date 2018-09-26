// Code generated by go-swagger; DO NOT EDIT.

package alerts_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new alerts policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
V2AlertsPoliciesJSONByIDPut updates

Update
*/
func (a *Client) V2AlertsPoliciesJSONByIDPut(params *V2AlertsPoliciesJSONByIDPutParams) (*V2AlertsPoliciesJSONByIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2AlertsPoliciesJSONByIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2AlertsPoliciesJsonByIdPut",
		Method:             "PUT",
		PathPattern:        "/v2/alerts_policies/{policy_id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V2AlertsPoliciesJSONByIDPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2AlertsPoliciesJSONByIDPutOK), nil

}

/*
V2AlertsPoliciesJSONByPolicyIDDelete deletes

Delete
*/
func (a *Client) V2AlertsPoliciesJSONByPolicyIDDelete(params *V2AlertsPoliciesJSONByPolicyIDDeleteParams) (*V2AlertsPoliciesJSONByPolicyIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2AlertsPoliciesJSONByPolicyIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2AlertsPoliciesJsonByPolicyIdDelete",
		Method:             "DELETE",
		PathPattern:        "/v2/alerts_policies/{policy_id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V2AlertsPoliciesJSONByPolicyIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2AlertsPoliciesJSONByPolicyIDDeleteOK), nil

}

/*
V2AlertsPoliciesJSONGet lists

List
*/
func (a *Client) V2AlertsPoliciesJSONGet(params *V2AlertsPoliciesJSONGetParams) (*V2AlertsPoliciesJSONGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2AlertsPoliciesJSONGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2AlertsPoliciesJsonGet",
		Method:             "GET",
		PathPattern:        "/v2/alerts_policies.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V2AlertsPoliciesJSONGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2AlertsPoliciesJSONGetOK), nil

}

/*
V2AlertsPoliciesJSONPost creates

Create
*/
func (a *Client) V2AlertsPoliciesJSONPost(params *V2AlertsPoliciesJSONPostParams) (*V2AlertsPoliciesJSONPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2AlertsPoliciesJSONPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2AlertsPoliciesJsonPost",
		Method:             "POST",
		PathPattern:        "/v2/alerts_policies.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V2AlertsPoliciesJSONPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2AlertsPoliciesJSONPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
