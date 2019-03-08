// Code generated by go-swagger; DO NOT EDIT.

package employees_utilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// EmployeesContactsCountsGetReader is a Reader for the EmployeesContactsCountsGet structure.
type EmployeesContactsCountsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeesContactsCountsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmployeesContactsCountsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEmployeesContactsCountsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeesContactsCountsGetOK creates a EmployeesContactsCountsGetOK with default headers values
func NewEmployeesContactsCountsGetOK() *EmployeesContactsCountsGetOK {
	return &EmployeesContactsCountsGetOK{}
}

/*EmployeesContactsCountsGetOK handles this case with default header values.

Success
*/
type EmployeesContactsCountsGetOK struct {
	Payload []*models.ContactsCount
}

func (o *EmployeesContactsCountsGetOK) Error() string {
	return fmt.Sprintf("[GET /employees/contacts_counts][%d] employeesContactsCountsGetOK  %+v", 200, o.Payload)
}

func (o *EmployeesContactsCountsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmployeesContactsCountsGetUnauthorized creates a EmployeesContactsCountsGetUnauthorized with default headers values
func NewEmployeesContactsCountsGetUnauthorized() *EmployeesContactsCountsGetUnauthorized {
	return &EmployeesContactsCountsGetUnauthorized{}
}

/*EmployeesContactsCountsGetUnauthorized handles this case with default header values.

Unauthorized
*/
type EmployeesContactsCountsGetUnauthorized struct {
}

func (o *EmployeesContactsCountsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /employees/contacts_counts][%d] employeesContactsCountsGetUnauthorized ", 401)
}

func (o *EmployeesContactsCountsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}