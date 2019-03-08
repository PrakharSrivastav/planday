// Code generated by go-swagger; DO NOT EDIT.

package employees_utilities

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

// NewEmployeesContactsCountsGetParams creates a new EmployeesContactsCountsGetParams object
// with the default values initialized.
func NewEmployeesContactsCountsGetParams() *EmployeesContactsCountsGetParams {
	var ()
	return &EmployeesContactsCountsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeesContactsCountsGetParamsWithTimeout creates a new EmployeesContactsCountsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeesContactsCountsGetParamsWithTimeout(timeout time.Duration) *EmployeesContactsCountsGetParams {
	var ()
	return &EmployeesContactsCountsGetParams{

		timeout: timeout,
	}
}

// NewEmployeesContactsCountsGetParamsWithContext creates a new EmployeesContactsCountsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeesContactsCountsGetParamsWithContext(ctx context.Context) *EmployeesContactsCountsGetParams {
	var ()
	return &EmployeesContactsCountsGetParams{

		Context: ctx,
	}
}

// NewEmployeesContactsCountsGetParamsWithHTTPClient creates a new EmployeesContactsCountsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeesContactsCountsGetParamsWithHTTPClient(client *http.Client) *EmployeesContactsCountsGetParams {
	var ()
	return &EmployeesContactsCountsGetParams{
		HTTPClient: client,
	}
}

/*EmployeesContactsCountsGetParams contains all the parameters to send to the API endpoint
for the employees contacts counts get operation typically these are written to a http.Request
*/
type EmployeesContactsCountsGetParams struct {

	/*DepartmentID*/
	DepartmentID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) WithTimeout(timeout time.Duration) *EmployeesContactsCountsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) WithContext(ctx context.Context) *EmployeesContactsCountsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) WithHTTPClient(client *http.Client) *EmployeesContactsCountsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDepartmentID adds the departmentID to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) WithDepartmentID(departmentID *int64) *EmployeesContactsCountsGetParams {
	o.SetDepartmentID(departmentID)
	return o
}

// SetDepartmentID adds the departmentId to the employees contacts counts get params
func (o *EmployeesContactsCountsGetParams) SetDepartmentID(departmentID *int64) {
	o.DepartmentID = departmentID
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeesContactsCountsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DepartmentID != nil {

		// query param department_id
		var qrDepartmentID int64
		if o.DepartmentID != nil {
			qrDepartmentID = *o.DepartmentID
		}
		qDepartmentID := swag.FormatInt64(qrDepartmentID)
		if qDepartmentID != "" {
			if err := r.SetQueryParam("department_id", qDepartmentID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
