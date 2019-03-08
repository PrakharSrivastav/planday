// Code generated by go-swagger; DO NOT EDIT.

package translations

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

// NewTranslationsMissingPostParams creates a new TranslationsMissingPostParams object
// with the default values initialized.
func NewTranslationsMissingPostParams() *TranslationsMissingPostParams {
	var ()
	return &TranslationsMissingPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTranslationsMissingPostParamsWithTimeout creates a new TranslationsMissingPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTranslationsMissingPostParamsWithTimeout(timeout time.Duration) *TranslationsMissingPostParams {
	var ()
	return &TranslationsMissingPostParams{

		timeout: timeout,
	}
}

// NewTranslationsMissingPostParamsWithContext creates a new TranslationsMissingPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewTranslationsMissingPostParamsWithContext(ctx context.Context) *TranslationsMissingPostParams {
	var ()
	return &TranslationsMissingPostParams{

		Context: ctx,
	}
}

// NewTranslationsMissingPostParamsWithHTTPClient creates a new TranslationsMissingPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTranslationsMissingPostParamsWithHTTPClient(client *http.Client) *TranslationsMissingPostParams {
	var ()
	return &TranslationsMissingPostParams{
		HTTPClient: client,
	}
}

/*TranslationsMissingPostParams contains all the parameters to send to the API endpoint
for the translations missing post operation typically these are written to a http.Request
*/
type TranslationsMissingPostParams struct {

	/*Culture*/
	Culture *string
	/*Key*/
	Key *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the translations missing post params
func (o *TranslationsMissingPostParams) WithTimeout(timeout time.Duration) *TranslationsMissingPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the translations missing post params
func (o *TranslationsMissingPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the translations missing post params
func (o *TranslationsMissingPostParams) WithContext(ctx context.Context) *TranslationsMissingPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the translations missing post params
func (o *TranslationsMissingPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the translations missing post params
func (o *TranslationsMissingPostParams) WithHTTPClient(client *http.Client) *TranslationsMissingPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the translations missing post params
func (o *TranslationsMissingPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCulture adds the culture to the translations missing post params
func (o *TranslationsMissingPostParams) WithCulture(culture *string) *TranslationsMissingPostParams {
	o.SetCulture(culture)
	return o
}

// SetCulture adds the culture to the translations missing post params
func (o *TranslationsMissingPostParams) SetCulture(culture *string) {
	o.Culture = culture
}

// WithKey adds the key to the translations missing post params
func (o *TranslationsMissingPostParams) WithKey(key *string) *TranslationsMissingPostParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the translations missing post params
func (o *TranslationsMissingPostParams) SetKey(key *string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *TranslationsMissingPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Culture != nil {

		// query param culture
		var qrCulture string
		if o.Culture != nil {
			qrCulture = *o.Culture
		}
		qCulture := qrCulture
		if qCulture != "" {
			if err := r.SetQueryParam("culture", qCulture); err != nil {
				return err
			}
		}

	}

	if o.Key != nil {

		// query param key
		var qrKey string
		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {
			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
