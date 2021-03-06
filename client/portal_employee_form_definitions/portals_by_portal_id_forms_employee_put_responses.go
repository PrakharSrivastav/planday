// Code generated by go-swagger; DO NOT EDIT.

package portal_employee_form_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PortalsByPortalIDFormsEmployeePutReader is a Reader for the PortalsByPortalIDFormsEmployeePut structure.
type PortalsByPortalIDFormsEmployeePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortalsByPortalIDFormsEmployeePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPortalsByPortalIDFormsEmployeePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPortalsByPortalIDFormsEmployeePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPortalsByPortalIDFormsEmployeePutOK creates a PortalsByPortalIDFormsEmployeePutOK with default headers values
func NewPortalsByPortalIDFormsEmployeePutOK() *PortalsByPortalIDFormsEmployeePutOK {
	return &PortalsByPortalIDFormsEmployeePutOK{}
}

/*PortalsByPortalIDFormsEmployeePutOK handles this case with default header values.

Success
*/
type PortalsByPortalIDFormsEmployeePutOK struct {
	Payload bool
}

func (o *PortalsByPortalIDFormsEmployeePutOK) Error() string {
	return fmt.Sprintf("[PUT /portals/{portal_id}/forms/employee][%d] portalsByPortalIdFormsEmployeePutOK  %+v", 200, o.Payload)
}

func (o *PortalsByPortalIDFormsEmployeePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortalsByPortalIDFormsEmployeePutUnauthorized creates a PortalsByPortalIDFormsEmployeePutUnauthorized with default headers values
func NewPortalsByPortalIDFormsEmployeePutUnauthorized() *PortalsByPortalIDFormsEmployeePutUnauthorized {
	return &PortalsByPortalIDFormsEmployeePutUnauthorized{}
}

/*PortalsByPortalIDFormsEmployeePutUnauthorized handles this case with default header values.

Unauthorized
*/
type PortalsByPortalIDFormsEmployeePutUnauthorized struct {
}

func (o *PortalsByPortalIDFormsEmployeePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /portals/{portal_id}/forms/employee][%d] portalsByPortalIdFormsEmployeePutUnauthorized ", 401)
}

func (o *PortalsByPortalIDFormsEmployeePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
