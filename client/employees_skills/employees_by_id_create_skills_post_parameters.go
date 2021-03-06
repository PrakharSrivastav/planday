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

// NewEmployeesByIDCreateSkillsPostParams creates a new EmployeesByIDCreateSkillsPostParams object
// with the default values initialized.
func NewEmployeesByIDCreateSkillsPostParams() *EmployeesByIDCreateSkillsPostParams {
	var ()
	return &EmployeesByIDCreateSkillsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeesByIDCreateSkillsPostParamsWithTimeout creates a new EmployeesByIDCreateSkillsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeesByIDCreateSkillsPostParamsWithTimeout(timeout time.Duration) *EmployeesByIDCreateSkillsPostParams {
	var ()
	return &EmployeesByIDCreateSkillsPostParams{

		timeout: timeout,
	}
}

// NewEmployeesByIDCreateSkillsPostParamsWithContext creates a new EmployeesByIDCreateSkillsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeesByIDCreateSkillsPostParamsWithContext(ctx context.Context) *EmployeesByIDCreateSkillsPostParams {
	var ()
	return &EmployeesByIDCreateSkillsPostParams{

		Context: ctx,
	}
}

// NewEmployeesByIDCreateSkillsPostParamsWithHTTPClient creates a new EmployeesByIDCreateSkillsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeesByIDCreateSkillsPostParamsWithHTTPClient(client *http.Client) *EmployeesByIDCreateSkillsPostParams {
	var ()
	return &EmployeesByIDCreateSkillsPostParams{
		HTTPClient: client,
	}
}

/*EmployeesByIDCreateSkillsPostParams contains all the parameters to send to the API endpoint
for the employees by Id create skills post operation typically these are written to a http.Request
*/
type EmployeesByIDCreateSkillsPostParams struct {

	/*ID*/
	ID int64
	/*SkillIds*/
	SkillIds []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) WithTimeout(timeout time.Duration) *EmployeesByIDCreateSkillsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) WithContext(ctx context.Context) *EmployeesByIDCreateSkillsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) WithHTTPClient(client *http.Client) *EmployeesByIDCreateSkillsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) WithID(id int64) *EmployeesByIDCreateSkillsPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) SetID(id int64) {
	o.ID = id
}

// WithSkillIds adds the skillIds to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) WithSkillIds(skillIds []int64) *EmployeesByIDCreateSkillsPostParams {
	o.SetSkillIds(skillIds)
	return o
}

// SetSkillIds adds the skillIds to the employees by Id create skills post params
func (o *EmployeesByIDCreateSkillsPostParams) SetSkillIds(skillIds []int64) {
	o.SkillIds = skillIds
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeesByIDCreateSkillsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
