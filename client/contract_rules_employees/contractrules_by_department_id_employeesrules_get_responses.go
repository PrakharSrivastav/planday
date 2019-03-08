// Code generated by go-swagger; DO NOT EDIT.

package contract_rules_employees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// ContractrulesByDepartmentIDEmployeesrulesGetReader is a Reader for the ContractrulesByDepartmentIDEmployeesrulesGet structure.
type ContractrulesByDepartmentIDEmployeesrulesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractrulesByDepartmentIDEmployeesrulesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContractrulesByDepartmentIDEmployeesrulesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewContractrulesByDepartmentIDEmployeesrulesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContractrulesByDepartmentIDEmployeesrulesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContractrulesByDepartmentIDEmployeesrulesGetOK creates a ContractrulesByDepartmentIDEmployeesrulesGetOK with default headers values
func NewContractrulesByDepartmentIDEmployeesrulesGetOK() *ContractrulesByDepartmentIDEmployeesrulesGetOK {
	return &ContractrulesByDepartmentIDEmployeesrulesGetOK{}
}

/*ContractrulesByDepartmentIDEmployeesrulesGetOK handles this case with default header values.

Success
*/
type ContractrulesByDepartmentIDEmployeesrulesGetOK struct {
	Payload []*models.EmployeeContractRuleAppliedResult
}

func (o *ContractrulesByDepartmentIDEmployeesrulesGetOK) Error() string {
	return fmt.Sprintf("[GET /contractrules/{departmentId}/employeesrules][%d] contractrulesByDepartmentIdEmployeesrulesGetOK  %+v", 200, o.Payload)
}

func (o *ContractrulesByDepartmentIDEmployeesrulesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContractrulesByDepartmentIDEmployeesrulesGetUnauthorized creates a ContractrulesByDepartmentIDEmployeesrulesGetUnauthorized with default headers values
func NewContractrulesByDepartmentIDEmployeesrulesGetUnauthorized() *ContractrulesByDepartmentIDEmployeesrulesGetUnauthorized {
	return &ContractrulesByDepartmentIDEmployeesrulesGetUnauthorized{}
}

/*ContractrulesByDepartmentIDEmployeesrulesGetUnauthorized handles this case with default header values.

Unauthorized
*/
type ContractrulesByDepartmentIDEmployeesrulesGetUnauthorized struct {
}

func (o *ContractrulesByDepartmentIDEmployeesrulesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /contractrules/{departmentId}/employeesrules][%d] contractrulesByDepartmentIdEmployeesrulesGetUnauthorized ", 401)
}

func (o *ContractrulesByDepartmentIDEmployeesrulesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContractrulesByDepartmentIDEmployeesrulesGetNotFound creates a ContractrulesByDepartmentIDEmployeesrulesGetNotFound with default headers values
func NewContractrulesByDepartmentIDEmployeesrulesGetNotFound() *ContractrulesByDepartmentIDEmployeesrulesGetNotFound {
	return &ContractrulesByDepartmentIDEmployeesrulesGetNotFound{}
}

/*ContractrulesByDepartmentIDEmployeesrulesGetNotFound handles this case with default header values.

Not Found
*/
type ContractrulesByDepartmentIDEmployeesrulesGetNotFound struct {
}

func (o *ContractrulesByDepartmentIDEmployeesrulesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /contractrules/{departmentId}/employeesrules][%d] contractrulesByDepartmentIdEmployeesrulesGetNotFound ", 404)
}

func (o *ContractrulesByDepartmentIDEmployeesrulesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
