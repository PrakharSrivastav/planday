// Code generated by go-swagger; DO NOT EDIT.

package employees_contract_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// EmployeesByIDIntervalByIntervalidGetReader is a Reader for the EmployeesByIDIntervalByIntervalidGet structure.
type EmployeesByIDIntervalByIntervalidGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeesByIDIntervalByIntervalidGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmployeesByIDIntervalByIntervalidGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEmployeesByIDIntervalByIntervalidGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeesByIDIntervalByIntervalidGetOK creates a EmployeesByIDIntervalByIntervalidGetOK with default headers values
func NewEmployeesByIDIntervalByIntervalidGetOK() *EmployeesByIDIntervalByIntervalidGetOK {
	return &EmployeesByIDIntervalByIntervalidGetOK{}
}

/*EmployeesByIDIntervalByIntervalidGetOK handles this case with default header values.

Success
*/
type EmployeesByIDIntervalByIntervalidGetOK struct {
	Payload []*models.EmployeeIntervalHoursViewModel
}

func (o *EmployeesByIDIntervalByIntervalidGetOK) Error() string {
	return fmt.Sprintf("[GET /employees/{id}/interval/{intervalid}][%d] employeesByIdIntervalByIntervalidGetOK  %+v", 200, o.Payload)
}

func (o *EmployeesByIDIntervalByIntervalidGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmployeesByIDIntervalByIntervalidGetUnauthorized creates a EmployeesByIDIntervalByIntervalidGetUnauthorized with default headers values
func NewEmployeesByIDIntervalByIntervalidGetUnauthorized() *EmployeesByIDIntervalByIntervalidGetUnauthorized {
	return &EmployeesByIDIntervalByIntervalidGetUnauthorized{}
}

/*EmployeesByIDIntervalByIntervalidGetUnauthorized handles this case with default header values.

Unauthorized
*/
type EmployeesByIDIntervalByIntervalidGetUnauthorized struct {
}

func (o *EmployeesByIDIntervalByIntervalidGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /employees/{id}/interval/{intervalid}][%d] employeesByIdIntervalByIntervalidGetUnauthorized ", 401)
}

func (o *EmployeesByIDIntervalByIntervalidGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
