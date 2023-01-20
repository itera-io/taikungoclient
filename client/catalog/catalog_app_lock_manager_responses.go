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

// CatalogAppLockManagerReader is a Reader for the CatalogAppLockManager structure.
type CatalogAppLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogAppLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogAppLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogAppLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogAppLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogAppLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogAppLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogAppLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogAppLockManagerOK creates a CatalogAppLockManagerOK with default headers values
func NewCatalogAppLockManagerOK() *CatalogAppLockManagerOK {
	return &CatalogAppLockManagerOK{}
}

/*
CatalogAppLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type CatalogAppLockManagerOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this catalog app lock manager o k response has a 2xx status code
func (o *CatalogAppLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog app lock manager o k response has a 3xx status code
func (o *CatalogAppLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog app lock manager o k response has a 4xx status code
func (o *CatalogAppLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog app lock manager o k response has a 5xx status code
func (o *CatalogAppLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog app lock manager o k response a status code equal to that given
func (o *CatalogAppLockManagerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog app lock manager o k response
func (o *CatalogAppLockManagerOK) Code() int {
	return 200
}

func (o *CatalogAppLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerOK  %+v", 200, o.Payload)
}

func (o *CatalogAppLockManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerOK  %+v", 200, o.Payload)
}

func (o *CatalogAppLockManagerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CatalogAppLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAppLockManagerBadRequest creates a CatalogAppLockManagerBadRequest with default headers values
func NewCatalogAppLockManagerBadRequest() *CatalogAppLockManagerBadRequest {
	return &CatalogAppLockManagerBadRequest{}
}

/*
CatalogAppLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogAppLockManagerBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog app lock manager bad request response has a 2xx status code
func (o *CatalogAppLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog app lock manager bad request response has a 3xx status code
func (o *CatalogAppLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog app lock manager bad request response has a 4xx status code
func (o *CatalogAppLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog app lock manager bad request response has a 5xx status code
func (o *CatalogAppLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog app lock manager bad request response a status code equal to that given
func (o *CatalogAppLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the catalog app lock manager bad request response
func (o *CatalogAppLockManagerBadRequest) Code() int {
	return 400
}

func (o *CatalogAppLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogAppLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogAppLockManagerBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogAppLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAppLockManagerUnauthorized creates a CatalogAppLockManagerUnauthorized with default headers values
func NewCatalogAppLockManagerUnauthorized() *CatalogAppLockManagerUnauthorized {
	return &CatalogAppLockManagerUnauthorized{}
}

/*
CatalogAppLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogAppLockManagerUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog app lock manager unauthorized response has a 2xx status code
func (o *CatalogAppLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog app lock manager unauthorized response has a 3xx status code
func (o *CatalogAppLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog app lock manager unauthorized response has a 4xx status code
func (o *CatalogAppLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog app lock manager unauthorized response has a 5xx status code
func (o *CatalogAppLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog app lock manager unauthorized response a status code equal to that given
func (o *CatalogAppLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the catalog app lock manager unauthorized response
func (o *CatalogAppLockManagerUnauthorized) Code() int {
	return 401
}

func (o *CatalogAppLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogAppLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogAppLockManagerUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogAppLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAppLockManagerForbidden creates a CatalogAppLockManagerForbidden with default headers values
func NewCatalogAppLockManagerForbidden() *CatalogAppLockManagerForbidden {
	return &CatalogAppLockManagerForbidden{}
}

/*
CatalogAppLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogAppLockManagerForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog app lock manager forbidden response has a 2xx status code
func (o *CatalogAppLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog app lock manager forbidden response has a 3xx status code
func (o *CatalogAppLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog app lock manager forbidden response has a 4xx status code
func (o *CatalogAppLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog app lock manager forbidden response has a 5xx status code
func (o *CatalogAppLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog app lock manager forbidden response a status code equal to that given
func (o *CatalogAppLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the catalog app lock manager forbidden response
func (o *CatalogAppLockManagerForbidden) Code() int {
	return 403
}

func (o *CatalogAppLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *CatalogAppLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *CatalogAppLockManagerForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogAppLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAppLockManagerNotFound creates a CatalogAppLockManagerNotFound with default headers values
func NewCatalogAppLockManagerNotFound() *CatalogAppLockManagerNotFound {
	return &CatalogAppLockManagerNotFound{}
}

/*
CatalogAppLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogAppLockManagerNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog app lock manager not found response has a 2xx status code
func (o *CatalogAppLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog app lock manager not found response has a 3xx status code
func (o *CatalogAppLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog app lock manager not found response has a 4xx status code
func (o *CatalogAppLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog app lock manager not found response has a 5xx status code
func (o *CatalogAppLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog app lock manager not found response a status code equal to that given
func (o *CatalogAppLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the catalog app lock manager not found response
func (o *CatalogAppLockManagerNotFound) Code() int {
	return 404
}

func (o *CatalogAppLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *CatalogAppLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *CatalogAppLockManagerNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogAppLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAppLockManagerInternalServerError creates a CatalogAppLockManagerInternalServerError with default headers values
func NewCatalogAppLockManagerInternalServerError() *CatalogAppLockManagerInternalServerError {
	return &CatalogAppLockManagerInternalServerError{}
}

/*
CatalogAppLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogAppLockManagerInternalServerError struct {
}

// IsSuccess returns true when this catalog app lock manager internal server error response has a 2xx status code
func (o *CatalogAppLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog app lock manager internal server error response has a 3xx status code
func (o *CatalogAppLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog app lock manager internal server error response has a 4xx status code
func (o *CatalogAppLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog app lock manager internal server error response has a 5xx status code
func (o *CatalogAppLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog app lock manager internal server error response a status code equal to that given
func (o *CatalogAppLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the catalog app lock manager internal server error response
func (o *CatalogAppLockManagerInternalServerError) Code() int {
	return 500
}

func (o *CatalogAppLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerInternalServerError ", 500)
}

func (o *CatalogAppLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/app-lockmanager][%d] catalogAppLockManagerInternalServerError ", 500)
}

func (o *CatalogAppLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
