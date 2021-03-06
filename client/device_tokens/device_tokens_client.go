// Code generated by go-swagger; DO NOT EDIT.

package device_tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new device tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for device tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeviceRegistrationRegisterPost registers a new device token
*/
func (a *Client) DeviceRegistrationRegisterPost(params *DeviceRegistrationRegisterPostParams, authInfo runtime.ClientAuthInfoWriter) (*DeviceRegistrationRegisterPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeviceRegistrationRegisterPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Device_registrationRegisterPost",
		Method:             "POST",
		PathPattern:        "/device_registration/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeviceRegistrationRegisterPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeviceRegistrationRegisterPostOK), nil

}

/*
DeviceRegistrationUnregisterPost unregisters a device with given token
*/
func (a *Client) DeviceRegistrationUnregisterPost(params *DeviceRegistrationUnregisterPostParams, authInfo runtime.ClientAuthInfoWriter) (*DeviceRegistrationUnregisterPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeviceRegistrationUnregisterPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Device_registrationUnregisterPost",
		Method:             "POST",
		PathPattern:        "/device_registration/unregister",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeviceRegistrationUnregisterPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeviceRegistrationUnregisterPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
