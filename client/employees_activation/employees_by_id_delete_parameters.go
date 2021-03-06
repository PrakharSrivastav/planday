// Code generated by go-swagger; DO NOT EDIT.

package employees_activation

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

// NewEmployeesByIDDeleteParams creates a new EmployeesByIDDeleteParams object
// with the default values initialized.
func NewEmployeesByIDDeleteParams() *EmployeesByIDDeleteParams {
	var (
		reasonDefault = string("")
	)
	return &EmployeesByIDDeleteParams{
		Reason: &reasonDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeesByIDDeleteParamsWithTimeout creates a new EmployeesByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeesByIDDeleteParamsWithTimeout(timeout time.Duration) *EmployeesByIDDeleteParams {
	var (
		reasonDefault = string("")
	)
	return &EmployeesByIDDeleteParams{
		Reason: &reasonDefault,

		timeout: timeout,
	}
}

// NewEmployeesByIDDeleteParamsWithContext creates a new EmployeesByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeesByIDDeleteParamsWithContext(ctx context.Context) *EmployeesByIDDeleteParams {
	var (
		reasonDefault = string("")
	)
	return &EmployeesByIDDeleteParams{
		Reason: &reasonDefault,

		Context: ctx,
	}
}

// NewEmployeesByIDDeleteParamsWithHTTPClient creates a new EmployeesByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeesByIDDeleteParamsWithHTTPClient(client *http.Client) *EmployeesByIDDeleteParams {
	var (
		reasonDefault = string("")
	)
	return &EmployeesByIDDeleteParams{
		Reason:     &reasonDefault,
		HTTPClient: client,
	}
}

/*EmployeesByIDDeleteParams contains all the parameters to send to the API endpoint
for the employees by Id delete operation typically these are written to a http.Request
*/
type EmployeesByIDDeleteParams struct {

	/*Date
	  The last date (in ISO-8601 date format 'YYYY-MM-DD') where the user is active. If {null}, the user will be deactivated immediately using the portal's current date.

	*/
	Date *string
	/*ID
	  The id of the employee that should be deactivated.

	*/
	ID int32
	/*KeepShifts
	  {true}, if the employee's shifts after the deactivation date should be set to open; {false}, the shifts should be leaft unchanged.

	*/
	KeepShifts *bool
	/*Reason
	  optional reason for deactivating

	*/
	Reason *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) WithTimeout(timeout time.Duration) *EmployeesByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) WithContext(ctx context.Context) *EmployeesByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) WithHTTPClient(client *http.Client) *EmployeesByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) WithDate(date *string) *EmployeesByIDDeleteParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) SetDate(date *string) {
	o.Date = date
}

// WithID adds the id to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) WithID(id int32) *EmployeesByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithKeepShifts adds the keepShifts to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) WithKeepShifts(keepShifts *bool) *EmployeesByIDDeleteParams {
	o.SetKeepShifts(keepShifts)
	return o
}

// SetKeepShifts adds the keepShifts to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) SetKeepShifts(keepShifts *bool) {
	o.KeepShifts = keepShifts
}

// WithReason adds the reason to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) WithReason(reason *string) *EmployeesByIDDeleteParams {
	o.SetReason(reason)
	return o
}

// SetReason adds the reason to the employees by Id delete params
func (o *EmployeesByIDDeleteParams) SetReason(reason *string) {
	o.Reason = reason
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeesByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Date != nil {

		// query param date
		var qrDate string
		if o.Date != nil {
			qrDate = *o.Date
		}
		qDate := qrDate
		if qDate != "" {
			if err := r.SetQueryParam("date", qDate); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.KeepShifts != nil {

		// query param keep_shifts
		var qrKeepShifts bool
		if o.KeepShifts != nil {
			qrKeepShifts = *o.KeepShifts
		}
		qKeepShifts := swag.FormatBool(qrKeepShifts)
		if qKeepShifts != "" {
			if err := r.SetQueryParam("keep_shifts", qKeepShifts); err != nil {
				return err
			}
		}

	}

	if o.Reason != nil {

		// query param reason
		var qrReason string
		if o.Reason != nil {
			qrReason = *o.Reason
		}
		qReason := qrReason
		if qReason != "" {
			if err := r.SetQueryParam("reason", qReason); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
