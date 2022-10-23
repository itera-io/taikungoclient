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
	Payload []*DNSServersListOKBodyItems0
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

func (o *DNSServersListOK) GetPayload() []*DNSServersListOKBodyItems0 {
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
	Payload []*DNSServersListBadRequestBodyItems0
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

func (o *DNSServersListBadRequest) GetPayload() []*DNSServersListBadRequestBodyItems0 {
	return o.Payload
}

func (o *DNSServersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
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
	Payload *DNSServersListUnauthorizedBody
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

func (o *DNSServersListUnauthorized) GetPayload() *DNSServersListUnauthorizedBody {
	return o.Payload
}

func (o *DNSServersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DNSServersListUnauthorizedBody)

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
	Payload *DNSServersListForbiddenBody
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

func (o *DNSServersListForbidden) GetPayload() *DNSServersListForbiddenBody {
	return o.Payload
}

func (o *DNSServersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DNSServersListForbiddenBody)

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
	Payload *DNSServersListNotFoundBody
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

func (o *DNSServersListNotFound) GetPayload() *DNSServersListNotFoundBody {
	return o.Payload
}

func (o *DNSServersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DNSServersListNotFoundBody)

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

/*
DNSServersListBadRequestBodyItems0 DNS servers list bad request body items0
swagger:model DNSServersListBadRequestBodyItems0
*/
type DNSServersListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this DNS servers list bad request body items0
func (o *DNSServersListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers list bad request body items0 based on context it is used
func (o *DNSServersListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res DNSServersListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersListForbiddenBody DNS servers list forbidden body
swagger:model DNSServersListForbiddenBody
*/
type DNSServersListForbiddenBody struct {

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

// Validate validates this DNS servers list forbidden body
func (o *DNSServersListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers list forbidden body based on context it is used
func (o *DNSServersListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DNSServersListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersListNotFoundBody DNS servers list not found body
swagger:model DNSServersListNotFoundBody
*/
type DNSServersListNotFoundBody struct {

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

// Validate validates this DNS servers list not found body
func (o *DNSServersListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers list not found body based on context it is used
func (o *DNSServersListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DNSServersListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersListOKBodyItems0 DNS servers list o k body items0
swagger:model DNSServersListOKBodyItems0
*/
type DNSServersListOKBodyItems0 struct {

	// access profile name
	AccessProfileName string `json:"accessProfileName,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this DNS servers list o k body items0
func (o *DNSServersListOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers list o k body items0 based on context it is used
func (o *DNSServersListOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersListOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersListOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res DNSServersListOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DNSServersListUnauthorizedBody DNS servers list unauthorized body
swagger:model DNSServersListUnauthorizedBody
*/
type DNSServersListUnauthorizedBody struct {

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

// Validate validates this DNS servers list unauthorized body
func (o *DNSServersListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this DNS servers list unauthorized body based on context it is used
func (o *DNSServersListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DNSServersListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DNSServersListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DNSServersListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
