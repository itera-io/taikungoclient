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

// DNSServersListReader is a Reader for the DNSServersList structure.
type DNSServersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSServersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSServersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDNSServersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDNSServersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDNSServersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDNSServersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDNSServersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDNSServersListOK creates a DNSServersListOK with default headers values
func NewDNSServersListOK() *DNSServersListOK {
	return &DNSServersListOK{}
}

/*
DNSServersListOK describes a response with status code 200, with default header values.

Success
*/
type DNSServersListOK struct {
	Payload []*models.DNSServersListDto
}

// IsSuccess returns true when this dns servers list o k response has a 2xx status code
func (o *DNSServersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns servers list o k response has a 3xx status code
func (o *DNSServersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers list o k response has a 4xx status code
func (o *DNSServersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers list o k response has a 5xx status code
func (o *DNSServersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers list o k response a status code equal to that given
func (o *DNSServersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DNSServersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListOK  %+v", 200, o.Payload)
}

func (o *DNSServersListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListOK  %+v", 200, o.Payload)
}

func (o *DNSServersListOK) GetPayload() []*models.DNSServersListDto {
	return o.Payload
}

func (o *DNSServersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersListBadRequest creates a DNSServersListBadRequest with default headers values
func NewDNSServersListBadRequest() *DNSServersListBadRequest {
	return &DNSServersListBadRequest{}
}

/*
DNSServersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DNSServersListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this dns servers list bad request response has a 2xx status code
func (o *DNSServersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers list bad request response has a 3xx status code
func (o *DNSServersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers list bad request response has a 4xx status code
func (o *DNSServersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers list bad request response has a 5xx status code
func (o *DNSServersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers list bad request response a status code equal to that given
func (o *DNSServersListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DNSServersListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *DNSServersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersListUnauthorized creates a DNSServersListUnauthorized with default headers values
func NewDNSServersListUnauthorized() *DNSServersListUnauthorized {
	return &DNSServersListUnauthorized{}
}

/*
DNSServersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DNSServersListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers list unauthorized response has a 2xx status code
func (o *DNSServersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers list unauthorized response has a 3xx status code
func (o *DNSServersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers list unauthorized response has a 4xx status code
func (o *DNSServersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers list unauthorized response has a 5xx status code
func (o *DNSServersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers list unauthorized response a status code equal to that given
func (o *DNSServersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DNSServersListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersListForbidden creates a DNSServersListForbidden with default headers values
func NewDNSServersListForbidden() *DNSServersListForbidden {
	return &DNSServersListForbidden{}
}

/*
DNSServersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DNSServersListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers list forbidden response has a 2xx status code
func (o *DNSServersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers list forbidden response has a 3xx status code
func (o *DNSServersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers list forbidden response has a 4xx status code
func (o *DNSServersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers list forbidden response has a 5xx status code
func (o *DNSServersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers list forbidden response a status code equal to that given
func (o *DNSServersListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DNSServersListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersListNotFound creates a DNSServersListNotFound with default headers values
func NewDNSServersListNotFound() *DNSServersListNotFound {
	return &DNSServersListNotFound{}
}

/*
DNSServersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DNSServersListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers list not found response has a 2xx status code
func (o *DNSServersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers list not found response has a 3xx status code
func (o *DNSServersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers list not found response has a 4xx status code
func (o *DNSServersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers list not found response has a 5xx status code
func (o *DNSServersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers list not found response a status code equal to that given
func (o *DNSServersListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DNSServersListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersListInternalServerError creates a DNSServersListInternalServerError with default headers values
func NewDNSServersListInternalServerError() *DNSServersListInternalServerError {
	return &DNSServersListInternalServerError{}
}

/*
DNSServersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DNSServersListInternalServerError struct {
}

// IsSuccess returns true when this dns servers list internal server error response has a 2xx status code
func (o *DNSServersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers list internal server error response has a 3xx status code
func (o *DNSServersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers list internal server error response has a 4xx status code
func (o *DNSServersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers list internal server error response has a 5xx status code
func (o *DNSServersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this dns servers list internal server error response a status code equal to that given
func (o *DNSServersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DNSServersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListInternalServerError ", 500)
}

func (o *DNSServersListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/DnsServers/list/{accessProfileId}][%d] dnsServersListInternalServerError ", 500)
}

func (o *DNSServersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
