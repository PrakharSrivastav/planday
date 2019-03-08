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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// NewContractrulesByIDPutParams creates a new ContractrulesByIDPutParams object
// with the default values initialized.
func NewContractrulesByIDPutParams() *ContractrulesByIDPutParams {
	var ()
	return &ContractrulesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContractrulesByIDPutParamsWithTimeout creates a new ContractrulesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContractrulesByIDPutParamsWithTimeout(timeout time.Duration) *ContractrulesByIDPutParams {
	var ()
	return &ContractrulesByIDPutParams{

		timeout: timeout,
	}
}

// NewContractrulesByIDPutParamsWithContext creates a new ContractrulesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewContractrulesByIDPutParamsWithContext(ctx context.Context) *ContractrulesByIDPutParams {
	var ()
	return &ContractrulesByIDPutParams{

		Context: ctx,
	}
}

// NewContractrulesByIDPutParamsWithHTTPClient creates a new ContractrulesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContractrulesByIDPutParamsWithHTTPClient(client *http.Client) *ContractrulesByIDPutParams {
	var ()
	return &ContractrulesByIDPutParams{
		HTTPClient: client,
	}
}

/*ContractrulesByIDPutParams contains all the parameters to send to the API endpoint
for the contractrules by Id put operation typically these are written to a http.Request
*/
type ContractrulesByIDPutParams struct {

	/*ID*/
	ID int64
	/*UpdateModel*/
	UpdateModel *models.CreateUpdateContractRulesRuleViewModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) WithTimeout(timeout time.Duration) *ContractrulesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) WithContext(ctx context.Context) *ContractrulesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) WithHTTPClient(client *http.Client) *ContractrulesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) WithID(id int64) *ContractrulesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) SetID(id int64) {
	o.ID = id
}

// WithUpdateModel adds the updateModel to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) WithUpdateModel(updateModel *models.CreateUpdateContractRulesRuleViewModel) *ContractrulesByIDPutParams {
	o.SetUpdateModel(updateModel)
	return o
}

// SetUpdateModel adds the updateModel to the contractrules by Id put params
func (o *ContractrulesByIDPutParams) SetUpdateModel(updateModel *models.CreateUpdateContractRulesRuleViewModel) {
	o.UpdateModel = updateModel
}

// WriteToRequest writes these params to a swagger request
func (o *ContractrulesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
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