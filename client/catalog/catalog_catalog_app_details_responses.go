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

// CatalogCatalogAppDetailsReader is a Reader for the CatalogCatalogAppDetails structure.
type CatalogCatalogAppDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCatalogAppDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCatalogAppDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCatalogAppDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogCatalogAppDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogCatalogAppDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogCatalogAppDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogCatalogAppDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogCatalogAppDetailsOK creates a CatalogCatalogAppDetailsOK with default headers values
func NewCatalogCatalogAppDetailsOK() *CatalogCatalogAppDetailsOK {
	return &CatalogCatalogAppDetailsOK{}
}

/*
CatalogCatalogAppDetailsOK describes a response with status code 200, with default header values.

Success
*/
type CatalogCatalogAppDetailsOK struct {
	Payload *models.CatalogAppDetailsDto
}

// IsSuccess returns true when this catalog catalog app details o k response has a 2xx status code
func (o *CatalogCatalogAppDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog catalog app details o k response has a 3xx status code
func (o *CatalogCatalogAppDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app details o k response has a 4xx status code
func (o *CatalogCatalogAppDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog app details o k response has a 5xx status code
func (o *CatalogCatalogAppDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app details o k response a status code equal to that given
func (o *CatalogCatalogAppDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CatalogCatalogAppDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogAppDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogAppDetailsOK) GetPayload() *models.CatalogAppDetailsDto {
	return o.Payload
}

func (o *CatalogCatalogAppDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogAppDetailsDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppDetailsBadRequest creates a CatalogCatalogAppDetailsBadRequest with default headers values
func NewCatalogCatalogAppDetailsBadRequest() *CatalogCatalogAppDetailsBadRequest {
	return &CatalogCatalogAppDetailsBadRequest{}
}

/*
CatalogCatalogAppDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogCatalogAppDetailsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this catalog catalog app details bad request response has a 2xx status code
func (o *CatalogCatalogAppDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app details bad request response has a 3xx status code
func (o *CatalogCatalogAppDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app details bad request response has a 4xx status code
func (o *CatalogCatalogAppDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app details bad request response has a 5xx status code
func (o *CatalogCatalogAppDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app details bad request response a status code equal to that given
func (o *CatalogCatalogAppDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CatalogCatalogAppDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogAppDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogAppDetailsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppDetailsUnauthorized creates a CatalogCatalogAppDetailsUnauthorized with default headers values
func NewCatalogCatalogAppDetailsUnauthorized() *CatalogCatalogAppDetailsUnauthorized {
	return &CatalogCatalogAppDetailsUnauthorized{}
}

/*
CatalogCatalogAppDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogCatalogAppDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog app details unauthorized response has a 2xx status code
func (o *CatalogCatalogAppDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app details unauthorized response has a 3xx status code
func (o *CatalogCatalogAppDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app details unauthorized response has a 4xx status code
func (o *CatalogCatalogAppDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app details unauthorized response has a 5xx status code
func (o *CatalogCatalogAppDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app details unauthorized response a status code equal to that given
func (o *CatalogCatalogAppDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CatalogCatalogAppDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogAppDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogAppDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppDetailsForbidden creates a CatalogCatalogAppDetailsForbidden with default headers values
func NewCatalogCatalogAppDetailsForbidden() *CatalogCatalogAppDetailsForbidden {
	return &CatalogCatalogAppDetailsForbidden{}
}

/*
CatalogCatalogAppDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogCatalogAppDetailsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog app details forbidden response has a 2xx status code
func (o *CatalogCatalogAppDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app details forbidden response has a 3xx status code
func (o *CatalogCatalogAppDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app details forbidden response has a 4xx status code
func (o *CatalogCatalogAppDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app details forbidden response has a 5xx status code
func (o *CatalogCatalogAppDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app details forbidden response a status code equal to that given
func (o *CatalogCatalogAppDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CatalogCatalogAppDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogAppDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogAppDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppDetailsNotFound creates a CatalogCatalogAppDetailsNotFound with default headers values
func NewCatalogCatalogAppDetailsNotFound() *CatalogCatalogAppDetailsNotFound {
	return &CatalogCatalogAppDetailsNotFound{}
}

/*
CatalogCatalogAppDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogCatalogAppDetailsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog app details not found response has a 2xx status code
func (o *CatalogCatalogAppDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app details not found response has a 3xx status code
func (o *CatalogCatalogAppDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app details not found response has a 4xx status code
func (o *CatalogCatalogAppDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app details not found response has a 5xx status code
func (o *CatalogCatalogAppDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app details not found response a status code equal to that given
func (o *CatalogCatalogAppDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CatalogCatalogAppDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogAppDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogAppDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppDetailsInternalServerError creates a CatalogCatalogAppDetailsInternalServerError with default headers values
func NewCatalogCatalogAppDetailsInternalServerError() *CatalogCatalogAppDetailsInternalServerError {
	return &CatalogCatalogAppDetailsInternalServerError{}
}

/*
CatalogCatalogAppDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogCatalogAppDetailsInternalServerError struct {
}

// IsSuccess returns true when this catalog catalog app details internal server error response has a 2xx status code
func (o *CatalogCatalogAppDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app details internal server error response has a 3xx status code
func (o *CatalogCatalogAppDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app details internal server error response has a 4xx status code
func (o *CatalogCatalogAppDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog app details internal server error response has a 5xx status code
func (o *CatalogCatalogAppDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog catalog app details internal server error response a status code equal to that given
func (o *CatalogCatalogAppDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CatalogCatalogAppDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsInternalServerError ", 500)
}

func (o *CatalogCatalogAppDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/catalog-app/{id}][%d] catalogCatalogAppDetailsInternalServerError ", 500)
}

func (o *CatalogCatalogAppDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
