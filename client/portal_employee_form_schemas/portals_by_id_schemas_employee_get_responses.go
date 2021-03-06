// Code generated by go-swagger; DO NOT EDIT.

package portal_employee_form_schemas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PortalsByIDSchemasEmployeeGetReader is a Reader for the PortalsByIDSchemasEmployeeGet structure.
type PortalsByIDSchemasEmployeeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortalsByIDSchemasEmployeeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 401:
		result := NewPortalsByIDSchemasEmployeeGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPortalsByIDSchemasEmployeeGetUnauthorized creates a PortalsByIDSchemasEmployeeGetUnauthorized with default headers values
func NewPortalsByIDSchemasEmployeeGetUnauthorized() *PortalsByIDSchemasEmployeeGetUnauthorized {
	return &PortalsByIDSchemasEmployeeGetUnauthorized{}
}

/*PortalsByIDSchemasEmployeeGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PortalsByIDSchemasEmployeeGetUnauthorized struct {
}

func (o *PortalsByIDSchemasEmployeeGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /portals/{id}/schemas/employee][%d] portalsByIdSchemasEmployeeGetUnauthorized ", 401)
}

func (o *PortalsByIDSchemasEmployeeGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
