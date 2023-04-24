// Code generated by go-swagger; DO NOT EDIT.

package dns_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DNSServersEditReader is a Reader for the DNSServersEdit structure.
type DNSServersEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSServersEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSServersEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDNSServersEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDNSServersEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDNSServersEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDNSServersEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDNSServersEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDNSServersEditOK creates a DNSServersEditOK with default headers values
func NewDNSServersEditOK() *DNSServersEditOK {
	return &DNSServersEditOK{}
}

/*
DNSServersEditOK describes a response with status code 200, with default header values.

Success
*/
type DNSServersEditOK struct {
}

// IsSuccess returns true when this dns servers edit o k response has a 2xx status code
func (o *DNSServersEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns servers edit o k response has a 3xx status code
func (o *DNSServersEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers edit o k response has a 4xx status code
func (o *DNSServersEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers edit o k response has a 5xx status code
func (o *DNSServersEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers edit o k response a status code equal to that given
func (o *DNSServersEditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dns servers edit o k response
func (o *DNSServersEditOK) Code() int {
	return 200
}

func (o *DNSServersEditOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditOK ", 200)
}

func (o *DNSServersEditOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditOK ", 200)
}

func (o *DNSServersEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDNSServersEditBadRequest creates a DNSServersEditBadRequest with default headers values
func NewDNSServersEditBadRequest() *DNSServersEditBadRequest {
	return &DNSServersEditBadRequest{}
}

/*
DNSServersEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DNSServersEditBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dns servers edit bad request response has a 2xx status code
func (o *DNSServersEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers edit bad request response has a 3xx status code
func (o *DNSServersEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers edit bad request response has a 4xx status code
func (o *DNSServersEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers edit bad request response has a 5xx status code
func (o *DNSServersEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers edit bad request response a status code equal to that given
func (o *DNSServersEditBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the dns servers edit bad request response
func (o *DNSServersEditBadRequest) Code() int {
	return 400
}

func (o *DNSServersEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersEditBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersEditBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DNSServersEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersEditUnauthorized creates a DNSServersEditUnauthorized with default headers values
func NewDNSServersEditUnauthorized() *DNSServersEditUnauthorized {
	return &DNSServersEditUnauthorized{}
}

/*
DNSServersEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DNSServersEditUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this dns servers edit unauthorized response has a 2xx status code
func (o *DNSServersEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers edit unauthorized response has a 3xx status code
func (o *DNSServersEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers edit unauthorized response has a 4xx status code
func (o *DNSServersEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers edit unauthorized response has a 5xx status code
func (o *DNSServersEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers edit unauthorized response a status code equal to that given
func (o *DNSServersEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the dns servers edit unauthorized response
func (o *DNSServersEditUnauthorized) Code() int {
	return 401
}

func (o *DNSServersEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersEditUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersEditUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DNSServersEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersEditForbidden creates a DNSServersEditForbidden with default headers values
func NewDNSServersEditForbidden() *DNSServersEditForbidden {
	return &DNSServersEditForbidden{}
}

/*
DNSServersEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DNSServersEditForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this dns servers edit forbidden response has a 2xx status code
func (o *DNSServersEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers edit forbidden response has a 3xx status code
func (o *DNSServersEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers edit forbidden response has a 4xx status code
func (o *DNSServersEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers edit forbidden response has a 5xx status code
func (o *DNSServersEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers edit forbidden response a status code equal to that given
func (o *DNSServersEditForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the dns servers edit forbidden response
func (o *DNSServersEditForbidden) Code() int {
	return 403
}

func (o *DNSServersEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersEditForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersEditForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DNSServersEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersEditNotFound creates a DNSServersEditNotFound with default headers values
func NewDNSServersEditNotFound() *DNSServersEditNotFound {
	return &DNSServersEditNotFound{}
}

/*
DNSServersEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DNSServersEditNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this dns servers edit not found response has a 2xx status code
func (o *DNSServersEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers edit not found response has a 3xx status code
func (o *DNSServersEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers edit not found response has a 4xx status code
func (o *DNSServersEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers edit not found response has a 5xx status code
func (o *DNSServersEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers edit not found response a status code equal to that given
func (o *DNSServersEditNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the dns servers edit not found response
func (o *DNSServersEditNotFound) Code() int {
	return 404
}

func (o *DNSServersEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersEditNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersEditNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DNSServersEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersEditInternalServerError creates a DNSServersEditInternalServerError with default headers values
func NewDNSServersEditInternalServerError() *DNSServersEditInternalServerError {
	return &DNSServersEditInternalServerError{}
}

/*
DNSServersEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DNSServersEditInternalServerError struct {
}

// IsSuccess returns true when this dns servers edit internal server error response has a 2xx status code
func (o *DNSServersEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers edit internal server error response has a 3xx status code
func (o *DNSServersEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers edit internal server error response has a 4xx status code
func (o *DNSServersEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers edit internal server error response has a 5xx status code
func (o *DNSServersEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this dns servers edit internal server error response a status code equal to that given
func (o *DNSServersEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the dns servers edit internal server error response
func (o *DNSServersEditInternalServerError) Code() int {
	return 500
}

func (o *DNSServersEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditInternalServerError ", 500)
}

func (o *DNSServersEditInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditInternalServerError ", 500)
}

func (o *DNSServersEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
