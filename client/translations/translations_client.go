// Code generated by go-swagger; DO NOT EDIT.

package translations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new translations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for translations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TranslationsGet returns the object representing the translations string set for the requested category
*/
func (a *Client) TranslationsGet(params *TranslationsGetParams, authInfo runtime.ClientAuthInfoWriter) (*TranslationsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTranslationsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TranslationsGet",
		Method:             "GET",
		PathPattern:        "/translations",
		ProducesMediaTypes: []string{"application/json", "json/plain"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TranslationsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TranslationsGetOK), nil

}

/*
TranslationsMissingPost complains about the missing translation string
*/
func (a *Client) TranslationsMissingPost(params *TranslationsMissingPostParams, authInfo runtime.ClientAuthInfoWriter) (*TranslationsMissingPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTranslationsMissingPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TranslationsMissingPost",
		Method:             "POST",
		PathPattern:        "/translations/missing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TranslationsMissingPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TranslationsMissingPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
