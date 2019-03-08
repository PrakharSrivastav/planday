// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProfilePostParams creates a new ProfilePostParams object
// with the default values initialized.
func NewProfilePostParams() *ProfilePostParams {
	var (
		ignoreWarningsDefault = bool(false)
	)
	return &ProfilePostParams{
		IgnoreWarnings: &ignoreWarningsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewProfilePostParamsWithTimeout creates a new ProfilePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProfilePostParamsWithTimeout(timeout time.Duration) *ProfilePostParams {
	var (
		ignoreWarningsDefault = bool(false)
	)
	return &ProfilePostParams{
		IgnoreWarnings: &ignoreWarningsDefault,

		timeout: timeout,
	}
}

// NewProfilePostParamsWithContext creates a new ProfilePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewProfilePostParamsWithContext(ctx context.Context) *ProfilePostParams {
	var (
		ignoreWarningsDefault = bool(false)
	)
	return &ProfilePostParams{
		IgnoreWarnings: &ignoreWarningsDefault,

		Context: ctx,
	}
}

// NewProfilePostParamsWithHTTPClient creates a new ProfilePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProfilePostParamsWithHTTPClient(client *http.Client) *ProfilePostParams {
	var (
		ignoreWarningsDefault = bool(false)
	)
	return &ProfilePostParams{
		IgnoreWarnings: &ignoreWarningsDefault,
		HTTPClient:     client,
	}
}

/*ProfilePostParams contains all the parameters to send to the API endpoint
for the profile post operation typically these are written to a http.Request
*/
type ProfilePostParams struct {

	/*IgnoreWarnings
	  if set to true, validation warnings will be ignored

	*/
	IgnoreWarnings *bool
	/*UpdateModel*/
	UpdateModel interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the profile post params
func (o *ProfilePostParams) WithTimeout(timeout time.Duration) *ProfilePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the profile post params
func (o *ProfilePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the profile post params
func (o *ProfilePostParams) WithContext(ctx context.Context) *ProfilePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the profile post params
func (o *ProfilePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the profile post params
func (o *ProfilePostParams) WithHTTPClient(client *http.Client) *ProfilePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the profile post params
func (o *ProfilePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIgnoreWarnings adds the ignoreWarnings to the profile post params
func (o *ProfilePostParams) WithIgnoreWarnings(ignoreWarnings *bool) *ProfilePostParams {
	o.SetIgnoreWarnings(ignoreWarnings)
	return o
}

// SetIgnoreWarnings adds the ignoreWarnings to the profile post params
func (o *ProfilePostParams) SetIgnoreWarnings(ignoreWarnings *bool) {
	o.IgnoreWarnings = ignoreWarnings
}

// WithUpdateModel adds the updateModel to the profile post params
func (o *ProfilePostParams) WithUpdateModel(updateModel interface{}) *ProfilePostParams {
	o.SetUpdateModel(updateModel)
	return o
}

// SetUpdateModel adds the updateModel to the profile post params
func (o *ProfilePostParams) SetUpdateModel(updateModel interface{}) {
	o.UpdateModel = updateModel
}

// WriteToRequest writes these params to a swagger request
func (o *ProfilePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IgnoreWarnings != nil {

		// query param ignore_warnings
		var qrIgnoreWarnings bool
		if o.IgnoreWarnings != nil {
			qrIgnoreWarnings = *o.IgnoreWarnings
		}
		qIgnoreWarnings := swag.FormatBool(qrIgnoreWarnings)
		if qIgnoreWarnings != "" {
			if err := r.SetQueryParam("ignore_warnings", qIgnoreWarnings); err != nil {
				return err
			}
		}

	}

	if o.UpdateModel != nil {
		if err := r.SetBodyParam(o.UpdateModel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
