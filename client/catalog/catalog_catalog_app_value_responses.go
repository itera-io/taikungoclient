// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CatalogCatalogAppValueReader is a Reader for the CatalogCatalogAppValue structure.
type CatalogCatalogAppValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCatalogAppValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCatalogAppValueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCatalogAppValueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogCatalogAppValueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogCatalogAppValueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogCatalogAppValueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogCatalogAppValueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogCatalogAppValueOK creates a CatalogCatalogAppValueOK with default headers values
func NewCatalogCatalogAppValueOK() *CatalogCatalogAppValueOK {
	return &CatalogCatalogAppValueOK{}
}

/*
CatalogCatalogAppValueOK describes a response with status code 200, with default header values.

Success
*/
type CatalogCatalogAppValueOK struct {
	Payload string
}

// IsSuccess returns true when this catalog catalog app value o k response has a 2xx status code
func (o *CatalogCatalogAppValueOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog catalog app value o k response has a 3xx status code
func (o *CatalogCatalogAppValueOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value o k response has a 4xx status code
func (o *CatalogCatalogAppValueOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog app value o k response has a 5xx status code
func (o *CatalogCatalogAppValueOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value o k response a status code equal to that given
func (o *CatalogCatalogAppValueOK) IsCode(code int) bool {
	return code == 200
}

func (o *CatalogCatalogAppValueOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogAppValueOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogAppValueOK) GetPayload() string {
	return o.Payload
}

func (o *CatalogCatalogAppValueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueBadRequest creates a CatalogCatalogAppValueBadRequest with default headers values
func NewCatalogCatalogAppValueBadRequest() *CatalogCatalogAppValueBadRequest {
	return &CatalogCatalogAppValueBadRequest{}
}

/*
CatalogCatalogAppValueBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogCatalogAppValueBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog catalog app value bad request response has a 2xx status code
func (o *CatalogCatalogAppValueBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value bad request response has a 3xx status code
func (o *CatalogCatalogAppValueBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value bad request response has a 4xx status code
func (o *CatalogCatalogAppValueBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value bad request response has a 5xx status code
func (o *CatalogCatalogAppValueBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value bad request response a status code equal to that given
func (o *CatalogCatalogAppValueBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CatalogCatalogAppValueBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogAppValueBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogAppValueBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogCatalogAppValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueUnauthorized creates a CatalogCatalogAppValueUnauthorized with default headers values
func NewCatalogCatalogAppValueUnauthorized() *CatalogCatalogAppValueUnauthorized {
	return &CatalogCatalogAppValueUnauthorized{}
}

/*
CatalogCatalogAppValueUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogCatalogAppValueUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog app value unauthorized response has a 2xx status code
func (o *CatalogCatalogAppValueUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value unauthorized response has a 3xx status code
func (o *CatalogCatalogAppValueUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value unauthorized response has a 4xx status code
func (o *CatalogCatalogAppValueUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value unauthorized response has a 5xx status code
func (o *CatalogCatalogAppValueUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value unauthorized response a status code equal to that given
func (o *CatalogCatalogAppValueUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CatalogCatalogAppValueUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogAppValueUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogAppValueUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppValueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueForbidden creates a CatalogCatalogAppValueForbidden with default headers values
func NewCatalogCatalogAppValueForbidden() *CatalogCatalogAppValueForbidden {
	return &CatalogCatalogAppValueForbidden{}
}

/*
CatalogCatalogAppValueForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogCatalogAppValueForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog app value forbidden response has a 2xx status code
func (o *CatalogCatalogAppValueForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value forbidden response has a 3xx status code
func (o *CatalogCatalogAppValueForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value forbidden response has a 4xx status code
func (o *CatalogCatalogAppValueForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value forbidden response has a 5xx status code
func (o *CatalogCatalogAppValueForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value forbidden response a status code equal to that given
func (o *CatalogCatalogAppValueForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CatalogCatalogAppValueForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogAppValueForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogAppValueForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppValueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueNotFound creates a CatalogCatalogAppValueNotFound with default headers values
func NewCatalogCatalogAppValueNotFound() *CatalogCatalogAppValueNotFound {
	return &CatalogCatalogAppValueNotFound{}
}

/*
CatalogCatalogAppValueNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogCatalogAppValueNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog app value not found response has a 2xx status code
func (o *CatalogCatalogAppValueNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value not found response has a 3xx status code
func (o *CatalogCatalogAppValueNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value not found response has a 4xx status code
func (o *CatalogCatalogAppValueNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value not found response has a 5xx status code
func (o *CatalogCatalogAppValueNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value not found response a status code equal to that given
func (o *CatalogCatalogAppValueNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CatalogCatalogAppValueNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogAppValueNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogAppValueNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppValueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueInternalServerError creates a CatalogCatalogAppValueInternalServerError with default headers values
func NewCatalogCatalogAppValueInternalServerError() *CatalogCatalogAppValueInternalServerError {
	return &CatalogCatalogAppValueInternalServerError{}
}

/*
CatalogCatalogAppValueInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogCatalogAppValueInternalServerError struct {
}

// IsSuccess returns true when this catalog catalog app value internal server error response has a 2xx status code
func (o *CatalogCatalogAppValueInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value internal server error response has a 3xx status code
func (o *CatalogCatalogAppValueInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value internal server error response has a 4xx status code
func (o *CatalogCatalogAppValueInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog app value internal server error response has a 5xx status code
func (o *CatalogCatalogAppValueInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog catalog app value internal server error response a status code equal to that given
func (o *CatalogCatalogAppValueInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CatalogCatalogAppValueInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueInternalServerError ", 500)
}

func (o *CatalogCatalogAppValueInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value][%d] catalogCatalogAppValueInternalServerError ", 500)
}

func (o *CatalogCatalogAppValueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
