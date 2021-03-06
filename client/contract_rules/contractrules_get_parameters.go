// Code generated by go-swagger; DO NOT EDIT.

package contract_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewContractrulesGetParams creates a new ContractrulesGetParams object
// with the default values initialized.
func NewContractrulesGetParams() *ContractrulesGetParams {

	return &ContractrulesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContractrulesGetParamsWithTimeout creates a new ContractrulesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContractrulesGetParamsWithTimeout(timeout time.Duration) *ContractrulesGetParams {

	return &ContractrulesGetParams{

		timeout: timeout,
	}
}

// NewContractrulesGetParamsWithContext creates a new ContractrulesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewContractrulesGetParamsWithContext(ctx context.Context) *ContractrulesGetParams {

	return &ContractrulesGetParams{

		Context: ctx,
	}
}

// NewContractrulesGetParamsWithHTTPClient creates a new ContractrulesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContractrulesGetParamsWithHTTPClient(client *http.Client) *ContractrulesGetParams {

	return &ContractrulesGetParams{
		HTTPClient: client,
	}
}

/*ContractrulesGetParams contains all the parameters to send to the API endpoint
for the contractrules get operation typically these are written to a http.Request
*/
type ContractrulesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contractrules get params
func (o *ContractrulesGetParams) WithTimeout(timeout time.Duration) *ContractrulesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contractrules get params
func (o *ContractrulesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contractrules get params
func (o *ContractrulesGetParams) WithContext(ctx context.Context) *ContractrulesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contractrules get params
func (o *ContractrulesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contractrules get params
func (o *ContractrulesGetParams) WithHTTPClient(client *http.Client) *ContractrulesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contractrules get params
func (o *ContractrulesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ContractrulesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
