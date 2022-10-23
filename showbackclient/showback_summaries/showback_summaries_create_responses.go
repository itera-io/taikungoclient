// Code generated by go-swagger; DO NOT EDIT.

package showback_summaries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ShowbackSummariesCreateReader is a Reader for the ShowbackSummariesCreate structure.
type ShowbackSummariesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackSummariesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackSummariesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackSummariesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackSummariesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackSummariesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackSummariesCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackSummariesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackSummariesCreateOK creates a ShowbackSummariesCreateOK with default headers values
func NewShowbackSummariesCreateOK() *ShowbackSummariesCreateOK {
	return &ShowbackSummariesCreateOK{}
}

/*
ShowbackSummariesCreateOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackSummariesCreateOK struct {
	Payload interface{}
}

// IsSuccess returns true when this showback summaries create o k response has a 2xx status code
func (o *ShowbackSummariesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback summaries create o k response has a 3xx status code
func (o *ShowbackSummariesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries create o k response has a 4xx status code
func (o *ShowbackSummariesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries create o k response has a 5xx status code
func (o *ShowbackSummariesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries create o k response a status code equal to that given
func (o *ShowbackSummariesCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ShowbackSummariesCreateOK) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateOK  %+v", 200, o.Payload)
}

func (o *ShowbackSummariesCreateOK) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateOK  %+v", 200, o.Payload)
}

func (o *ShowbackSummariesCreateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackSummariesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesCreateBadRequest creates a ShowbackSummariesCreateBadRequest with default headers values
func NewShowbackSummariesCreateBadRequest() *ShowbackSummariesCreateBadRequest {
	return &ShowbackSummariesCreateBadRequest{}
}

/*
ShowbackSummariesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackSummariesCreateBadRequest struct {
	Payload *ShowbackSummariesCreateBadRequestBody
}

// IsSuccess returns true when this showback summaries create bad request response has a 2xx status code
func (o *ShowbackSummariesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries create bad request response has a 3xx status code
func (o *ShowbackSummariesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries create bad request response has a 4xx status code
func (o *ShowbackSummariesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries create bad request response has a 5xx status code
func (o *ShowbackSummariesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries create bad request response a status code equal to that given
func (o *ShowbackSummariesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ShowbackSummariesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesCreateBadRequest) GetPayload() *ShowbackSummariesCreateBadRequestBody {
	return o.Payload
}

func (o *ShowbackSummariesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackSummariesCreateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesCreateUnauthorized creates a ShowbackSummariesCreateUnauthorized with default headers values
func NewShowbackSummariesCreateUnauthorized() *ShowbackSummariesCreateUnauthorized {
	return &ShowbackSummariesCreateUnauthorized{}
}

/*
ShowbackSummariesCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackSummariesCreateUnauthorized struct {
	Payload *ShowbackSummariesCreateUnauthorizedBody
}

// IsSuccess returns true when this showback summaries create unauthorized response has a 2xx status code
func (o *ShowbackSummariesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries create unauthorized response has a 3xx status code
func (o *ShowbackSummariesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries create unauthorized response has a 4xx status code
func (o *ShowbackSummariesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries create unauthorized response has a 5xx status code
func (o *ShowbackSummariesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries create unauthorized response a status code equal to that given
func (o *ShowbackSummariesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ShowbackSummariesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesCreateUnauthorized) GetPayload() *ShowbackSummariesCreateUnauthorizedBody {
	return o.Payload
}

func (o *ShowbackSummariesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackSummariesCreateUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesCreateForbidden creates a ShowbackSummariesCreateForbidden with default headers values
func NewShowbackSummariesCreateForbidden() *ShowbackSummariesCreateForbidden {
	return &ShowbackSummariesCreateForbidden{}
}

/*
ShowbackSummariesCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackSummariesCreateForbidden struct {
	Payload *ShowbackSummariesCreateForbiddenBody
}

// IsSuccess returns true when this showback summaries create forbidden response has a 2xx status code
func (o *ShowbackSummariesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries create forbidden response has a 3xx status code
func (o *ShowbackSummariesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries create forbidden response has a 4xx status code
func (o *ShowbackSummariesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries create forbidden response has a 5xx status code
func (o *ShowbackSummariesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries create forbidden response a status code equal to that given
func (o *ShowbackSummariesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ShowbackSummariesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesCreateForbidden) GetPayload() *ShowbackSummariesCreateForbiddenBody {
	return o.Payload
}

func (o *ShowbackSummariesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackSummariesCreateForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesCreateNotFound creates a ShowbackSummariesCreateNotFound with default headers values
func NewShowbackSummariesCreateNotFound() *ShowbackSummariesCreateNotFound {
	return &ShowbackSummariesCreateNotFound{}
}

/*
ShowbackSummariesCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackSummariesCreateNotFound struct {
	Payload *ShowbackSummariesCreateNotFoundBody
}

// IsSuccess returns true when this showback summaries create not found response has a 2xx status code
func (o *ShowbackSummariesCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries create not found response has a 3xx status code
func (o *ShowbackSummariesCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries create not found response has a 4xx status code
func (o *ShowbackSummariesCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries create not found response has a 5xx status code
func (o *ShowbackSummariesCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries create not found response a status code equal to that given
func (o *ShowbackSummariesCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ShowbackSummariesCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesCreateNotFound) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesCreateNotFound) GetPayload() *ShowbackSummariesCreateNotFoundBody {
	return o.Payload
}

func (o *ShowbackSummariesCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackSummariesCreateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesCreateInternalServerError creates a ShowbackSummariesCreateInternalServerError with default headers values
func NewShowbackSummariesCreateInternalServerError() *ShowbackSummariesCreateInternalServerError {
	return &ShowbackSummariesCreateInternalServerError{}
}

/*
ShowbackSummariesCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackSummariesCreateInternalServerError struct {
}

// IsSuccess returns true when this showback summaries create internal server error response has a 2xx status code
func (o *ShowbackSummariesCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries create internal server error response has a 3xx status code
func (o *ShowbackSummariesCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries create internal server error response has a 4xx status code
func (o *ShowbackSummariesCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries create internal server error response has a 5xx status code
func (o *ShowbackSummariesCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback summaries create internal server error response a status code equal to that given
func (o *ShowbackSummariesCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ShowbackSummariesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateInternalServerError ", 500)
}

func (o *ShowbackSummariesCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackSummaries/create][%d] showbackSummariesCreateInternalServerError ", 500)
}

func (o *ShowbackSummariesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ShowbackSummariesCreateBadRequestBody showback summaries create bad request body
swagger:model ShowbackSummariesCreateBadRequestBody
*/
type ShowbackSummariesCreateBadRequestBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// errors
	// Read Only: true
	Errors map[string][]string `json:"errors,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this showback summaries create bad request body
func (o *ShowbackSummariesCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this showback summaries create bad request body based on the context it is used
func (o *ShowbackSummariesCreateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ShowbackSummariesCreateBadRequestBody) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackSummariesCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackSummariesCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ShowbackSummariesCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackSummariesCreateBody showback summaries create body
swagger:model ShowbackSummariesCreateBody
*/
type ShowbackSummariesCreateBody struct {

	// begin apply
	// Format: date-time
	BeginApply *strfmt.DateTime `json:"beginApply,omitempty"`

	// by label value
	ByLabelValue string `json:"byLabelValue,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// showback rule Id
	ShowbackRuleID int32 `json:"showbackRuleId,omitempty"`
}

// Validate validates this showback summaries create body
func (o *ShowbackSummariesCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBeginApply(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ShowbackSummariesCreateBody) validateBeginApply(formats strfmt.Registry) error {
	if swag.IsZero(o.BeginApply) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"beginApply", "body", "date-time", o.BeginApply.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this showback summaries create body based on context it is used
func (o *ShowbackSummariesCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackSummariesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackSummariesCreateBody) UnmarshalBinary(b []byte) error {
	var res ShowbackSummariesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackSummariesCreateForbiddenBody showback summaries create forbidden body
swagger:model ShowbackSummariesCreateForbiddenBody
*/
type ShowbackSummariesCreateForbiddenBody struct {

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

// Validate validates this showback summaries create forbidden body
func (o *ShowbackSummariesCreateForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback summaries create forbidden body based on context it is used
func (o *ShowbackSummariesCreateForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackSummariesCreateForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackSummariesCreateForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ShowbackSummariesCreateForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackSummariesCreateNotFoundBody showback summaries create not found body
swagger:model ShowbackSummariesCreateNotFoundBody
*/
type ShowbackSummariesCreateNotFoundBody struct {

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

// Validate validates this showback summaries create not found body
func (o *ShowbackSummariesCreateNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback summaries create not found body based on context it is used
func (o *ShowbackSummariesCreateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackSummariesCreateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackSummariesCreateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ShowbackSummariesCreateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackSummariesCreateUnauthorizedBody showback summaries create unauthorized body
swagger:model ShowbackSummariesCreateUnauthorizedBody
*/
type ShowbackSummariesCreateUnauthorizedBody struct {

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

// Validate validates this showback summaries create unauthorized body
func (o *ShowbackSummariesCreateUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback summaries create unauthorized body based on context it is used
func (o *ShowbackSummariesCreateUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackSummariesCreateUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackSummariesCreateUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ShowbackSummariesCreateUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
