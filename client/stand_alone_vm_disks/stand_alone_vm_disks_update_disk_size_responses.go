// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_vm_disks

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

// StandAloneVMDisksUpdateDiskSizeReader is a Reader for the StandAloneVMDisksUpdateDiskSize structure.
type StandAloneVMDisksUpdateDiskSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneVMDisksUpdateDiskSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneVMDisksUpdateDiskSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneVMDisksUpdateDiskSizeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneVMDisksUpdateDiskSizeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneVMDisksUpdateDiskSizeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneVMDisksUpdateDiskSizeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneVMDisksUpdateDiskSizeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneVMDisksUpdateDiskSizeOK creates a StandAloneVMDisksUpdateDiskSizeOK with default headers values
func NewStandAloneVMDisksUpdateDiskSizeOK() *StandAloneVMDisksUpdateDiskSizeOK {
	return &StandAloneVMDisksUpdateDiskSizeOK{}
}

/*
StandAloneVMDisksUpdateDiskSizeOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneVMDisksUpdateDiskSizeOK struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks update disk size o k response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone Vm disks update disk size o k response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size o k response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks update disk size o k response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size o k response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeBadRequest creates a StandAloneVMDisksUpdateDiskSizeBadRequest with default headers values
func NewStandAloneVMDisksUpdateDiskSizeBadRequest() *StandAloneVMDisksUpdateDiskSizeBadRequest {
	return &StandAloneVMDisksUpdateDiskSizeBadRequest{}
}

/*
StandAloneVMDisksUpdateDiskSizeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneVMDisksUpdateDiskSizeBadRequest struct {
	Payload []*StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0
}

// IsSuccess returns true when this stand alone Vm disks update disk size bad request response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size bad request response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size bad request response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size bad request response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size bad request response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) GetPayload() []*StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0 {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeUnauthorized creates a StandAloneVMDisksUpdateDiskSizeUnauthorized with default headers values
func NewStandAloneVMDisksUpdateDiskSizeUnauthorized() *StandAloneVMDisksUpdateDiskSizeUnauthorized {
	return &StandAloneVMDisksUpdateDiskSizeUnauthorized{}
}

/*
StandAloneVMDisksUpdateDiskSizeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneVMDisksUpdateDiskSizeUnauthorized struct {
	Payload *StandAloneVMDisksUpdateDiskSizeUnauthorizedBody
}

// IsSuccess returns true when this stand alone Vm disks update disk size unauthorized response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size unauthorized response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size unauthorized response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size unauthorized response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size unauthorized response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) GetPayload() *StandAloneVMDisksUpdateDiskSizeUnauthorizedBody {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StandAloneVMDisksUpdateDiskSizeUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeForbidden creates a StandAloneVMDisksUpdateDiskSizeForbidden with default headers values
func NewStandAloneVMDisksUpdateDiskSizeForbidden() *StandAloneVMDisksUpdateDiskSizeForbidden {
	return &StandAloneVMDisksUpdateDiskSizeForbidden{}
}

/*
StandAloneVMDisksUpdateDiskSizeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneVMDisksUpdateDiskSizeForbidden struct {
	Payload *StandAloneVMDisksUpdateDiskSizeForbiddenBody
}

// IsSuccess returns true when this stand alone Vm disks update disk size forbidden response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size forbidden response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size forbidden response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size forbidden response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size forbidden response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) GetPayload() *StandAloneVMDisksUpdateDiskSizeForbiddenBody {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StandAloneVMDisksUpdateDiskSizeForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeNotFound creates a StandAloneVMDisksUpdateDiskSizeNotFound with default headers values
func NewStandAloneVMDisksUpdateDiskSizeNotFound() *StandAloneVMDisksUpdateDiskSizeNotFound {
	return &StandAloneVMDisksUpdateDiskSizeNotFound{}
}

/*
StandAloneVMDisksUpdateDiskSizeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneVMDisksUpdateDiskSizeNotFound struct {
	Payload *StandAloneVMDisksUpdateDiskSizeNotFoundBody
}

// IsSuccess returns true when this stand alone Vm disks update disk size not found response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size not found response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size not found response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size not found response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size not found response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) GetPayload() *StandAloneVMDisksUpdateDiskSizeNotFoundBody {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StandAloneVMDisksUpdateDiskSizeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeInternalServerError creates a StandAloneVMDisksUpdateDiskSizeInternalServerError with default headers values
func NewStandAloneVMDisksUpdateDiskSizeInternalServerError() *StandAloneVMDisksUpdateDiskSizeInternalServerError {
	return &StandAloneVMDisksUpdateDiskSizeInternalServerError{}
}

/*
StandAloneVMDisksUpdateDiskSizeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneVMDisksUpdateDiskSizeInternalServerError struct {
}

// IsSuccess returns true when this stand alone Vm disks update disk size internal server error response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size internal server error response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size internal server error response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks update disk size internal server error response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone Vm disks update disk size internal server error response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeInternalServerError ", 500)
}

func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeInternalServerError ", 500)
}

func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0 stand alone VM disks update disk size bad request body items0
swagger:model StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0
*/
type StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this stand alone VM disks update disk size bad request body items0
func (o *StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone VM disks update disk size bad request body items0 based on context it is used
func (o *StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res StandAloneVMDisksUpdateDiskSizeBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StandAloneVMDisksUpdateDiskSizeBody stand alone VM disks update disk size body
swagger:model StandAloneVMDisksUpdateDiskSizeBody
*/
type StandAloneVMDisksUpdateDiskSizeBody struct {

	// id
	ID int32 `json:"id,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this stand alone VM disks update disk size body
func (o *StandAloneVMDisksUpdateDiskSizeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone VM disks update disk size body based on context it is used
func (o *StandAloneVMDisksUpdateDiskSizeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeBody) UnmarshalBinary(b []byte) error {
	var res StandAloneVMDisksUpdateDiskSizeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StandAloneVMDisksUpdateDiskSizeForbiddenBody stand alone VM disks update disk size forbidden body
swagger:model StandAloneVMDisksUpdateDiskSizeForbiddenBody
*/
type StandAloneVMDisksUpdateDiskSizeForbiddenBody struct {

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

// Validate validates this stand alone VM disks update disk size forbidden body
func (o *StandAloneVMDisksUpdateDiskSizeForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone VM disks update disk size forbidden body based on context it is used
func (o *StandAloneVMDisksUpdateDiskSizeForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeForbiddenBody) UnmarshalBinary(b []byte) error {
	var res StandAloneVMDisksUpdateDiskSizeForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StandAloneVMDisksUpdateDiskSizeNotFoundBody stand alone VM disks update disk size not found body
swagger:model StandAloneVMDisksUpdateDiskSizeNotFoundBody
*/
type StandAloneVMDisksUpdateDiskSizeNotFoundBody struct {

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

// Validate validates this stand alone VM disks update disk size not found body
func (o *StandAloneVMDisksUpdateDiskSizeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone VM disks update disk size not found body based on context it is used
func (o *StandAloneVMDisksUpdateDiskSizeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StandAloneVMDisksUpdateDiskSizeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StandAloneVMDisksUpdateDiskSizeUnauthorizedBody stand alone VM disks update disk size unauthorized body
swagger:model StandAloneVMDisksUpdateDiskSizeUnauthorizedBody
*/
type StandAloneVMDisksUpdateDiskSizeUnauthorizedBody struct {

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

// Validate validates this stand alone VM disks update disk size unauthorized body
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone VM disks update disk size unauthorized body based on context it is used
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res StandAloneVMDisksUpdateDiskSizeUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
