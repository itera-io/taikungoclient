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

// KubernetesPatchCrdReader is a Reader for the KubernetesPatchCrd structure.
type KubernetesPatchCrdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchCrdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchCrdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchCrdBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchCrdUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchCrdForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchCrdNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchCrdInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchCrdOK creates a KubernetesPatchCrdOK with default headers values
func NewKubernetesPatchCrdOK() *KubernetesPatchCrdOK {
	return &KubernetesPatchCrdOK{}
}

/*
KubernetesPatchCrdOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchCrdOK struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch crd o k response has a 2xx status code
func (o *KubernetesPatchCrdOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch crd o k response has a 3xx status code
func (o *KubernetesPatchCrdOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd o k response has a 4xx status code
func (o *KubernetesPatchCrdOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch crd o k response has a 5xx status code
func (o *KubernetesPatchCrdOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd o k response a status code equal to that given
func (o *KubernetesPatchCrdOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchCrdOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchCrdOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchCrdOK) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchCrdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdBadRequest creates a KubernetesPatchCrdBadRequest with default headers values
func NewKubernetesPatchCrdBadRequest() *KubernetesPatchCrdBadRequest {
	return &KubernetesPatchCrdBadRequest{}
}

/*
KubernetesPatchCrdBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchCrdBadRequest struct {
	Payload []*KubernetesPatchCrdBadRequestBodyItems0
}

// IsSuccess returns true when this kubernetes patch crd bad request response has a 2xx status code
func (o *KubernetesPatchCrdBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd bad request response has a 3xx status code
func (o *KubernetesPatchCrdBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd bad request response has a 4xx status code
func (o *KubernetesPatchCrdBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd bad request response has a 5xx status code
func (o *KubernetesPatchCrdBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd bad request response a status code equal to that given
func (o *KubernetesPatchCrdBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchCrdBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchCrdBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchCrdBadRequest) GetPayload() []*KubernetesPatchCrdBadRequestBodyItems0 {
	return o.Payload
}

func (o *KubernetesPatchCrdBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdUnauthorized creates a KubernetesPatchCrdUnauthorized with default headers values
func NewKubernetesPatchCrdUnauthorized() *KubernetesPatchCrdUnauthorized {
	return &KubernetesPatchCrdUnauthorized{}
}

/*
KubernetesPatchCrdUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchCrdUnauthorized struct {
	Payload *KubernetesPatchCrdUnauthorizedBody
}

// IsSuccess returns true when this kubernetes patch crd unauthorized response has a 2xx status code
func (o *KubernetesPatchCrdUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd unauthorized response has a 3xx status code
func (o *KubernetesPatchCrdUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd unauthorized response has a 4xx status code
func (o *KubernetesPatchCrdUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd unauthorized response has a 5xx status code
func (o *KubernetesPatchCrdUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd unauthorized response a status code equal to that given
func (o *KubernetesPatchCrdUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchCrdUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchCrdUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchCrdUnauthorized) GetPayload() *KubernetesPatchCrdUnauthorizedBody {
	return o.Payload
}

func (o *KubernetesPatchCrdUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesPatchCrdUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdForbidden creates a KubernetesPatchCrdForbidden with default headers values
func NewKubernetesPatchCrdForbidden() *KubernetesPatchCrdForbidden {
	return &KubernetesPatchCrdForbidden{}
}

/*
KubernetesPatchCrdForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchCrdForbidden struct {
	Payload *KubernetesPatchCrdForbiddenBody
}

// IsSuccess returns true when this kubernetes patch crd forbidden response has a 2xx status code
func (o *KubernetesPatchCrdForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd forbidden response has a 3xx status code
func (o *KubernetesPatchCrdForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd forbidden response has a 4xx status code
func (o *KubernetesPatchCrdForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd forbidden response has a 5xx status code
func (o *KubernetesPatchCrdForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd forbidden response a status code equal to that given
func (o *KubernetesPatchCrdForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchCrdForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchCrdForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchCrdForbidden) GetPayload() *KubernetesPatchCrdForbiddenBody {
	return o.Payload
}

func (o *KubernetesPatchCrdForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesPatchCrdForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdNotFound creates a KubernetesPatchCrdNotFound with default headers values
func NewKubernetesPatchCrdNotFound() *KubernetesPatchCrdNotFound {
	return &KubernetesPatchCrdNotFound{}
}

/*
KubernetesPatchCrdNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchCrdNotFound struct {
	Payload *KubernetesPatchCrdNotFoundBody
}

// IsSuccess returns true when this kubernetes patch crd not found response has a 2xx status code
func (o *KubernetesPatchCrdNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd not found response has a 3xx status code
func (o *KubernetesPatchCrdNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd not found response has a 4xx status code
func (o *KubernetesPatchCrdNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd not found response has a 5xx status code
func (o *KubernetesPatchCrdNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd not found response a status code equal to that given
func (o *KubernetesPatchCrdNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchCrdNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchCrdNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchCrdNotFound) GetPayload() *KubernetesPatchCrdNotFoundBody {
	return o.Payload
}

func (o *KubernetesPatchCrdNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesPatchCrdNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdInternalServerError creates a KubernetesPatchCrdInternalServerError with default headers values
func NewKubernetesPatchCrdInternalServerError() *KubernetesPatchCrdInternalServerError {
	return &KubernetesPatchCrdInternalServerError{}
}

/*
KubernetesPatchCrdInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchCrdInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch crd internal server error response has a 2xx status code
func (o *KubernetesPatchCrdInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd internal server error response has a 3xx status code
func (o *KubernetesPatchCrdInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd internal server error response has a 4xx status code
func (o *KubernetesPatchCrdInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch crd internal server error response has a 5xx status code
func (o *KubernetesPatchCrdInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch crd internal server error response a status code equal to that given
func (o *KubernetesPatchCrdInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchCrdInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdInternalServerError ", 500)
}

func (o *KubernetesPatchCrdInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdInternalServerError ", 500)
}

func (o *KubernetesPatchCrdInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
KubernetesPatchCrdBadRequestBodyItems0 kubernetes patch crd bad request body items0
swagger:model KubernetesPatchCrdBadRequestBodyItems0
*/
type KubernetesPatchCrdBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this kubernetes patch crd bad request body items0
func (o *KubernetesPatchCrdBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch crd bad request body items0 based on context it is used
func (o *KubernetesPatchCrdBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCrdBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCrdBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCrdBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCrdBody kubernetes patch crd body
swagger:model KubernetesPatchCrdBody
*/
type KubernetesPatchCrdBody struct {

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// yaml
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this kubernetes patch crd body
func (o *KubernetesPatchCrdBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch crd body based on context it is used
func (o *KubernetesPatchCrdBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCrdBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCrdBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCrdBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCrdForbiddenBody kubernetes patch crd forbidden body
swagger:model KubernetesPatchCrdForbiddenBody
*/
type KubernetesPatchCrdForbiddenBody struct {

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

// Validate validates this kubernetes patch crd forbidden body
func (o *KubernetesPatchCrdForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch crd forbidden body based on context it is used
func (o *KubernetesPatchCrdForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCrdForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCrdForbiddenBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCrdForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCrdNotFoundBody kubernetes patch crd not found body
swagger:model KubernetesPatchCrdNotFoundBody
*/
type KubernetesPatchCrdNotFoundBody struct {

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

// Validate validates this kubernetes patch crd not found body
func (o *KubernetesPatchCrdNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch crd not found body based on context it is used
func (o *KubernetesPatchCrdNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCrdNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCrdNotFoundBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCrdNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesPatchCrdUnauthorizedBody kubernetes patch crd unauthorized body
swagger:model KubernetesPatchCrdUnauthorizedBody
*/
type KubernetesPatchCrdUnauthorizedBody struct {

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

// Validate validates this kubernetes patch crd unauthorized body
func (o *KubernetesPatchCrdUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes patch crd unauthorized body based on context it is used
func (o *KubernetesPatchCrdUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesPatchCrdUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesPatchCrdUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res KubernetesPatchCrdUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
