// Code generated by go-swagger; DO NOT EDIT.

package employees_user_accounts

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

// NewEmployeesByIDSendSmsWelcomeMessagePutParams creates a new EmployeesByIDSendSmsWelcomeMessagePutParams object
// with the default values initialized.
func NewEmployeesByIDSendSmsWelcomeMessagePutParams() *EmployeesByIDSendSmsWelcomeMessagePutParams {
	var ()
	return &EmployeesByIDSendSmsWelcomeMessagePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeesByIDSendSmsWelcomeMessagePutParamsWithTimeout creates a new EmployeesByIDSendSmsWelcomeMessagePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeesByIDSendSmsWelcomeMessagePutParamsWithTimeout(timeout time.Duration) *EmployeesByIDSendSmsWelcomeMessagePutParams {
	var ()
	return &EmployeesByIDSendSmsWelcomeMessagePutParams{

		timeout: timeout,
	}
}

// NewEmployeesByIDSendSmsWelcomeMessagePutParamsWithContext creates a new EmployeesByIDSendSmsWelcomeMessagePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeesByIDSendSmsWelcomeMessagePutParamsWithContext(ctx context.Context) *EmployeesByIDSendSmsWelcomeMessagePutParams {
	var ()
	return &EmployeesByIDSendSmsWelcomeMessagePutParams{

		Context: ctx,
	}
}

// NewEmployeesByIDSendSmsWelcomeMessagePutParamsWithHTTPClient creates a new EmployeesByIDSendSmsWelcomeMessagePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeesByIDSendSmsWelcomeMessagePutParamsWithHTTPClient(client *http.Client) *EmployeesByIDSendSmsWelcomeMessagePutParams {
	var ()
	return &EmployeesByIDSendSmsWelcomeMessagePutParams{
		HTTPClient: client,
	}
}

/*EmployeesByIDSendSmsWelcomeMessagePutParams contains all the parameters to send to the API endpoint
for the employees by Id send sms welcome message put operation typically these are written to a http.Request
*/
type EmployeesByIDSendSmsWelcomeMessagePutParams struct {

	/*ID
	  The id of employee to send the welcome SMS

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) WithTimeout(timeout time.Duration) *EmployeesByIDSendSmsWelcomeMessagePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) WithContext(ctx context.Context) *EmployeesByIDSendSmsWelcomeMessagePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) WithHTTPClient(client *http.Client) *EmployeesByIDSendSmsWelcomeMessagePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) WithID(id int64) *EmployeesByIDSendSmsWelcomeMessagePutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employees by Id send sms welcome message put params
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeesByIDSendSmsWelcomeMessagePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
