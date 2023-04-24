// Code generated by go-swagger; DO NOT EDIT.

package partner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PartnerAddWhiteListDomainReader is a Reader for the PartnerAddWhiteListDomain structure.
type PartnerAddWhiteListDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerAddWhiteListDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerAddWhiteListDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartnerAddWhiteListDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartnerAddWhiteListDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPartnerAddWhiteListDomainForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartnerAddWhiteListDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPartnerAddWhiteListDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartnerAddWhiteListDomainOK creates a PartnerAddWhiteListDomainOK with default headers values
func NewPartnerAddWhiteListDomainOK() *PartnerAddWhiteListDomainOK {
	return &PartnerAddWhiteListDomainOK{}
}

/*
PartnerAddWhiteListDomainOK describes a response with status code 200, with default header values.

Success
*/
type PartnerAddWhiteListDomainOK struct {
}

// IsSuccess returns true when this partner add white list domain o k response has a 2xx status code
func (o *PartnerAddWhiteListDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner add white list domain o k response has a 3xx status code
func (o *PartnerAddWhiteListDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner add white list domain o k response has a 4xx status code
func (o *PartnerAddWhiteListDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner add white list domain o k response has a 5xx status code
func (o *PartnerAddWhiteListDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner add white list domain o k response a status code equal to that given
func (o *PartnerAddWhiteListDomainOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner add white list domain o k response
func (o *PartnerAddWhiteListDomainOK) Code() int {
	return 200
}

func (o *PartnerAddWhiteListDomainOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainOK ", 200)
}

func (o *PartnerAddWhiteListDomainOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainOK ", 200)
}

func (o *PartnerAddWhiteListDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPartnerAddWhiteListDomainBadRequest creates a PartnerAddWhiteListDomainBadRequest with default headers values
func NewPartnerAddWhiteListDomainBadRequest() *PartnerAddWhiteListDomainBadRequest {
	return &PartnerAddWhiteListDomainBadRequest{}
}

/*
PartnerAddWhiteListDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PartnerAddWhiteListDomainBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this partner add white list domain bad request response has a 2xx status code
func (o *PartnerAddWhiteListDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner add white list domain bad request response has a 3xx status code
func (o *PartnerAddWhiteListDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner add white list domain bad request response has a 4xx status code
func (o *PartnerAddWhiteListDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner add white list domain bad request response has a 5xx status code
func (o *PartnerAddWhiteListDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this partner add white list domain bad request response a status code equal to that given
func (o *PartnerAddWhiteListDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the partner add white list domain bad request response
func (o *PartnerAddWhiteListDomainBadRequest) Code() int {
	return 400
}

func (o *PartnerAddWhiteListDomainBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerAddWhiteListDomainBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerAddWhiteListDomainBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerAddWhiteListDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerAddWhiteListDomainUnauthorized creates a PartnerAddWhiteListDomainUnauthorized with default headers values
func NewPartnerAddWhiteListDomainUnauthorized() *PartnerAddWhiteListDomainUnauthorized {
	return &PartnerAddWhiteListDomainUnauthorized{}
}

/*
PartnerAddWhiteListDomainUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PartnerAddWhiteListDomainUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this partner add white list domain unauthorized response has a 2xx status code
func (o *PartnerAddWhiteListDomainUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner add white list domain unauthorized response has a 3xx status code
func (o *PartnerAddWhiteListDomainUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner add white list domain unauthorized response has a 4xx status code
func (o *PartnerAddWhiteListDomainUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner add white list domain unauthorized response has a 5xx status code
func (o *PartnerAddWhiteListDomainUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this partner add white list domain unauthorized response a status code equal to that given
func (o *PartnerAddWhiteListDomainUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the partner add white list domain unauthorized response
func (o *PartnerAddWhiteListDomainUnauthorized) Code() int {
	return 401
}

func (o *PartnerAddWhiteListDomainUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerAddWhiteListDomainUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerAddWhiteListDomainUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerAddWhiteListDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerAddWhiteListDomainForbidden creates a PartnerAddWhiteListDomainForbidden with default headers values
func NewPartnerAddWhiteListDomainForbidden() *PartnerAddWhiteListDomainForbidden {
	return &PartnerAddWhiteListDomainForbidden{}
}

/*
PartnerAddWhiteListDomainForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PartnerAddWhiteListDomainForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this partner add white list domain forbidden response has a 2xx status code
func (o *PartnerAddWhiteListDomainForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner add white list domain forbidden response has a 3xx status code
func (o *PartnerAddWhiteListDomainForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner add white list domain forbidden response has a 4xx status code
func (o *PartnerAddWhiteListDomainForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner add white list domain forbidden response has a 5xx status code
func (o *PartnerAddWhiteListDomainForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this partner add white list domain forbidden response a status code equal to that given
func (o *PartnerAddWhiteListDomainForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the partner add white list domain forbidden response
func (o *PartnerAddWhiteListDomainForbidden) Code() int {
	return 403
}

func (o *PartnerAddWhiteListDomainForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainForbidden  %+v", 403, o.Payload)
}

func (o *PartnerAddWhiteListDomainForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainForbidden  %+v", 403, o.Payload)
}

func (o *PartnerAddWhiteListDomainForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerAddWhiteListDomainForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerAddWhiteListDomainNotFound creates a PartnerAddWhiteListDomainNotFound with default headers values
func NewPartnerAddWhiteListDomainNotFound() *PartnerAddWhiteListDomainNotFound {
	return &PartnerAddWhiteListDomainNotFound{}
}

/*
PartnerAddWhiteListDomainNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PartnerAddWhiteListDomainNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this partner add white list domain not found response has a 2xx status code
func (o *PartnerAddWhiteListDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner add white list domain not found response has a 3xx status code
func (o *PartnerAddWhiteListDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner add white list domain not found response has a 4xx status code
func (o *PartnerAddWhiteListDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner add white list domain not found response has a 5xx status code
func (o *PartnerAddWhiteListDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this partner add white list domain not found response a status code equal to that given
func (o *PartnerAddWhiteListDomainNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the partner add white list domain not found response
func (o *PartnerAddWhiteListDomainNotFound) Code() int {
	return 404
}

func (o *PartnerAddWhiteListDomainNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainNotFound  %+v", 404, o.Payload)
}

func (o *PartnerAddWhiteListDomainNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainNotFound  %+v", 404, o.Payload)
}

func (o *PartnerAddWhiteListDomainNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerAddWhiteListDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerAddWhiteListDomainInternalServerError creates a PartnerAddWhiteListDomainInternalServerError with default headers values
func NewPartnerAddWhiteListDomainInternalServerError() *PartnerAddWhiteListDomainInternalServerError {
	return &PartnerAddWhiteListDomainInternalServerError{}
}

/*
PartnerAddWhiteListDomainInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PartnerAddWhiteListDomainInternalServerError struct {
}

// IsSuccess returns true when this partner add white list domain internal server error response has a 2xx status code
func (o *PartnerAddWhiteListDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner add white list domain internal server error response has a 3xx status code
func (o *PartnerAddWhiteListDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner add white list domain internal server error response has a 4xx status code
func (o *PartnerAddWhiteListDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner add white list domain internal server error response has a 5xx status code
func (o *PartnerAddWhiteListDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this partner add white list domain internal server error response a status code equal to that given
func (o *PartnerAddWhiteListDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the partner add white list domain internal server error response
func (o *PartnerAddWhiteListDomainInternalServerError) Code() int {
	return 500
}

func (o *PartnerAddWhiteListDomainInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainInternalServerError ", 500)
}

func (o *PartnerAddWhiteListDomainInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/add/whitelist/domain][%d] partnerAddWhiteListDomainInternalServerError ", 500)
}

func (o *PartnerAddWhiteListDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
