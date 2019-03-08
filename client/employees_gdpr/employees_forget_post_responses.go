// Code generated by go-swagger; DO NOT EDIT.

package employees_gdpr

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// EmployeesForgetPostReader is a Reader for the EmployeesForgetPost structure.
type EmployeesForgetPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeesForgetPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmployeesForgetPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEmployeesForgetPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeesForgetPostOK creates a EmployeesForgetPostOK with default headers values
func NewEmployeesForgetPostOK() *EmployeesForgetPostOK {
	return &EmployeesForgetPostOK{}
}

/*EmployeesForgetPostOK handles this case with default header values.

Success
*/
type EmployeesForgetPostOK struct {
	Payload []*models.Employee
}

func (o *EmployeesForgetPostOK) Error() string {
	return fmt.Sprintf("[POST /employees/forget][%d] employeesForgetPostOK  %+v", 200, o.Payload)
}

func (o *EmployeesForgetPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmployeesForgetPostUnauthorized creates a EmployeesForgetPostUnauthorized with default headers values
func NewEmployeesForgetPostUnauthorized() *EmployeesForgetPostUnauthorized {
	return &EmployeesForgetPostUnauthorized{}
}

/*EmployeesForgetPostUnauthorized handles this case with default header values.

Unauthorized
*/
type EmployeesForgetPostUnauthorized struct {
}

func (o *EmployeesForgetPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /employees/forget][%d] employeesForgetPostUnauthorized ", 401)
}

func (o *EmployeesForgetPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}