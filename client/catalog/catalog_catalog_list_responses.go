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

// CatalogCatalogListReader is a Reader for the CatalogCatalogList structure.
type CatalogCatalogListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCatalogListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCatalogListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCatalogListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogCatalogListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogCatalogListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogCatalogListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogCatalogListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogCatalogListOK creates a CatalogCatalogListOK with default headers values
func NewCatalogCatalogListOK() *CatalogCatalogListOK {
	return &CatalogCatalogListOK{}
}

/*
CatalogCatalogListOK describes a response with status code 200, with default header values.

Success
*/
type CatalogCatalogListOK struct {
	Payload *models.CatalogList
}

// IsSuccess returns true when this catalog catalog list o k response has a 2xx status code
func (o *CatalogCatalogListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog catalog list o k response has a 3xx status code
func (o *CatalogCatalogListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog list o k response has a 4xx status code
func (o *CatalogCatalogListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog list o k response has a 5xx status code
func (o *CatalogCatalogListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog list o k response a status code equal to that given
func (o *CatalogCatalogListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog catalog list o k response
func (o *CatalogCatalogListOK) Code() int {
	return 200
}

func (o *CatalogCatalogListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListOK  %+v", 200, o.Payload)
}

func (o *CatalogCatalogListOK) GetPayload() *models.CatalogList {
	return o.Payload
}

func (o *CatalogCatalogListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogListBadRequest creates a CatalogCatalogListBadRequest with default headers values
func NewCatalogCatalogListBadRequest() *CatalogCatalogListBadRequest {
	return &CatalogCatalogListBadRequest{}
}

/*
CatalogCatalogListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogCatalogListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog list bad request response has a 2xx status code
func (o *CatalogCatalogListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog list bad request response has a 3xx status code
func (o *CatalogCatalogListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog list bad request response has a 4xx status code
func (o *CatalogCatalogListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog list bad request response has a 5xx status code
func (o *CatalogCatalogListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog list bad request response a status code equal to that given
func (o *CatalogCatalogListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the catalog catalog list bad request response
func (o *CatalogCatalogListBadRequest) Code() int {
	return 400
}

func (o *CatalogCatalogListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogCatalogListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogListUnauthorized creates a CatalogCatalogListUnauthorized with default headers values
func NewCatalogCatalogListUnauthorized() *CatalogCatalogListUnauthorized {
	return &CatalogCatalogListUnauthorized{}
}

/*
CatalogCatalogListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogCatalogListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog list unauthorized response has a 2xx status code
func (o *CatalogCatalogListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog list unauthorized response has a 3xx status code
func (o *CatalogCatalogListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog list unauthorized response has a 4xx status code
func (o *CatalogCatalogListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog list unauthorized response has a 5xx status code
func (o *CatalogCatalogListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog list unauthorized response a status code equal to that given
func (o *CatalogCatalogListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the catalog catalog list unauthorized response
func (o *CatalogCatalogListUnauthorized) Code() int {
	return 401
}

func (o *CatalogCatalogListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogCatalogListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogListForbidden creates a CatalogCatalogListForbidden with default headers values
func NewCatalogCatalogListForbidden() *CatalogCatalogListForbidden {
	return &CatalogCatalogListForbidden{}
}

/*
CatalogCatalogListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogCatalogListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog list forbidden response has a 2xx status code
func (o *CatalogCatalogListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog list forbidden response has a 3xx status code
func (o *CatalogCatalogListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog list forbidden response has a 4xx status code
func (o *CatalogCatalogListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog list forbidden response has a 5xx status code
func (o *CatalogCatalogListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog list forbidden response a status code equal to that given
func (o *CatalogCatalogListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the catalog catalog list forbidden response
func (o *CatalogCatalogListForbidden) Code() int {
	return 403
}

func (o *CatalogCatalogListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListForbidden  %+v", 403, o.Payload)
}

func (o *CatalogCatalogListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogListNotFound creates a CatalogCatalogListNotFound with default headers values
func NewCatalogCatalogListNotFound() *CatalogCatalogListNotFound {
	return &CatalogCatalogListNotFound{}
}

/*
CatalogCatalogListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogCatalogListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog catalog list not found response has a 2xx status code
func (o *CatalogCatalogListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog list not found response has a 3xx status code
func (o *CatalogCatalogListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog list not found response has a 4xx status code
func (o *CatalogCatalogListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog catalog list not found response has a 5xx status code
func (o *CatalogCatalogListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog catalog list not found response a status code equal to that given
func (o *CatalogCatalogListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the catalog catalog list not found response
func (o *CatalogCatalogListNotFound) Code() int {
	return 404
}

func (o *CatalogCatalogListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListNotFound  %+v", 404, o.Payload)
}

func (o *CatalogCatalogListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogListInternalServerError creates a CatalogCatalogListInternalServerError with default headers values
func NewCatalogCatalogListInternalServerError() *CatalogCatalogListInternalServerError {
	return &CatalogCatalogListInternalServerError{}
}

/*
CatalogCatalogListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogCatalogListInternalServerError struct {
}

// IsSuccess returns true when this catalog catalog list internal server error response has a 2xx status code
func (o *CatalogCatalogListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog catalog list internal server error response has a 3xx status code
func (o *CatalogCatalogListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog catalog list internal server error response has a 4xx status code
func (o *CatalogCatalogListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog catalog list internal server error response has a 5xx status code
func (o *CatalogCatalogListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog catalog list internal server error response a status code equal to that given
func (o *CatalogCatalogListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the catalog catalog list internal server error response
func (o *CatalogCatalogListInternalServerError) Code() int {
	return 500
}

func (o *CatalogCatalogListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListInternalServerError ", 500)
}

func (o *CatalogCatalogListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/list][%d] catalogCatalogListInternalServerError ", 500)
}

func (o *CatalogCatalogListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
