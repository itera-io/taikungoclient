// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SearchSecretListReader is a Reader for the SearchSecretList structure.
type SearchSecretListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchSecretListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchSecretListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchSecretListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchSecretListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchSecretListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchSecretListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchSecretListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchSecretListOK creates a SearchSecretListOK with default headers values
func NewSearchSecretListOK() *SearchSecretListOK {
	return &SearchSecretListOK{}
}

/*
SearchSecretListOK describes a response with status code 200, with default header values.

Success
*/
type SearchSecretListOK struct {
	Payload *models.SecretSearchList
}

// IsSuccess returns true when this search secret list o k response has a 2xx status code
func (o *SearchSecretListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search secret list o k response has a 3xx status code
func (o *SearchSecretListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search secret list o k response has a 4xx status code
func (o *SearchSecretListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search secret list o k response has a 5xx status code
func (o *SearchSecretListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search secret list o k response a status code equal to that given
func (o *SearchSecretListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchSecretListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListOK  %+v", 200, o.Payload)
}

func (o *SearchSecretListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListOK  %+v", 200, o.Payload)
}

func (o *SearchSecretListOK) GetPayload() *models.SecretSearchList {
	return o.Payload
}

func (o *SearchSecretListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSecretListBadRequest creates a SearchSecretListBadRequest with default headers values
func NewSearchSecretListBadRequest() *SearchSecretListBadRequest {
	return &SearchSecretListBadRequest{}
}

/*
SearchSecretListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchSecretListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search secret list bad request response has a 2xx status code
func (o *SearchSecretListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search secret list bad request response has a 3xx status code
func (o *SearchSecretListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search secret list bad request response has a 4xx status code
func (o *SearchSecretListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search secret list bad request response has a 5xx status code
func (o *SearchSecretListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search secret list bad request response a status code equal to that given
func (o *SearchSecretListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchSecretListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchSecretListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchSecretListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchSecretListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSecretListUnauthorized creates a SearchSecretListUnauthorized with default headers values
func NewSearchSecretListUnauthorized() *SearchSecretListUnauthorized {
	return &SearchSecretListUnauthorized{}
}

/*
SearchSecretListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchSecretListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search secret list unauthorized response has a 2xx status code
func (o *SearchSecretListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search secret list unauthorized response has a 3xx status code
func (o *SearchSecretListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search secret list unauthorized response has a 4xx status code
func (o *SearchSecretListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search secret list unauthorized response has a 5xx status code
func (o *SearchSecretListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search secret list unauthorized response a status code equal to that given
func (o *SearchSecretListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchSecretListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchSecretListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchSecretListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchSecretListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSecretListForbidden creates a SearchSecretListForbidden with default headers values
func NewSearchSecretListForbidden() *SearchSecretListForbidden {
	return &SearchSecretListForbidden{}
}

/*
SearchSecretListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchSecretListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search secret list forbidden response has a 2xx status code
func (o *SearchSecretListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search secret list forbidden response has a 3xx status code
func (o *SearchSecretListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search secret list forbidden response has a 4xx status code
func (o *SearchSecretListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search secret list forbidden response has a 5xx status code
func (o *SearchSecretListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search secret list forbidden response a status code equal to that given
func (o *SearchSecretListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchSecretListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListForbidden  %+v", 403, o.Payload)
}

func (o *SearchSecretListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListForbidden  %+v", 403, o.Payload)
}

func (o *SearchSecretListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchSecretListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSecretListNotFound creates a SearchSecretListNotFound with default headers values
func NewSearchSecretListNotFound() *SearchSecretListNotFound {
	return &SearchSecretListNotFound{}
}

/*
SearchSecretListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchSecretListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search secret list not found response has a 2xx status code
func (o *SearchSecretListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search secret list not found response has a 3xx status code
func (o *SearchSecretListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search secret list not found response has a 4xx status code
func (o *SearchSecretListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search secret list not found response has a 5xx status code
func (o *SearchSecretListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search secret list not found response a status code equal to that given
func (o *SearchSecretListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchSecretListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListNotFound  %+v", 404, o.Payload)
}

func (o *SearchSecretListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListNotFound  %+v", 404, o.Payload)
}

func (o *SearchSecretListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchSecretListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSecretListInternalServerError creates a SearchSecretListInternalServerError with default headers values
func NewSearchSecretListInternalServerError() *SearchSecretListInternalServerError {
	return &SearchSecretListInternalServerError{}
}

/*
SearchSecretListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchSecretListInternalServerError struct {
}

// IsSuccess returns true when this search secret list internal server error response has a 2xx status code
func (o *SearchSecretListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search secret list internal server error response has a 3xx status code
func (o *SearchSecretListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search secret list internal server error response has a 4xx status code
func (o *SearchSecretListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search secret list internal server error response has a 5xx status code
func (o *SearchSecretListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search secret list internal server error response a status code equal to that given
func (o *SearchSecretListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchSecretListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListInternalServerError ", 500)
}

func (o *SearchSecretListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/secrets][%d] searchSecretListInternalServerError ", 500)
}

func (o *SearchSecretListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
