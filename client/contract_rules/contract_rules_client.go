// Code generated by go-swagger; DO NOT EDIT.

package contract_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new contract rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contract rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ContractrulesByIDDelete deletes a contract rule
*/
func (a *Client) ContractrulesByIDDelete(params *ContractrulesByIDDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ContractrulesByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractrulesByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContractrulesByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/contractrules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContractrulesByIDDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContractrulesByIDDeleteOK), nil

}

/*
ContractrulesByIDPut updates a contract rule
*/
func (a *Client) ContractrulesByIDPut(params *ContractrulesByIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*ContractrulesByIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractrulesByIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContractrulesByIdPut",
		Method:             "PUT",
		PathPattern:        "/contractrules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContractrulesByIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContractrulesByIDPutOK), nil

}

/*
ContractrulesGet gets the list of contract rules
*/
func (a *Client) ContractrulesGet(params *ContractrulesGetParams, authInfo runtime.ClientAuthInfoWriter) (*ContractrulesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractrulesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContractrulesGet",
		Method:             "GET",
		PathPattern:        "/contractrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContractrulesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContractrulesGetOK), nil

}

/*
ContractrulesPost creates a contract rule
*/
func (a *Client) ContractrulesPost(params *ContractrulesPostParams, authInfo runtime.ClientAuthInfoWriter) (*ContractrulesPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractrulesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContractrulesPost",
		Method:             "POST",
		PathPattern:        "/contractrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContractrulesPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContractrulesPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}