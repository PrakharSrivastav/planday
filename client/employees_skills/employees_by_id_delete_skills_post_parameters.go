// Code generated by go-swagger; DO NOT EDIT.

package employees_skills

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

// NewEmployeesByIDDeleteSkillsPostParams creates a new EmployeesByIDDeleteSkillsPostParams object
// with the default values initialized.
func NewEmployeesByIDDeleteSkillsPostParams() *EmployeesByIDDeleteSkillsPostParams {
	var ()
	return &EmployeesByIDDeleteSkillsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeesByIDDeleteSkillsPostParamsWithTimeout creates a new EmployeesByIDDeleteSkillsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeesByIDDeleteSkillsPostParamsWithTimeout(timeout time.Duration) *EmployeesByIDDeleteSkillsPostParams {
	var ()
	return &EmployeesByIDDeleteSkillsPostParams{

		timeout: timeout,
	}
}

// NewEmployeesByIDDeleteSkillsPostParamsWithContext creates a new EmployeesByIDDeleteSkillsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeesByIDDeleteSkillsPostParamsWithContext(ctx context.Context) *EmployeesByIDDeleteSkillsPostParams {
	var ()
	return &EmployeesByIDDeleteSkillsPostParams{

		Context: ctx,
	}
}

// NewEmployeesByIDDeleteSkillsPostParamsWithHTTPClient creates a new EmployeesByIDDeleteSkillsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeesByIDDeleteSkillsPostParamsWithHTTPClient(client *http.Client) *EmployeesByIDDeleteSkillsPostParams {
	var ()
	return &EmployeesByIDDeleteSkillsPostParams{
		HTTPClient: client,
	}
}

/*EmployeesByIDDeleteSkillsPostParams contains all the parameters to send to the API endpoint
for the employees by Id delete skills post operation typically these are written to a http.Request
*/
type EmployeesByIDDeleteSkillsPostParams struct {

	/*ID*/
	ID int64
	/*SkillIds*/
	SkillIds []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) WithTimeout(timeout time.Duration) *EmployeesByIDDeleteSkillsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) WithContext(ctx context.Context) *EmployeesByIDDeleteSkillsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) WithHTTPClient(client *http.Client) *EmployeesByIDDeleteSkillsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) WithID(id int64) *EmployeesByIDDeleteSkillsPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) SetID(id int64) {
	o.ID = id
}

// WithSkillIds adds the skillIds to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) WithSkillIds(skillIds []int64) *EmployeesByIDDeleteSkillsPostParams {
	o.SetSkillIds(skillIds)
	return o
}

// SetSkillIds adds the skillIds to the employees by Id delete skills post params
func (o *EmployeesByIDDeleteSkillsPostParams) SetSkillIds(skillIds []int64) {
	o.SkillIds = skillIds
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeesByIDDeleteSkillsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.SkillIds != nil {
		if err := r.SetBodyParam(o.SkillIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}