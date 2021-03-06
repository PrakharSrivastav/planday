// Code generated by go-swagger; DO NOT EDIT.

package gdpr

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

// NewGdprPortalsByPortalAliasGdprStatusGetParams creates a new GdprPortalsByPortalAliasGdprStatusGetParams object
// with the default values initialized.
func NewGdprPortalsByPortalAliasGdprStatusGetParams() *GdprPortalsByPortalAliasGdprStatusGetParams {
	var ()
	return &GdprPortalsByPortalAliasGdprStatusGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGdprPortalsByPortalAliasGdprStatusGetParamsWithTimeout creates a new GdprPortalsByPortalAliasGdprStatusGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGdprPortalsByPortalAliasGdprStatusGetParamsWithTimeout(timeout time.Duration) *GdprPortalsByPortalAliasGdprStatusGetParams {
	var ()
	return &GdprPortalsByPortalAliasGdprStatusGetParams{

		timeout: timeout,
	}
}

// NewGdprPortalsByPortalAliasGdprStatusGetParamsWithContext creates a new GdprPortalsByPortalAliasGdprStatusGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGdprPortalsByPortalAliasGdprStatusGetParamsWithContext(ctx context.Context) *GdprPortalsByPortalAliasGdprStatusGetParams {
	var ()
	return &GdprPortalsByPortalAliasGdprStatusGetParams{

		Context: ctx,
	}
}

// NewGdprPortalsByPortalAliasGdprStatusGetParamsWithHTTPClient creates a new GdprPortalsByPortalAliasGdprStatusGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGdprPortalsByPortalAliasGdprStatusGetParamsWithHTTPClient(client *http.Client) *GdprPortalsByPortalAliasGdprStatusGetParams {
	var ()
	return &GdprPortalsByPortalAliasGdprStatusGetParams{
		HTTPClient: client,
	}
}

/*GdprPortalsByPortalAliasGdprStatusGetParams contains all the parameters to send to the API endpoint
for the gdpr portals by portal alias gdpr status get operation typically these are written to a http.Request
*/
type GdprPortalsByPortalAliasGdprStatusGetParams struct {

	/*PortalAlias*/
	PortalAlias string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) WithTimeout(timeout time.Duration) *GdprPortalsByPortalAliasGdprStatusGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) WithContext(ctx context.Context) *GdprPortalsByPortalAliasGdprStatusGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) WithHTTPClient(client *http.Client) *GdprPortalsByPortalAliasGdprStatusGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPortalAlias adds the portalAlias to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) WithPortalAlias(portalAlias string) *GdprPortalsByPortalAliasGdprStatusGetParams {
	o.SetPortalAlias(portalAlias)
	return o
}

// SetPortalAlias adds the portalAlias to the gdpr portals by portal alias gdpr status get params
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) SetPortalAlias(portalAlias string) {
	o.PortalAlias = portalAlias
}

// WriteToRequest writes these params to a swagger request
func (o *GdprPortalsByPortalAliasGdprStatusGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param portal_alias
	if err := r.SetPathParam("portal_alias", o.PortalAlias); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
