// Code generated by go-swagger; DO NOT EDIT.

package contract_rules_shift_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// ContractrulesByRuleIDShifttypesByAllSelectedPutReader is a Reader for the ContractrulesByRuleIDShifttypesByAllSelectedPut structure.
type ContractrulesByRuleIDShifttypesByAllSelectedPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractrulesByRuleIDShifttypesByAllSelectedPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContractrulesByRuleIDShifttypesByAllSelectedPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContractrulesByRuleIDShifttypesByAllSelectedPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContractrulesByRuleIDShifttypesByAllSelectedPutOK creates a ContractrulesByRuleIDShifttypesByAllSelectedPutOK with default headers values
func NewContractrulesByRuleIDShifttypesByAllSelectedPutOK() *ContractrulesByRuleIDShifttypesByAllSelectedPutOK {
	return &ContractrulesByRuleIDShifttypesByAllSelectedPutOK{}
}

/*ContractrulesByRuleIDShifttypesByAllSelectedPutOK handles this case with default header values.

Success
*/
type ContractrulesByRuleIDShifttypesByAllSelectedPutOK struct {
	Payload *models.ContractRulesShiftTypeResultViewModel
}

func (o *ContractrulesByRuleIDShifttypesByAllSelectedPutOK) Error() string {
	return fmt.Sprintf("[PUT /contractrules/{ruleId}/shifttypes/{allSelected}][%d] contractrulesByRuleIdShifttypesByAllSelectedPutOK  %+v", 200, o.Payload)
}

func (o *ContractrulesByRuleIDShifttypesByAllSelectedPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContractRulesShiftTypeResultViewModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized creates a ContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized with default headers values
func NewContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized() *ContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized {
	return &ContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized{}
}

/*ContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized handles this case with default header values.

Unauthorized
*/
type ContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized struct {
}

func (o *ContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /contractrules/{ruleId}/shifttypes/{allSelected}][%d] contractrulesByRuleIdShifttypesByAllSelectedPutUnauthorized ", 401)
}

func (o *ContractrulesByRuleIDShifttypesByAllSelectedPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContractrulesByRuleIDShifttypesByAllSelectedPutNotFound creates a ContractrulesByRuleIDShifttypesByAllSelectedPutNotFound with default headers values
func NewContractrulesByRuleIDShifttypesByAllSelectedPutNotFound() *ContractrulesByRuleIDShifttypesByAllSelectedPutNotFound {
	return &ContractrulesByRuleIDShifttypesByAllSelectedPutNotFound{}
}

/*ContractrulesByRuleIDShifttypesByAllSelectedPutNotFound handles this case with default header values.

Not Found
*/
type ContractrulesByRuleIDShifttypesByAllSelectedPutNotFound struct {
}

func (o *ContractrulesByRuleIDShifttypesByAllSelectedPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /contractrules/{ruleId}/shifttypes/{allSelected}][%d] contractrulesByRuleIdShifttypesByAllSelectedPutNotFound ", 404)
}

func (o *ContractrulesByRuleIDShifttypesByAllSelectedPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}