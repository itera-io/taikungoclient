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

// KubernetesDescribeNodeReader is a Reader for the KubernetesDescribeNode structure.
type KubernetesDescribeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeNodeOK creates a KubernetesDescribeNodeOK with default headers values
func NewKubernetesDescribeNodeOK() *KubernetesDescribeNodeOK {
	return &KubernetesDescribeNodeOK{}
}

/*
KubernetesDescribeNodeOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeNodeOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe node o k response has a 2xx status code
func (o *KubernetesDescribeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe node o k response has a 3xx status code
func (o *KubernetesDescribeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe node o k response has a 4xx status code
func (o *KubernetesDescribeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe node o k response has a 5xx status code
func (o *KubernetesDescribeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe node o k response a status code equal to that given
func (o *KubernetesDescribeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeNodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeNodeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeNodeOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNodeBadRequest creates a KubernetesDescribeNodeBadRequest with default headers values
func NewKubernetesDescribeNodeBadRequest() *KubernetesDescribeNodeBadRequest {
	return &KubernetesDescribeNodeBadRequest{}
}

/*
KubernetesDescribeNodeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeNodeBadRequest struct {
	Payload []*KubernetesDescribeNodeBadRequestBodyItems0
}

// IsSuccess returns true when this kubernetes describe node bad request response has a 2xx status code
func (o *KubernetesDescribeNodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe node bad request response has a 3xx status code
func (o *KubernetesDescribeNodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe node bad request response has a 4xx status code
func (o *KubernetesDescribeNodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe node bad request response has a 5xx status code
func (o *KubernetesDescribeNodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe node bad request response a status code equal to that given
func (o *KubernetesDescribeNodeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeNodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeNodeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeNodeBadRequest) GetPayload() []*KubernetesDescribeNodeBadRequestBodyItems0 {
	return o.Payload
}

func (o *KubernetesDescribeNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNodeUnauthorized creates a KubernetesDescribeNodeUnauthorized with default headers values
func NewKubernetesDescribeNodeUnauthorized() *KubernetesDescribeNodeUnauthorized {
	return &KubernetesDescribeNodeUnauthorized{}
}

/*
KubernetesDescribeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeNodeUnauthorized struct {
	Payload *KubernetesDescribeNodeUnauthorizedBody
}

// IsSuccess returns true when this kubernetes describe node unauthorized response has a 2xx status code
func (o *KubernetesDescribeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe node unauthorized response has a 3xx status code
func (o *KubernetesDescribeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe node unauthorized response has a 4xx status code
func (o *KubernetesDescribeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe node unauthorized response has a 5xx status code
func (o *KubernetesDescribeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe node unauthorized response a status code equal to that given
func (o *KubernetesDescribeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeNodeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeNodeUnauthorized) GetPayload() *KubernetesDescribeNodeUnauthorizedBody {
	return o.Payload
}

func (o *KubernetesDescribeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesDescribeNodeUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNodeForbidden creates a KubernetesDescribeNodeForbidden with default headers values
func NewKubernetesDescribeNodeForbidden() *KubernetesDescribeNodeForbidden {
	return &KubernetesDescribeNodeForbidden{}
}

/*
KubernetesDescribeNodeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeNodeForbidden struct {
	Payload *KubernetesDescribeNodeForbiddenBody
}

// IsSuccess returns true when this kubernetes describe node forbidden response has a 2xx status code
func (o *KubernetesDescribeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe node forbidden response has a 3xx status code
func (o *KubernetesDescribeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe node forbidden response has a 4xx status code
func (o *KubernetesDescribeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe node forbidden response has a 5xx status code
func (o *KubernetesDescribeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe node forbidden response a status code equal to that given
func (o *KubernetesDescribeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeNodeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeNodeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeNodeForbidden) GetPayload() *KubernetesDescribeNodeForbiddenBody {
	return o.Payload
}

func (o *KubernetesDescribeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesDescribeNodeForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNodeNotFound creates a KubernetesDescribeNodeNotFound with default headers values
func NewKubernetesDescribeNodeNotFound() *KubernetesDescribeNodeNotFound {
	return &KubernetesDescribeNodeNotFound{}
}

/*
KubernetesDescribeNodeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeNodeNotFound struct {
	Payload *KubernetesDescribeNodeNotFoundBody
}

// IsSuccess returns true when this kubernetes describe node not found response has a 2xx status code
func (o *KubernetesDescribeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe node not found response has a 3xx status code
func (o *KubernetesDescribeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe node not found response has a 4xx status code
func (o *KubernetesDescribeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe node not found response has a 5xx status code
func (o *KubernetesDescribeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe node not found response a status code equal to that given
func (o *KubernetesDescribeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeNodeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeNodeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeNodeNotFound) GetPayload() *KubernetesDescribeNodeNotFoundBody {
	return o.Payload
}

func (o *KubernetesDescribeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesDescribeNodeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNodeInternalServerError creates a KubernetesDescribeNodeInternalServerError with default headers values
func NewKubernetesDescribeNodeInternalServerError() *KubernetesDescribeNodeInternalServerError {
	return &KubernetesDescribeNodeInternalServerError{}
}

/*
KubernetesDescribeNodeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeNodeInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe node internal server error response has a 2xx status code
func (o *KubernetesDescribeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe node internal server error response has a 3xx status code
func (o *KubernetesDescribeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe node internal server error response has a 4xx status code
func (o *KubernetesDescribeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe node internal server error response has a 5xx status code
func (o *KubernetesDescribeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe node internal server error response a status code equal to that given
func (o *KubernetesDescribeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeInternalServerError ", 500)
}

func (o *KubernetesDescribeNodeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/node][%d] kubernetesDescribeNodeInternalServerError ", 500)
}

func (o *KubernetesDescribeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
KubernetesDescribeNodeBadRequestBodyItems0 kubernetes describe node bad request body items0
swagger:model KubernetesDescribeNodeBadRequestBodyItems0
*/
type KubernetesDescribeNodeBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this kubernetes describe node bad request body items0
func (o *KubernetesDescribeNodeBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes describe node bad request body items0 based on context it is used
func (o *KubernetesDescribeNodeBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesDescribeNodeBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesDescribeNodeBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesDescribeNodeBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesDescribeNodeBody kubernetes describe node body
swagger:model KubernetesDescribeNodeBody
*/
type KubernetesDescribeNodeBody struct {

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this kubernetes describe node body
func (o *KubernetesDescribeNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes describe node body based on context it is used
func (o *KubernetesDescribeNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesDescribeNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesDescribeNodeBody) UnmarshalBinary(b []byte) error {
	var res KubernetesDescribeNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesDescribeNodeForbiddenBody kubernetes describe node forbidden body
swagger:model KubernetesDescribeNodeForbiddenBody
*/
type KubernetesDescribeNodeForbiddenBody struct {

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

// Validate validates this kubernetes describe node forbidden body
func (o *KubernetesDescribeNodeForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes describe node forbidden body based on context it is used
func (o *KubernetesDescribeNodeForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesDescribeNodeForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesDescribeNodeForbiddenBody) UnmarshalBinary(b []byte) error {
	var res KubernetesDescribeNodeForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesDescribeNodeNotFoundBody kubernetes describe node not found body
swagger:model KubernetesDescribeNodeNotFoundBody
*/
type KubernetesDescribeNodeNotFoundBody struct {

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

// Validate validates this kubernetes describe node not found body
func (o *KubernetesDescribeNodeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes describe node not found body based on context it is used
func (o *KubernetesDescribeNodeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesDescribeNodeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesDescribeNodeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res KubernetesDescribeNodeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesDescribeNodeUnauthorizedBody kubernetes describe node unauthorized body
swagger:model KubernetesDescribeNodeUnauthorizedBody
*/
type KubernetesDescribeNodeUnauthorizedBody struct {

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

// Validate validates this kubernetes describe node unauthorized body
func (o *KubernetesDescribeNodeUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes describe node unauthorized body based on context it is used
func (o *KubernetesDescribeNodeUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesDescribeNodeUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesDescribeNodeUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res KubernetesDescribeNodeUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
