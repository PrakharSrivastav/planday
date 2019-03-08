// Code generated by go-swagger; DO NOT EDIT.

package portal_employee_form_definitions

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

// NewPortalsByPortalIDFormsEmployeePutParams creates a new PortalsByPortalIDFormsEmployeePutParams object
// with the default values initialized.
func NewPortalsByPortalIDFormsEmployeePutParams() *PortalsByPortalIDFormsEmployeePutParams {
	var (
		isNewDefault = bool(false)
	)
	return &PortalsByPortalIDFormsEmployeePutParams{
		IsNew: &isNewDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPortalsByPortalIDFormsEmployeePutParamsWithTimeout creates a new PortalsByPortalIDFormsEmployeePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPortalsByPortalIDFormsEmployeePutParamsWithTimeout(timeout time.Duration) *PortalsByPortalIDFormsEmployeePutParams {
	var (
		isNewDefault = bool(false)
	)
	return &PortalsByPortalIDFormsEmployeePutParams{
		IsNew: &isNewDefault,

		timeout: timeout,
	}
}

// NewPortalsByPortalIDFormsEmployeePutParamsWithContext creates a new PortalsByPortalIDFormsEmployeePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPortalsByPortalIDFormsEmployeePutParamsWithContext(ctx context.Context) *PortalsByPortalIDFormsEmployeePutParams {
	var (
		isNewDefault = bool(false)
	)
	return &PortalsByPortalIDFormsEmployeePutParams{
		IsNew: &isNewDefault,

		Context: ctx,
	}
}

// NewPortalsByPortalIDFormsEmployeePutParamsWithHTTPClient creates a new PortalsByPortalIDFormsEmployeePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPortalsByPortalIDFormsEmployeePutParamsWithHTTPClient(client *http.Client) *PortalsByPortalIDFormsEmployeePutParams {
	var (
		isNewDefault = bool(false)
	)
	return &PortalsByPortalIDFormsEmployeePutParams{
		IsNew:      &isNewDefault,
		HTTPClient: client,
	}
}

/*PortalsByPortalIDFormsEmployeePutParams contains all the parameters to send to the API endpoint
for the portals by portal id forms employee put operation typically these are written to a http.Request
*/
type PortalsByPortalIDFormsEmployeePutParams struct {

	/*IsNew*/
	IsNew *bool
	/*Model*/
	Model *models.FormDefinitionInputModel
	/*PortalID*/
	PortalID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) WithTimeout(timeout time.Duration) *PortalsByPortalIDFormsEmployeePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) WithContext(ctx context.Context) *PortalsByPortalIDFormsEmployeePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) WithHTTPClient(client *http.Client) *PortalsByPortalIDFormsEmployeePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsNew adds the isNew to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) WithIsNew(isNew *bool) *PortalsByPortalIDFormsEmployeePutParams {
	o.SetIsNew(isNew)
	return o
}

// SetIsNew adds the isNew to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) SetIsNew(isNew *bool) {
	o.IsNew = isNew
}

// WithModel adds the model to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) WithModel(model *models.FormDefinitionInputModel) *PortalsByPortalIDFormsEmployeePutParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) SetModel(model *models.FormDefinitionInputModel) {
	o.Model = model
}

// WithPortalID adds the portalID to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) WithPortalID(portalID int64) *PortalsByPortalIDFormsEmployeePutParams {
	o.SetPortalID(portalID)
	return o
}

// SetPortalID adds the portalId to the portals by portal id forms employee put params
func (o *PortalsByPortalIDFormsEmployeePutParams) SetPortalID(portalID int64) {
	o.PortalID = portalID
}

// WriteToRequest writes these params to a swagger request
func (o *PortalsByPortalIDFormsEmployeePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsNew != nil {

		// query param is_new
		var qrIsNew bool
		if o.IsNew != nil {
			qrIsNew = *o.IsNew
		}
		qIsNew := swag.FormatBool(qrIsNew)
		if qIsNew != "" {
			if err := r.SetQueryParam("is_new", qIsNew); err != nil {
				return err
			}
		}

	}

	if o.Model != nil {
		if err := r.SetBodyParam(o.Model); err != nil {
			return err
		}
	}

	// path param portal_id
	if err := r.SetPathParam("portal_id", swag.FormatInt64(o.PortalID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}