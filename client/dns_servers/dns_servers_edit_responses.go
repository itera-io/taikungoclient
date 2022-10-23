// Code generated by go-swagger; DO NOT EDIT.

package dns_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
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
	Payload interface{}
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

func (o *DNSServersEditOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditOK  %+v", 200, o.Payload)
}

func (o *DNSServersEditOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditOK  %+v", 200, o.Payload)
}

func (o *DNSServersEditOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DNSServersEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

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
	Payload []*DNSServersEditBadRequestBodyItems0
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

func (o *DNSServersEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersEditBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditBadRequest  %+v", 400, o.Payload)
}

func (o *DNSServersEditBadRequest) GetPayload() []*DNSServersEditBadRequestBodyItems0 {
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
	Payload *DNSServersEditUnauthorizedBody
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

func (o *DNSServersEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersEditUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditUnauthorized  %+v", 401, o.Payload)
}

func (o *DNSServersEditUnauthorized) GetPayload() *DNSServersEditUnauthorizedBody {
	return o.Payload
}

func (o *DNSServersEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DNSServersEditUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
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
	Payload *DNSServersEditForbiddenBody
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

func (o *DNSServersEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersEditForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditForbidden  %+v", 403, o.Payload)
}

func (o *DNSServersEditForbidden) GetPayload() *DNSServersEditForbiddenBody {
	return o.Payload
}

func (o *DNSServersEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DNSServersEditForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
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
	Payload *DNSServersEditNotFoundBody
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

func (o *DNSServersEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersEditNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditNotFound  %+v", 404, o.Payload)
}

func (o *DNSServersEditNotFound) GetPayload() *DNSServersEditNotFoundBody {
	return o.Payload
}

func (o *DNSServersEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DNSServersEditNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
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

func (o *DNSServersEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditInternalServerError ", 500)
}

func (o *DNSServersEditInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/DnsServers/edit/{id}][%d] dnsServersEditInternalServerError ", 500)
}

func (o *DNSServersEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
DNSServersEditBadRequestBodyItems0 DNS servers edit bad request body items0
swagger:model DNSServersEditBadRequestBodyItems0
*/
type DNSServersEditBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this DNS servers edit bad request body items0
func (o *DNSServersEditBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers edit bad request body items0 based on context it is used
func (o *DNSServersEditBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersEditBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersEditBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res DNSServersEditBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersEditBody DNS servers edit body
swagger:model DNSServersEditBody
*/
type DNSServersEditBody struct {

	// address
	Address string `json:"address,omitempty"`
}

// Validate validates this DNS servers edit body
func (o *DNSServersEditBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers edit body based on context it is used
func (o *DNSServersEditBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersEditBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersEditBody) UnmarshalBinary(b []byte) error {
	var res DNSServersEditBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersEditForbiddenBody DNS servers edit forbidden body
swagger:model DNSServersEditForbiddenBody
*/
type DNSServersEditForbiddenBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this DNS servers edit forbidden body
func (o *DNSServersEditForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers edit forbidden body based on context it is used
func (o *DNSServersEditForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersEditForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersEditForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DNSServersEditForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersEditNotFoundBody DNS servers edit not found body
swagger:model DNSServersEditNotFoundBody
*/
type DNSServersEditNotFoundBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this DNS servers edit not found body
func (o *DNSServersEditNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers edit not found body based on context it is used
func (o *DNSServersEditNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersEditNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersEditNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DNSServersEditNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersEditUnauthorizedBody DNS servers edit unauthorized body
swagger:model DNSServersEditUnauthorizedBody
*/
type DNSServersEditUnauthorizedBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this DNS servers edit unauthorized body
func (o *DNSServersEditUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers edit unauthorized body based on context it is used
func (o *DNSServersEditUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersEditUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersEditUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DNSServersEditUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
