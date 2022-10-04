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

// CatalogCatalogAppValueAutocompleteReader is a Reader for the CatalogCatalogAppValueAutocomplete structure.
type CatalogCatalogAppValueAutocompleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCatalogAppValueAutocompleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCatalogAppValueAutocompleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCatalogAppValueAutocompleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogCatalogAppValueAutocompleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogCatalogAppValueAutocompleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogCatalogAppValueAutocompleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogCatalogAppValueAutocompleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogCatalogAppValueAutocompleteOK creates a CatalogCatalogAppValueAutocompleteOK with default headers values
func NewCatalogCatalogAppValueAutocompleteOK() *CatalogCatalogAppValueAutocompleteOK {
	return &CatalogCatalogAppValueAutocompleteOK{}
}

/*
CatalogCatalogAppValueAutocompleteOK describes a response with status code 200, with default header values.

Success
*/
type CatalogCatalogAppValueAutocompleteOK struct {
	Payload []*models.PackageAutocompleteDto
}

// IsSuccess returns true when this catalog catalog app value autocomplete o k response has a 2xx status code
func (o *CatalogCatalogAppValueAutocompleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog catalog app value autocomplete o k response has a 3xx status code
func (o *CatalogCatalogAppValueAutocompleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value autocomplete o k response has a 4xx status code
func (o *CatalogCatalogAppValueAutocompleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog app value autocomplete o k response has a 5xx status code
func (o *CatalogCatalogAppValueAutocompleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value autocomplete o k response a status code equal to that given
func (o *CatalogCatalogAppValueAutocompleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *CatalogCatalogAppValueAutocompleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteOK) GetPayload() []*models.PackageAutocompleteDto {
	return o.Payload
}

func (o *CatalogCatalogAppValueAutocompleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueAutocompleteBadRequest creates a CatalogCatalogAppValueAutocompleteBadRequest with default headers values
func NewCatalogCatalogAppValueAutocompleteBadRequest() *CatalogCatalogAppValueAutocompleteBadRequest {
	return &CatalogCatalogAppValueAutocompleteBadRequest{}
}

/*
CatalogCatalogAppValueAutocompleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogCatalogAppValueAutocompleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this catalog catalog app value autocomplete bad request response has a 2xx status code
func (o *CatalogCatalogAppValueAutocompleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value autocomplete bad request response has a 3xx status code
func (o *CatalogCatalogAppValueAutocompleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value autocomplete bad request response has a 4xx status code
func (o *CatalogCatalogAppValueAutocompleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value autocomplete bad request response has a 5xx status code
func (o *CatalogCatalogAppValueAutocompleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value autocomplete bad request response a status code equal to that given
func (o *CatalogCatalogAppValueAutocompleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CatalogCatalogAppValueAutocompleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogAppValueAutocompleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueAutocompleteUnauthorized creates a CatalogCatalogAppValueAutocompleteUnauthorized with default headers values
func NewCatalogCatalogAppValueAutocompleteUnauthorized() *CatalogCatalogAppValueAutocompleteUnauthorized {
	return &CatalogCatalogAppValueAutocompleteUnauthorized{}
}

/*
CatalogCatalogAppValueAutocompleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogCatalogAppValueAutocompleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog catalog app value autocomplete unauthorized response has a 2xx status code
func (o *CatalogCatalogAppValueAutocompleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value autocomplete unauthorized response has a 3xx status code
func (o *CatalogCatalogAppValueAutocompleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value autocomplete unauthorized response has a 4xx status code
func (o *CatalogCatalogAppValueAutocompleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value autocomplete unauthorized response has a 5xx status code
func (o *CatalogCatalogAppValueAutocompleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value autocomplete unauthorized response a status code equal to that given
func (o *CatalogCatalogAppValueAutocompleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CatalogCatalogAppValueAutocompleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogCatalogAppValueAutocompleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueAutocompleteForbidden creates a CatalogCatalogAppValueAutocompleteForbidden with default headers values
func NewCatalogCatalogAppValueAutocompleteForbidden() *CatalogCatalogAppValueAutocompleteForbidden {
	return &CatalogCatalogAppValueAutocompleteForbidden{}
}

/*
CatalogCatalogAppValueAutocompleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogCatalogAppValueAutocompleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog catalog app value autocomplete forbidden response has a 2xx status code
func (o *CatalogCatalogAppValueAutocompleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value autocomplete forbidden response has a 3xx status code
func (o *CatalogCatalogAppValueAutocompleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value autocomplete forbidden response has a 4xx status code
func (o *CatalogCatalogAppValueAutocompleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value autocomplete forbidden response has a 5xx status code
func (o *CatalogCatalogAppValueAutocompleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value autocomplete forbidden response a status code equal to that given
func (o *CatalogCatalogAppValueAutocompleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CatalogCatalogAppValueAutocompleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogCatalogAppValueAutocompleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueAutocompleteNotFound creates a CatalogCatalogAppValueAutocompleteNotFound with default headers values
func NewCatalogCatalogAppValueAutocompleteNotFound() *CatalogCatalogAppValueAutocompleteNotFound {
	return &CatalogCatalogAppValueAutocompleteNotFound{}
}

/*
CatalogCatalogAppValueAutocompleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogCatalogAppValueAutocompleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog catalog app value autocomplete not found response has a 2xx status code
func (o *CatalogCatalogAppValueAutocompleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value autocomplete not found response has a 3xx status code
func (o *CatalogCatalogAppValueAutocompleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value autocomplete not found response has a 4xx status code
func (o *CatalogCatalogAppValueAutocompleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog app value autocomplete not found response has a 5xx status code
func (o *CatalogCatalogAppValueAutocompleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog app value autocomplete not found response a status code equal to that given
func (o *CatalogCatalogAppValueAutocompleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CatalogCatalogAppValueAutocompleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogAppValueAutocompleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogCatalogAppValueAutocompleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogAppValueAutocompleteInternalServerError creates a CatalogCatalogAppValueAutocompleteInternalServerError with default headers values
func NewCatalogCatalogAppValueAutocompleteInternalServerError() *CatalogCatalogAppValueAutocompleteInternalServerError {
	return &CatalogCatalogAppValueAutocompleteInternalServerError{}
}

/*
CatalogCatalogAppValueAutocompleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogCatalogAppValueAutocompleteInternalServerError struct {
}

// IsSuccess returns true when this catalog catalog app value autocomplete internal server error response has a 2xx status code
func (o *CatalogCatalogAppValueAutocompleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog app value autocomplete internal server error response has a 3xx status code
func (o *CatalogCatalogAppValueAutocompleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog app value autocomplete internal server error response has a 4xx status code
func (o *CatalogCatalogAppValueAutocompleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog app value autocomplete internal server error response has a 5xx status code
func (o *CatalogCatalogAppValueAutocompleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog catalog app value autocomplete internal server error response a status code equal to that given
func (o *CatalogCatalogAppValueAutocompleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CatalogCatalogAppValueAutocompleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteInternalServerError ", 500)
}

func (o *CatalogCatalogAppValueAutocompleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/package-value-autocomplete][%d] catalogCatalogAppValueAutocompleteInternalServerError ", 500)
}

func (o *CatalogCatalogAppValueAutocompleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
