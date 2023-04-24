// Code generated by go-swagger; DO NOT EDIT.

package proxmox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProxmoxBridgeListReader is a Reader for the ProxmoxBridgeList structure.
type ProxmoxBridgeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProxmoxBridgeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProxmoxBridgeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProxmoxBridgeListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProxmoxBridgeListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProxmoxBridgeListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProxmoxBridgeListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProxmoxBridgeListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProxmoxBridgeListOK creates a ProxmoxBridgeListOK with default headers values
func NewProxmoxBridgeListOK() *ProxmoxBridgeListOK {
	return &ProxmoxBridgeListOK{}
}

/*
ProxmoxBridgeListOK describes a response with status code 200, with default header values.

Success
*/
type ProxmoxBridgeListOK struct {
	Payload []string
}

// IsSuccess returns true when this proxmox bridge list o k response has a 2xx status code
func (o *ProxmoxBridgeListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this proxmox bridge list o k response has a 3xx status code
func (o *ProxmoxBridgeListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this proxmox bridge list o k response has a 4xx status code
func (o *ProxmoxBridgeListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this proxmox bridge list o k response has a 5xx status code
func (o *ProxmoxBridgeListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this proxmox bridge list o k response a status code equal to that given
func (o *ProxmoxBridgeListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the proxmox bridge list o k response
func (o *ProxmoxBridgeListOK) Code() int {
	return 200
}

func (o *ProxmoxBridgeListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListOK  %+v", 200, o.Payload)
}

func (o *ProxmoxBridgeListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListOK  %+v", 200, o.Payload)
}

func (o *ProxmoxBridgeListOK) GetPayload() []string {
	return o.Payload
}

func (o *ProxmoxBridgeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProxmoxBridgeListBadRequest creates a ProxmoxBridgeListBadRequest with default headers values
func NewProxmoxBridgeListBadRequest() *ProxmoxBridgeListBadRequest {
	return &ProxmoxBridgeListBadRequest{}
}

/*
ProxmoxBridgeListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProxmoxBridgeListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this proxmox bridge list bad request response has a 2xx status code
func (o *ProxmoxBridgeListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this proxmox bridge list bad request response has a 3xx status code
func (o *ProxmoxBridgeListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this proxmox bridge list bad request response has a 4xx status code
func (o *ProxmoxBridgeListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this proxmox bridge list bad request response has a 5xx status code
func (o *ProxmoxBridgeListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this proxmox bridge list bad request response a status code equal to that given
func (o *ProxmoxBridgeListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the proxmox bridge list bad request response
func (o *ProxmoxBridgeListBadRequest) Code() int {
	return 400
}

func (o *ProxmoxBridgeListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListBadRequest  %+v", 400, o.Payload)
}

func (o *ProxmoxBridgeListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListBadRequest  %+v", 400, o.Payload)
}

func (o *ProxmoxBridgeListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProxmoxBridgeListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProxmoxBridgeListUnauthorized creates a ProxmoxBridgeListUnauthorized with default headers values
func NewProxmoxBridgeListUnauthorized() *ProxmoxBridgeListUnauthorized {
	return &ProxmoxBridgeListUnauthorized{}
}

/*
ProxmoxBridgeListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProxmoxBridgeListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this proxmox bridge list unauthorized response has a 2xx status code
func (o *ProxmoxBridgeListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this proxmox bridge list unauthorized response has a 3xx status code
func (o *ProxmoxBridgeListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this proxmox bridge list unauthorized response has a 4xx status code
func (o *ProxmoxBridgeListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this proxmox bridge list unauthorized response has a 5xx status code
func (o *ProxmoxBridgeListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this proxmox bridge list unauthorized response a status code equal to that given
func (o *ProxmoxBridgeListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the proxmox bridge list unauthorized response
func (o *ProxmoxBridgeListUnauthorized) Code() int {
	return 401
}

func (o *ProxmoxBridgeListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProxmoxBridgeListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProxmoxBridgeListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProxmoxBridgeListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProxmoxBridgeListForbidden creates a ProxmoxBridgeListForbidden with default headers values
func NewProxmoxBridgeListForbidden() *ProxmoxBridgeListForbidden {
	return &ProxmoxBridgeListForbidden{}
}

/*
ProxmoxBridgeListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProxmoxBridgeListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this proxmox bridge list forbidden response has a 2xx status code
func (o *ProxmoxBridgeListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this proxmox bridge list forbidden response has a 3xx status code
func (o *ProxmoxBridgeListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this proxmox bridge list forbidden response has a 4xx status code
func (o *ProxmoxBridgeListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this proxmox bridge list forbidden response has a 5xx status code
func (o *ProxmoxBridgeListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this proxmox bridge list forbidden response a status code equal to that given
func (o *ProxmoxBridgeListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the proxmox bridge list forbidden response
func (o *ProxmoxBridgeListForbidden) Code() int {
	return 403
}

func (o *ProxmoxBridgeListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListForbidden  %+v", 403, o.Payload)
}

func (o *ProxmoxBridgeListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListForbidden  %+v", 403, o.Payload)
}

func (o *ProxmoxBridgeListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProxmoxBridgeListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProxmoxBridgeListNotFound creates a ProxmoxBridgeListNotFound with default headers values
func NewProxmoxBridgeListNotFound() *ProxmoxBridgeListNotFound {
	return &ProxmoxBridgeListNotFound{}
}

/*
ProxmoxBridgeListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProxmoxBridgeListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this proxmox bridge list not found response has a 2xx status code
func (o *ProxmoxBridgeListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this proxmox bridge list not found response has a 3xx status code
func (o *ProxmoxBridgeListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this proxmox bridge list not found response has a 4xx status code
func (o *ProxmoxBridgeListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this proxmox bridge list not found response has a 5xx status code
func (o *ProxmoxBridgeListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this proxmox bridge list not found response a status code equal to that given
func (o *ProxmoxBridgeListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the proxmox bridge list not found response
func (o *ProxmoxBridgeListNotFound) Code() int {
	return 404
}

func (o *ProxmoxBridgeListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListNotFound  %+v", 404, o.Payload)
}

func (o *ProxmoxBridgeListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListNotFound  %+v", 404, o.Payload)
}

func (o *ProxmoxBridgeListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProxmoxBridgeListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProxmoxBridgeListInternalServerError creates a ProxmoxBridgeListInternalServerError with default headers values
func NewProxmoxBridgeListInternalServerError() *ProxmoxBridgeListInternalServerError {
	return &ProxmoxBridgeListInternalServerError{}
}

/*
ProxmoxBridgeListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProxmoxBridgeListInternalServerError struct {
}

// IsSuccess returns true when this proxmox bridge list internal server error response has a 2xx status code
func (o *ProxmoxBridgeListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this proxmox bridge list internal server error response has a 3xx status code
func (o *ProxmoxBridgeListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this proxmox bridge list internal server error response has a 4xx status code
func (o *ProxmoxBridgeListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this proxmox bridge list internal server error response has a 5xx status code
func (o *ProxmoxBridgeListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this proxmox bridge list internal server error response a status code equal to that given
func (o *ProxmoxBridgeListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the proxmox bridge list internal server error response
func (o *ProxmoxBridgeListInternalServerError) Code() int {
	return 500
}

func (o *ProxmoxBridgeListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListInternalServerError ", 500)
}

func (o *ProxmoxBridgeListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Proxmox/bridge-list][%d] proxmoxBridgeListInternalServerError ", 500)
}

func (o *ProxmoxBridgeListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
