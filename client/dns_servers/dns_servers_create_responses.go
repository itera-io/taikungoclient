// Code generated by go-swagger; DO NOT EDIT.

package dns_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// DNSServersCreateReader is a Reader for the DNSServersCreate structure.
type DNSServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSServersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDNSServersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDNSServersCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDNSServersCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDNSServersCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDNSServersCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDNSServersCreateOK creates a DNSServersCreateOK with default headers values
func NewDNSServersCreateOK() *DNSServersCreateOK {
	return &DNSServersCreateOK{}
}

/*
DNSServersCreateOK describes a response with status code 200, with default header values.

Success
*/
type DNSServersCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this dns servers create o k response has a 2xx status code
func (o *DNSServersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns servers create o k response has a 3xx status code
func (o *DNSServersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers create o k response has a 4xx status code
func (o *DNSServersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers create o k response has a 5xx status code
func (o *DNSServersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers create o k response a status code equal to that given
func (o *DNSServersCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DNSServersCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateOK  %+v", 200, o.Payload)
}

func (o *DNSServersCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateOK  %+v", 200, o.Payload)
}

func (o *DNSServersCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *DNSServersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersCreateBadRequest creates a DNSServersCreateBadRequest with default headers values
func NewDNSServersCreateBadRequest() *DNSServersCreateBadRequest {
	return &DNSServersCreateBadRequest{}
}

/*
DNSServersCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DNSServersCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this dns servers create bad request response has a 2xx status code
func (o *DNSServersCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers create bad request response has a 3xx status code
func (o *DNSServersCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers create bad request response has a 4xx status code
func (o *DNSServersCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers create bad request response has a 5xx status code
func (o *DNSServersCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers create bad request response a status code equal to that given
func (o *DNSServersCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DNSServersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *DNSServersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersCreateUnauthorized creates a DNSServersCreateUnauthorized with default headers values
func NewDNSServersCreateUnauthorized() *DNSServersCreateUnauthorized {
	return &DNSServersCreateUnauthorized{}
}

/*
DNSServersCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DNSServersCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers create unauthorized response has a 2xx status code
func (o *DNSServersCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers create unauthorized response has a 3xx status code
func (o *DNSServersCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers create unauthorized response has a 4xx status code
func (o *DNSServersCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers create unauthorized response has a 5xx status code
func (o *DNSServersCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers create unauthorized response a status code equal to that given
func (o *DNSServersCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DNSServersCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersCreateForbidden creates a DNSServersCreateForbidden with default headers values
func NewDNSServersCreateForbidden() *DNSServersCreateForbidden {
	return &DNSServersCreateForbidden{}
}

/*
DNSServersCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DNSServersCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers create forbidden response has a 2xx status code
func (o *DNSServersCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers create forbidden response has a 3xx status code
func (o *DNSServersCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers create forbidden response has a 4xx status code
func (o *DNSServersCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers create forbidden response has a 5xx status code
func (o *DNSServersCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers create forbidden response a status code equal to that given
func (o *DNSServersCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DNSServersCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersCreateNotFound creates a DNSServersCreateNotFound with default headers values
func NewDNSServersCreateNotFound() *DNSServersCreateNotFound {
	return &DNSServersCreateNotFound{}
}

/*
DNSServersCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DNSServersCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers create not found response has a 2xx status code
func (o *DNSServersCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers create not found response has a 3xx status code
func (o *DNSServersCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers create not found response has a 4xx status code
func (o *DNSServersCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers create not found response has a 5xx status code
func (o *DNSServersCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers create not found response a status code equal to that given
func (o *DNSServersCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DNSServersCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersCreateInternalServerError creates a DNSServersCreateInternalServerError with default headers values
func NewDNSServersCreateInternalServerError() *DNSServersCreateInternalServerError {
	return &DNSServersCreateInternalServerError{}
}

/*
DNSServersCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DNSServersCreateInternalServerError struct {
}

// IsSuccess returns true when this dns servers create internal server error response has a 2xx status code
func (o *DNSServersCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers create internal server error response has a 3xx status code
func (o *DNSServersCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers create internal server error response has a 4xx status code
func (o *DNSServersCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers create internal server error response has a 5xx status code
func (o *DNSServersCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this dns servers create internal server error response a status code equal to that given
func (o *DNSServersCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DNSServersCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateInternalServerError ", 500)
}

func (o *DNSServersCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/DnsServers/create][%d] dnsServersCreateInternalServerError ", 500)
}

func (o *DNSServersCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
