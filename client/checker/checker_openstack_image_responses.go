// Code generated by go-swagger; DO NOT EDIT.

package checker

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

// CheckerOpenstackImageReader is a Reader for the CheckerOpenstackImage structure.
type CheckerOpenstackImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerOpenstackImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerOpenstackImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerOpenstackImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerOpenstackImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerOpenstackImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerOpenstackImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerOpenstackImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerOpenstackImageOK creates a CheckerOpenstackImageOK with default headers values
func NewCheckerOpenstackImageOK() *CheckerOpenstackImageOK {
	return &CheckerOpenstackImageOK{}
}

/*
CheckerOpenstackImageOK describes a response with status code 200, with default header values.

Success
*/
type CheckerOpenstackImageOK struct {
	Payload interface{}
}

// IsSuccess returns true when this checker openstack image o k response has a 2xx status code
func (o *CheckerOpenstackImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker openstack image o k response has a 3xx status code
func (o *CheckerOpenstackImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack image o k response has a 4xx status code
func (o *CheckerOpenstackImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker openstack image o k response has a 5xx status code
func (o *CheckerOpenstackImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack image o k response a status code equal to that given
func (o *CheckerOpenstackImageOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerOpenstackImageOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageOK  %+v", 200, o.Payload)
}

func (o *CheckerOpenstackImageOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageOK  %+v", 200, o.Payload)
}

func (o *CheckerOpenstackImageOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerOpenstackImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackImageBadRequest creates a CheckerOpenstackImageBadRequest with default headers values
func NewCheckerOpenstackImageBadRequest() *CheckerOpenstackImageBadRequest {
	return &CheckerOpenstackImageBadRequest{}
}

/*
CheckerOpenstackImageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerOpenstackImageBadRequest struct {
	Payload []*CheckerOpenstackImageBadRequestBodyItems0
}

// IsSuccess returns true when this checker openstack image bad request response has a 2xx status code
func (o *CheckerOpenstackImageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack image bad request response has a 3xx status code
func (o *CheckerOpenstackImageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack image bad request response has a 4xx status code
func (o *CheckerOpenstackImageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack image bad request response has a 5xx status code
func (o *CheckerOpenstackImageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack image bad request response a status code equal to that given
func (o *CheckerOpenstackImageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerOpenstackImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerOpenstackImageBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerOpenstackImageBadRequest) GetPayload() []*CheckerOpenstackImageBadRequestBodyItems0 {
	return o.Payload
}

func (o *CheckerOpenstackImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackImageUnauthorized creates a CheckerOpenstackImageUnauthorized with default headers values
func NewCheckerOpenstackImageUnauthorized() *CheckerOpenstackImageUnauthorized {
	return &CheckerOpenstackImageUnauthorized{}
}

/*
CheckerOpenstackImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerOpenstackImageUnauthorized struct {
	Payload *CheckerOpenstackImageUnauthorizedBody
}

// IsSuccess returns true when this checker openstack image unauthorized response has a 2xx status code
func (o *CheckerOpenstackImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack image unauthorized response has a 3xx status code
func (o *CheckerOpenstackImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack image unauthorized response has a 4xx status code
func (o *CheckerOpenstackImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack image unauthorized response has a 5xx status code
func (o *CheckerOpenstackImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack image unauthorized response a status code equal to that given
func (o *CheckerOpenstackImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerOpenstackImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerOpenstackImageUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerOpenstackImageUnauthorized) GetPayload() *CheckerOpenstackImageUnauthorizedBody {
	return o.Payload
}

func (o *CheckerOpenstackImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerOpenstackImageUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackImageForbidden creates a CheckerOpenstackImageForbidden with default headers values
func NewCheckerOpenstackImageForbidden() *CheckerOpenstackImageForbidden {
	return &CheckerOpenstackImageForbidden{}
}

/*
CheckerOpenstackImageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerOpenstackImageForbidden struct {
	Payload *CheckerOpenstackImageForbiddenBody
}

// IsSuccess returns true when this checker openstack image forbidden response has a 2xx status code
func (o *CheckerOpenstackImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack image forbidden response has a 3xx status code
func (o *CheckerOpenstackImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack image forbidden response has a 4xx status code
func (o *CheckerOpenstackImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack image forbidden response has a 5xx status code
func (o *CheckerOpenstackImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack image forbidden response a status code equal to that given
func (o *CheckerOpenstackImageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerOpenstackImageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageForbidden  %+v", 403, o.Payload)
}

func (o *CheckerOpenstackImageForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageForbidden  %+v", 403, o.Payload)
}

func (o *CheckerOpenstackImageForbidden) GetPayload() *CheckerOpenstackImageForbiddenBody {
	return o.Payload
}

func (o *CheckerOpenstackImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerOpenstackImageForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackImageNotFound creates a CheckerOpenstackImageNotFound with default headers values
func NewCheckerOpenstackImageNotFound() *CheckerOpenstackImageNotFound {
	return &CheckerOpenstackImageNotFound{}
}

/*
CheckerOpenstackImageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerOpenstackImageNotFound struct {
	Payload *CheckerOpenstackImageNotFoundBody
}

// IsSuccess returns true when this checker openstack image not found response has a 2xx status code
func (o *CheckerOpenstackImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack image not found response has a 3xx status code
func (o *CheckerOpenstackImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack image not found response has a 4xx status code
func (o *CheckerOpenstackImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker openstack image not found response has a 5xx status code
func (o *CheckerOpenstackImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker openstack image not found response a status code equal to that given
func (o *CheckerOpenstackImageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerOpenstackImageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageNotFound  %+v", 404, o.Payload)
}

func (o *CheckerOpenstackImageNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageNotFound  %+v", 404, o.Payload)
}

func (o *CheckerOpenstackImageNotFound) GetPayload() *CheckerOpenstackImageNotFoundBody {
	return o.Payload
}

func (o *CheckerOpenstackImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerOpenstackImageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackImageInternalServerError creates a CheckerOpenstackImageInternalServerError with default headers values
func NewCheckerOpenstackImageInternalServerError() *CheckerOpenstackImageInternalServerError {
	return &CheckerOpenstackImageInternalServerError{}
}

/*
CheckerOpenstackImageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerOpenstackImageInternalServerError struct {
}

// IsSuccess returns true when this checker openstack image internal server error response has a 2xx status code
func (o *CheckerOpenstackImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker openstack image internal server error response has a 3xx status code
func (o *CheckerOpenstackImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker openstack image internal server error response has a 4xx status code
func (o *CheckerOpenstackImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker openstack image internal server error response has a 5xx status code
func (o *CheckerOpenstackImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker openstack image internal server error response a status code equal to that given
func (o *CheckerOpenstackImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerOpenstackImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageInternalServerError ", 500)
}

func (o *CheckerOpenstackImageInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack-image/{id}][%d] checkerOpenstackImageInternalServerError ", 500)
}

func (o *CheckerOpenstackImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
CheckerOpenstackImageBadRequestBodyItems0 checker openstack image bad request body items0
swagger:model CheckerOpenstackImageBadRequestBodyItems0
*/
type CheckerOpenstackImageBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this checker openstack image bad request body items0
func (o *CheckerOpenstackImageBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker openstack image bad request body items0 based on context it is used
func (o *CheckerOpenstackImageBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerOpenstackImageBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerOpenstackImageBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res CheckerOpenstackImageBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerOpenstackImageForbiddenBody checker openstack image forbidden body
swagger:model CheckerOpenstackImageForbiddenBody
*/
type CheckerOpenstackImageForbiddenBody struct {

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

// Validate validates this checker openstack image forbidden body
func (o *CheckerOpenstackImageForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker openstack image forbidden body based on context it is used
func (o *CheckerOpenstackImageForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerOpenstackImageForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerOpenstackImageForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CheckerOpenstackImageForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerOpenstackImageNotFoundBody checker openstack image not found body
swagger:model CheckerOpenstackImageNotFoundBody
*/
type CheckerOpenstackImageNotFoundBody struct {

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

// Validate validates this checker openstack image not found body
func (o *CheckerOpenstackImageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker openstack image not found body based on context it is used
func (o *CheckerOpenstackImageNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerOpenstackImageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerOpenstackImageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CheckerOpenstackImageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerOpenstackImageUnauthorizedBody checker openstack image unauthorized body
swagger:model CheckerOpenstackImageUnauthorizedBody
*/
type CheckerOpenstackImageUnauthorizedBody struct {

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

// Validate validates this checker openstack image unauthorized body
func (o *CheckerOpenstackImageUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker openstack image unauthorized body based on context it is used
func (o *CheckerOpenstackImageUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerOpenstackImageUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerOpenstackImageUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CheckerOpenstackImageUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
