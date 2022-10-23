// Code generated by go-swagger; DO NOT EDIT.

package ssh_users

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

// SSHUsersCreateReader is a Reader for the SSHUsersCreate structure.
type SSHUsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSHUsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSHUsersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSSHUsersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSSHUsersCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSSHUsersCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSSHUsersCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSSHUsersCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSSHUsersCreateOK creates a SSHUsersCreateOK with default headers values
func NewSSHUsersCreateOK() *SSHUsersCreateOK {
	return &SSHUsersCreateOK{}
}

/*
SSHUsersCreateOK describes a response with status code 200, with default header values.

Success
*/
type SSHUsersCreateOK struct {
	Payload *SSHUsersCreateOKBody
}

// IsSuccess returns true when this ssh users create o k response has a 2xx status code
func (o *SSHUsersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ssh users create o k response has a 3xx status code
func (o *SSHUsersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create o k response has a 4xx status code
func (o *SSHUsersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ssh users create o k response has a 5xx status code
func (o *SSHUsersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create o k response a status code equal to that given
func (o *SSHUsersCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SSHUsersCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateOK  %+v", 200, o.Payload)
}

func (o *SSHUsersCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateOK  %+v", 200, o.Payload)
}

func (o *SSHUsersCreateOK) GetPayload() *SSHUsersCreateOKBody {
	return o.Payload
}

func (o *SSHUsersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SSHUsersCreateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateBadRequest creates a SSHUsersCreateBadRequest with default headers values
func NewSSHUsersCreateBadRequest() *SSHUsersCreateBadRequest {
	return &SSHUsersCreateBadRequest{}
}

/*
SSHUsersCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SSHUsersCreateBadRequest struct {
	Payload []*SSHUsersCreateBadRequestBodyItems0
}

// IsSuccess returns true when this ssh users create bad request response has a 2xx status code
func (o *SSHUsersCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create bad request response has a 3xx status code
func (o *SSHUsersCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create bad request response has a 4xx status code
func (o *SSHUsersCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create bad request response has a 5xx status code
func (o *SSHUsersCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create bad request response a status code equal to that given
func (o *SSHUsersCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SSHUsersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SSHUsersCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SSHUsersCreateBadRequest) GetPayload() []*SSHUsersCreateBadRequestBodyItems0 {
	return o.Payload
}

func (o *SSHUsersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateUnauthorized creates a SSHUsersCreateUnauthorized with default headers values
func NewSSHUsersCreateUnauthorized() *SSHUsersCreateUnauthorized {
	return &SSHUsersCreateUnauthorized{}
}

/*
SSHUsersCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SSHUsersCreateUnauthorized struct {
	Payload *SSHUsersCreateUnauthorizedBody
}

// IsSuccess returns true when this ssh users create unauthorized response has a 2xx status code
func (o *SSHUsersCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create unauthorized response has a 3xx status code
func (o *SSHUsersCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create unauthorized response has a 4xx status code
func (o *SSHUsersCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create unauthorized response has a 5xx status code
func (o *SSHUsersCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create unauthorized response a status code equal to that given
func (o *SSHUsersCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SSHUsersCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SSHUsersCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SSHUsersCreateUnauthorized) GetPayload() *SSHUsersCreateUnauthorizedBody {
	return o.Payload
}

func (o *SSHUsersCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SSHUsersCreateUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateForbidden creates a SSHUsersCreateForbidden with default headers values
func NewSSHUsersCreateForbidden() *SSHUsersCreateForbidden {
	return &SSHUsersCreateForbidden{}
}

/*
SSHUsersCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SSHUsersCreateForbidden struct {
	Payload *SSHUsersCreateForbiddenBody
}

// IsSuccess returns true when this ssh users create forbidden response has a 2xx status code
func (o *SSHUsersCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create forbidden response has a 3xx status code
func (o *SSHUsersCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create forbidden response has a 4xx status code
func (o *SSHUsersCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create forbidden response has a 5xx status code
func (o *SSHUsersCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create forbidden response a status code equal to that given
func (o *SSHUsersCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SSHUsersCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateForbidden  %+v", 403, o.Payload)
}

func (o *SSHUsersCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateForbidden  %+v", 403, o.Payload)
}

func (o *SSHUsersCreateForbidden) GetPayload() *SSHUsersCreateForbiddenBody {
	return o.Payload
}

func (o *SSHUsersCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SSHUsersCreateForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateNotFound creates a SSHUsersCreateNotFound with default headers values
func NewSSHUsersCreateNotFound() *SSHUsersCreateNotFound {
	return &SSHUsersCreateNotFound{}
}

/*
SSHUsersCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SSHUsersCreateNotFound struct {
	Payload *SSHUsersCreateNotFoundBody
}

// IsSuccess returns true when this ssh users create not found response has a 2xx status code
func (o *SSHUsersCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create not found response has a 3xx status code
func (o *SSHUsersCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create not found response has a 4xx status code
func (o *SSHUsersCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create not found response has a 5xx status code
func (o *SSHUsersCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create not found response a status code equal to that given
func (o *SSHUsersCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SSHUsersCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateNotFound  %+v", 404, o.Payload)
}

func (o *SSHUsersCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateNotFound  %+v", 404, o.Payload)
}

func (o *SSHUsersCreateNotFound) GetPayload() *SSHUsersCreateNotFoundBody {
	return o.Payload
}

func (o *SSHUsersCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SSHUsersCreateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateInternalServerError creates a SSHUsersCreateInternalServerError with default headers values
func NewSSHUsersCreateInternalServerError() *SSHUsersCreateInternalServerError {
	return &SSHUsersCreateInternalServerError{}
}

/*
SSHUsersCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SSHUsersCreateInternalServerError struct {
}

// IsSuccess returns true when this ssh users create internal server error response has a 2xx status code
func (o *SSHUsersCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create internal server error response has a 3xx status code
func (o *SSHUsersCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create internal server error response has a 4xx status code
func (o *SSHUsersCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ssh users create internal server error response has a 5xx status code
func (o *SSHUsersCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ssh users create internal server error response a status code equal to that given
func (o *SSHUsersCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SSHUsersCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateInternalServerError ", 500)
}

func (o *SSHUsersCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateInternalServerError ", 500)
}

func (o *SSHUsersCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
SSHUsersCreateBadRequestBodyItems0 SSH users create bad request body items0
swagger:model SSHUsersCreateBadRequestBodyItems0
*/
type SSHUsersCreateBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this SSH users create bad request body items0
func (o *SSHUsersCreateBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH users create bad request body items0 based on context it is used
func (o *SSHUsersCreateBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SSHUsersCreateBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSHUsersCreateBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res SSHUsersCreateBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SSHUsersCreateBody SSH users create body
swagger:model SSHUsersCreateBody
*/
type SSHUsersCreateBody struct {

	// access profile Id
	AccessProfileID int32 `json:"accessProfileId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ssh public key
	SSHPublicKey string `json:"sshPublicKey,omitempty"`
}

// Validate validates this SSH users create body
func (o *SSHUsersCreateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH users create body based on context it is used
func (o *SSHUsersCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SSHUsersCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSHUsersCreateBody) UnmarshalBinary(b []byte) error {
	var res SSHUsersCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SSHUsersCreateForbiddenBody SSH users create forbidden body
swagger:model SSHUsersCreateForbiddenBody
*/
type SSHUsersCreateForbiddenBody struct {

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

// Validate validates this SSH users create forbidden body
func (o *SSHUsersCreateForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH users create forbidden body based on context it is used
func (o *SSHUsersCreateForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SSHUsersCreateForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSHUsersCreateForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SSHUsersCreateForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SSHUsersCreateNotFoundBody SSH users create not found body
swagger:model SSHUsersCreateNotFoundBody
*/
type SSHUsersCreateNotFoundBody struct {

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

// Validate validates this SSH users create not found body
func (o *SSHUsersCreateNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH users create not found body based on context it is used
func (o *SSHUsersCreateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SSHUsersCreateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSHUsersCreateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SSHUsersCreateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SSHUsersCreateOKBody SSH users create o k body
swagger:model SSHUsersCreateOKBody
*/
type SSHUsersCreateOKBody struct {

	// id
	ID string `json:"id,omitempty"`

	// is error
	IsError bool `json:"isError"`

	// message
	Message string `json:"message,omitempty"`

	// result
	Result interface{} `json:"result,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this SSH users create o k body
func (o *SSHUsersCreateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH users create o k body based on context it is used
func (o *SSHUsersCreateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SSHUsersCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSHUsersCreateOKBody) UnmarshalBinary(b []byte) error {
	var res SSHUsersCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SSHUsersCreateUnauthorizedBody SSH users create unauthorized body
swagger:model SSHUsersCreateUnauthorizedBody
*/
type SSHUsersCreateUnauthorizedBody struct {

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

// Validate validates this SSH users create unauthorized body
func (o *SSHUsersCreateUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH users create unauthorized body based on context it is used
func (o *SSHUsersCreateUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SSHUsersCreateUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SSHUsersCreateUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res SSHUsersCreateUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
