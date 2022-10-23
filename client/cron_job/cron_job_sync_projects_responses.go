// Code generated by go-swagger; DO NOT EDIT.

package cron_job

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

// CronJobSyncProjectsReader is a Reader for the CronJobSyncProjects structure.
type CronJobSyncProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobSyncProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobSyncProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobSyncProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobSyncProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobSyncProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobSyncProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobSyncProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobSyncProjectsOK creates a CronJobSyncProjectsOK with default headers values
func NewCronJobSyncProjectsOK() *CronJobSyncProjectsOK {
	return &CronJobSyncProjectsOK{}
}

/*
CronJobSyncProjectsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobSyncProjectsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job sync projects o k response has a 2xx status code
func (o *CronJobSyncProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job sync projects o k response has a 3xx status code
func (o *CronJobSyncProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync projects o k response has a 4xx status code
func (o *CronJobSyncProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync projects o k response has a 5xx status code
func (o *CronJobSyncProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync projects o k response a status code equal to that given
func (o *CronJobSyncProjectsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobSyncProjectsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncProjectsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncProjectsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobSyncProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectsBadRequest creates a CronJobSyncProjectsBadRequest with default headers values
func NewCronJobSyncProjectsBadRequest() *CronJobSyncProjectsBadRequest {
	return &CronJobSyncProjectsBadRequest{}
}

/*
CronJobSyncProjectsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobSyncProjectsBadRequest struct {
	Payload []*CronJobSyncProjectsBadRequestBodyItems0
}

// IsSuccess returns true when this cron job sync projects bad request response has a 2xx status code
func (o *CronJobSyncProjectsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync projects bad request response has a 3xx status code
func (o *CronJobSyncProjectsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync projects bad request response has a 4xx status code
func (o *CronJobSyncProjectsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync projects bad request response has a 5xx status code
func (o *CronJobSyncProjectsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync projects bad request response a status code equal to that given
func (o *CronJobSyncProjectsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobSyncProjectsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncProjectsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncProjectsBadRequest) GetPayload() []*CronJobSyncProjectsBadRequestBodyItems0 {
	return o.Payload
}

func (o *CronJobSyncProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectsUnauthorized creates a CronJobSyncProjectsUnauthorized with default headers values
func NewCronJobSyncProjectsUnauthorized() *CronJobSyncProjectsUnauthorized {
	return &CronJobSyncProjectsUnauthorized{}
}

/*
CronJobSyncProjectsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobSyncProjectsUnauthorized struct {
	Payload *CronJobSyncProjectsUnauthorizedBody
}

// IsSuccess returns true when this cron job sync projects unauthorized response has a 2xx status code
func (o *CronJobSyncProjectsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync projects unauthorized response has a 3xx status code
func (o *CronJobSyncProjectsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync projects unauthorized response has a 4xx status code
func (o *CronJobSyncProjectsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync projects unauthorized response has a 5xx status code
func (o *CronJobSyncProjectsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync projects unauthorized response a status code equal to that given
func (o *CronJobSyncProjectsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobSyncProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncProjectsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncProjectsUnauthorized) GetPayload() *CronJobSyncProjectsUnauthorizedBody {
	return o.Payload
}

func (o *CronJobSyncProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CronJobSyncProjectsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectsForbidden creates a CronJobSyncProjectsForbidden with default headers values
func NewCronJobSyncProjectsForbidden() *CronJobSyncProjectsForbidden {
	return &CronJobSyncProjectsForbidden{}
}

/*
CronJobSyncProjectsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobSyncProjectsForbidden struct {
	Payload *CronJobSyncProjectsForbiddenBody
}

// IsSuccess returns true when this cron job sync projects forbidden response has a 2xx status code
func (o *CronJobSyncProjectsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync projects forbidden response has a 3xx status code
func (o *CronJobSyncProjectsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync projects forbidden response has a 4xx status code
func (o *CronJobSyncProjectsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync projects forbidden response has a 5xx status code
func (o *CronJobSyncProjectsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync projects forbidden response a status code equal to that given
func (o *CronJobSyncProjectsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobSyncProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncProjectsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncProjectsForbidden) GetPayload() *CronJobSyncProjectsForbiddenBody {
	return o.Payload
}

func (o *CronJobSyncProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CronJobSyncProjectsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectsNotFound creates a CronJobSyncProjectsNotFound with default headers values
func NewCronJobSyncProjectsNotFound() *CronJobSyncProjectsNotFound {
	return &CronJobSyncProjectsNotFound{}
}

/*
CronJobSyncProjectsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobSyncProjectsNotFound struct {
	Payload *CronJobSyncProjectsNotFoundBody
}

// IsSuccess returns true when this cron job sync projects not found response has a 2xx status code
func (o *CronJobSyncProjectsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync projects not found response has a 3xx status code
func (o *CronJobSyncProjectsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync projects not found response has a 4xx status code
func (o *CronJobSyncProjectsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync projects not found response has a 5xx status code
func (o *CronJobSyncProjectsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync projects not found response a status code equal to that given
func (o *CronJobSyncProjectsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobSyncProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncProjectsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncProjectsNotFound) GetPayload() *CronJobSyncProjectsNotFoundBody {
	return o.Payload
}

func (o *CronJobSyncProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CronJobSyncProjectsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectsInternalServerError creates a CronJobSyncProjectsInternalServerError with default headers values
func NewCronJobSyncProjectsInternalServerError() *CronJobSyncProjectsInternalServerError {
	return &CronJobSyncProjectsInternalServerError{}
}

/*
CronJobSyncProjectsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobSyncProjectsInternalServerError struct {
}

// IsSuccess returns true when this cron job sync projects internal server error response has a 2xx status code
func (o *CronJobSyncProjectsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync projects internal server error response has a 3xx status code
func (o *CronJobSyncProjectsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync projects internal server error response has a 4xx status code
func (o *CronJobSyncProjectsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync projects internal server error response has a 5xx status code
func (o *CronJobSyncProjectsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job sync projects internal server error response a status code equal to that given
func (o *CronJobSyncProjectsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobSyncProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsInternalServerError ", 500)
}

func (o *CronJobSyncProjectsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-projects][%d] cronJobSyncProjectsInternalServerError ", 500)
}

func (o *CronJobSyncProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
CronJobSyncProjectsBadRequestBodyItems0 cron job sync projects bad request body items0
swagger:model CronJobSyncProjectsBadRequestBodyItems0
*/
type CronJobSyncProjectsBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this cron job sync projects bad request body items0
func (o *CronJobSyncProjectsBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cron job sync projects bad request body items0 based on context it is used
func (o *CronJobSyncProjectsBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CronJobSyncProjectsBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CronJobSyncProjectsBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res CronJobSyncProjectsBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CronJobSyncProjectsForbiddenBody cron job sync projects forbidden body
swagger:model CronJobSyncProjectsForbiddenBody
*/
type CronJobSyncProjectsForbiddenBody struct {

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

// Validate validates this cron job sync projects forbidden body
func (o *CronJobSyncProjectsForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cron job sync projects forbidden body based on context it is used
func (o *CronJobSyncProjectsForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CronJobSyncProjectsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CronJobSyncProjectsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CronJobSyncProjectsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CronJobSyncProjectsNotFoundBody cron job sync projects not found body
swagger:model CronJobSyncProjectsNotFoundBody
*/
type CronJobSyncProjectsNotFoundBody struct {

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

// Validate validates this cron job sync projects not found body
func (o *CronJobSyncProjectsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cron job sync projects not found body based on context it is used
func (o *CronJobSyncProjectsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CronJobSyncProjectsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CronJobSyncProjectsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CronJobSyncProjectsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CronJobSyncProjectsUnauthorizedBody cron job sync projects unauthorized body
swagger:model CronJobSyncProjectsUnauthorizedBody
*/
type CronJobSyncProjectsUnauthorizedBody struct {

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

// Validate validates this cron job sync projects unauthorized body
func (o *CronJobSyncProjectsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cron job sync projects unauthorized body based on context it is used
func (o *CronJobSyncProjectsUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CronJobSyncProjectsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CronJobSyncProjectsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CronJobSyncProjectsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
