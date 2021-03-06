// Code generated by go-swagger; DO NOT EDIT.

package employees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// EmployeesPostReader is a Reader for the EmployeesPost structure.
type EmployeesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmployeesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEmployeesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeesPostOK creates a EmployeesPostOK with default headers values
func NewEmployeesPostOK() *EmployeesPostOK {
	return &EmployeesPostOK{}
}

/*EmployeesPostOK handles this case with default header values.

Success
*/
type EmployeesPostOK struct {
	Payload *models.EmployeeViewModel
}

func (o *EmployeesPostOK) Error() string {
	return fmt.Sprintf("[POST /employees][%d] employeesPostOK  %+v", 200, o.Payload)
}

func (o *EmployeesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmployeeViewModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmployeesPostUnauthorized creates a EmployeesPostUnauthorized with default headers values
func NewEmployeesPostUnauthorized() *EmployeesPostUnauthorized {
	return &EmployeesPostUnauthorized{}
}

/*EmployeesPostUnauthorized handles this case with default header values.

Unauthorized
*/
type EmployeesPostUnauthorized struct {
}

func (o *EmployeesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /employees][%d] employeesPostUnauthorized ", 401)
}

func (o *EmployeesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
