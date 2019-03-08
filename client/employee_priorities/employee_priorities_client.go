// Code generated by go-swagger; DO NOT EDIT.

package employee_priorities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new employee priorities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for employee priorities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EmployeePrioritiesByIDDelete employee priorities by Id delete API
*/
func (a *Client) EmployeePrioritiesByIDDelete(params *EmployeePrioritiesByIDDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeePrioritiesByIDDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Employee_prioritiesByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/employee_priorities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EmployeePrioritiesByIDDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
EmployeePrioritiesByIDGet employee priorities by Id get API
*/
func (a *Client) EmployeePrioritiesByIDGet(params *EmployeePrioritiesByIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeePrioritiesByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeePrioritiesByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Employee_prioritiesByIdGet",
		Method:             "GET",
		PathPattern:        "/employee_priorities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EmployeePrioritiesByIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmployeePrioritiesByIDGetOK), nil

}

/*
EmployeePrioritiesByIDPut employee priorities by Id put API
*/
func (a *Client) EmployeePrioritiesByIDPut(params *EmployeePrioritiesByIDPutParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeePrioritiesByIDPutParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Employee_prioritiesByIdPut",
		Method:             "PUT",
		PathPattern:        "/employee_priorities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EmployeePrioritiesByIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
EmployeePrioritiesGet employee priorities get API
*/
func (a *Client) EmployeePrioritiesGet(params *EmployeePrioritiesGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeePrioritiesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeePrioritiesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Employee_prioritiesGet",
		Method:             "GET",
		PathPattern:        "/employee_priorities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EmployeePrioritiesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmployeePrioritiesGetOK), nil

}

/*
EmployeePrioritiesPost employee priorities post API
*/
func (a *Client) EmployeePrioritiesPost(params *EmployeePrioritiesPostParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeePrioritiesPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeePrioritiesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Employee_prioritiesPost",
		Method:             "POST",
		PathPattern:        "/employee_priorities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EmployeePrioritiesPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmployeePrioritiesPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
