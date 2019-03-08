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

// NewContractrulesByRuleIDIntervalsGetParams creates a new ContractrulesByRuleIDIntervalsGetParams object
// with the default values initialized.
func NewContractrulesByRuleIDIntervalsGetParams() *ContractrulesByRuleIDIntervalsGetParams {
	var ()
	return &ContractrulesByRuleIDIntervalsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContractrulesByRuleIDIntervalsGetParamsWithTimeout creates a new ContractrulesByRuleIDIntervalsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContractrulesByRuleIDIntervalsGetParamsWithTimeout(timeout time.Duration) *ContractrulesByRuleIDIntervalsGetParams {
	var ()
	return &ContractrulesByRuleIDIntervalsGetParams{

		timeout: timeout,
	}
}

// NewContractrulesByRuleIDIntervalsGetParamsWithContext creates a new ContractrulesByRuleIDIntervalsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewContractrulesByRuleIDIntervalsGetParamsWithContext(ctx context.Context) *ContractrulesByRuleIDIntervalsGetParams {
	var ()
	return &ContractrulesByRuleIDIntervalsGetParams{

		Context: ctx,
	}
}

// NewContractrulesByRuleIDIntervalsGetParamsWithHTTPClient creates a new ContractrulesByRuleIDIntervalsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContractrulesByRuleIDIntervalsGetParamsWithHTTPClient(client *http.Client) *ContractrulesByRuleIDIntervalsGetParams {
	var ()
	return &ContractrulesByRuleIDIntervalsGetParams{
		HTTPClient: client,
	}
}

/*ContractrulesByRuleIDIntervalsGetParams contains all the parameters to send to the API endpoint
for the contractrules by rule Id intervals get operation typically these are written to a http.Request
*/
type ContractrulesByRuleIDIntervalsGetParams struct {

	/*RuleID*/
	RuleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) WithTimeout(timeout time.Duration) *ContractrulesByRuleIDIntervalsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) WithContext(ctx context.Context) *ContractrulesByRuleIDIntervalsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) WithHTTPClient(client *http.Client) *ContractrulesByRuleIDIntervalsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) WithRuleID(ruleID int64) *ContractrulesByRuleIDIntervalsGetParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the contractrules by rule Id intervals get params
func (o *ContractrulesByRuleIDIntervalsGetParams) SetRuleID(ruleID int64) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *ContractrulesByRuleIDIntervalsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleId
	if err := r.SetPathParam("ruleId", swag.FormatInt64(o.RuleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}