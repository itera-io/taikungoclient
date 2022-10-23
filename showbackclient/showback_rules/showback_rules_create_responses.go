// Code generated by go-swagger; DO NOT EDIT.

package showback_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ShowbackRulesCreateReader is a Reader for the ShowbackRulesCreate structure.
type ShowbackRulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackRulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackRulesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackRulesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackRulesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackRulesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackRulesCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackRulesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackRulesCreateOK creates a ShowbackRulesCreateOK with default headers values
func NewShowbackRulesCreateOK() *ShowbackRulesCreateOK {
	return &ShowbackRulesCreateOK{}
}

/*
ShowbackRulesCreateOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackRulesCreateOK struct {
	Payload *ShowbackRulesCreateOKBody
}

// IsSuccess returns true when this showback rules create o k response has a 2xx status code
func (o *ShowbackRulesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback rules create o k response has a 3xx status code
func (o *ShowbackRulesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules create o k response has a 4xx status code
func (o *ShowbackRulesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules create o k response has a 5xx status code
func (o *ShowbackRulesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules create o k response a status code equal to that given
func (o *ShowbackRulesCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ShowbackRulesCreateOK) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesCreateOK) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesCreateOK) GetPayload() *ShowbackRulesCreateOKBody {
	return o.Payload
}

func (o *ShowbackRulesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackRulesCreateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesCreateBadRequest creates a ShowbackRulesCreateBadRequest with default headers values
func NewShowbackRulesCreateBadRequest() *ShowbackRulesCreateBadRequest {
	return &ShowbackRulesCreateBadRequest{}
}

/*
ShowbackRulesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackRulesCreateBadRequest struct {
	Payload *ShowbackRulesCreateBadRequestBody
}

// IsSuccess returns true when this showback rules create bad request response has a 2xx status code
func (o *ShowbackRulesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules create bad request response has a 3xx status code
func (o *ShowbackRulesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules create bad request response has a 4xx status code
func (o *ShowbackRulesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules create bad request response has a 5xx status code
func (o *ShowbackRulesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules create bad request response a status code equal to that given
func (o *ShowbackRulesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ShowbackRulesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesCreateBadRequest) GetPayload() *ShowbackRulesCreateBadRequestBody {
	return o.Payload
}

func (o *ShowbackRulesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackRulesCreateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesCreateUnauthorized creates a ShowbackRulesCreateUnauthorized with default headers values
func NewShowbackRulesCreateUnauthorized() *ShowbackRulesCreateUnauthorized {
	return &ShowbackRulesCreateUnauthorized{}
}

/*
ShowbackRulesCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackRulesCreateUnauthorized struct {
	Payload *ShowbackRulesCreateUnauthorizedBody
}

// IsSuccess returns true when this showback rules create unauthorized response has a 2xx status code
func (o *ShowbackRulesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules create unauthorized response has a 3xx status code
func (o *ShowbackRulesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules create unauthorized response has a 4xx status code
func (o *ShowbackRulesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules create unauthorized response has a 5xx status code
func (o *ShowbackRulesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules create unauthorized response a status code equal to that given
func (o *ShowbackRulesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ShowbackRulesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesCreateUnauthorized) GetPayload() *ShowbackRulesCreateUnauthorizedBody {
	return o.Payload
}

func (o *ShowbackRulesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackRulesCreateUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesCreateForbidden creates a ShowbackRulesCreateForbidden with default headers values
func NewShowbackRulesCreateForbidden() *ShowbackRulesCreateForbidden {
	return &ShowbackRulesCreateForbidden{}
}

/*
ShowbackRulesCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackRulesCreateForbidden struct {
	Payload *ShowbackRulesCreateForbiddenBody
}

// IsSuccess returns true when this showback rules create forbidden response has a 2xx status code
func (o *ShowbackRulesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules create forbidden response has a 3xx status code
func (o *ShowbackRulesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules create forbidden response has a 4xx status code
func (o *ShowbackRulesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules create forbidden response has a 5xx status code
func (o *ShowbackRulesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules create forbidden response a status code equal to that given
func (o *ShowbackRulesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ShowbackRulesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesCreateForbidden) GetPayload() *ShowbackRulesCreateForbiddenBody {
	return o.Payload
}

func (o *ShowbackRulesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackRulesCreateForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesCreateNotFound creates a ShowbackRulesCreateNotFound with default headers values
func NewShowbackRulesCreateNotFound() *ShowbackRulesCreateNotFound {
	return &ShowbackRulesCreateNotFound{}
}

/*
ShowbackRulesCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackRulesCreateNotFound struct {
	Payload *ShowbackRulesCreateNotFoundBody
}

// IsSuccess returns true when this showback rules create not found response has a 2xx status code
func (o *ShowbackRulesCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules create not found response has a 3xx status code
func (o *ShowbackRulesCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules create not found response has a 4xx status code
func (o *ShowbackRulesCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules create not found response has a 5xx status code
func (o *ShowbackRulesCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules create not found response a status code equal to that given
func (o *ShowbackRulesCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ShowbackRulesCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesCreateNotFound) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesCreateNotFound) GetPayload() *ShowbackRulesCreateNotFoundBody {
	return o.Payload
}

func (o *ShowbackRulesCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ShowbackRulesCreateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesCreateInternalServerError creates a ShowbackRulesCreateInternalServerError with default headers values
func NewShowbackRulesCreateInternalServerError() *ShowbackRulesCreateInternalServerError {
	return &ShowbackRulesCreateInternalServerError{}
}

/*
ShowbackRulesCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackRulesCreateInternalServerError struct {
}

// IsSuccess returns true when this showback rules create internal server error response has a 2xx status code
func (o *ShowbackRulesCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules create internal server error response has a 3xx status code
func (o *ShowbackRulesCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules create internal server error response has a 4xx status code
func (o *ShowbackRulesCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules create internal server error response has a 5xx status code
func (o *ShowbackRulesCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback rules create internal server error response a status code equal to that given
func (o *ShowbackRulesCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ShowbackRulesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateInternalServerError ", 500)
}

func (o *ShowbackRulesCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackRules/create][%d] showbackRulesCreateInternalServerError ", 500)
}

func (o *ShowbackRulesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ShowbackRulesCreateBadRequestBody showback rules create bad request body
swagger:model ShowbackRulesCreateBadRequestBody
*/
type ShowbackRulesCreateBadRequestBody struct {

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

// Validate validates this showback rules create bad request body
func (o *ShowbackRulesCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this showback rules create bad request body based on the context it is used
func (o *ShowbackRulesCreateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ShowbackRulesCreateBadRequestBody) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackRulesCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackRulesCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ShowbackRulesCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackRulesCreateBody showback rules create body
swagger:model ShowbackRulesCreateBody
*/
type ShowbackRulesCreateBody struct {

	// by label
	ByLabel string `json:"byLabel,omitempty"`

	// global alert limit
	GlobalAlertLimit int32 `json:"globalAlertLimit,omitempty"`

	// kind
	// Enum: [100 200]
	Kind int32 `json:"kind,omitempty"`

	// labels
	Labels []*ShowbackRulesCreateParamsBodyLabelsItems0 `json:"labels"`

	// metric name
	MetricName string `json:"metricName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// project alert limit
	ProjectAlertLimit int32 `json:"projectAlertLimit,omitempty"`

	// showback credential Id
	ShowbackCredentialID *int32 `json:"showbackCredentialId,omitempty"`

	// type
	// Enum: [100 200]
	Type int32 `json:"type,omitempty"`
}

// Validate validates this showback rules create body
func (o *ShowbackRulesCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var showbackRulesCreateBodyTypeKindPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[100,200]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		showbackRulesCreateBodyTypeKindPropEnum = append(showbackRulesCreateBodyTypeKindPropEnum, v)
	}
}

// prop value enum
func (o *ShowbackRulesCreateBody) validateKindEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, showbackRulesCreateBodyTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ShowbackRulesCreateBody) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(o.Kind) { // not required
		return nil
	}

	// value enum
	if err := o.validateKindEnum("body"+"."+"kind", "body", o.Kind); err != nil {
		return err
	}

	return nil
}

func (o *ShowbackRulesCreateBody) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for i := 0; i < len(o.Labels); i++ {
		if swag.IsZero(o.Labels[i]) { // not required
			continue
		}

		if o.Labels[i] != nil {
			if err := o.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var showbackRulesCreateBodyTypeTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[100,200]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		showbackRulesCreateBodyTypeTypePropEnum = append(showbackRulesCreateBodyTypeTypePropEnum, v)
	}
}

// prop value enum
func (o *ShowbackRulesCreateBody) validateTypeEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, showbackRulesCreateBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ShowbackRulesCreateBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this showback rules create body based on the context it is used
func (o *ShowbackRulesCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ShowbackRulesCreateBody) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Labels); i++ {

		if o.Labels[i] != nil {
			if err := o.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackRulesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackRulesCreateBody) UnmarshalBinary(b []byte) error {
	var res ShowbackRulesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackRulesCreateForbiddenBody showback rules create forbidden body
swagger:model ShowbackRulesCreateForbiddenBody
*/
type ShowbackRulesCreateForbiddenBody struct {

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

// Validate validates this showback rules create forbidden body
func (o *ShowbackRulesCreateForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback rules create forbidden body based on context it is used
func (o *ShowbackRulesCreateForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackRulesCreateForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackRulesCreateForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ShowbackRulesCreateForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackRulesCreateNotFoundBody showback rules create not found body
swagger:model ShowbackRulesCreateNotFoundBody
*/
type ShowbackRulesCreateNotFoundBody struct {

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

// Validate validates this showback rules create not found body
func (o *ShowbackRulesCreateNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback rules create not found body based on context it is used
func (o *ShowbackRulesCreateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackRulesCreateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackRulesCreateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ShowbackRulesCreateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackRulesCreateOKBody showback rules create o k body
swagger:model ShowbackRulesCreateOKBody
*/
type ShowbackRulesCreateOKBody struct {

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

// Validate validates this showback rules create o k body
func (o *ShowbackRulesCreateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback rules create o k body based on context it is used
func (o *ShowbackRulesCreateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackRulesCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackRulesCreateOKBody) UnmarshalBinary(b []byte) error {
	var res ShowbackRulesCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackRulesCreateParamsBodyLabelsItems0 showback rules create params body labels items0
swagger:model ShowbackRulesCreateParamsBodyLabelsItems0
*/
type ShowbackRulesCreateParamsBodyLabelsItems0 struct {

	// label
	Label string `json:"label,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this showback rules create params body labels items0
func (o *ShowbackRulesCreateParamsBodyLabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback rules create params body labels items0 based on context it is used
func (o *ShowbackRulesCreateParamsBodyLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackRulesCreateParamsBodyLabelsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackRulesCreateParamsBodyLabelsItems0) UnmarshalBinary(b []byte) error {
	var res ShowbackRulesCreateParamsBodyLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ShowbackRulesCreateUnauthorizedBody showback rules create unauthorized body
swagger:model ShowbackRulesCreateUnauthorizedBody
*/
type ShowbackRulesCreateUnauthorizedBody struct {

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

// Validate validates this showback rules create unauthorized body
func (o *ShowbackRulesCreateUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback rules create unauthorized body based on context it is used
func (o *ShowbackRulesCreateUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ShowbackRulesCreateUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ShowbackRulesCreateUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ShowbackRulesCreateUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
