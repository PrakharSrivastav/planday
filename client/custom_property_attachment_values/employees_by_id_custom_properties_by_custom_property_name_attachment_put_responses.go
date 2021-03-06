// Code generated by go-swagger; DO NOT EDIT.

package custom_property_attachment_values

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutReader is a Reader for the EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPut structure.
type EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewEmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK creates a EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK with default headers values
func NewEmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK() *EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK {
	return &EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK{}
}

/*EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK handles this case with default header values.

Success
*/
type EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK struct {
	Payload string
}

func (o *EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK) Error() string {
	return fmt.Sprintf("[PUT /employees/{id}/custom_properties/{custom_property_name}/attachment][%d] employeesByIdCustomPropertiesByCustomPropertyNameAttachmentPutOK  %+v", 200, o.Payload)
}

func (o *EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized creates a EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized with default headers values
func NewEmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized() *EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized {
	return &EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized{}
}

/*EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized handles this case with default header values.

Unauthorized
*/
type EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized struct {
}

func (o *EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /employees/{id}/custom_properties/{custom_property_name}/attachment][%d] employeesByIdCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized ", 401)
}

func (o *EmployeesByIDCustomPropertiesByCustomPropertyNameAttachmentPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
