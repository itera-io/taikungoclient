// Code generated by go-swagger; DO NOT EDIT.

package common

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

// CommonGetSortingElementsReader is a Reader for the CommonGetSortingElements structure.
type CommonGetSortingElementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommonGetSortingElementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommonGetSortingElementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCommonGetSortingElementsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCommonGetSortingElementsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCommonGetSortingElementsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCommonGetSortingElementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCommonGetSortingElementsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCommonGetSortingElementsOK creates a CommonGetSortingElementsOK with default headers values
func NewCommonGetSortingElementsOK() *CommonGetSortingElementsOK {
	return &CommonGetSortingElementsOK{}
}

/*
CommonGetSortingElementsOK describes a response with status code 200, with default header values.

Success
*/
type CommonGetSortingElementsOK struct {
	Payload []string
}

// IsSuccess returns true when this common get sorting elements o k response has a 2xx status code
func (o *CommonGetSortingElementsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this common get sorting elements o k response has a 3xx status code
func (o *CommonGetSortingElementsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get sorting elements o k response has a 4xx status code
func (o *CommonGetSortingElementsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this common get sorting elements o k response has a 5xx status code
func (o *CommonGetSortingElementsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this common get sorting elements o k response a status code equal to that given
func (o *CommonGetSortingElementsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CommonGetSortingElementsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsOK  %+v", 200, o.Payload)
}

func (o *CommonGetSortingElementsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsOK  %+v", 200, o.Payload)
}

func (o *CommonGetSortingElementsOK) GetPayload() []string {
	return o.Payload
}

func (o *CommonGetSortingElementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetSortingElementsBadRequest creates a CommonGetSortingElementsBadRequest with default headers values
func NewCommonGetSortingElementsBadRequest() *CommonGetSortingElementsBadRequest {
	return &CommonGetSortingElementsBadRequest{}
}

/*
CommonGetSortingElementsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CommonGetSortingElementsBadRequest struct {
	Payload []*CommonGetSortingElementsBadRequestBodyItems0
}

// IsSuccess returns true when this common get sorting elements bad request response has a 2xx status code
func (o *CommonGetSortingElementsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get sorting elements bad request response has a 3xx status code
func (o *CommonGetSortingElementsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get sorting elements bad request response has a 4xx status code
func (o *CommonGetSortingElementsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get sorting elements bad request response has a 5xx status code
func (o *CommonGetSortingElementsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this common get sorting elements bad request response a status code equal to that given
func (o *CommonGetSortingElementsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CommonGetSortingElementsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsBadRequest  %+v", 400, o.Payload)
}

func (o *CommonGetSortingElementsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsBadRequest  %+v", 400, o.Payload)
}

func (o *CommonGetSortingElementsBadRequest) GetPayload() []*CommonGetSortingElementsBadRequestBodyItems0 {
	return o.Payload
}

func (o *CommonGetSortingElementsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetSortingElementsUnauthorized creates a CommonGetSortingElementsUnauthorized with default headers values
func NewCommonGetSortingElementsUnauthorized() *CommonGetSortingElementsUnauthorized {
	return &CommonGetSortingElementsUnauthorized{}
}

/*
CommonGetSortingElementsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CommonGetSortingElementsUnauthorized struct {
	Payload *CommonGetSortingElementsUnauthorizedBody
}

// IsSuccess returns true when this common get sorting elements unauthorized response has a 2xx status code
func (o *CommonGetSortingElementsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get sorting elements unauthorized response has a 3xx status code
func (o *CommonGetSortingElementsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get sorting elements unauthorized response has a 4xx status code
func (o *CommonGetSortingElementsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get sorting elements unauthorized response has a 5xx status code
func (o *CommonGetSortingElementsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this common get sorting elements unauthorized response a status code equal to that given
func (o *CommonGetSortingElementsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CommonGetSortingElementsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsUnauthorized  %+v", 401, o.Payload)
}

func (o *CommonGetSortingElementsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsUnauthorized  %+v", 401, o.Payload)
}

func (o *CommonGetSortingElementsUnauthorized) GetPayload() *CommonGetSortingElementsUnauthorizedBody {
	return o.Payload
}

func (o *CommonGetSortingElementsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CommonGetSortingElementsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetSortingElementsForbidden creates a CommonGetSortingElementsForbidden with default headers values
func NewCommonGetSortingElementsForbidden() *CommonGetSortingElementsForbidden {
	return &CommonGetSortingElementsForbidden{}
}

/*
CommonGetSortingElementsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CommonGetSortingElementsForbidden struct {
	Payload *CommonGetSortingElementsForbiddenBody
}

// IsSuccess returns true when this common get sorting elements forbidden response has a 2xx status code
func (o *CommonGetSortingElementsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get sorting elements forbidden response has a 3xx status code
func (o *CommonGetSortingElementsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get sorting elements forbidden response has a 4xx status code
func (o *CommonGetSortingElementsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get sorting elements forbidden response has a 5xx status code
func (o *CommonGetSortingElementsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this common get sorting elements forbidden response a status code equal to that given
func (o *CommonGetSortingElementsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CommonGetSortingElementsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsForbidden  %+v", 403, o.Payload)
}

func (o *CommonGetSortingElementsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsForbidden  %+v", 403, o.Payload)
}

func (o *CommonGetSortingElementsForbidden) GetPayload() *CommonGetSortingElementsForbiddenBody {
	return o.Payload
}

func (o *CommonGetSortingElementsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CommonGetSortingElementsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetSortingElementsNotFound creates a CommonGetSortingElementsNotFound with default headers values
func NewCommonGetSortingElementsNotFound() *CommonGetSortingElementsNotFound {
	return &CommonGetSortingElementsNotFound{}
}

/*
CommonGetSortingElementsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CommonGetSortingElementsNotFound struct {
	Payload *CommonGetSortingElementsNotFoundBody
}

// IsSuccess returns true when this common get sorting elements not found response has a 2xx status code
func (o *CommonGetSortingElementsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get sorting elements not found response has a 3xx status code
func (o *CommonGetSortingElementsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get sorting elements not found response has a 4xx status code
func (o *CommonGetSortingElementsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this common get sorting elements not found response has a 5xx status code
func (o *CommonGetSortingElementsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this common get sorting elements not found response a status code equal to that given
func (o *CommonGetSortingElementsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CommonGetSortingElementsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsNotFound  %+v", 404, o.Payload)
}

func (o *CommonGetSortingElementsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsNotFound  %+v", 404, o.Payload)
}

func (o *CommonGetSortingElementsNotFound) GetPayload() *CommonGetSortingElementsNotFoundBody {
	return o.Payload
}

func (o *CommonGetSortingElementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CommonGetSortingElementsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommonGetSortingElementsInternalServerError creates a CommonGetSortingElementsInternalServerError with default headers values
func NewCommonGetSortingElementsInternalServerError() *CommonGetSortingElementsInternalServerError {
	return &CommonGetSortingElementsInternalServerError{}
}

/*
CommonGetSortingElementsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CommonGetSortingElementsInternalServerError struct {
}

// IsSuccess returns true when this common get sorting elements internal server error response has a 2xx status code
func (o *CommonGetSortingElementsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this common get sorting elements internal server error response has a 3xx status code
func (o *CommonGetSortingElementsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this common get sorting elements internal server error response has a 4xx status code
func (o *CommonGetSortingElementsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this common get sorting elements internal server error response has a 5xx status code
func (o *CommonGetSortingElementsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this common get sorting elements internal server error response a status code equal to that given
func (o *CommonGetSortingElementsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CommonGetSortingElementsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsInternalServerError ", 500)
}

func (o *CommonGetSortingElementsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Common/sorting-elements/{type}][%d] commonGetSortingElementsInternalServerError ", 500)
}

func (o *CommonGetSortingElementsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
CommonGetSortingElementsBadRequestBodyItems0 common get sorting elements bad request body items0
swagger:model CommonGetSortingElementsBadRequestBodyItems0
*/
type CommonGetSortingElementsBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this common get sorting elements bad request body items0
func (o *CommonGetSortingElementsBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common get sorting elements bad request body items0 based on context it is used
func (o *CommonGetSortingElementsBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CommonGetSortingElementsBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CommonGetSortingElementsBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res CommonGetSortingElementsBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CommonGetSortingElementsForbiddenBody common get sorting elements forbidden body
swagger:model CommonGetSortingElementsForbiddenBody
*/
type CommonGetSortingElementsForbiddenBody struct {

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

// Validate validates this common get sorting elements forbidden body
func (o *CommonGetSortingElementsForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common get sorting elements forbidden body based on context it is used
func (o *CommonGetSortingElementsForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CommonGetSortingElementsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CommonGetSortingElementsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CommonGetSortingElementsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CommonGetSortingElementsNotFoundBody common get sorting elements not found body
swagger:model CommonGetSortingElementsNotFoundBody
*/
type CommonGetSortingElementsNotFoundBody struct {

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

// Validate validates this common get sorting elements not found body
func (o *CommonGetSortingElementsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common get sorting elements not found body based on context it is used
func (o *CommonGetSortingElementsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CommonGetSortingElementsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CommonGetSortingElementsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CommonGetSortingElementsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CommonGetSortingElementsUnauthorizedBody common get sorting elements unauthorized body
swagger:model CommonGetSortingElementsUnauthorizedBody
*/
type CommonGetSortingElementsUnauthorizedBody struct {

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

// Validate validates this common get sorting elements unauthorized body
func (o *CommonGetSortingElementsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common get sorting elements unauthorized body based on context it is used
func (o *CommonGetSortingElementsUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CommonGetSortingElementsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CommonGetSortingElementsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CommonGetSortingElementsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
