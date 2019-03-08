// Code generated by go-swagger; DO NOT EDIT.

package contract_rules_intervals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ContractrulesByRuleIDIntervalsByIntervalIDDeleteReader is a Reader for the ContractrulesByRuleIDIntervalsByIntervalIDDelete structure.
type ContractrulesByRuleIDIntervalsByIntervalIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContractrulesByRuleIDIntervalsByIntervalIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContractrulesByRuleIDIntervalsByIntervalIDDeleteOK creates a ContractrulesByRuleIDIntervalsByIntervalIDDeleteOK with default headers values
func NewContractrulesByRuleIDIntervalsByIntervalIDDeleteOK() *ContractrulesByRuleIDIntervalsByIntervalIDDeleteOK {
	return &ContractrulesByRuleIDIntervalsByIntervalIDDeleteOK{}
}

/*ContractrulesByRuleIDIntervalsByIntervalIDDeleteOK handles this case with default header values.

Success
*/
type ContractrulesByRuleIDIntervalsByIntervalIDDeleteOK struct {
}

func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /contractrules/{ruleId}/intervals/{intervalId}][%d] contractrulesByRuleIdIntervalsByIntervalIdDeleteOK ", 200)
}

func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized creates a ContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized with default headers values
func NewContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized() *ContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized {
	return &ContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized{}
}

/*ContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized handles this case with default header values.

Unauthorized
*/
type ContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized struct {
}

func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /contractrules/{ruleId}/intervals/{intervalId}][%d] contractrulesByRuleIdIntervalsByIntervalIdDeleteUnauthorized ", 401)
}

func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound creates a ContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound with default headers values
func NewContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound() *ContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound {
	return &ContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound{}
}

/*ContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound handles this case with default header values.

Not Found
*/
type ContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound struct {
}

func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /contractrules/{ruleId}/intervals/{intervalId}][%d] contractrulesByRuleIdIntervalsByIntervalIdDeleteNotFound ", 404)
}

func (o *ContractrulesByRuleIDIntervalsByIntervalIDDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}