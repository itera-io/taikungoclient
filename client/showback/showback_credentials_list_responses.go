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

// ShowbackCredentialsListReader is a Reader for the ShowbackCredentialsList structure.
type ShowbackCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackCredentialsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackCredentialsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackCredentialsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackCredentialsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewShowbackCredentialsListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackCredentialsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackCredentialsListOK creates a ShowbackCredentialsListOK with default headers values
func NewShowbackCredentialsListOK() *ShowbackCredentialsListOK {
	return &ShowbackCredentialsListOK{}
}

/* ShowbackCredentialsListOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackCredentialsListOK struct {
	Payload *models.ShowbackCredentialsList
}

func (o *ShowbackCredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials][%d] showbackCredentialsListOK  %+v", 200, o.Payload)
}
func (o *ShowbackCredentialsListOK) GetPayload() *models.ShowbackCredentialsList {
	return o.Payload
}

func (o *ShowbackCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShowbackCredentialsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListBadRequest creates a ShowbackCredentialsListBadRequest with default headers values
func NewShowbackCredentialsListBadRequest() *ShowbackCredentialsListBadRequest {
	return &ShowbackCredentialsListBadRequest{}
}

/* ShowbackCredentialsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackCredentialsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ShowbackCredentialsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials][%d] showbackCredentialsListBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackCredentialsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListUnauthorized creates a ShowbackCredentialsListUnauthorized with default headers values
func NewShowbackCredentialsListUnauthorized() *ShowbackCredentialsListUnauthorized {
	return &ShowbackCredentialsListUnauthorized{}
}

/* ShowbackCredentialsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackCredentialsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCredentialsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials][%d] showbackCredentialsListUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackCredentialsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListForbidden creates a ShowbackCredentialsListForbidden with default headers values
func NewShowbackCredentialsListForbidden() *ShowbackCredentialsListForbidden {
	return &ShowbackCredentialsListForbidden{}
}

/* ShowbackCredentialsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackCredentialsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCredentialsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials][%d] showbackCredentialsListForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackCredentialsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListNotFound creates a ShowbackCredentialsListNotFound with default headers values
func NewShowbackCredentialsListNotFound() *ShowbackCredentialsListNotFound {
	return &ShowbackCredentialsListNotFound{}
}

/* ShowbackCredentialsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackCredentialsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCredentialsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials][%d] showbackCredentialsListNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackCredentialsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListTooManyRequests creates a ShowbackCredentialsListTooManyRequests with default headers values
func NewShowbackCredentialsListTooManyRequests() *ShowbackCredentialsListTooManyRequests {
	return &ShowbackCredentialsListTooManyRequests{}
}

/* ShowbackCredentialsListTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type ShowbackCredentialsListTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCredentialsListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials][%d] showbackCredentialsListTooManyRequests  %+v", 429, o.Payload)
}
func (o *ShowbackCredentialsListTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListInternalServerError creates a ShowbackCredentialsListInternalServerError with default headers values
func NewShowbackCredentialsListInternalServerError() *ShowbackCredentialsListInternalServerError {
	return &ShowbackCredentialsListInternalServerError{}
}

/* ShowbackCredentialsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackCredentialsListInternalServerError struct {
}

func (o *ShowbackCredentialsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/credentials][%d] showbackCredentialsListInternalServerError ", 500)
}

func (o *ShowbackCredentialsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
