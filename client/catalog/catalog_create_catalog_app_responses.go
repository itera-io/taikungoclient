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

// CatalogCreateCatalogAppReader is a Reader for the CatalogCreateCatalogApp structure.
type CatalogCreateCatalogAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCreateCatalogAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCreateCatalogAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCreateCatalogAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogCreateCatalogAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogCreateCatalogAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogCreateCatalogAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogCreateCatalogAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogCreateCatalogAppOK creates a CatalogCreateCatalogAppOK with default headers values
func NewCatalogCreateCatalogAppOK() *CatalogCreateCatalogAppOK {
	return &CatalogCreateCatalogAppOK{}
}

/*
CatalogCreateCatalogAppOK describes a response with status code 200, with default header values.

Success
*/
type CatalogCreateCatalogAppOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this catalog create catalog app o k response has a 2xx status code
func (o *CatalogCreateCatalogAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog create catalog app o k response has a 3xx status code
func (o *CatalogCreateCatalogAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog app o k response has a 4xx status code
func (o *CatalogCreateCatalogAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog create catalog app o k response has a 5xx status code
func (o *CatalogCreateCatalogAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog app o k response a status code equal to that given
func (o *CatalogCreateCatalogAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog create catalog app o k response
func (o *CatalogCreateCatalogAppOK) Code() int {
	return 200
}

func (o *CatalogCreateCatalogAppOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppOK  %+v", 200, o.Payload)
}

func (o *CatalogCreateCatalogAppOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppOK  %+v", 200, o.Payload)
}

func (o *CatalogCreateCatalogAppOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CatalogCreateCatalogAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogAppBadRequest creates a CatalogCreateCatalogAppBadRequest with default headers values
func NewCatalogCreateCatalogAppBadRequest() *CatalogCreateCatalogAppBadRequest {
	return &CatalogCreateCatalogAppBadRequest{}
}

/*
CatalogCreateCatalogAppBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogCreateCatalogAppBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog app bad request response has a 2xx status code
func (o *CatalogCreateCatalogAppBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog app bad request response has a 3xx status code
func (o *CatalogCreateCatalogAppBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog app bad request response has a 4xx status code
func (o *CatalogCreateCatalogAppBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog app bad request response has a 5xx status code
func (o *CatalogCreateCatalogAppBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog app bad request response a status code equal to that given
func (o *CatalogCreateCatalogAppBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the catalog create catalog app bad request response
func (o *CatalogCreateCatalogAppBadRequest) Code() int {
	return 400
}

func (o *CatalogCreateCatalogAppBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCreateCatalogAppBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCreateCatalogAppBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogAppUnauthorized creates a CatalogCreateCatalogAppUnauthorized with default headers values
func NewCatalogCreateCatalogAppUnauthorized() *CatalogCreateCatalogAppUnauthorized {
	return &CatalogCreateCatalogAppUnauthorized{}
}

/*
CatalogCreateCatalogAppUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogCreateCatalogAppUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog app unauthorized response has a 2xx status code
func (o *CatalogCreateCatalogAppUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog app unauthorized response has a 3xx status code
func (o *CatalogCreateCatalogAppUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog app unauthorized response has a 4xx status code
func (o *CatalogCreateCatalogAppUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog app unauthorized response has a 5xx status code
func (o *CatalogCreateCatalogAppUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog app unauthorized response a status code equal to that given
func (o *CatalogCreateCatalogAppUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the catalog create catalog app unauthorized response
func (o *CatalogCreateCatalogAppUnauthorized) Code() int {
	return 401
}

func (o *CatalogCreateCatalogAppUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCreateCatalogAppUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCreateCatalogAppUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogAppForbidden creates a CatalogCreateCatalogAppForbidden with default headers values
func NewCatalogCreateCatalogAppForbidden() *CatalogCreateCatalogAppForbidden {
	return &CatalogCreateCatalogAppForbidden{}
}

/*
CatalogCreateCatalogAppForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogCreateCatalogAppForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog app forbidden response has a 2xx status code
func (o *CatalogCreateCatalogAppForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog app forbidden response has a 3xx status code
func (o *CatalogCreateCatalogAppForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog app forbidden response has a 4xx status code
func (o *CatalogCreateCatalogAppForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog app forbidden response has a 5xx status code
func (o *CatalogCreateCatalogAppForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog app forbidden response a status code equal to that given
func (o *CatalogCreateCatalogAppForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the catalog create catalog app forbidden response
func (o *CatalogCreateCatalogAppForbidden) Code() int {
	return 403
}

func (o *CatalogCreateCatalogAppForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCreateCatalogAppForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCreateCatalogAppForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogAppNotFound creates a CatalogCreateCatalogAppNotFound with default headers values
func NewCatalogCreateCatalogAppNotFound() *CatalogCreateCatalogAppNotFound {
	return &CatalogCreateCatalogAppNotFound{}
}

/*
CatalogCreateCatalogAppNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogCreateCatalogAppNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog create catalog app not found response has a 2xx status code
func (o *CatalogCreateCatalogAppNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog app not found response has a 3xx status code
func (o *CatalogCreateCatalogAppNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog app not found response has a 4xx status code
func (o *CatalogCreateCatalogAppNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog create catalog app not found response has a 5xx status code
func (o *CatalogCreateCatalogAppNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog create catalog app not found response a status code equal to that given
func (o *CatalogCreateCatalogAppNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the catalog create catalog app not found response
func (o *CatalogCreateCatalogAppNotFound) Code() int {
	return 404
}

func (o *CatalogCreateCatalogAppNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCreateCatalogAppNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCreateCatalogAppNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCreateCatalogAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCreateCatalogAppInternalServerError creates a CatalogCreateCatalogAppInternalServerError with default headers values
func NewCatalogCreateCatalogAppInternalServerError() *CatalogCreateCatalogAppInternalServerError {
	return &CatalogCreateCatalogAppInternalServerError{}
}

/*
CatalogCreateCatalogAppInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogCreateCatalogAppInternalServerError struct {
}

// IsSuccess returns true when this catalog create catalog app internal server error response has a 2xx status code
func (o *CatalogCreateCatalogAppInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog create catalog app internal server error response has a 3xx status code
func (o *CatalogCreateCatalogAppInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog create catalog app internal server error response has a 4xx status code
func (o *CatalogCreateCatalogAppInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog create catalog app internal server error response has a 5xx status code
func (o *CatalogCreateCatalogAppInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog create catalog app internal server error response a status code equal to that given
func (o *CatalogCreateCatalogAppInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the catalog create catalog app internal server error response
func (o *CatalogCreateCatalogAppInternalServerError) Code() int {
	return 500
}

func (o *CatalogCreateCatalogAppInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppInternalServerError ", 500)
}

func (o *CatalogCreateCatalogAppInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-to-catalog][%d] catalogCreateCatalogAppInternalServerError ", 500)
}

func (o *CatalogCreateCatalogAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
