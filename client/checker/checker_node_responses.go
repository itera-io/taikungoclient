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

// CheckerNodeReader is a Reader for the CheckerNode structure.
type CheckerNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerNodeOK creates a CheckerNodeOK with default headers values
func NewCheckerNodeOK() *CheckerNodeOK {
	return &CheckerNodeOK{}
}

/*
CheckerNodeOK describes a response with status code 200, with default header values.

Success
*/
type CheckerNodeOK struct {
	Payload interface{}
}

// IsSuccess returns true when this checker node o k response has a 2xx status code
func (o *CheckerNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker node o k response has a 3xx status code
func (o *CheckerNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker node o k response has a 4xx status code
func (o *CheckerNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker node o k response has a 5xx status code
func (o *CheckerNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker node o k response a status code equal to that given
func (o *CheckerNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerNodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeOK  %+v", 200, o.Payload)
}

func (o *CheckerNodeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeOK  %+v", 200, o.Payload)
}

func (o *CheckerNodeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNodeBadRequest creates a CheckerNodeBadRequest with default headers values
func NewCheckerNodeBadRequest() *CheckerNodeBadRequest {
	return &CheckerNodeBadRequest{}
}

/*
CheckerNodeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerNodeBadRequest struct {
	Payload []*CheckerNodeBadRequestBodyItems0
}

// IsSuccess returns true when this checker node bad request response has a 2xx status code
func (o *CheckerNodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker node bad request response has a 3xx status code
func (o *CheckerNodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker node bad request response has a 4xx status code
func (o *CheckerNodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker node bad request response has a 5xx status code
func (o *CheckerNodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker node bad request response a status code equal to that given
func (o *CheckerNodeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerNodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerNodeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerNodeBadRequest) GetPayload() []*CheckerNodeBadRequestBodyItems0 {
	return o.Payload
}

func (o *CheckerNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNodeUnauthorized creates a CheckerNodeUnauthorized with default headers values
func NewCheckerNodeUnauthorized() *CheckerNodeUnauthorized {
	return &CheckerNodeUnauthorized{}
}

/*
CheckerNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerNodeUnauthorized struct {
	Payload *CheckerNodeUnauthorizedBody
}

// IsSuccess returns true when this checker node unauthorized response has a 2xx status code
func (o *CheckerNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker node unauthorized response has a 3xx status code
func (o *CheckerNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker node unauthorized response has a 4xx status code
func (o *CheckerNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker node unauthorized response has a 5xx status code
func (o *CheckerNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker node unauthorized response a status code equal to that given
func (o *CheckerNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerNodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerNodeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerNodeUnauthorized) GetPayload() *CheckerNodeUnauthorizedBody {
	return o.Payload
}

func (o *CheckerNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerNodeUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNodeForbidden creates a CheckerNodeForbidden with default headers values
func NewCheckerNodeForbidden() *CheckerNodeForbidden {
	return &CheckerNodeForbidden{}
}

/*
CheckerNodeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerNodeForbidden struct {
	Payload *CheckerNodeForbiddenBody
}

// IsSuccess returns true when this checker node forbidden response has a 2xx status code
func (o *CheckerNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker node forbidden response has a 3xx status code
func (o *CheckerNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker node forbidden response has a 4xx status code
func (o *CheckerNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker node forbidden response has a 5xx status code
func (o *CheckerNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker node forbidden response a status code equal to that given
func (o *CheckerNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerNodeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeForbidden  %+v", 403, o.Payload)
}

func (o *CheckerNodeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeForbidden  %+v", 403, o.Payload)
}

func (o *CheckerNodeForbidden) GetPayload() *CheckerNodeForbiddenBody {
	return o.Payload
}

func (o *CheckerNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerNodeForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNodeNotFound creates a CheckerNodeNotFound with default headers values
func NewCheckerNodeNotFound() *CheckerNodeNotFound {
	return &CheckerNodeNotFound{}
}

/*
CheckerNodeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerNodeNotFound struct {
	Payload *CheckerNodeNotFoundBody
}

// IsSuccess returns true when this checker node not found response has a 2xx status code
func (o *CheckerNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker node not found response has a 3xx status code
func (o *CheckerNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker node not found response has a 4xx status code
func (o *CheckerNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker node not found response has a 5xx status code
func (o *CheckerNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker node not found response a status code equal to that given
func (o *CheckerNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerNodeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeNotFound  %+v", 404, o.Payload)
}

func (o *CheckerNodeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeNotFound  %+v", 404, o.Payload)
}

func (o *CheckerNodeNotFound) GetPayload() *CheckerNodeNotFoundBody {
	return o.Payload
}

func (o *CheckerNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerNodeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNodeInternalServerError creates a CheckerNodeInternalServerError with default headers values
func NewCheckerNodeInternalServerError() *CheckerNodeInternalServerError {
	return &CheckerNodeInternalServerError{}
}

/*
CheckerNodeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerNodeInternalServerError struct {
}

// IsSuccess returns true when this checker node internal server error response has a 2xx status code
func (o *CheckerNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker node internal server error response has a 3xx status code
func (o *CheckerNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker node internal server error response has a 4xx status code
func (o *CheckerNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker node internal server error response has a 5xx status code
func (o *CheckerNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker node internal server error response a status code equal to that given
func (o *CheckerNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeInternalServerError ", 500)
}

func (o *CheckerNodeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/node][%d] checkerNodeInternalServerError ", 500)
}

func (o *CheckerNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
CheckerNodeBadRequestBodyItems0 checker node bad request body items0
swagger:model CheckerNodeBadRequestBodyItems0
*/
type CheckerNodeBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this checker node bad request body items0
func (o *CheckerNodeBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker node bad request body items0 based on context it is used
func (o *CheckerNodeBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerNodeBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerNodeBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res CheckerNodeBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerNodeBody checker node body
swagger:model CheckerNodeBody
*/
type CheckerNodeBody struct {

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this checker node body
func (o *CheckerNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker node body based on context it is used
func (o *CheckerNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerNodeBody) UnmarshalBinary(b []byte) error {
	var res CheckerNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerNodeForbiddenBody checker node forbidden body
swagger:model CheckerNodeForbiddenBody
*/
type CheckerNodeForbiddenBody struct {

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

// Validate validates this checker node forbidden body
func (o *CheckerNodeForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker node forbidden body based on context it is used
func (o *CheckerNodeForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerNodeForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerNodeForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CheckerNodeForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerNodeNotFoundBody checker node not found body
swagger:model CheckerNodeNotFoundBody
*/
type CheckerNodeNotFoundBody struct {

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

// Validate validates this checker node not found body
func (o *CheckerNodeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker node not found body based on context it is used
func (o *CheckerNodeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerNodeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerNodeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CheckerNodeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerNodeUnauthorizedBody checker node unauthorized body
swagger:model CheckerNodeUnauthorizedBody
*/
type CheckerNodeUnauthorizedBody struct {

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

// Validate validates this checker node unauthorized body
func (o *CheckerNodeUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker node unauthorized body based on context it is used
func (o *CheckerNodeUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerNodeUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerNodeUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CheckerNodeUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
