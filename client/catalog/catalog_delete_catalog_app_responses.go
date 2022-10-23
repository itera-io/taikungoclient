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

// CatalogDeleteCatalogAppReader is a Reader for the CatalogDeleteCatalogApp structure.
type CatalogDeleteCatalogAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogDeleteCatalogAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogDeleteCatalogAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogDeleteCatalogAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogDeleteCatalogAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogDeleteCatalogAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogDeleteCatalogAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogDeleteCatalogAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogDeleteCatalogAppOK creates a CatalogDeleteCatalogAppOK with default headers values
func NewCatalogDeleteCatalogAppOK() *CatalogDeleteCatalogAppOK {
	return &CatalogDeleteCatalogAppOK{}
}

/*
CatalogDeleteCatalogAppOK describes a response with status code 200, with default header values.

Success
*/
type CatalogDeleteCatalogAppOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this catalog delete catalog app o k response has a 2xx status code
func (o *CatalogDeleteCatalogAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog delete catalog app o k response has a 3xx status code
func (o *CatalogDeleteCatalogAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog delete catalog app o k response has a 4xx status code
func (o *CatalogDeleteCatalogAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog delete catalog app o k response has a 5xx status code
func (o *CatalogDeleteCatalogAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog delete catalog app o k response a status code equal to that given
func (o *CatalogDeleteCatalogAppOK) IsCode(code int) bool {
	return code == 200
}

func (o *CatalogDeleteCatalogAppOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppOK  %+v", 200, o.Payload)
}

func (o *CatalogDeleteCatalogAppOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppOK  %+v", 200, o.Payload)
}

func (o *CatalogDeleteCatalogAppOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CatalogDeleteCatalogAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogAppBadRequest creates a CatalogDeleteCatalogAppBadRequest with default headers values
func NewCatalogDeleteCatalogAppBadRequest() *CatalogDeleteCatalogAppBadRequest {
	return &CatalogDeleteCatalogAppBadRequest{}
}

/*
CatalogDeleteCatalogAppBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogDeleteCatalogAppBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this catalog delete catalog app bad request response has a 2xx status code
func (o *CatalogDeleteCatalogAppBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog delete catalog app bad request response has a 3xx status code
func (o *CatalogDeleteCatalogAppBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog delete catalog app bad request response has a 4xx status code
func (o *CatalogDeleteCatalogAppBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog delete catalog app bad request response has a 5xx status code
func (o *CatalogDeleteCatalogAppBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog delete catalog app bad request response a status code equal to that given
func (o *CatalogDeleteCatalogAppBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CatalogDeleteCatalogAppBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogDeleteCatalogAppBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogDeleteCatalogAppBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CatalogDeleteCatalogAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogAppUnauthorized creates a CatalogDeleteCatalogAppUnauthorized with default headers values
func NewCatalogDeleteCatalogAppUnauthorized() *CatalogDeleteCatalogAppUnauthorized {
	return &CatalogDeleteCatalogAppUnauthorized{}
}

/*
CatalogDeleteCatalogAppUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogDeleteCatalogAppUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog delete catalog app unauthorized response has a 2xx status code
func (o *CatalogDeleteCatalogAppUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog delete catalog app unauthorized response has a 3xx status code
func (o *CatalogDeleteCatalogAppUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog delete catalog app unauthorized response has a 4xx status code
func (o *CatalogDeleteCatalogAppUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog delete catalog app unauthorized response has a 5xx status code
func (o *CatalogDeleteCatalogAppUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog delete catalog app unauthorized response a status code equal to that given
func (o *CatalogDeleteCatalogAppUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CatalogDeleteCatalogAppUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogDeleteCatalogAppUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogDeleteCatalogAppUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogDeleteCatalogAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogAppForbidden creates a CatalogDeleteCatalogAppForbidden with default headers values
func NewCatalogDeleteCatalogAppForbidden() *CatalogDeleteCatalogAppForbidden {
	return &CatalogDeleteCatalogAppForbidden{}
}

/*
CatalogDeleteCatalogAppForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogDeleteCatalogAppForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog delete catalog app forbidden response has a 2xx status code
func (o *CatalogDeleteCatalogAppForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog delete catalog app forbidden response has a 3xx status code
func (o *CatalogDeleteCatalogAppForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog delete catalog app forbidden response has a 4xx status code
func (o *CatalogDeleteCatalogAppForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog delete catalog app forbidden response has a 5xx status code
func (o *CatalogDeleteCatalogAppForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog delete catalog app forbidden response a status code equal to that given
func (o *CatalogDeleteCatalogAppForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CatalogDeleteCatalogAppForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppForbidden  %+v", 403, o.Payload)
}

func (o *CatalogDeleteCatalogAppForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppForbidden  %+v", 403, o.Payload)
}

func (o *CatalogDeleteCatalogAppForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogDeleteCatalogAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogAppNotFound creates a CatalogDeleteCatalogAppNotFound with default headers values
func NewCatalogDeleteCatalogAppNotFound() *CatalogDeleteCatalogAppNotFound {
	return &CatalogDeleteCatalogAppNotFound{}
}

/*
CatalogDeleteCatalogAppNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogDeleteCatalogAppNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog delete catalog app not found response has a 2xx status code
func (o *CatalogDeleteCatalogAppNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog delete catalog app not found response has a 3xx status code
func (o *CatalogDeleteCatalogAppNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog delete catalog app not found response has a 4xx status code
func (o *CatalogDeleteCatalogAppNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog delete catalog app not found response has a 5xx status code
func (o *CatalogDeleteCatalogAppNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog delete catalog app not found response a status code equal to that given
func (o *CatalogDeleteCatalogAppNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CatalogDeleteCatalogAppNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppNotFound  %+v", 404, o.Payload)
}

func (o *CatalogDeleteCatalogAppNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppNotFound  %+v", 404, o.Payload)
}

func (o *CatalogDeleteCatalogAppNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogDeleteCatalogAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogAppInternalServerError creates a CatalogDeleteCatalogAppInternalServerError with default headers values
func NewCatalogDeleteCatalogAppInternalServerError() *CatalogDeleteCatalogAppInternalServerError {
	return &CatalogDeleteCatalogAppInternalServerError{}
}

/*
CatalogDeleteCatalogAppInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogDeleteCatalogAppInternalServerError struct {
}

// IsSuccess returns true when this catalog delete catalog app internal server error response has a 2xx status code
func (o *CatalogDeleteCatalogAppInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog delete catalog app internal server error response has a 3xx status code
func (o *CatalogDeleteCatalogAppInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog delete catalog app internal server error response has a 4xx status code
func (o *CatalogDeleteCatalogAppInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog delete catalog app internal server error response has a 5xx status code
func (o *CatalogDeleteCatalogAppInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog delete catalog app internal server error response a status code equal to that given
func (o *CatalogDeleteCatalogAppInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CatalogDeleteCatalogAppInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppInternalServerError ", 500)
}

func (o *CatalogDeleteCatalogAppInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete-app/{id}][%d] catalogDeleteCatalogAppInternalServerError ", 500)
}

func (o *CatalogDeleteCatalogAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
