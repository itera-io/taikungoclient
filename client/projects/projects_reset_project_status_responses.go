// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectsResetProjectStatusReader is a Reader for the ProjectsResetProjectStatus structure.
type ProjectsResetProjectStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsResetProjectStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsResetProjectStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsResetProjectStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsResetProjectStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsResetProjectStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsResetProjectStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsResetProjectStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsResetProjectStatusOK creates a ProjectsResetProjectStatusOK with default headers values
func NewProjectsResetProjectStatusOK() *ProjectsResetProjectStatusOK {
	return &ProjectsResetProjectStatusOK{}
}

/*
ProjectsResetProjectStatusOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsResetProjectStatusOK struct {
	Payload interface{}
}

// IsSuccess returns true when this projects reset project status o k response has a 2xx status code
func (o *ProjectsResetProjectStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects reset project status o k response has a 3xx status code
func (o *ProjectsResetProjectStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status o k response has a 4xx status code
func (o *ProjectsResetProjectStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects reset project status o k response has a 5xx status code
func (o *ProjectsResetProjectStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status o k response a status code equal to that given
func (o *ProjectsResetProjectStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsResetProjectStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusOK  %+v", 200, o.Payload)
}

func (o *ProjectsResetProjectStatusOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusOK  %+v", 200, o.Payload)
}

func (o *ProjectsResetProjectStatusOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsResetProjectStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusBadRequest creates a ProjectsResetProjectStatusBadRequest with default headers values
func NewProjectsResetProjectStatusBadRequest() *ProjectsResetProjectStatusBadRequest {
	return &ProjectsResetProjectStatusBadRequest{}
}

/*
ProjectsResetProjectStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsResetProjectStatusBadRequest struct {
	Payload []*ProjectsResetProjectStatusBadRequestBodyItems0
}

// IsSuccess returns true when this projects reset project status bad request response has a 2xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status bad request response has a 3xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status bad request response has a 4xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status bad request response has a 5xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status bad request response a status code equal to that given
func (o *ProjectsResetProjectStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsResetProjectStatusBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsResetProjectStatusBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsResetProjectStatusBadRequest) GetPayload() []*ProjectsResetProjectStatusBadRequestBodyItems0 {
	return o.Payload
}

func (o *ProjectsResetProjectStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusUnauthorized creates a ProjectsResetProjectStatusUnauthorized with default headers values
func NewProjectsResetProjectStatusUnauthorized() *ProjectsResetProjectStatusUnauthorized {
	return &ProjectsResetProjectStatusUnauthorized{}
}

/*
ProjectsResetProjectStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsResetProjectStatusUnauthorized struct {
	Payload *ProjectsResetProjectStatusUnauthorizedBody
}

// IsSuccess returns true when this projects reset project status unauthorized response has a 2xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status unauthorized response has a 3xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status unauthorized response has a 4xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status unauthorized response has a 5xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status unauthorized response a status code equal to that given
func (o *ProjectsResetProjectStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsResetProjectStatusUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsResetProjectStatusUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsResetProjectStatusUnauthorized) GetPayload() *ProjectsResetProjectStatusUnauthorizedBody {
	return o.Payload
}

func (o *ProjectsResetProjectStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectsResetProjectStatusUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusForbidden creates a ProjectsResetProjectStatusForbidden with default headers values
func NewProjectsResetProjectStatusForbidden() *ProjectsResetProjectStatusForbidden {
	return &ProjectsResetProjectStatusForbidden{}
}

/*
ProjectsResetProjectStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsResetProjectStatusForbidden struct {
	Payload *ProjectsResetProjectStatusForbiddenBody
}

// IsSuccess returns true when this projects reset project status forbidden response has a 2xx status code
func (o *ProjectsResetProjectStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status forbidden response has a 3xx status code
func (o *ProjectsResetProjectStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status forbidden response has a 4xx status code
func (o *ProjectsResetProjectStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status forbidden response has a 5xx status code
func (o *ProjectsResetProjectStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status forbidden response a status code equal to that given
func (o *ProjectsResetProjectStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsResetProjectStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsResetProjectStatusForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsResetProjectStatusForbidden) GetPayload() *ProjectsResetProjectStatusForbiddenBody {
	return o.Payload
}

func (o *ProjectsResetProjectStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectsResetProjectStatusForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusNotFound creates a ProjectsResetProjectStatusNotFound with default headers values
func NewProjectsResetProjectStatusNotFound() *ProjectsResetProjectStatusNotFound {
	return &ProjectsResetProjectStatusNotFound{}
}

/*
ProjectsResetProjectStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsResetProjectStatusNotFound struct {
	Payload *ProjectsResetProjectStatusNotFoundBody
}

// IsSuccess returns true when this projects reset project status not found response has a 2xx status code
func (o *ProjectsResetProjectStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status not found response has a 3xx status code
func (o *ProjectsResetProjectStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status not found response has a 4xx status code
func (o *ProjectsResetProjectStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status not found response has a 5xx status code
func (o *ProjectsResetProjectStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status not found response a status code equal to that given
func (o *ProjectsResetProjectStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsResetProjectStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsResetProjectStatusNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsResetProjectStatusNotFound) GetPayload() *ProjectsResetProjectStatusNotFoundBody {
	return o.Payload
}

func (o *ProjectsResetProjectStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectsResetProjectStatusNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusInternalServerError creates a ProjectsResetProjectStatusInternalServerError with default headers values
func NewProjectsResetProjectStatusInternalServerError() *ProjectsResetProjectStatusInternalServerError {
	return &ProjectsResetProjectStatusInternalServerError{}
}

/*
ProjectsResetProjectStatusInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsResetProjectStatusInternalServerError struct {
}

// IsSuccess returns true when this projects reset project status internal server error response has a 2xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status internal server error response has a 3xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status internal server error response has a 4xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects reset project status internal server error response has a 5xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects reset project status internal server error response a status code equal to that given
func (o *ProjectsResetProjectStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsResetProjectStatusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusInternalServerError ", 500)
}

func (o *ProjectsResetProjectStatusInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusInternalServerError ", 500)
}

func (o *ProjectsResetProjectStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ProjectsResetProjectStatusBadRequestBodyItems0 projects reset project status bad request body items0
swagger:model ProjectsResetProjectStatusBadRequestBodyItems0
*/
type ProjectsResetProjectStatusBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this projects reset project status bad request body items0
func (o *ProjectsResetProjectStatusBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects reset project status bad request body items0 based on context it is used
func (o *ProjectsResetProjectStatusBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsResetProjectStatusBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsResetProjectStatusBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res ProjectsResetProjectStatusBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsResetProjectStatusBody projects reset project status body
swagger:model ProjectsResetProjectStatusBody
*/
type ProjectsResetProjectStatusBody struct {

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// status
	// Enum: [100 145 150 154 155 156 160 165 200 250 300 400 500 550 600 700 800 900 1000 1100]
	Status int32 `json:"status,omitempty"`
}

// Validate validates this projects reset project status body
func (o *ProjectsResetProjectStatusBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var projectsResetProjectStatusBodyTypeStatusPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[100,145,150,154,155,156,160,165,200,250,300,400,500,550,600,700,800,900,1000,1100]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectsResetProjectStatusBodyTypeStatusPropEnum = append(projectsResetProjectStatusBodyTypeStatusPropEnum, v)
	}
}

// prop value enum
func (o *ProjectsResetProjectStatusBody) validateStatusEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, projectsResetProjectStatusBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ProjectsResetProjectStatusBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("body"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this projects reset project status body based on context it is used
func (o *ProjectsResetProjectStatusBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsResetProjectStatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsResetProjectStatusBody) UnmarshalBinary(b []byte) error {
	var res ProjectsResetProjectStatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsResetProjectStatusForbiddenBody projects reset project status forbidden body
swagger:model ProjectsResetProjectStatusForbiddenBody
*/
type ProjectsResetProjectStatusForbiddenBody struct {

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

// Validate validates this projects reset project status forbidden body
func (o *ProjectsResetProjectStatusForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects reset project status forbidden body based on context it is used
func (o *ProjectsResetProjectStatusForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsResetProjectStatusForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsResetProjectStatusForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ProjectsResetProjectStatusForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsResetProjectStatusNotFoundBody projects reset project status not found body
swagger:model ProjectsResetProjectStatusNotFoundBody
*/
type ProjectsResetProjectStatusNotFoundBody struct {

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

// Validate validates this projects reset project status not found body
func (o *ProjectsResetProjectStatusNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects reset project status not found body based on context it is used
func (o *ProjectsResetProjectStatusNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsResetProjectStatusNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsResetProjectStatusNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ProjectsResetProjectStatusNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsResetProjectStatusUnauthorizedBody projects reset project status unauthorized body
swagger:model ProjectsResetProjectStatusUnauthorizedBody
*/
type ProjectsResetProjectStatusUnauthorizedBody struct {

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

// Validate validates this projects reset project status unauthorized body
func (o *ProjectsResetProjectStatusUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects reset project status unauthorized body based on context it is used
func (o *ProjectsResetProjectStatusUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsResetProjectStatusUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsResetProjectStatusUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ProjectsResetProjectStatusUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
