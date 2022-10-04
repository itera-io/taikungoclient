// Code generated by go-swagger; DO NOT EDIT.

package ntp_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// NtpServersCreateReader is a Reader for the NtpServersCreate structure.
type NtpServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NtpServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNtpServersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNtpServersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNtpServersCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNtpServersCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNtpServersCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNtpServersCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNtpServersCreateOK creates a NtpServersCreateOK with default headers values
func NewNtpServersCreateOK() *NtpServersCreateOK {
	return &NtpServersCreateOK{}
}

/*
NtpServersCreateOK describes a response with status code 200, with default header values.

Success
*/
type NtpServersCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this ntp servers create o k response has a 2xx status code
func (o *NtpServersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ntp servers create o k response has a 3xx status code
func (o *NtpServersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers create o k response has a 4xx status code
func (o *NtpServersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ntp servers create o k response has a 5xx status code
func (o *NtpServersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers create o k response a status code equal to that given
func (o *NtpServersCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *NtpServersCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateOK  %+v", 200, o.Payload)
}

func (o *NtpServersCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateOK  %+v", 200, o.Payload)
}

func (o *NtpServersCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *NtpServersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersCreateBadRequest creates a NtpServersCreateBadRequest with default headers values
func NewNtpServersCreateBadRequest() *NtpServersCreateBadRequest {
	return &NtpServersCreateBadRequest{}
}

/*
NtpServersCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NtpServersCreateBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ntp servers create bad request response has a 2xx status code
func (o *NtpServersCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers create bad request response has a 3xx status code
func (o *NtpServersCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers create bad request response has a 4xx status code
func (o *NtpServersCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers create bad request response has a 5xx status code
func (o *NtpServersCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers create bad request response a status code equal to that given
func (o *NtpServersCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *NtpServersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *NtpServersCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *NtpServersCreateBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *NtpServersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersCreateUnauthorized creates a NtpServersCreateUnauthorized with default headers values
func NewNtpServersCreateUnauthorized() *NtpServersCreateUnauthorized {
	return &NtpServersCreateUnauthorized{}
}

/*
NtpServersCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NtpServersCreateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ntp servers create unauthorized response has a 2xx status code
func (o *NtpServersCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers create unauthorized response has a 3xx status code
func (o *NtpServersCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers create unauthorized response has a 4xx status code
func (o *NtpServersCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers create unauthorized response has a 5xx status code
func (o *NtpServersCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers create unauthorized response a status code equal to that given
func (o *NtpServersCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *NtpServersCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *NtpServersCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *NtpServersCreateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *NtpServersCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersCreateForbidden creates a NtpServersCreateForbidden with default headers values
func NewNtpServersCreateForbidden() *NtpServersCreateForbidden {
	return &NtpServersCreateForbidden{}
}

/*
NtpServersCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NtpServersCreateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ntp servers create forbidden response has a 2xx status code
func (o *NtpServersCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers create forbidden response has a 3xx status code
func (o *NtpServersCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers create forbidden response has a 4xx status code
func (o *NtpServersCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers create forbidden response has a 5xx status code
func (o *NtpServersCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers create forbidden response a status code equal to that given
func (o *NtpServersCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NtpServersCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateForbidden  %+v", 403, o.Payload)
}

func (o *NtpServersCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateForbidden  %+v", 403, o.Payload)
}

func (o *NtpServersCreateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *NtpServersCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersCreateNotFound creates a NtpServersCreateNotFound with default headers values
func NewNtpServersCreateNotFound() *NtpServersCreateNotFound {
	return &NtpServersCreateNotFound{}
}

/*
NtpServersCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NtpServersCreateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ntp servers create not found response has a 2xx status code
func (o *NtpServersCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers create not found response has a 3xx status code
func (o *NtpServersCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers create not found response has a 4xx status code
func (o *NtpServersCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers create not found response has a 5xx status code
func (o *NtpServersCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers create not found response a status code equal to that given
func (o *NtpServersCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NtpServersCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateNotFound  %+v", 404, o.Payload)
}

func (o *NtpServersCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateNotFound  %+v", 404, o.Payload)
}

func (o *NtpServersCreateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *NtpServersCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersCreateInternalServerError creates a NtpServersCreateInternalServerError with default headers values
func NewNtpServersCreateInternalServerError() *NtpServersCreateInternalServerError {
	return &NtpServersCreateInternalServerError{}
}

/*
NtpServersCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NtpServersCreateInternalServerError struct {
}

// IsSuccess returns true when this ntp servers create internal server error response has a 2xx status code
func (o *NtpServersCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers create internal server error response has a 3xx status code
func (o *NtpServersCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers create internal server error response has a 4xx status code
func (o *NtpServersCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ntp servers create internal server error response has a 5xx status code
func (o *NtpServersCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ntp servers create internal server error response a status code equal to that given
func (o *NtpServersCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NtpServersCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateInternalServerError ", 500)
}

func (o *NtpServersCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/NtpServers/create][%d] ntpServersCreateInternalServerError ", 500)
}

func (o *NtpServersCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
