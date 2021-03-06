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

// EmployeesFilterPresetsGetReader is a Reader for the EmployeesFilterPresetsGet structure.
type EmployeesFilterPresetsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeesFilterPresetsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmployeesFilterPresetsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEmployeesFilterPresetsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeesFilterPresetsGetOK creates a EmployeesFilterPresetsGetOK with default headers values
func NewEmployeesFilterPresetsGetOK() *EmployeesFilterPresetsGetOK {
	return &EmployeesFilterPresetsGetOK{}
}

/*EmployeesFilterPresetsGetOK handles this case with default header values.

Success
*/
type EmployeesFilterPresetsGetOK struct {
	Payload []*models.EmployeeFilterPreset
}

func (o *EmployeesFilterPresetsGetOK) Error() string {
	return fmt.Sprintf("[GET /employees/filter_presets][%d] employeesFilterPresetsGetOK  %+v", 200, o.Payload)
}

func (o *EmployeesFilterPresetsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmployeesFilterPresetsGetUnauthorized creates a EmployeesFilterPresetsGetUnauthorized with default headers values
func NewEmployeesFilterPresetsGetUnauthorized() *EmployeesFilterPresetsGetUnauthorized {
	return &EmployeesFilterPresetsGetUnauthorized{}
}

/*EmployeesFilterPresetsGetUnauthorized handles this case with default header values.

Unauthorized
*/
type EmployeesFilterPresetsGetUnauthorized struct {
}

func (o *EmployeesFilterPresetsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /employees/filter_presets][%d] employeesFilterPresetsGetUnauthorized ", 401)
}

func (o *EmployeesFilterPresetsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
