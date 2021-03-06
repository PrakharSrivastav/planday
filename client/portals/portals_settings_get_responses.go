// Code generated by go-swagger; DO NOT EDIT.

package portals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/PrakharSrivastav/planday/models"
)

// PortalsSettingsGetReader is a Reader for the PortalsSettingsGet structure.
type PortalsSettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortalsSettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPortalsSettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPortalsSettingsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPortalsSettingsGetOK creates a PortalsSettingsGetOK with default headers values
func NewPortalsSettingsGetOK() *PortalsSettingsGetOK {
	return &PortalsSettingsGetOK{}
}

/*PortalsSettingsGetOK handles this case with default header values.

Success
*/
type PortalsSettingsGetOK struct {
	Payload *models.PortalSettingsViewModel
}

func (o *PortalsSettingsGetOK) Error() string {
	return fmt.Sprintf("[GET /portals/settings][%d] portalsSettingsGetOK  %+v", 200, o.Payload)
}

func (o *PortalsSettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortalSettingsViewModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortalsSettingsGetUnauthorized creates a PortalsSettingsGetUnauthorized with default headers values
func NewPortalsSettingsGetUnauthorized() *PortalsSettingsGetUnauthorized {
	return &PortalsSettingsGetUnauthorized{}
}

/*PortalsSettingsGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PortalsSettingsGetUnauthorized struct {
}

func (o *PortalsSettingsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /portals/settings][%d] portalsSettingsGetUnauthorized ", 401)
}

func (o *PortalsSettingsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
