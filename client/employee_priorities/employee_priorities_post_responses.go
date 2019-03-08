// Code generated by go-swagger; DO NOT EDIT.

package employee_priorities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// EmployeePrioritiesPostReader is a Reader for the EmployeePrioritiesPost structure.
type EmployeePrioritiesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeePrioritiesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmployeePrioritiesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEmployeePrioritiesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeePrioritiesPostOK creates a EmployeePrioritiesPostOK with default headers values
func NewEmployeePrioritiesPostOK() *EmployeePrioritiesPostOK {
	return &EmployeePrioritiesPostOK{}
}

/*EmployeePrioritiesPostOK handles this case with default header values.

Success
*/
type EmployeePrioritiesPostOK struct {
	Payload *models.EmployeePriorityViewModel
}

func (o *EmployeePrioritiesPostOK) Error() string {
	return fmt.Sprintf("[POST /employee_priorities][%d] employeePrioritiesPostOK  %+v", 200, o.Payload)
}

func (o *EmployeePrioritiesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmployeePriorityViewModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmployeePrioritiesPostUnauthorized creates a EmployeePrioritiesPostUnauthorized with default headers values
func NewEmployeePrioritiesPostUnauthorized() *EmployeePrioritiesPostUnauthorized {
	return &EmployeePrioritiesPostUnauthorized{}
}

/*EmployeePrioritiesPostUnauthorized handles this case with default header values.

Unauthorized
*/
type EmployeePrioritiesPostUnauthorized struct {
}

func (o *EmployeePrioritiesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /employee_priorities][%d] employeePrioritiesPostUnauthorized ", 401)
}

func (o *EmployeePrioritiesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}