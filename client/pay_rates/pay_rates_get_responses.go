// Code generated by go-swagger; DO NOT EDIT.

package pay_rates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// PayRatesGetReader is a Reader for the PayRatesGet structure.
type PayRatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PayRatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPayRatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPayRatesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPayRatesGetOK creates a PayRatesGetOK with default headers values
func NewPayRatesGetOK() *PayRatesGetOK {
	return &PayRatesGetOK{}
}

/*PayRatesGetOK handles this case with default header values.

Success
*/
type PayRatesGetOK struct {
	Payload []*models.PayRateViewModelExtended
}

func (o *PayRatesGetOK) Error() string {
	return fmt.Sprintf("[GET /pay_rates][%d] payRatesGetOK  %+v", 200, o.Payload)
}

func (o *PayRatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPayRatesGetUnauthorized creates a PayRatesGetUnauthorized with default headers values
func NewPayRatesGetUnauthorized() *PayRatesGetUnauthorized {
	return &PayRatesGetUnauthorized{}
}

/*PayRatesGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PayRatesGetUnauthorized struct {
}

func (o *PayRatesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pay_rates][%d] payRatesGetUnauthorized ", 401)
}

func (o *PayRatesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
