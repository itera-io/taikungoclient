// Code generated by go-swagger; DO NOT EDIT.

package showback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackShowbackCredentialsDropdownReader is a Reader for the ShowbackShowbackCredentialsDropdown structure.
type ShowbackShowbackCredentialsDropdownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackShowbackCredentialsDropdownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackShowbackCredentialsDropdownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackShowbackCredentialsDropdownBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackShowbackCredentialsDropdownUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackShowbackCredentialsDropdownForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackShowbackCredentialsDropdownNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackShowbackCredentialsDropdownInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackShowbackCredentialsDropdownOK creates a ShowbackShowbackCredentialsDropdownOK with default headers values
func NewShowbackShowbackCredentialsDropdownOK() *ShowbackShowbackCredentialsDropdownOK {
	return &ShowbackShowbackCredentialsDropdownOK{}
}

/* ShowbackShowbackCredentialsDropdownOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackShowbackCredentialsDropdownOK struct {
	Payload []*models.ShowbackCredentialsDetailsDto
}

func (o *ShowbackShowbackCredentialsDropdownOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials/list][%d] showbackShowbackCredentialsDropdownOK  %+v", 200, o.Payload)
}
func (o *ShowbackShowbackCredentialsDropdownOK) GetPayload() []*models.ShowbackCredentialsDetailsDto {
	return o.Payload
}

func (o *ShowbackShowbackCredentialsDropdownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackShowbackCredentialsDropdownBadRequest creates a ShowbackShowbackCredentialsDropdownBadRequest with default headers values
func NewShowbackShowbackCredentialsDropdownBadRequest() *ShowbackShowbackCredentialsDropdownBadRequest {
	return &ShowbackShowbackCredentialsDropdownBadRequest{}
}

/* ShowbackShowbackCredentialsDropdownBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackShowbackCredentialsDropdownBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ShowbackShowbackCredentialsDropdownBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials/list][%d] showbackShowbackCredentialsDropdownBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackShowbackCredentialsDropdownBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackShowbackCredentialsDropdownBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackShowbackCredentialsDropdownUnauthorized creates a ShowbackShowbackCredentialsDropdownUnauthorized with default headers values
func NewShowbackShowbackCredentialsDropdownUnauthorized() *ShowbackShowbackCredentialsDropdownUnauthorized {
	return &ShowbackShowbackCredentialsDropdownUnauthorized{}
}

/* ShowbackShowbackCredentialsDropdownUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackShowbackCredentialsDropdownUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackShowbackCredentialsDropdownUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials/list][%d] showbackShowbackCredentialsDropdownUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackShowbackCredentialsDropdownUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackShowbackCredentialsDropdownUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackShowbackCredentialsDropdownForbidden creates a ShowbackShowbackCredentialsDropdownForbidden with default headers values
func NewShowbackShowbackCredentialsDropdownForbidden() *ShowbackShowbackCredentialsDropdownForbidden {
	return &ShowbackShowbackCredentialsDropdownForbidden{}
}

/* ShowbackShowbackCredentialsDropdownForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackShowbackCredentialsDropdownForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackShowbackCredentialsDropdownForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials/list][%d] showbackShowbackCredentialsDropdownForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackShowbackCredentialsDropdownForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackShowbackCredentialsDropdownForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackShowbackCredentialsDropdownNotFound creates a ShowbackShowbackCredentialsDropdownNotFound with default headers values
func NewShowbackShowbackCredentialsDropdownNotFound() *ShowbackShowbackCredentialsDropdownNotFound {
	return &ShowbackShowbackCredentialsDropdownNotFound{}
}

/* ShowbackShowbackCredentialsDropdownNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackShowbackCredentialsDropdownNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackShowbackCredentialsDropdownNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials/list][%d] showbackShowbackCredentialsDropdownNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackShowbackCredentialsDropdownNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackShowbackCredentialsDropdownNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackShowbackCredentialsDropdownInternalServerError creates a ShowbackShowbackCredentialsDropdownInternalServerError with default headers values
func NewShowbackShowbackCredentialsDropdownInternalServerError() *ShowbackShowbackCredentialsDropdownInternalServerError {
	return &ShowbackShowbackCredentialsDropdownInternalServerError{}
}

/* ShowbackShowbackCredentialsDropdownInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackShowbackCredentialsDropdownInternalServerError struct {
}

func (o *ShowbackShowbackCredentialsDropdownInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials/list][%d] showbackShowbackCredentialsDropdownInternalServerError ", 500)
}

func (o *ShowbackShowbackCredentialsDropdownInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
