// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// KubernetesPatchCronJobReader is a Reader for the KubernetesPatchCronJob structure.
type KubernetesPatchCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchCronJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchCronJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchCronJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchCronJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchCronJobOK creates a KubernetesPatchCronJobOK with default headers values
func NewKubernetesPatchCronJobOK() *KubernetesPatchCronJobOK {
	return &KubernetesPatchCronJobOK{}
}

/*
KubernetesPatchCronJobOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchCronJobOK struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch cron job o k response has a 2xx status code
func (o *KubernetesPatchCronJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch cron job o k response has a 3xx status code
func (o *KubernetesPatchCronJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch cron job o k response has a 4xx status code
func (o *KubernetesPatchCronJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch cron job o k response has a 5xx status code
func (o *KubernetesPatchCronJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch cron job o k response a status code equal to that given
func (o *KubernetesPatchCronJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchCronJobOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchCronJobOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchCronJobOK) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCronJobBadRequest creates a KubernetesPatchCronJobBadRequest with default headers values
func NewKubernetesPatchCronJobBadRequest() *KubernetesPatchCronJobBadRequest {
	return &KubernetesPatchCronJobBadRequest{}
}

/*
KubernetesPatchCronJobBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchCronJobBadRequest struct {
	Payload []*KubernetesPatchCronJobBadRequestBodyItems0
}

// IsSuccess returns true when this kubernetes patch cron job bad request response has a 2xx status code
func (o *KubernetesPatchCronJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch cron job bad request response has a 3xx status code
func (o *KubernetesPatchCronJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch cron job bad request response has a 4xx status code
func (o *KubernetesPatchCronJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch cron job bad request response has a 5xx status code
func (o *KubernetesPatchCronJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch cron job bad request response a status code equal to that given
func (o *KubernetesPatchCronJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchCronJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchCronJobBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchCronJobBadRequest) GetPayload() []*KubernetesPatchCronJobBadRequestBodyItems0 {
	return o.Payload
}

func (o *KubernetesPatchCronJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCronJobUnauthorized creates a KubernetesPatchCronJobUnauthorized with default headers values
func NewKubernetesPatchCronJobUnauthorized() *KubernetesPatchCronJobUnauthorized {
	return &KubernetesPatchCronJobUnauthorized{}
}

/*
KubernetesPatchCronJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchCronJobUnauthorized struct {
	Payload *KubernetesPatchCronJobUnauthorizedBody
}

// IsSuccess returns true when this kubernetes patch cron job unauthorized response has a 2xx status code
func (o *KubernetesPatchCronJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch cron job unauthorized response has a 3xx status code
func (o *KubernetesPatchCronJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch cron job unauthorized response has a 4xx status code
func (o *KubernetesPatchCronJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch cron job unauthorized response has a 5xx status code
func (o *KubernetesPatchCronJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch cron job unauthorized response a status code equal to that given
func (o *KubernetesPatchCronJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchCronJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchCronJobUnauthorized) GetPayload() *KubernetesPatchCronJobUnauthorizedBody {
	return o.Payload
}

func (o *KubernetesPatchCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesPatchCronJobUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCronJobForbidden creates a KubernetesPatchCronJobForbidden with default headers values
func NewKubernetesPatchCronJobForbidden() *KubernetesPatchCronJobForbidden {
	return &KubernetesPatchCronJobForbidden{}
}

/*
KubernetesPatchCronJobForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchCronJobForbidden struct {
	Payload *KubernetesPatchCronJobForbiddenBody
}

// IsSuccess returns true when this kubernetes patch cron job forbidden response has a 2xx status code
func (o *KubernetesPatchCronJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch cron job forbidden response has a 3xx status code
func (o *KubernetesPatchCronJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch cron job forbidden response has a 4xx status code
func (o *KubernetesPatchCronJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch cron job forbidden response has a 5xx status code
func (o *KubernetesPatchCronJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch cron job forbidden response a status code equal to that given
func (o *KubernetesPatchCronJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchCronJobForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchCronJobForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchCronJobForbidden) GetPayload() *KubernetesPatchCronJobForbiddenBody {
	return o.Payload
}

func (o *KubernetesPatchCronJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesPatchCronJobForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCronJobNotFound creates a KubernetesPatchCronJobNotFound with default headers values
func NewKubernetesPatchCronJobNotFound() *KubernetesPatchCronJobNotFound {
	return &KubernetesPatchCronJobNotFound{}
}

/*
KubernetesPatchCronJobNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchCronJobNotFound struct {
	Payload *KubernetesPatchCronJobNotFoundBody
}

// IsSuccess returns true when this kubernetes patch cron job not found response has a 2xx status code
func (o *KubernetesPatchCronJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch cron job not found response has a 3xx status code
func (o *KubernetesPatchCronJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch cron job not found response has a 4xx status code
func (o *KubernetesPatchCronJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch cron job not found response has a 5xx status code
func (o *KubernetesPatchCronJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch cron job not found response a status code equal to that given
func (o *KubernetesPatchCronJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchCronJobNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchCronJobNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchCronJobNotFound) GetPayload() *KubernetesPatchCronJobNotFoundBody {
	return o.Payload
}

func (o *KubernetesPatchCronJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesPatchCronJobNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCronJobInternalServerError creates a KubernetesPatchCronJobInternalServerError with default headers values
func NewKubernetesPatchCronJobInternalServerError() *KubernetesPatchCronJobInternalServerError {
	return &KubernetesPatchCronJobInternalServerError{}
}

/*
KubernetesPatchCronJobInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchCronJobInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch cron job internal server error response has a 2xx status code
func (o *KubernetesPatchCronJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch cron job internal server error response has a 3xx status code
func (o *KubernetesPatchCronJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch cron job internal server error response has a 4xx status code
func (o *KubernetesPatchCronJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch cron job internal server error response has a 5xx status code
func (o *KubernetesPatchCronJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch cron job internal server error response a status code equal to that given
func (o *KubernetesPatchCronJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchCronJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobInternalServerError ", 500)
}

func (o *KubernetesPatchCronJobInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/cronjob][%d] kubernetesPatchCronJobInternalServerError ", 500)
}

func (o *KubernetesPatchCronJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
KubernetesPatchCronJobBadRequestBodyItems0 kubernetes patch cron job bad request body items0
swagger:model KubernetesPatchCronJobBadRequestBodyItems0
*/
type KubernetesPatchCronJobBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this kubernetes patch cron job bad request body items0
func (o *KubernetesPatchCronJobBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch cron job bad request body items0 based on context it is used
func (o *KubernetesPatchCronJobBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCronJobBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCronJobBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCronJobBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCronJobBody kubernetes patch cron job body
swagger:model KubernetesPatchCronJobBody
*/
type KubernetesPatchCronJobBody struct {

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// yaml
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this kubernetes patch cron job body
func (o *KubernetesPatchCronJobBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch cron job body based on context it is used
func (o *KubernetesPatchCronJobBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCronJobBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCronJobBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCronJobBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCronJobForbiddenBody kubernetes patch cron job forbidden body
swagger:model KubernetesPatchCronJobForbiddenBody
*/
type KubernetesPatchCronJobForbiddenBody struct {

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

// Validate validates this kubernetes patch cron job forbidden body
func (o *KubernetesPatchCronJobForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch cron job forbidden body based on context it is used
func (o *KubernetesPatchCronJobForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCronJobForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCronJobForbiddenBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCronJobForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCronJobNotFoundBody kubernetes patch cron job not found body
swagger:model KubernetesPatchCronJobNotFoundBody
*/
type KubernetesPatchCronJobNotFoundBody struct {

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

// Validate validates this kubernetes patch cron job not found body
func (o *KubernetesPatchCronJobNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch cron job not found body based on context it is used
func (o *KubernetesPatchCronJobNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCronJobNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCronJobNotFoundBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCronJobNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCronJobUnauthorizedBody kubernetes patch cron job unauthorized body
swagger:model KubernetesPatchCronJobUnauthorizedBody
*/
type KubernetesPatchCronJobUnauthorizedBody struct {

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

// Validate validates this kubernetes patch cron job unauthorized body
func (o *KubernetesPatchCronJobUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch cron job unauthorized body based on context it is used
func (o *KubernetesPatchCronJobUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCronJobUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCronJobUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCronJobUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
