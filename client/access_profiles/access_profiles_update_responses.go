// Code generated by go-swagger; DO NOT EDIT.

package access_profiles

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

// AccessProfilesUpdateReader is a Reader for the AccessProfilesUpdate structure.
type AccessProfilesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessProfilesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessProfilesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccessProfilesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccessProfilesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccessProfilesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessProfilesUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessProfilesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccessProfilesUpdateOK creates a AccessProfilesUpdateOK with default headers values
func NewAccessProfilesUpdateOK() *AccessProfilesUpdateOK {
	return &AccessProfilesUpdateOK{}
}

/*
AccessProfilesUpdateOK describes a response with status code 200, with default header values.

Success
*/
type AccessProfilesUpdateOK struct {
	Payload interface{}
}

// IsSuccess returns true when this access profiles update o k response has a 2xx status code
func (o *AccessProfilesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this access profiles update o k response has a 3xx status code
func (o *AccessProfilesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update o k response has a 4xx status code
func (o *AccessProfilesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles update o k response has a 5xx status code
func (o *AccessProfilesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update o k response a status code equal to that given
func (o *AccessProfilesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AccessProfilesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesUpdateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessProfilesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateBadRequest creates a AccessProfilesUpdateBadRequest with default headers values
func NewAccessProfilesUpdateBadRequest() *AccessProfilesUpdateBadRequest {
	return &AccessProfilesUpdateBadRequest{}
}

/*
AccessProfilesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AccessProfilesUpdateBadRequest struct {
	Payload []*AccessProfilesUpdateBadRequestBodyItems0
}

// IsSuccess returns true when this access profiles update bad request response has a 2xx status code
func (o *AccessProfilesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update bad request response has a 3xx status code
func (o *AccessProfilesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update bad request response has a 4xx status code
func (o *AccessProfilesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update bad request response has a 5xx status code
func (o *AccessProfilesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update bad request response a status code equal to that given
func (o *AccessProfilesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AccessProfilesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesUpdateBadRequest) GetPayload() []*AccessProfilesUpdateBadRequestBodyItems0 {
	return o.Payload
}

func (o *AccessProfilesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateUnauthorized creates a AccessProfilesUpdateUnauthorized with default headers values
func NewAccessProfilesUpdateUnauthorized() *AccessProfilesUpdateUnauthorized {
	return &AccessProfilesUpdateUnauthorized{}
}

/*
AccessProfilesUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AccessProfilesUpdateUnauthorized struct {
	Payload *AccessProfilesUpdateUnauthorizedBody
}

// IsSuccess returns true when this access profiles update unauthorized response has a 2xx status code
func (o *AccessProfilesUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update unauthorized response has a 3xx status code
func (o *AccessProfilesUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update unauthorized response has a 4xx status code
func (o *AccessProfilesUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update unauthorized response has a 5xx status code
func (o *AccessProfilesUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update unauthorized response a status code equal to that given
func (o *AccessProfilesUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AccessProfilesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesUpdateUnauthorized) GetPayload() *AccessProfilesUpdateUnauthorizedBody {
	return o.Payload
}

func (o *AccessProfilesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessProfilesUpdateUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateForbidden creates a AccessProfilesUpdateForbidden with default headers values
func NewAccessProfilesUpdateForbidden() *AccessProfilesUpdateForbidden {
	return &AccessProfilesUpdateForbidden{}
}

/*
AccessProfilesUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AccessProfilesUpdateForbidden struct {
	Payload *AccessProfilesUpdateForbiddenBody
}

// IsSuccess returns true when this access profiles update forbidden response has a 2xx status code
func (o *AccessProfilesUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update forbidden response has a 3xx status code
func (o *AccessProfilesUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update forbidden response has a 4xx status code
func (o *AccessProfilesUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update forbidden response has a 5xx status code
func (o *AccessProfilesUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update forbidden response a status code equal to that given
func (o *AccessProfilesUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AccessProfilesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesUpdateForbidden) GetPayload() *AccessProfilesUpdateForbiddenBody {
	return o.Payload
}

func (o *AccessProfilesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessProfilesUpdateForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateNotFound creates a AccessProfilesUpdateNotFound with default headers values
func NewAccessProfilesUpdateNotFound() *AccessProfilesUpdateNotFound {
	return &AccessProfilesUpdateNotFound{}
}

/*
AccessProfilesUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AccessProfilesUpdateNotFound struct {
	Payload *AccessProfilesUpdateNotFoundBody
}

// IsSuccess returns true when this access profiles update not found response has a 2xx status code
func (o *AccessProfilesUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update not found response has a 3xx status code
func (o *AccessProfilesUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update not found response has a 4xx status code
func (o *AccessProfilesUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update not found response has a 5xx status code
func (o *AccessProfilesUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update not found response a status code equal to that given
func (o *AccessProfilesUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AccessProfilesUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesUpdateNotFound) GetPayload() *AccessProfilesUpdateNotFoundBody {
	return o.Payload
}

func (o *AccessProfilesUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AccessProfilesUpdateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateInternalServerError creates a AccessProfilesUpdateInternalServerError with default headers values
func NewAccessProfilesUpdateInternalServerError() *AccessProfilesUpdateInternalServerError {
	return &AccessProfilesUpdateInternalServerError{}
}

/*
AccessProfilesUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AccessProfilesUpdateInternalServerError struct {
}

// IsSuccess returns true when this access profiles update internal server error response has a 2xx status code
func (o *AccessProfilesUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update internal server error response has a 3xx status code
func (o *AccessProfilesUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update internal server error response has a 4xx status code
func (o *AccessProfilesUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles update internal server error response has a 5xx status code
func (o *AccessProfilesUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this access profiles update internal server error response a status code equal to that given
func (o *AccessProfilesUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AccessProfilesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateInternalServerError ", 500)
}

func (o *AccessProfilesUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateInternalServerError ", 500)
}

func (o *AccessProfilesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
AccessProfilesUpdateBadRequestBodyItems0 access profiles update bad request body items0
swagger:model AccessProfilesUpdateBadRequestBodyItems0
*/
type AccessProfilesUpdateBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this access profiles update bad request body items0
func (o *AccessProfilesUpdateBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access profiles update bad request body items0 based on context it is used
func (o *AccessProfilesUpdateBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessProfilesUpdateBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessProfilesUpdateBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res AccessProfilesUpdateBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessProfilesUpdateBody access profiles update body
swagger:model AccessProfilesUpdateBody
*/
type AccessProfilesUpdateBody struct {

	// http proxy
	HTTPProxy string `json:"httpProxy,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this access profiles update body
func (o *AccessProfilesUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access profiles update body based on context it is used
func (o *AccessProfilesUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessProfilesUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessProfilesUpdateBody) UnmarshalBinary(b []byte) error {
	var res AccessProfilesUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessProfilesUpdateForbiddenBody access profiles update forbidden body
swagger:model AccessProfilesUpdateForbiddenBody
*/
type AccessProfilesUpdateForbiddenBody struct {

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

// Validate validates this access profiles update forbidden body
func (o *AccessProfilesUpdateForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access profiles update forbidden body based on context it is used
func (o *AccessProfilesUpdateForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessProfilesUpdateForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessProfilesUpdateForbiddenBody) UnmarshalBinary(b []byte) error {
	var res AccessProfilesUpdateForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessProfilesUpdateNotFoundBody access profiles update not found body
swagger:model AccessProfilesUpdateNotFoundBody
*/
type AccessProfilesUpdateNotFoundBody struct {

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

// Validate validates this access profiles update not found body
func (o *AccessProfilesUpdateNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access profiles update not found body based on context it is used
func (o *AccessProfilesUpdateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessProfilesUpdateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessProfilesUpdateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AccessProfilesUpdateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AccessProfilesUpdateUnauthorizedBody access profiles update unauthorized body
swagger:model AccessProfilesUpdateUnauthorizedBody
*/
type AccessProfilesUpdateUnauthorizedBody struct {

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

// Validate validates this access profiles update unauthorized body
func (o *AccessProfilesUpdateUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access profiles update unauthorized body based on context it is used
func (o *AccessProfilesUpdateUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessProfilesUpdateUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessProfilesUpdateUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res AccessProfilesUpdateUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
