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

// DNSServersDeleteReader is a Reader for the DNSServersDelete structure.
type DNSServersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSServersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSServersDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDNSServersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDNSServersDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDNSServersDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDNSServersDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDNSServersDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDNSServersDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDNSServersDeleteOK creates a DNSServersDeleteOK with default headers values
func NewDNSServersDeleteOK() *DNSServersDeleteOK {
	return &DNSServersDeleteOK{}
}

/*
DNSServersDeleteOK describes a response with status code 200, with default header values.

Success
*/
type DNSServersDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this dns servers delete o k response has a 2xx status code
func (o *DNSServersDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns servers delete o k response has a 3xx status code
func (o *DNSServersDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers delete o k response has a 4xx status code
func (o *DNSServersDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers delete o k response has a 5xx status code
func (o *DNSServersDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers delete o k response a status code equal to that given
func (o *DNSServersDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *DNSServersDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteOK  %+v", 200, o.Payload)
}

func (o *DNSServersDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteOK  %+v", 200, o.Payload)
}

func (o *DNSServersDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *DNSServersDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersDeleteNoContent creates a DNSServersDeleteNoContent with default headers values
func NewDNSServersDeleteNoContent() *DNSServersDeleteNoContent {
	return &DNSServersDeleteNoContent{}
}

/*
DNSServersDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type DNSServersDeleteNoContent struct {
}

// IsSuccess returns true when this dns servers delete no content response has a 2xx status code
func (o *DNSServersDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns servers delete no content response has a 3xx status code
func (o *DNSServersDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers delete no content response has a 4xx status code
func (o *DNSServersDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers delete no content response has a 5xx status code
func (o *DNSServersDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers delete no content response a status code equal to that given
func (o *DNSServersDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DNSServersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteNoContent ", 204)
}

func (o *DNSServersDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteNoContent ", 204)
}

func (o *DNSServersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDNSServersDeleteBadRequest creates a DNSServersDeleteBadRequest with default headers values
func NewDNSServersDeleteBadRequest() *DNSServersDeleteBadRequest {
	return &DNSServersDeleteBadRequest{}
}

/*
DNSServersDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DNSServersDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dns servers delete bad request response has a 2xx status code
func (o *DNSServersDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers delete bad request response has a 3xx status code
func (o *DNSServersDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers delete bad request response has a 4xx status code
func (o *DNSServersDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers delete bad request response has a 5xx status code
func (o *DNSServersDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers delete bad request response a status code equal to that given
func (o *DNSServersDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DNSServersDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DNSServersDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersDeleteUnauthorized creates a DNSServersDeleteUnauthorized with default headers values
func NewDNSServersDeleteUnauthorized() *DNSServersDeleteUnauthorized {
	return &DNSServersDeleteUnauthorized{}
}

/*
DNSServersDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DNSServersDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers delete unauthorized response has a 2xx status code
func (o *DNSServersDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers delete unauthorized response has a 3xx status code
func (o *DNSServersDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers delete unauthorized response has a 4xx status code
func (o *DNSServersDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers delete unauthorized response has a 5xx status code
func (o *DNSServersDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers delete unauthorized response a status code equal to that given
func (o *DNSServersDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DNSServersDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersDeleteForbidden creates a DNSServersDeleteForbidden with default headers values
func NewDNSServersDeleteForbidden() *DNSServersDeleteForbidden {
	return &DNSServersDeleteForbidden{}
}

/*
DNSServersDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DNSServersDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers delete forbidden response has a 2xx status code
func (o *DNSServersDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers delete forbidden response has a 3xx status code
func (o *DNSServersDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers delete forbidden response has a 4xx status code
func (o *DNSServersDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers delete forbidden response has a 5xx status code
func (o *DNSServersDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers delete forbidden response a status code equal to that given
func (o *DNSServersDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DNSServersDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersDeleteNotFound creates a DNSServersDeleteNotFound with default headers values
func NewDNSServersDeleteNotFound() *DNSServersDeleteNotFound {
	return &DNSServersDeleteNotFound{}
}

/*
DNSServersDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DNSServersDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this dns servers delete not found response has a 2xx status code
func (o *DNSServersDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers delete not found response has a 3xx status code
func (o *DNSServersDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers delete not found response has a 4xx status code
func (o *DNSServersDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this dns servers delete not found response has a 5xx status code
func (o *DNSServersDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this dns servers delete not found response a status code equal to that given
func (o *DNSServersDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DNSServersDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *DNSServersDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServersDeleteInternalServerError creates a DNSServersDeleteInternalServerError with default headers values
func NewDNSServersDeleteInternalServerError() *DNSServersDeleteInternalServerError {
	return &DNSServersDeleteInternalServerError{}
}

/*
DNSServersDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DNSServersDeleteInternalServerError struct {
}

// IsSuccess returns true when this dns servers delete internal server error response has a 2xx status code
func (o *DNSServersDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dns servers delete internal server error response has a 3xx status code
func (o *DNSServersDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns servers delete internal server error response has a 4xx status code
func (o *DNSServersDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns servers delete internal server error response has a 5xx status code
func (o *DNSServersDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this dns servers delete internal server error response a status code equal to that given
func (o *DNSServersDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DNSServersDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteInternalServerError ", 500)
}

func (o *DNSServersDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/DnsServers/{id}][%d] dnsServersDeleteInternalServerError ", 500)
}

func (o *DNSServersDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
