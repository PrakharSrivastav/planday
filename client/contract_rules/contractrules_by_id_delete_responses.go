// Code generated by go-swagger; DO NOT EDIT.

package contract_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ContractrulesByIDDeleteReader is a Reader for the ContractrulesByIDDelete structure.
type ContractrulesByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractrulesByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContractrulesByIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewContractrulesByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContractrulesByIDDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContractrulesByIDDeleteOK creates a ContractrulesByIDDeleteOK with default headers values
func NewContractrulesByIDDeleteOK() *ContractrulesByIDDeleteOK {
	return &ContractrulesByIDDeleteOK{}
}

/*ContractrulesByIDDeleteOK handles this case with default header values.

Success
*/
type ContractrulesByIDDeleteOK struct {
}

func (o *ContractrulesByIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /contractrules/{id}][%d] contractrulesByIdDeleteOK ", 200)
}

func (o *ContractrulesByIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContractrulesByIDDeleteUnauthorized creates a ContractrulesByIDDeleteUnauthorized with default headers values
func NewContractrulesByIDDeleteUnauthorized() *ContractrulesByIDDeleteUnauthorized {
	return &ContractrulesByIDDeleteUnauthorized{}
}

/*ContractrulesByIDDeleteUnauthorized handles this case with default header values.

Unauthorized
*/
type ContractrulesByIDDeleteUnauthorized struct {
}

func (o *ContractrulesByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /contractrules/{id}][%d] contractrulesByIdDeleteUnauthorized ", 401)
}

func (o *ContractrulesByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContractrulesByIDDeleteNotFound creates a ContractrulesByIDDeleteNotFound with default headers values
func NewContractrulesByIDDeleteNotFound() *ContractrulesByIDDeleteNotFound {
	return &ContractrulesByIDDeleteNotFound{}
}

/*ContractrulesByIDDeleteNotFound handles this case with default header values.

Not Found
*/
type ContractrulesByIDDeleteNotFound struct {
}

func (o *ContractrulesByIDDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /contractrules/{id}][%d] contractrulesByIdDeleteNotFound ", 404)
}

func (o *ContractrulesByIDDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
