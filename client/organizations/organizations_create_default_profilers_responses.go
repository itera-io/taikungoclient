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

// OrganizationsCreateDefaultProfilersReader is a Reader for the OrganizationsCreateDefaultProfilers structure.
type OrganizationsCreateDefaultProfilersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsCreateDefaultProfilersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsCreateDefaultProfilersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsCreateDefaultProfilersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsCreateDefaultProfilersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsCreateDefaultProfilersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsCreateDefaultProfilersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsCreateDefaultProfilersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsCreateDefaultProfilersOK creates a OrganizationsCreateDefaultProfilersOK with default headers values
func NewOrganizationsCreateDefaultProfilersOK() *OrganizationsCreateDefaultProfilersOK {
	return &OrganizationsCreateDefaultProfilersOK{}
}

/*
OrganizationsCreateDefaultProfilersOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsCreateDefaultProfilersOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this organizations create default profilers o k response has a 2xx status code
func (o *OrganizationsCreateDefaultProfilersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations create default profilers o k response has a 3xx status code
func (o *OrganizationsCreateDefaultProfilersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create default profilers o k response has a 4xx status code
func (o *OrganizationsCreateDefaultProfilersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations create default profilers o k response has a 5xx status code
func (o *OrganizationsCreateDefaultProfilersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create default profilers o k response a status code equal to that given
func (o *OrganizationsCreateDefaultProfilersOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsCreateDefaultProfilersOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersOK  %+v", 200, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersOK  %+v", 200, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OrganizationsCreateDefaultProfilersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateDefaultProfilersBadRequest creates a OrganizationsCreateDefaultProfilersBadRequest with default headers values
func NewOrganizationsCreateDefaultProfilersBadRequest() *OrganizationsCreateDefaultProfilersBadRequest {
	return &OrganizationsCreateDefaultProfilersBadRequest{}
}

/*
OrganizationsCreateDefaultProfilersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsCreateDefaultProfilersBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this organizations create default profilers bad request response has a 2xx status code
func (o *OrganizationsCreateDefaultProfilersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create default profilers bad request response has a 3xx status code
func (o *OrganizationsCreateDefaultProfilersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create default profilers bad request response has a 4xx status code
func (o *OrganizationsCreateDefaultProfilersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create default profilers bad request response has a 5xx status code
func (o *OrganizationsCreateDefaultProfilersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create default profilers bad request response a status code equal to that given
func (o *OrganizationsCreateDefaultProfilersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OrganizationsCreateDefaultProfilersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OrganizationsCreateDefaultProfilersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateDefaultProfilersUnauthorized creates a OrganizationsCreateDefaultProfilersUnauthorized with default headers values
func NewOrganizationsCreateDefaultProfilersUnauthorized() *OrganizationsCreateDefaultProfilersUnauthorized {
	return &OrganizationsCreateDefaultProfilersUnauthorized{}
}

/*
OrganizationsCreateDefaultProfilersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsCreateDefaultProfilersUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations create default profilers unauthorized response has a 2xx status code
func (o *OrganizationsCreateDefaultProfilersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create default profilers unauthorized response has a 3xx status code
func (o *OrganizationsCreateDefaultProfilersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create default profilers unauthorized response has a 4xx status code
func (o *OrganizationsCreateDefaultProfilersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create default profilers unauthorized response has a 5xx status code
func (o *OrganizationsCreateDefaultProfilersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create default profilers unauthorized response a status code equal to that given
func (o *OrganizationsCreateDefaultProfilersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OrganizationsCreateDefaultProfilersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsCreateDefaultProfilersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateDefaultProfilersForbidden creates a OrganizationsCreateDefaultProfilersForbidden with default headers values
func NewOrganizationsCreateDefaultProfilersForbidden() *OrganizationsCreateDefaultProfilersForbidden {
	return &OrganizationsCreateDefaultProfilersForbidden{}
}

/*
OrganizationsCreateDefaultProfilersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsCreateDefaultProfilersForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations create default profilers forbidden response has a 2xx status code
func (o *OrganizationsCreateDefaultProfilersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create default profilers forbidden response has a 3xx status code
func (o *OrganizationsCreateDefaultProfilersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create default profilers forbidden response has a 4xx status code
func (o *OrganizationsCreateDefaultProfilersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create default profilers forbidden response has a 5xx status code
func (o *OrganizationsCreateDefaultProfilersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create default profilers forbidden response a status code equal to that given
func (o *OrganizationsCreateDefaultProfilersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsCreateDefaultProfilersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsCreateDefaultProfilersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateDefaultProfilersNotFound creates a OrganizationsCreateDefaultProfilersNotFound with default headers values
func NewOrganizationsCreateDefaultProfilersNotFound() *OrganizationsCreateDefaultProfilersNotFound {
	return &OrganizationsCreateDefaultProfilersNotFound{}
}

/*
OrganizationsCreateDefaultProfilersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsCreateDefaultProfilersNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations create default profilers not found response has a 2xx status code
func (o *OrganizationsCreateDefaultProfilersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create default profilers not found response has a 3xx status code
func (o *OrganizationsCreateDefaultProfilersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create default profilers not found response has a 4xx status code
func (o *OrganizationsCreateDefaultProfilersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create default profilers not found response has a 5xx status code
func (o *OrganizationsCreateDefaultProfilersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create default profilers not found response a status code equal to that given
func (o *OrganizationsCreateDefaultProfilersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OrganizationsCreateDefaultProfilersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsCreateDefaultProfilersNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsCreateDefaultProfilersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateDefaultProfilersInternalServerError creates a OrganizationsCreateDefaultProfilersInternalServerError with default headers values
func NewOrganizationsCreateDefaultProfilersInternalServerError() *OrganizationsCreateDefaultProfilersInternalServerError {
	return &OrganizationsCreateDefaultProfilersInternalServerError{}
}

/*
OrganizationsCreateDefaultProfilersInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsCreateDefaultProfilersInternalServerError struct {
}

// IsSuccess returns true when this organizations create default profilers internal server error response has a 2xx status code
func (o *OrganizationsCreateDefaultProfilersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create default profilers internal server error response has a 3xx status code
func (o *OrganizationsCreateDefaultProfilersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create default profilers internal server error response has a 4xx status code
func (o *OrganizationsCreateDefaultProfilersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations create default profilers internal server error response has a 5xx status code
func (o *OrganizationsCreateDefaultProfilersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations create default profilers internal server error response a status code equal to that given
func (o *OrganizationsCreateDefaultProfilersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OrganizationsCreateDefaultProfilersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersInternalServerError ", 500)
}

func (o *OrganizationsCreateDefaultProfilersInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/createdefaultprofilers][%d] organizationsCreateDefaultProfilersInternalServerError ", 500)
}

func (o *OrganizationsCreateDefaultProfilersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
