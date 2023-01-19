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

// CatalogCreateCatalogReader is a Reader for the CatalogCreateCatalog structure.
type CatalogCreateCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCreateCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCreateCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCreateCatalogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogCreateCatalogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogCreateCatalogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogCreateCatalogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogCreateCatalogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogCreateCatalogOK creates a CatalogCreateCatalogOK with default headers values
func NewCatalogCreateCatalogOK() *CatalogCreateCatalogOK {
	return &CatalogCreateCatalogOK{}
}

/*
CatalogCreateCatalogOK describes a response with status code 200, with default header values.

Success
*/
type CatalogCreateCatalogOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this catalog create catalog o k response has a 2xx status code
func (o *CatalogCreateCatalogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog create catalog o k response has a 3xx status code
func (o *CatalogCreateCatalogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog o k response has a 4xx status code
func (o *CatalogCreateCatalogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog create catalog o k response has a 5xx status code
func (o *CatalogCreateCatalogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog o k response a status code equal to that given
func (o *CatalogCreateCatalogOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog create catalog o k response
func (o *CatalogCreateCatalogOK) Code() int {
	return 200
}

func (o *CatalogCreateCatalogOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogOK  %+v", 200, o.Payload)
}

func (o *CatalogCreateCatalogOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogOK  %+v", 200, o.Payload)
}

func (o *CatalogCreateCatalogOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CatalogCreateCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogBadRequest creates a CatalogCreateCatalogBadRequest with default headers values
func NewCatalogCreateCatalogBadRequest() *CatalogCreateCatalogBadRequest {
	return &CatalogCreateCatalogBadRequest{}
}

/*
CatalogCreateCatalogBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogCreateCatalogBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog bad request response has a 2xx status code
func (o *CatalogCreateCatalogBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog bad request response has a 3xx status code
func (o *CatalogCreateCatalogBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog bad request response has a 4xx status code
func (o *CatalogCreateCatalogBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog bad request response has a 5xx status code
func (o *CatalogCreateCatalogBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog bad request response a status code equal to that given
func (o *CatalogCreateCatalogBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the catalog create catalog bad request response
func (o *CatalogCreateCatalogBadRequest) Code() int {
	return 400
}

func (o *CatalogCreateCatalogBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCreateCatalogBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCreateCatalogBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogUnauthorized creates a CatalogCreateCatalogUnauthorized with default headers values
func NewCatalogCreateCatalogUnauthorized() *CatalogCreateCatalogUnauthorized {
	return &CatalogCreateCatalogUnauthorized{}
}

/*
CatalogCreateCatalogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogCreateCatalogUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog unauthorized response has a 2xx status code
func (o *CatalogCreateCatalogUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog unauthorized response has a 3xx status code
func (o *CatalogCreateCatalogUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog unauthorized response has a 4xx status code
func (o *CatalogCreateCatalogUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog unauthorized response has a 5xx status code
func (o *CatalogCreateCatalogUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog unauthorized response a status code equal to that given
func (o *CatalogCreateCatalogUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the catalog create catalog unauthorized response
func (o *CatalogCreateCatalogUnauthorized) Code() int {
	return 401
}

func (o *CatalogCreateCatalogUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCreateCatalogUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCreateCatalogUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogForbidden creates a CatalogCreateCatalogForbidden with default headers values
func NewCatalogCreateCatalogForbidden() *CatalogCreateCatalogForbidden {
	return &CatalogCreateCatalogForbidden{}
}

/*
CatalogCreateCatalogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogCreateCatalogForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog forbidden response has a 2xx status code
func (o *CatalogCreateCatalogForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog forbidden response has a 3xx status code
func (o *CatalogCreateCatalogForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog forbidden response has a 4xx status code
func (o *CatalogCreateCatalogForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog forbidden response has a 5xx status code
func (o *CatalogCreateCatalogForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog forbidden response a status code equal to that given
func (o *CatalogCreateCatalogForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the catalog create catalog forbidden response
func (o *CatalogCreateCatalogForbidden) Code() int {
	return 403
}

func (o *CatalogCreateCatalogForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCreateCatalogForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCreateCatalogForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogNotFound creates a CatalogCreateCatalogNotFound with default headers values
func NewCatalogCreateCatalogNotFound() *CatalogCreateCatalogNotFound {
	return &CatalogCreateCatalogNotFound{}
}

/*
CatalogCreateCatalogNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogCreateCatalogNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog not found response has a 2xx status code
func (o *CatalogCreateCatalogNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog not found response has a 3xx status code
func (o *CatalogCreateCatalogNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog not found response has a 4xx status code
func (o *CatalogCreateCatalogNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog not found response has a 5xx status code
func (o *CatalogCreateCatalogNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog not found response a status code equal to that given
func (o *CatalogCreateCatalogNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the catalog create catalog not found response
func (o *CatalogCreateCatalogNotFound) Code() int {
	return 404
}

func (o *CatalogCreateCatalogNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCreateCatalogNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCreateCatalogNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogInternalServerError creates a CatalogCreateCatalogInternalServerError with default headers values
func NewCatalogCreateCatalogInternalServerError() *CatalogCreateCatalogInternalServerError {
	return &CatalogCreateCatalogInternalServerError{}
}

/*
CatalogCreateCatalogInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogCreateCatalogInternalServerError struct {
}

// IsSuccess returns true when this catalog create catalog internal server error response has a 2xx status code
func (o *CatalogCreateCatalogInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog internal server error response has a 3xx status code
func (o *CatalogCreateCatalogInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog internal server error response has a 4xx status code
func (o *CatalogCreateCatalogInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog create catalog internal server error response has a 5xx status code
func (o *CatalogCreateCatalogInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog create catalog internal server error response a status code equal to that given
func (o *CatalogCreateCatalogInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the catalog create catalog internal server error response
func (o *CatalogCreateCatalogInternalServerError) Code() int {
	return 500
}

func (o *CatalogCreateCatalogInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogInternalServerError ", 500)
}

func (o *CatalogCreateCatalogInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/create-catalog][%d] catalogCreateCatalogInternalServerError ", 500)
}

func (o *CatalogCreateCatalogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
