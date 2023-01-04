// Code generated by go-swagger; DO NOT EDIT.

package allowed_host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AllowedHostListReader is a Reader for the AllowedHostList structure.
type AllowedHostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllowedHostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllowedHostListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllowedHostListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAllowedHostListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAllowedHostListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllowedHostListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAllowedHostListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllowedHostListOK creates a AllowedHostListOK with default headers values
func NewAllowedHostListOK() *AllowedHostListOK {
	return &AllowedHostListOK{}
}

/*
AllowedHostListOK describes a response with status code 200, with default header values.

Success
*/
type AllowedHostListOK struct {
	Payload *models.AllowedHostList
}

// IsSuccess returns true when this allowed host list o k response has a 2xx status code
func (o *AllowedHostListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this allowed host list o k response has a 3xx status code
func (o *AllowedHostListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host list o k response has a 4xx status code
func (o *AllowedHostListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this allowed host list o k response has a 5xx status code
func (o *AllowedHostListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host list o k response a status code equal to that given
func (o *AllowedHostListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AllowedHostListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListOK  %+v", 200, o.Payload)
}

func (o *AllowedHostListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListOK  %+v", 200, o.Payload)
}

func (o *AllowedHostListOK) GetPayload() *models.AllowedHostList {
	return o.Payload
}

func (o *AllowedHostListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AllowedHostList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostListBadRequest creates a AllowedHostListBadRequest with default headers values
func NewAllowedHostListBadRequest() *AllowedHostListBadRequest {
	return &AllowedHostListBadRequest{}
}

/*
AllowedHostListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllowedHostListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this allowed host list bad request response has a 2xx status code
func (o *AllowedHostListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host list bad request response has a 3xx status code
func (o *AllowedHostListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host list bad request response has a 4xx status code
func (o *AllowedHostListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host list bad request response has a 5xx status code
func (o *AllowedHostListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host list bad request response a status code equal to that given
func (o *AllowedHostListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AllowedHostListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListBadRequest  %+v", 400, o.Payload)
}

func (o *AllowedHostListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListBadRequest  %+v", 400, o.Payload)
}

func (o *AllowedHostListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *AllowedHostListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostListUnauthorized creates a AllowedHostListUnauthorized with default headers values
func NewAllowedHostListUnauthorized() *AllowedHostListUnauthorized {
	return &AllowedHostListUnauthorized{}
}

/*
AllowedHostListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AllowedHostListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this allowed host list unauthorized response has a 2xx status code
func (o *AllowedHostListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host list unauthorized response has a 3xx status code
func (o *AllowedHostListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host list unauthorized response has a 4xx status code
func (o *AllowedHostListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host list unauthorized response has a 5xx status code
func (o *AllowedHostListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host list unauthorized response a status code equal to that given
func (o *AllowedHostListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AllowedHostListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListUnauthorized  %+v", 401, o.Payload)
}

func (o *AllowedHostListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListUnauthorized  %+v", 401, o.Payload)
}

func (o *AllowedHostListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AllowedHostListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostListForbidden creates a AllowedHostListForbidden with default headers values
func NewAllowedHostListForbidden() *AllowedHostListForbidden {
	return &AllowedHostListForbidden{}
}

/*
AllowedHostListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AllowedHostListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this allowed host list forbidden response has a 2xx status code
func (o *AllowedHostListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host list forbidden response has a 3xx status code
func (o *AllowedHostListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host list forbidden response has a 4xx status code
func (o *AllowedHostListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host list forbidden response has a 5xx status code
func (o *AllowedHostListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host list forbidden response a status code equal to that given
func (o *AllowedHostListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AllowedHostListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListForbidden  %+v", 403, o.Payload)
}

func (o *AllowedHostListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListForbidden  %+v", 403, o.Payload)
}

func (o *AllowedHostListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AllowedHostListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostListNotFound creates a AllowedHostListNotFound with default headers values
func NewAllowedHostListNotFound() *AllowedHostListNotFound {
	return &AllowedHostListNotFound{}
}

/*
AllowedHostListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllowedHostListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this allowed host list not found response has a 2xx status code
func (o *AllowedHostListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host list not found response has a 3xx status code
func (o *AllowedHostListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host list not found response has a 4xx status code
func (o *AllowedHostListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host list not found response has a 5xx status code
func (o *AllowedHostListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host list not found response a status code equal to that given
func (o *AllowedHostListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AllowedHostListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListNotFound  %+v", 404, o.Payload)
}

func (o *AllowedHostListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListNotFound  %+v", 404, o.Payload)
}

func (o *AllowedHostListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AllowedHostListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostListInternalServerError creates a AllowedHostListInternalServerError with default headers values
func NewAllowedHostListInternalServerError() *AllowedHostListInternalServerError {
	return &AllowedHostListInternalServerError{}
}

/*
AllowedHostListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AllowedHostListInternalServerError struct {
}

// IsSuccess returns true when this allowed host list internal server error response has a 2xx status code
func (o *AllowedHostListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host list internal server error response has a 3xx status code
func (o *AllowedHostListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host list internal server error response has a 4xx status code
func (o *AllowedHostListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this allowed host list internal server error response has a 5xx status code
func (o *AllowedHostListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this allowed host list internal server error response a status code equal to that given
func (o *AllowedHostListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AllowedHostListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListInternalServerError ", 500)
}

func (o *AllowedHostListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AllowedHost/list/{accessProfileId}][%d] allowedHostListInternalServerError ", 500)
}

func (o *AllowedHostListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
