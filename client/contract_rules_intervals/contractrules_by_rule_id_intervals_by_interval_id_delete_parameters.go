// Code generated by go-swagger; DO NOT EDIT.

package contract_rules_intervals

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

// NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParams creates a new ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams object
// with the default values initialized.
func NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParams() *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	var ()
	return &ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParamsWithTimeout creates a new ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParamsWithTimeout(timeout time.Duration) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	var ()
	return &ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams{

		timeout: timeout,
	}
}

// NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParamsWithContext creates a new ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParamsWithContext(ctx context.Context) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	var ()
	return &ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams{

		Context: ctx,
	}
}

// NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParamsWithHTTPClient creates a new ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContractrulesByRuleIDIntervalsByIntervalIDDeleteParamsWithHTTPClient(client *http.Client) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	var ()
	return &ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams{
		HTTPClient: client,
	}
}

/*ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams contains all the parameters to send to the API endpoint
for the contractrules by rule Id intervals by interval Id delete operation typically these are written to a http.Request
*/
type ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams struct {

	/*IntervalID*/
	IntervalID int64
	/*RuleID*/
	RuleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) WithTimeout(timeout time.Duration) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) WithContext(ctx context.Context) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) WithHTTPClient(client *http.Client) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntervalID adds the intervalID to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) WithIntervalID(intervalID int64) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	o.SetIntervalID(intervalID)
	return o
}

// SetIntervalID adds the intervalId to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) SetIntervalID(intervalID int64) {
	o.IntervalID = intervalID
}

// WithRuleID adds the ruleID to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) WithRuleID(ruleID int64) *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the contractrules by rule Id intervals by interval Id delete params
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) SetRuleID(ruleID int64) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param intervalId
	if err := r.SetPathParam("intervalId", swag.FormatInt64(o.IntervalID)); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", swag.FormatInt64(o.RuleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}