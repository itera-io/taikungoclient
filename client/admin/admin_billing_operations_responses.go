// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminBillingOperationsReader is a Reader for the AdminBillingOperations structure.
type AdminBillingOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminBillingOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminBillingOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminBillingOperationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminBillingOperationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminBillingOperationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminBillingOperationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminBillingOperationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminBillingOperationsOK creates a AdminBillingOperationsOK with default headers values
func NewAdminBillingOperationsOK() *AdminBillingOperationsOK {
	return &AdminBillingOperationsOK{}
}

/*
AdminBillingOperationsOK describes a response with status code 200, with default header values.

Success
*/
type AdminBillingOperationsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this admin billing operations o k response has a 2xx status code
func (o *AdminBillingOperationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin billing operations o k response has a 3xx status code
func (o *AdminBillingOperationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin billing operations o k response has a 4xx status code
func (o *AdminBillingOperationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin billing operations o k response has a 5xx status code
func (o *AdminBillingOperationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin billing operations o k response a status code equal to that given
func (o *AdminBillingOperationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin billing operations o k response
func (o *AdminBillingOperationsOK) Code() int {
	return 200
}

func (o *AdminBillingOperationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsOK  %+v", 200, o.Payload)
}

func (o *AdminBillingOperationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsOK  %+v", 200, o.Payload)
}

func (o *AdminBillingOperationsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminBillingOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminBillingOperationsBadRequest creates a AdminBillingOperationsBadRequest with default headers values
func NewAdminBillingOperationsBadRequest() *AdminBillingOperationsBadRequest {
	return &AdminBillingOperationsBadRequest{}
}

/*
AdminBillingOperationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminBillingOperationsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin billing operations bad request response has a 2xx status code
func (o *AdminBillingOperationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin billing operations bad request response has a 3xx status code
func (o *AdminBillingOperationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin billing operations bad request response has a 4xx status code
func (o *AdminBillingOperationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin billing operations bad request response has a 5xx status code
func (o *AdminBillingOperationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin billing operations bad request response a status code equal to that given
func (o *AdminBillingOperationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin billing operations bad request response
func (o *AdminBillingOperationsBadRequest) Code() int {
	return 400
}

func (o *AdminBillingOperationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *AdminBillingOperationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *AdminBillingOperationsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminBillingOperationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminBillingOperationsUnauthorized creates a AdminBillingOperationsUnauthorized with default headers values
func NewAdminBillingOperationsUnauthorized() *AdminBillingOperationsUnauthorized {
	return &AdminBillingOperationsUnauthorized{}
}

/*
AdminBillingOperationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminBillingOperationsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin billing operations unauthorized response has a 2xx status code
func (o *AdminBillingOperationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin billing operations unauthorized response has a 3xx status code
func (o *AdminBillingOperationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin billing operations unauthorized response has a 4xx status code
func (o *AdminBillingOperationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin billing operations unauthorized response has a 5xx status code
func (o *AdminBillingOperationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin billing operations unauthorized response a status code equal to that given
func (o *AdminBillingOperationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin billing operations unauthorized response
func (o *AdminBillingOperationsUnauthorized) Code() int {
	return 401
}

func (o *AdminBillingOperationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminBillingOperationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminBillingOperationsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminBillingOperationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminBillingOperationsForbidden creates a AdminBillingOperationsForbidden with default headers values
func NewAdminBillingOperationsForbidden() *AdminBillingOperationsForbidden {
	return &AdminBillingOperationsForbidden{}
}

/*
AdminBillingOperationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminBillingOperationsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin billing operations forbidden response has a 2xx status code
func (o *AdminBillingOperationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin billing operations forbidden response has a 3xx status code
func (o *AdminBillingOperationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin billing operations forbidden response has a 4xx status code
func (o *AdminBillingOperationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin billing operations forbidden response has a 5xx status code
func (o *AdminBillingOperationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin billing operations forbidden response a status code equal to that given
func (o *AdminBillingOperationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin billing operations forbidden response
func (o *AdminBillingOperationsForbidden) Code() int {
	return 403
}

func (o *AdminBillingOperationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsForbidden  %+v", 403, o.Payload)
}

func (o *AdminBillingOperationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsForbidden  %+v", 403, o.Payload)
}

func (o *AdminBillingOperationsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminBillingOperationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminBillingOperationsNotFound creates a AdminBillingOperationsNotFound with default headers values
func NewAdminBillingOperationsNotFound() *AdminBillingOperationsNotFound {
	return &AdminBillingOperationsNotFound{}
}

/*
AdminBillingOperationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminBillingOperationsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin billing operations not found response has a 2xx status code
func (o *AdminBillingOperationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin billing operations not found response has a 3xx status code
func (o *AdminBillingOperationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin billing operations not found response has a 4xx status code
func (o *AdminBillingOperationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin billing operations not found response has a 5xx status code
func (o *AdminBillingOperationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin billing operations not found response a status code equal to that given
func (o *AdminBillingOperationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin billing operations not found response
func (o *AdminBillingOperationsNotFound) Code() int {
	return 404
}

func (o *AdminBillingOperationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsNotFound  %+v", 404, o.Payload)
}

func (o *AdminBillingOperationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsNotFound  %+v", 404, o.Payload)
}

func (o *AdminBillingOperationsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminBillingOperationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminBillingOperationsInternalServerError creates a AdminBillingOperationsInternalServerError with default headers values
func NewAdminBillingOperationsInternalServerError() *AdminBillingOperationsInternalServerError {
	return &AdminBillingOperationsInternalServerError{}
}

/*
AdminBillingOperationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminBillingOperationsInternalServerError struct {
}

// IsSuccess returns true when this admin billing operations internal server error response has a 2xx status code
func (o *AdminBillingOperationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin billing operations internal server error response has a 3xx status code
func (o *AdminBillingOperationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin billing operations internal server error response has a 4xx status code
func (o *AdminBillingOperationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin billing operations internal server error response has a 5xx status code
func (o *AdminBillingOperationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin billing operations internal server error response a status code equal to that given
func (o *AdminBillingOperationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin billing operations internal server error response
func (o *AdminBillingOperationsInternalServerError) Code() int {
	return 500
}

func (o *AdminBillingOperationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsInternalServerError ", 500)
}

func (o *AdminBillingOperationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/cloudcredentials/billing][%d] adminBillingOperationsInternalServerError ", 500)
}

func (o *AdminBillingOperationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
