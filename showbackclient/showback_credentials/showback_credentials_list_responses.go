// Code generated by go-swagger; DO NOT EDIT.

package showback_credentials

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

/*
ShowbackCredentialsListOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackCredentialsListOK struct {
	Payload *models.ShowbackCredentialsList
}

// IsSuccess returns true when this showback credentials list o k response has a 2xx status code
func (o *ShowbackCredentialsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback credentials list o k response has a 3xx status code
func (o *ShowbackCredentialsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials list o k response has a 4xx status code
func (o *ShowbackCredentialsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback credentials list o k response has a 5xx status code
func (o *ShowbackCredentialsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials list o k response a status code equal to that given
func (o *ShowbackCredentialsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ShowbackCredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListOK  %+v", 200, o.Payload)
}

func (o *ShowbackCredentialsListOK) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListOK  %+v", 200, o.Payload)
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

/*
ShowbackCredentialsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackCredentialsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this showback credentials list bad request response has a 2xx status code
func (o *ShowbackCredentialsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials list bad request response has a 3xx status code
func (o *ShowbackCredentialsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials list bad request response has a 4xx status code
func (o *ShowbackCredentialsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials list bad request response has a 5xx status code
func (o *ShowbackCredentialsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials list bad request response a status code equal to that given
func (o *ShowbackCredentialsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ShowbackCredentialsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackCredentialsListBadRequest) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListBadRequest  %+v", 400, o.Payload)
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

/*
ShowbackCredentialsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackCredentialsListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials list unauthorized response has a 2xx status code
func (o *ShowbackCredentialsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials list unauthorized response has a 3xx status code
func (o *ShowbackCredentialsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials list unauthorized response has a 4xx status code
func (o *ShowbackCredentialsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials list unauthorized response has a 5xx status code
func (o *ShowbackCredentialsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials list unauthorized response a status code equal to that given
func (o *ShowbackCredentialsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ShowbackCredentialsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackCredentialsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackCredentialsListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListForbidden creates a ShowbackCredentialsListForbidden with default headers values
func NewShowbackCredentialsListForbidden() *ShowbackCredentialsListForbidden {
	return &ShowbackCredentialsListForbidden{}
}

/*
ShowbackCredentialsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackCredentialsListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials list forbidden response has a 2xx status code
func (o *ShowbackCredentialsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials list forbidden response has a 3xx status code
func (o *ShowbackCredentialsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials list forbidden response has a 4xx status code
func (o *ShowbackCredentialsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials list forbidden response has a 5xx status code
func (o *ShowbackCredentialsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials list forbidden response a status code equal to that given
func (o *ShowbackCredentialsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ShowbackCredentialsListForbidden) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackCredentialsListForbidden) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackCredentialsListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListNotFound creates a ShowbackCredentialsListNotFound with default headers values
func NewShowbackCredentialsListNotFound() *ShowbackCredentialsListNotFound {
	return &ShowbackCredentialsListNotFound{}
}

/*
ShowbackCredentialsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackCredentialsListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials list not found response has a 2xx status code
func (o *ShowbackCredentialsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials list not found response has a 3xx status code
func (o *ShowbackCredentialsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials list not found response has a 4xx status code
func (o *ShowbackCredentialsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials list not found response has a 5xx status code
func (o *ShowbackCredentialsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials list not found response a status code equal to that given
func (o *ShowbackCredentialsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ShowbackCredentialsListNotFound) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackCredentialsListNotFound) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackCredentialsListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsListInternalServerError creates a ShowbackCredentialsListInternalServerError with default headers values
func NewShowbackCredentialsListInternalServerError() *ShowbackCredentialsListInternalServerError {
	return &ShowbackCredentialsListInternalServerError{}
}

/*
ShowbackCredentialsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackCredentialsListInternalServerError struct {
}

// IsSuccess returns true when this showback credentials list internal server error response has a 2xx status code
func (o *ShowbackCredentialsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials list internal server error response has a 3xx status code
func (o *ShowbackCredentialsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials list internal server error response has a 4xx status code
func (o *ShowbackCredentialsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback credentials list internal server error response has a 5xx status code
func (o *ShowbackCredentialsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback credentials list internal server error response a status code equal to that given
func (o *ShowbackCredentialsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ShowbackCredentialsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListInternalServerError ", 500)
}

func (o *ShowbackCredentialsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackCredentials][%d] showbackCredentialsListInternalServerError ", 500)
}

func (o *ShowbackCredentialsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
