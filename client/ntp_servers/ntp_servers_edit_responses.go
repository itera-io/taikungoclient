// Code generated by go-swagger; DO NOT EDIT.

package ntp_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NtpServersEditReader is a Reader for the NtpServersEdit structure.
type NtpServersEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NtpServersEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNtpServersEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNtpServersEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNtpServersEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNtpServersEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNtpServersEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNtpServersEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNtpServersEditOK creates a NtpServersEditOK with default headers values
func NewNtpServersEditOK() *NtpServersEditOK {
	return &NtpServersEditOK{}
}

/*
NtpServersEditOK describes a response with status code 200, with default header values.

Success
*/
type NtpServersEditOK struct {
}

// IsSuccess returns true when this ntp servers edit o k response has a 2xx status code
func (o *NtpServersEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ntp servers edit o k response has a 3xx status code
func (o *NtpServersEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers edit o k response has a 4xx status code
func (o *NtpServersEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ntp servers edit o k response has a 5xx status code
func (o *NtpServersEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers edit o k response a status code equal to that given
func (o *NtpServersEditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ntp servers edit o k response
func (o *NtpServersEditOK) Code() int {
	return 200
}

func (o *NtpServersEditOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditOK ", 200)
}

func (o *NtpServersEditOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditOK ", 200)
}

func (o *NtpServersEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNtpServersEditBadRequest creates a NtpServersEditBadRequest with default headers values
func NewNtpServersEditBadRequest() *NtpServersEditBadRequest {
	return &NtpServersEditBadRequest{}
}

/*
NtpServersEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NtpServersEditBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ntp servers edit bad request response has a 2xx status code
func (o *NtpServersEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers edit bad request response has a 3xx status code
func (o *NtpServersEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers edit bad request response has a 4xx status code
func (o *NtpServersEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers edit bad request response has a 5xx status code
func (o *NtpServersEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers edit bad request response a status code equal to that given
func (o *NtpServersEditBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ntp servers edit bad request response
func (o *NtpServersEditBadRequest) Code() int {
	return 400
}

func (o *NtpServersEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditBadRequest  %+v", 400, o.Payload)
}

func (o *NtpServersEditBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditBadRequest  %+v", 400, o.Payload)
}

func (o *NtpServersEditBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *NtpServersEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersEditUnauthorized creates a NtpServersEditUnauthorized with default headers values
func NewNtpServersEditUnauthorized() *NtpServersEditUnauthorized {
	return &NtpServersEditUnauthorized{}
}

/*
NtpServersEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NtpServersEditUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this ntp servers edit unauthorized response has a 2xx status code
func (o *NtpServersEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers edit unauthorized response has a 3xx status code
func (o *NtpServersEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers edit unauthorized response has a 4xx status code
func (o *NtpServersEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers edit unauthorized response has a 5xx status code
func (o *NtpServersEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers edit unauthorized response a status code equal to that given
func (o *NtpServersEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ntp servers edit unauthorized response
func (o *NtpServersEditUnauthorized) Code() int {
	return 401
}

func (o *NtpServersEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditUnauthorized  %+v", 401, o.Payload)
}

func (o *NtpServersEditUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditUnauthorized  %+v", 401, o.Payload)
}

func (o *NtpServersEditUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *NtpServersEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersEditForbidden creates a NtpServersEditForbidden with default headers values
func NewNtpServersEditForbidden() *NtpServersEditForbidden {
	return &NtpServersEditForbidden{}
}

/*
NtpServersEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NtpServersEditForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this ntp servers edit forbidden response has a 2xx status code
func (o *NtpServersEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers edit forbidden response has a 3xx status code
func (o *NtpServersEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers edit forbidden response has a 4xx status code
func (o *NtpServersEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers edit forbidden response has a 5xx status code
func (o *NtpServersEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers edit forbidden response a status code equal to that given
func (o *NtpServersEditForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ntp servers edit forbidden response
func (o *NtpServersEditForbidden) Code() int {
	return 403
}

func (o *NtpServersEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditForbidden  %+v", 403, o.Payload)
}

func (o *NtpServersEditForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditForbidden  %+v", 403, o.Payload)
}

func (o *NtpServersEditForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *NtpServersEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersEditNotFound creates a NtpServersEditNotFound with default headers values
func NewNtpServersEditNotFound() *NtpServersEditNotFound {
	return &NtpServersEditNotFound{}
}

/*
NtpServersEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NtpServersEditNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this ntp servers edit not found response has a 2xx status code
func (o *NtpServersEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers edit not found response has a 3xx status code
func (o *NtpServersEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers edit not found response has a 4xx status code
func (o *NtpServersEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers edit not found response has a 5xx status code
func (o *NtpServersEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers edit not found response a status code equal to that given
func (o *NtpServersEditNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the ntp servers edit not found response
func (o *NtpServersEditNotFound) Code() int {
	return 404
}

func (o *NtpServersEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditNotFound  %+v", 404, o.Payload)
}

func (o *NtpServersEditNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditNotFound  %+v", 404, o.Payload)
}

func (o *NtpServersEditNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *NtpServersEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersEditInternalServerError creates a NtpServersEditInternalServerError with default headers values
func NewNtpServersEditInternalServerError() *NtpServersEditInternalServerError {
	return &NtpServersEditInternalServerError{}
}

/*
NtpServersEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NtpServersEditInternalServerError struct {
}

// IsSuccess returns true when this ntp servers edit internal server error response has a 2xx status code
func (o *NtpServersEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers edit internal server error response has a 3xx status code
func (o *NtpServersEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers edit internal server error response has a 4xx status code
func (o *NtpServersEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ntp servers edit internal server error response has a 5xx status code
func (o *NtpServersEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ntp servers edit internal server error response a status code equal to that given
func (o *NtpServersEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ntp servers edit internal server error response
func (o *NtpServersEditInternalServerError) Code() int {
	return 500
}

func (o *NtpServersEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditInternalServerError ", 500)
}

func (o *NtpServersEditInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/NtpServers/edit/{id}][%d] ntpServersEditInternalServerError ", 500)
}

func (o *NtpServersEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
