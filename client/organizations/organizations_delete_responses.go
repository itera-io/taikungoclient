// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OrganizationsDeleteReader is a Reader for the OrganizationsDelete structure.
type OrganizationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewOrganizationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsDeleteOK creates a OrganizationsDeleteOK with default headers values
func NewOrganizationsDeleteOK() *OrganizationsDeleteOK {
	return &OrganizationsDeleteOK{}
}

/*
OrganizationsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this organizations delete o k response has a 2xx status code
func (o *OrganizationsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations delete o k response has a 3xx status code
func (o *OrganizationsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations delete o k response has a 4xx status code
func (o *OrganizationsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations delete o k response has a 5xx status code
func (o *OrganizationsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations delete o k response a status code equal to that given
func (o *OrganizationsDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteOK  %+v", 200, o.Payload)
}

func (o *OrganizationsDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteOK  %+v", 200, o.Payload)
}

func (o *OrganizationsDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OrganizationsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDeleteNoContent creates a OrganizationsDeleteNoContent with default headers values
func NewOrganizationsDeleteNoContent() *OrganizationsDeleteNoContent {
	return &OrganizationsDeleteNoContent{}
}

/*
OrganizationsDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type OrganizationsDeleteNoContent struct {
}

// IsSuccess returns true when this organizations delete no content response has a 2xx status code
func (o *OrganizationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations delete no content response has a 3xx status code
func (o *OrganizationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations delete no content response has a 4xx status code
func (o *OrganizationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations delete no content response has a 5xx status code
func (o *OrganizationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations delete no content response a status code equal to that given
func (o *OrganizationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *OrganizationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteNoContent ", 204)
}

func (o *OrganizationsDeleteNoContent) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteNoContent ", 204)
}

func (o *OrganizationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationsDeleteBadRequest creates a OrganizationsDeleteBadRequest with default headers values
func NewOrganizationsDeleteBadRequest() *OrganizationsDeleteBadRequest {
	return &OrganizationsDeleteBadRequest{}
}

/*
OrganizationsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations delete bad request response has a 2xx status code
func (o *OrganizationsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations delete bad request response has a 3xx status code
func (o *OrganizationsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations delete bad request response has a 4xx status code
func (o *OrganizationsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations delete bad request response has a 5xx status code
func (o *OrganizationsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations delete bad request response a status code equal to that given
func (o *OrganizationsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OrganizationsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDeleteUnauthorized creates a OrganizationsDeleteUnauthorized with default headers values
func NewOrganizationsDeleteUnauthorized() *OrganizationsDeleteUnauthorized {
	return &OrganizationsDeleteUnauthorized{}
}

/*
OrganizationsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations delete unauthorized response has a 2xx status code
func (o *OrganizationsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations delete unauthorized response has a 3xx status code
func (o *OrganizationsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations delete unauthorized response has a 4xx status code
func (o *OrganizationsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations delete unauthorized response has a 5xx status code
func (o *OrganizationsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations delete unauthorized response a status code equal to that given
func (o *OrganizationsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OrganizationsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDeleteForbidden creates a OrganizationsDeleteForbidden with default headers values
func NewOrganizationsDeleteForbidden() *OrganizationsDeleteForbidden {
	return &OrganizationsDeleteForbidden{}
}

/*
OrganizationsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations delete forbidden response has a 2xx status code
func (o *OrganizationsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations delete forbidden response has a 3xx status code
func (o *OrganizationsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations delete forbidden response has a 4xx status code
func (o *OrganizationsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations delete forbidden response has a 5xx status code
func (o *OrganizationsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations delete forbidden response a status code equal to that given
func (o *OrganizationsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDeleteNotFound creates a OrganizationsDeleteNotFound with default headers values
func NewOrganizationsDeleteNotFound() *OrganizationsDeleteNotFound {
	return &OrganizationsDeleteNotFound{}
}

/*
OrganizationsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations delete not found response has a 2xx status code
func (o *OrganizationsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations delete not found response has a 3xx status code
func (o *OrganizationsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations delete not found response has a 4xx status code
func (o *OrganizationsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations delete not found response has a 5xx status code
func (o *OrganizationsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations delete not found response a status code equal to that given
func (o *OrganizationsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OrganizationsDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDeleteInternalServerError creates a OrganizationsDeleteInternalServerError with default headers values
func NewOrganizationsDeleteInternalServerError() *OrganizationsDeleteInternalServerError {
	return &OrganizationsDeleteInternalServerError{}
}

/*
OrganizationsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsDeleteInternalServerError struct {
}

// IsSuccess returns true when this organizations delete internal server error response has a 2xx status code
func (o *OrganizationsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations delete internal server error response has a 3xx status code
func (o *OrganizationsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations delete internal server error response has a 4xx status code
func (o *OrganizationsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations delete internal server error response has a 5xx status code
func (o *OrganizationsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations delete internal server error response a status code equal to that given
func (o *OrganizationsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OrganizationsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteInternalServerError ", 500)
}

func (o *OrganizationsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/delete/{organizationId}][%d] organizationsDeleteInternalServerError ", 500)
}

func (o *OrganizationsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
