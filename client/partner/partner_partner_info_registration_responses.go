// Code generated by go-swagger; DO NOT EDIT.

package partner

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

// PartnerPartnerInfoRegistrationReader is a Reader for the PartnerPartnerInfoRegistration structure.
type PartnerPartnerInfoRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerPartnerInfoRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerPartnerInfoRegistrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartnerPartnerInfoRegistrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartnerPartnerInfoRegistrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPartnerPartnerInfoRegistrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartnerPartnerInfoRegistrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPartnerPartnerInfoRegistrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartnerPartnerInfoRegistrationOK creates a PartnerPartnerInfoRegistrationOK with default headers values
func NewPartnerPartnerInfoRegistrationOK() *PartnerPartnerInfoRegistrationOK {
	return &PartnerPartnerInfoRegistrationOK{}
}

/*
PartnerPartnerInfoRegistrationOK describes a response with status code 200, with default header values.

Success
*/
type PartnerPartnerInfoRegistrationOK struct {
	Payload *PartnerPartnerInfoRegistrationOKBody
}

// IsSuccess returns true when this partner partner info registration o k response has a 2xx status code
func (o *PartnerPartnerInfoRegistrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner partner info registration o k response has a 3xx status code
func (o *PartnerPartnerInfoRegistrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner partner info registration o k response has a 4xx status code
func (o *PartnerPartnerInfoRegistrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner partner info registration o k response has a 5xx status code
func (o *PartnerPartnerInfoRegistrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner partner info registration o k response a status code equal to that given
func (o *PartnerPartnerInfoRegistrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PartnerPartnerInfoRegistrationOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationOK  %+v", 200, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationOK  %+v", 200, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationOK) GetPayload() *PartnerPartnerInfoRegistrationOKBody {
	return o.Payload
}

func (o *PartnerPartnerInfoRegistrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PartnerPartnerInfoRegistrationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerPartnerInfoRegistrationBadRequest creates a PartnerPartnerInfoRegistrationBadRequest with default headers values
func NewPartnerPartnerInfoRegistrationBadRequest() *PartnerPartnerInfoRegistrationBadRequest {
	return &PartnerPartnerInfoRegistrationBadRequest{}
}

/*
PartnerPartnerInfoRegistrationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PartnerPartnerInfoRegistrationBadRequest struct {
	Payload []*PartnerPartnerInfoRegistrationBadRequestBodyItems0
}

// IsSuccess returns true when this partner partner info registration bad request response has a 2xx status code
func (o *PartnerPartnerInfoRegistrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner partner info registration bad request response has a 3xx status code
func (o *PartnerPartnerInfoRegistrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner partner info registration bad request response has a 4xx status code
func (o *PartnerPartnerInfoRegistrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner partner info registration bad request response has a 5xx status code
func (o *PartnerPartnerInfoRegistrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this partner partner info registration bad request response a status code equal to that given
func (o *PartnerPartnerInfoRegistrationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PartnerPartnerInfoRegistrationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationBadRequest) GetPayload() []*PartnerPartnerInfoRegistrationBadRequestBodyItems0 {
	return o.Payload
}

func (o *PartnerPartnerInfoRegistrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerPartnerInfoRegistrationUnauthorized creates a PartnerPartnerInfoRegistrationUnauthorized with default headers values
func NewPartnerPartnerInfoRegistrationUnauthorized() *PartnerPartnerInfoRegistrationUnauthorized {
	return &PartnerPartnerInfoRegistrationUnauthorized{}
}

/*
PartnerPartnerInfoRegistrationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PartnerPartnerInfoRegistrationUnauthorized struct {
	Payload *PartnerPartnerInfoRegistrationUnauthorizedBody
}

// IsSuccess returns true when this partner partner info registration unauthorized response has a 2xx status code
func (o *PartnerPartnerInfoRegistrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner partner info registration unauthorized response has a 3xx status code
func (o *PartnerPartnerInfoRegistrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner partner info registration unauthorized response has a 4xx status code
func (o *PartnerPartnerInfoRegistrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner partner info registration unauthorized response has a 5xx status code
func (o *PartnerPartnerInfoRegistrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this partner partner info registration unauthorized response a status code equal to that given
func (o *PartnerPartnerInfoRegistrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PartnerPartnerInfoRegistrationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationUnauthorized) GetPayload() *PartnerPartnerInfoRegistrationUnauthorizedBody {
	return o.Payload
}

func (o *PartnerPartnerInfoRegistrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PartnerPartnerInfoRegistrationUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerPartnerInfoRegistrationForbidden creates a PartnerPartnerInfoRegistrationForbidden with default headers values
func NewPartnerPartnerInfoRegistrationForbidden() *PartnerPartnerInfoRegistrationForbidden {
	return &PartnerPartnerInfoRegistrationForbidden{}
}

/*
PartnerPartnerInfoRegistrationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PartnerPartnerInfoRegistrationForbidden struct {
	Payload *PartnerPartnerInfoRegistrationForbiddenBody
}

// IsSuccess returns true when this partner partner info registration forbidden response has a 2xx status code
func (o *PartnerPartnerInfoRegistrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner partner info registration forbidden response has a 3xx status code
func (o *PartnerPartnerInfoRegistrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner partner info registration forbidden response has a 4xx status code
func (o *PartnerPartnerInfoRegistrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner partner info registration forbidden response has a 5xx status code
func (o *PartnerPartnerInfoRegistrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this partner partner info registration forbidden response a status code equal to that given
func (o *PartnerPartnerInfoRegistrationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PartnerPartnerInfoRegistrationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationForbidden  %+v", 403, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationForbidden  %+v", 403, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationForbidden) GetPayload() *PartnerPartnerInfoRegistrationForbiddenBody {
	return o.Payload
}

func (o *PartnerPartnerInfoRegistrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PartnerPartnerInfoRegistrationForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerPartnerInfoRegistrationNotFound creates a PartnerPartnerInfoRegistrationNotFound with default headers values
func NewPartnerPartnerInfoRegistrationNotFound() *PartnerPartnerInfoRegistrationNotFound {
	return &PartnerPartnerInfoRegistrationNotFound{}
}

/*
PartnerPartnerInfoRegistrationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PartnerPartnerInfoRegistrationNotFound struct {
	Payload *PartnerPartnerInfoRegistrationNotFoundBody
}

// IsSuccess returns true when this partner partner info registration not found response has a 2xx status code
func (o *PartnerPartnerInfoRegistrationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner partner info registration not found response has a 3xx status code
func (o *PartnerPartnerInfoRegistrationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner partner info registration not found response has a 4xx status code
func (o *PartnerPartnerInfoRegistrationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner partner info registration not found response has a 5xx status code
func (o *PartnerPartnerInfoRegistrationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this partner partner info registration not found response a status code equal to that given
func (o *PartnerPartnerInfoRegistrationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PartnerPartnerInfoRegistrationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationNotFound  %+v", 404, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationNotFound  %+v", 404, o.Payload)
}

func (o *PartnerPartnerInfoRegistrationNotFound) GetPayload() *PartnerPartnerInfoRegistrationNotFoundBody {
	return o.Payload
}

func (o *PartnerPartnerInfoRegistrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PartnerPartnerInfoRegistrationNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerPartnerInfoRegistrationInternalServerError creates a PartnerPartnerInfoRegistrationInternalServerError with default headers values
func NewPartnerPartnerInfoRegistrationInternalServerError() *PartnerPartnerInfoRegistrationInternalServerError {
	return &PartnerPartnerInfoRegistrationInternalServerError{}
}

/*
PartnerPartnerInfoRegistrationInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PartnerPartnerInfoRegistrationInternalServerError struct {
}

// IsSuccess returns true when this partner partner info registration internal server error response has a 2xx status code
func (o *PartnerPartnerInfoRegistrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner partner info registration internal server error response has a 3xx status code
func (o *PartnerPartnerInfoRegistrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner partner info registration internal server error response has a 4xx status code
func (o *PartnerPartnerInfoRegistrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner partner info registration internal server error response has a 5xx status code
func (o *PartnerPartnerInfoRegistrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this partner partner info registration internal server error response a status code equal to that given
func (o *PartnerPartnerInfoRegistrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PartnerPartnerInfoRegistrationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationInternalServerError ", 500)
}

func (o *PartnerPartnerInfoRegistrationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Partner/info][%d] partnerPartnerInfoRegistrationInternalServerError ", 500)
}

func (o *PartnerPartnerInfoRegistrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
PartnerPartnerInfoRegistrationBadRequestBodyItems0 partner partner info registration bad request body items0
swagger:model PartnerPartnerInfoRegistrationBadRequestBodyItems0
*/
type PartnerPartnerInfoRegistrationBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this partner partner info registration bad request body items0
func (o *PartnerPartnerInfoRegistrationBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partner partner info registration bad request body items0 based on context it is used
func (o *PartnerPartnerInfoRegistrationBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res PartnerPartnerInfoRegistrationBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PartnerPartnerInfoRegistrationForbiddenBody partner partner info registration forbidden body
swagger:model PartnerPartnerInfoRegistrationForbiddenBody
*/
type PartnerPartnerInfoRegistrationForbiddenBody struct {

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

// Validate validates this partner partner info registration forbidden body
func (o *PartnerPartnerInfoRegistrationForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partner partner info registration forbidden body based on context it is used
func (o *PartnerPartnerInfoRegistrationForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PartnerPartnerInfoRegistrationForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PartnerPartnerInfoRegistrationNotFoundBody partner partner info registration not found body
swagger:model PartnerPartnerInfoRegistrationNotFoundBody
*/
type PartnerPartnerInfoRegistrationNotFoundBody struct {

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

// Validate validates this partner partner info registration not found body
func (o *PartnerPartnerInfoRegistrationNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partner partner info registration not found body based on context it is used
func (o *PartnerPartnerInfoRegistrationNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PartnerPartnerInfoRegistrationNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PartnerPartnerInfoRegistrationOKBody partner partner info registration o k body
swagger:model PartnerPartnerInfoRegistrationOKBody
*/
type PartnerPartnerInfoRegistrationOKBody struct {

	// allow registration
	AllowRegistration bool `json:"allowRegistration"`

	// id
	ID int32 `json:"id,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// payment enabled
	PaymentEnabled bool `json:"paymentEnabled"`
}

// Validate validates this partner partner info registration o k body
func (o *PartnerPartnerInfoRegistrationOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partner partner info registration o k body based on context it is used
func (o *PartnerPartnerInfoRegistrationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationOKBody) UnmarshalBinary(b []byte) error {
	var res PartnerPartnerInfoRegistrationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PartnerPartnerInfoRegistrationUnauthorizedBody partner partner info registration unauthorized body
swagger:model PartnerPartnerInfoRegistrationUnauthorizedBody
*/
type PartnerPartnerInfoRegistrationUnauthorizedBody struct {

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

// Validate validates this partner partner info registration unauthorized body
func (o *PartnerPartnerInfoRegistrationUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partner partner info registration unauthorized body based on context it is used
func (o *PartnerPartnerInfoRegistrationUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PartnerPartnerInfoRegistrationUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PartnerPartnerInfoRegistrationUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
