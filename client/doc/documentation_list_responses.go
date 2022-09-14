// Code generated by go-swagger; DO NOT EDIT.

package doc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// DocumentationListReader is a Reader for the DocumentationList structure.
type DocumentationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDocumentationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDocumentationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDocumentationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDocumentationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDocumentationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentationListOK creates a DocumentationListOK with default headers values
func NewDocumentationListOK() *DocumentationListOK {
	return &DocumentationListOK{}
}

/*
DocumentationListOK describes a response with status code 200, with default header values.

Success
*/
type DocumentationListOK struct {
	Payload *models.DocumentationsList
}

// IsSuccess returns true when this documentation list o k response has a 2xx status code
func (o *DocumentationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this documentation list o k response has a 3xx status code
func (o *DocumentationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this documentation list o k response has a 4xx status code
func (o *DocumentationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this documentation list o k response has a 5xx status code
func (o *DocumentationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this documentation list o k response a status code equal to that given
func (o *DocumentationListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DocumentationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListOK  %+v", 200, o.Payload)
}

func (o *DocumentationListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListOK  %+v", 200, o.Payload)
}

func (o *DocumentationListOK) GetPayload() *models.DocumentationsList {
	return o.Payload
}

func (o *DocumentationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentationsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocumentationListBadRequest creates a DocumentationListBadRequest with default headers values
func NewDocumentationListBadRequest() *DocumentationListBadRequest {
	return &DocumentationListBadRequest{}
}

/*
DocumentationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DocumentationListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this documentation list bad request response has a 2xx status code
func (o *DocumentationListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this documentation list bad request response has a 3xx status code
func (o *DocumentationListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this documentation list bad request response has a 4xx status code
func (o *DocumentationListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this documentation list bad request response has a 5xx status code
func (o *DocumentationListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this documentation list bad request response a status code equal to that given
func (o *DocumentationListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DocumentationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListBadRequest  %+v", 400, o.Payload)
}

func (o *DocumentationListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListBadRequest  %+v", 400, o.Payload)
}

func (o *DocumentationListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *DocumentationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocumentationListUnauthorized creates a DocumentationListUnauthorized with default headers values
func NewDocumentationListUnauthorized() *DocumentationListUnauthorized {
	return &DocumentationListUnauthorized{}
}

/*
DocumentationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DocumentationListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this documentation list unauthorized response has a 2xx status code
func (o *DocumentationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this documentation list unauthorized response has a 3xx status code
func (o *DocumentationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this documentation list unauthorized response has a 4xx status code
func (o *DocumentationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this documentation list unauthorized response has a 5xx status code
func (o *DocumentationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this documentation list unauthorized response a status code equal to that given
func (o *DocumentationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DocumentationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListUnauthorized  %+v", 401, o.Payload)
}

func (o *DocumentationListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListUnauthorized  %+v", 401, o.Payload)
}

func (o *DocumentationListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DocumentationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocumentationListForbidden creates a DocumentationListForbidden with default headers values
func NewDocumentationListForbidden() *DocumentationListForbidden {
	return &DocumentationListForbidden{}
}

/*
DocumentationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DocumentationListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this documentation list forbidden response has a 2xx status code
func (o *DocumentationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this documentation list forbidden response has a 3xx status code
func (o *DocumentationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this documentation list forbidden response has a 4xx status code
func (o *DocumentationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this documentation list forbidden response has a 5xx status code
func (o *DocumentationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this documentation list forbidden response a status code equal to that given
func (o *DocumentationListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DocumentationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListForbidden  %+v", 403, o.Payload)
}

func (o *DocumentationListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListForbidden  %+v", 403, o.Payload)
}

func (o *DocumentationListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DocumentationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocumentationListNotFound creates a DocumentationListNotFound with default headers values
func NewDocumentationListNotFound() *DocumentationListNotFound {
	return &DocumentationListNotFound{}
}

/*
DocumentationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DocumentationListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this documentation list not found response has a 2xx status code
func (o *DocumentationListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this documentation list not found response has a 3xx status code
func (o *DocumentationListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this documentation list not found response has a 4xx status code
func (o *DocumentationListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this documentation list not found response has a 5xx status code
func (o *DocumentationListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this documentation list not found response a status code equal to that given
func (o *DocumentationListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DocumentationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListNotFound  %+v", 404, o.Payload)
}

func (o *DocumentationListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListNotFound  %+v", 404, o.Payload)
}

func (o *DocumentationListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DocumentationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocumentationListInternalServerError creates a DocumentationListInternalServerError with default headers values
func NewDocumentationListInternalServerError() *DocumentationListInternalServerError {
	return &DocumentationListInternalServerError{}
}

/*
DocumentationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DocumentationListInternalServerError struct {
}

// IsSuccess returns true when this documentation list internal server error response has a 2xx status code
func (o *DocumentationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this documentation list internal server error response has a 3xx status code
func (o *DocumentationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this documentation list internal server error response has a 4xx status code
func (o *DocumentationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this documentation list internal server error response has a 5xx status code
func (o *DocumentationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this documentation list internal server error response a status code equal to that given
func (o *DocumentationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DocumentationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListInternalServerError ", 500)
}

func (o *DocumentationListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Documentation][%d] documentationListInternalServerError ", 500)
}

func (o *DocumentationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
