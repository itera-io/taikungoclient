// Code generated by go-swagger; DO NOT EDIT.

package ntp_servers

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

// NtpServersListReader is a Reader for the NtpServersList structure.
type NtpServersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NtpServersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNtpServersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNtpServersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNtpServersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNtpServersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNtpServersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNtpServersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNtpServersListOK creates a NtpServersListOK with default headers values
func NewNtpServersListOK() *NtpServersListOK {
	return &NtpServersListOK{}
}

/*
NtpServersListOK describes a response with status code 200, with default header values.

Success
*/
type NtpServersListOK struct {
	Payload []*NtpServersListOKBodyItems0
}

// IsSuccess returns true when this ntp servers list o k response has a 2xx status code
func (o *NtpServersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ntp servers list o k response has a 3xx status code
func (o *NtpServersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers list o k response has a 4xx status code
func (o *NtpServersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ntp servers list o k response has a 5xx status code
func (o *NtpServersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers list o k response a status code equal to that given
func (o *NtpServersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *NtpServersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListOK  %+v", 200, o.Payload)
}

func (o *NtpServersListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListOK  %+v", 200, o.Payload)
}

func (o *NtpServersListOK) GetPayload() []*NtpServersListOKBodyItems0 {
	return o.Payload
}

func (o *NtpServersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersListBadRequest creates a NtpServersListBadRequest with default headers values
func NewNtpServersListBadRequest() *NtpServersListBadRequest {
	return &NtpServersListBadRequest{}
}

/*
NtpServersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NtpServersListBadRequest struct {
	Payload []*NtpServersListBadRequestBodyItems0
}

// IsSuccess returns true when this ntp servers list bad request response has a 2xx status code
func (o *NtpServersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers list bad request response has a 3xx status code
func (o *NtpServersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers list bad request response has a 4xx status code
func (o *NtpServersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers list bad request response has a 5xx status code
func (o *NtpServersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers list bad request response a status code equal to that given
func (o *NtpServersListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *NtpServersListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListBadRequest  %+v", 400, o.Payload)
}

func (o *NtpServersListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListBadRequest  %+v", 400, o.Payload)
}

func (o *NtpServersListBadRequest) GetPayload() []*NtpServersListBadRequestBodyItems0 {
	return o.Payload
}

func (o *NtpServersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersListUnauthorized creates a NtpServersListUnauthorized with default headers values
func NewNtpServersListUnauthorized() *NtpServersListUnauthorized {
	return &NtpServersListUnauthorized{}
}

/*
NtpServersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NtpServersListUnauthorized struct {
	Payload *NtpServersListUnauthorizedBody
}

// IsSuccess returns true when this ntp servers list unauthorized response has a 2xx status code
func (o *NtpServersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers list unauthorized response has a 3xx status code
func (o *NtpServersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers list unauthorized response has a 4xx status code
func (o *NtpServersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers list unauthorized response has a 5xx status code
func (o *NtpServersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers list unauthorized response a status code equal to that given
func (o *NtpServersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *NtpServersListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListUnauthorized  %+v", 401, o.Payload)
}

func (o *NtpServersListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListUnauthorized  %+v", 401, o.Payload)
}

func (o *NtpServersListUnauthorized) GetPayload() *NtpServersListUnauthorizedBody {
	return o.Payload
}

func (o *NtpServersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NtpServersListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersListForbidden creates a NtpServersListForbidden with default headers values
func NewNtpServersListForbidden() *NtpServersListForbidden {
	return &NtpServersListForbidden{}
}

/*
NtpServersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NtpServersListForbidden struct {
	Payload *NtpServersListForbiddenBody
}

// IsSuccess returns true when this ntp servers list forbidden response has a 2xx status code
func (o *NtpServersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers list forbidden response has a 3xx status code
func (o *NtpServersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers list forbidden response has a 4xx status code
func (o *NtpServersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers list forbidden response has a 5xx status code
func (o *NtpServersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers list forbidden response a status code equal to that given
func (o *NtpServersListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NtpServersListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListForbidden  %+v", 403, o.Payload)
}

func (o *NtpServersListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListForbidden  %+v", 403, o.Payload)
}

func (o *NtpServersListForbidden) GetPayload() *NtpServersListForbiddenBody {
	return o.Payload
}

func (o *NtpServersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NtpServersListForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersListNotFound creates a NtpServersListNotFound with default headers values
func NewNtpServersListNotFound() *NtpServersListNotFound {
	return &NtpServersListNotFound{}
}

/*
NtpServersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NtpServersListNotFound struct {
	Payload *NtpServersListNotFoundBody
}

// IsSuccess returns true when this ntp servers list not found response has a 2xx status code
func (o *NtpServersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers list not found response has a 3xx status code
func (o *NtpServersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers list not found response has a 4xx status code
func (o *NtpServersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ntp servers list not found response has a 5xx status code
func (o *NtpServersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ntp servers list not found response a status code equal to that given
func (o *NtpServersListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NtpServersListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListNotFound  %+v", 404, o.Payload)
}

func (o *NtpServersListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListNotFound  %+v", 404, o.Payload)
}

func (o *NtpServersListNotFound) GetPayload() *NtpServersListNotFoundBody {
	return o.Payload
}

func (o *NtpServersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NtpServersListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNtpServersListInternalServerError creates a NtpServersListInternalServerError with default headers values
func NewNtpServersListInternalServerError() *NtpServersListInternalServerError {
	return &NtpServersListInternalServerError{}
}

/*
NtpServersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NtpServersListInternalServerError struct {
}

// IsSuccess returns true when this ntp servers list internal server error response has a 2xx status code
func (o *NtpServersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ntp servers list internal server error response has a 3xx status code
func (o *NtpServersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ntp servers list internal server error response has a 4xx status code
func (o *NtpServersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ntp servers list internal server error response has a 5xx status code
func (o *NtpServersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ntp servers list internal server error response a status code equal to that given
func (o *NtpServersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NtpServersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListInternalServerError ", 500)
}

func (o *NtpServersListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/NtpServers/list/{accessProfileId}][%d] ntpServersListInternalServerError ", 500)
}

func (o *NtpServersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
NtpServersListBadRequestBodyItems0 ntp servers list bad request body items0
swagger:model NtpServersListBadRequestBodyItems0
*/
type NtpServersListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this ntp servers list bad request body items0
func (o *NtpServersListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ntp servers list bad request body items0 based on context it is used
func (o *NtpServersListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NtpServersListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NtpServersListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res NtpServersListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NtpServersListForbiddenBody ntp servers list forbidden body
swagger:model NtpServersListForbiddenBody
*/
type NtpServersListForbiddenBody struct {

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

// Validate validates this ntp servers list forbidden body
func (o *NtpServersListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ntp servers list forbidden body based on context it is used
func (o *NtpServersListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NtpServersListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NtpServersListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res NtpServersListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NtpServersListNotFoundBody ntp servers list not found body
swagger:model NtpServersListNotFoundBody
*/
type NtpServersListNotFoundBody struct {

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

// Validate validates this ntp servers list not found body
func (o *NtpServersListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ntp servers list not found body based on context it is used
func (o *NtpServersListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NtpServersListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NtpServersListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NtpServersListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NtpServersListOKBodyItems0 ntp servers list o k body items0
swagger:model NtpServersListOKBodyItems0
*/
type NtpServersListOKBodyItems0 struct {

	// access profile name
	AccessProfileName string `json:"accessProfileName,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this ntp servers list o k body items0
func (o *NtpServersListOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ntp servers list o k body items0 based on context it is used
func (o *NtpServersListOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NtpServersListOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NtpServersListOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res NtpServersListOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NtpServersListUnauthorizedBody ntp servers list unauthorized body
swagger:model NtpServersListUnauthorizedBody
*/
type NtpServersListUnauthorizedBody struct {

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

// Validate validates this ntp servers list unauthorized body
func (o *NtpServersListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ntp servers list unauthorized body based on context it is used
func (o *NtpServersListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NtpServersListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NtpServersListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res NtpServersListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
