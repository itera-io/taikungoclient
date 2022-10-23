// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesGetNetworkPoliciesReader is a Reader for the KubernetesGetNetworkPolicies structure.
type KubernetesGetNetworkPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetNetworkPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetNetworkPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetNetworkPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetNetworkPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetNetworkPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetNetworkPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetNetworkPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetNetworkPoliciesOK creates a KubernetesGetNetworkPoliciesOK with default headers values
func NewKubernetesGetNetworkPoliciesOK() *KubernetesGetNetworkPoliciesOK {
	return &KubernetesGetNetworkPoliciesOK{}
}

/*
KubernetesGetNetworkPoliciesOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetNetworkPoliciesOK struct {
	Payload *KubernetesGetNetworkPoliciesOKBody
}

// IsSuccess returns true when this kubernetes get network policies o k response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get network policies o k response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies o k response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get network policies o k response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies o k response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetNetworkPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesOK) GetPayload() *KubernetesGetNetworkPoliciesOKBody {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesGetNetworkPoliciesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesBadRequest creates a KubernetesGetNetworkPoliciesBadRequest with default headers values
func NewKubernetesGetNetworkPoliciesBadRequest() *KubernetesGetNetworkPoliciesBadRequest {
	return &KubernetesGetNetworkPoliciesBadRequest{}
}

/*
KubernetesGetNetworkPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetNetworkPoliciesBadRequest struct {
	Payload []*KubernetesGetNetworkPoliciesBadRequestBodyItems0
}

// IsSuccess returns true when this kubernetes get network policies bad request response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies bad request response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies bad request response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies bad request response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies bad request response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetNetworkPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesBadRequest) GetPayload() []*KubernetesGetNetworkPoliciesBadRequestBodyItems0 {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesUnauthorized creates a KubernetesGetNetworkPoliciesUnauthorized with default headers values
func NewKubernetesGetNetworkPoliciesUnauthorized() *KubernetesGetNetworkPoliciesUnauthorized {
	return &KubernetesGetNetworkPoliciesUnauthorized{}
}

/*
KubernetesGetNetworkPoliciesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetNetworkPoliciesUnauthorized struct {
	Payload *KubernetesGetNetworkPoliciesUnauthorizedBody
}

// IsSuccess returns true when this kubernetes get network policies unauthorized response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies unauthorized response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies unauthorized response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies unauthorized response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies unauthorized response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) GetPayload() *KubernetesGetNetworkPoliciesUnauthorizedBody {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesGetNetworkPoliciesUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesForbidden creates a KubernetesGetNetworkPoliciesForbidden with default headers values
func NewKubernetesGetNetworkPoliciesForbidden() *KubernetesGetNetworkPoliciesForbidden {
	return &KubernetesGetNetworkPoliciesForbidden{}
}

/*
KubernetesGetNetworkPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetNetworkPoliciesForbidden struct {
	Payload *KubernetesGetNetworkPoliciesForbiddenBody
}

// IsSuccess returns true when this kubernetes get network policies forbidden response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies forbidden response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies forbidden response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies forbidden response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies forbidden response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetNetworkPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesForbidden) GetPayload() *KubernetesGetNetworkPoliciesForbiddenBody {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesGetNetworkPoliciesForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesNotFound creates a KubernetesGetNetworkPoliciesNotFound with default headers values
func NewKubernetesGetNetworkPoliciesNotFound() *KubernetesGetNetworkPoliciesNotFound {
	return &KubernetesGetNetworkPoliciesNotFound{}
}

/*
KubernetesGetNetworkPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetNetworkPoliciesNotFound struct {
	Payload *KubernetesGetNetworkPoliciesNotFoundBody
}

// IsSuccess returns true when this kubernetes get network policies not found response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies not found response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies not found response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies not found response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies not found response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetNetworkPoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesNotFound) GetPayload() *KubernetesGetNetworkPoliciesNotFoundBody {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(KubernetesGetNetworkPoliciesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesInternalServerError creates a KubernetesGetNetworkPoliciesInternalServerError with default headers values
func NewKubernetesGetNetworkPoliciesInternalServerError() *KubernetesGetNetworkPoliciesInternalServerError {
	return &KubernetesGetNetworkPoliciesInternalServerError{}
}

/*
KubernetesGetNetworkPoliciesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetNetworkPoliciesInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get network policies internal server error response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies internal server error response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies internal server error response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get network policies internal server error response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get network policies internal server error response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetNetworkPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesInternalServerError ", 500)
}

func (o *KubernetesGetNetworkPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesInternalServerError ", 500)
}

func (o *KubernetesGetNetworkPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
KubernetesGetNetworkPoliciesBadRequestBodyItems0 kubernetes get network policies bad request body items0
swagger:model KubernetesGetNetworkPoliciesBadRequestBodyItems0
*/
type KubernetesGetNetworkPoliciesBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this kubernetes get network policies bad request body items0
func (o *KubernetesGetNetworkPoliciesBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes get network policies bad request body items0 based on context it is used
func (o *KubernetesGetNetworkPoliciesBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesGetNetworkPoliciesBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesGetNetworkPoliciesForbiddenBody kubernetes get network policies forbidden body
swagger:model KubernetesGetNetworkPoliciesForbiddenBody
*/
type KubernetesGetNetworkPoliciesForbiddenBody struct {

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

// Validate validates this kubernetes get network policies forbidden body
func (o *KubernetesGetNetworkPoliciesForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes get network policies forbidden body based on context it is used
func (o *KubernetesGetNetworkPoliciesForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesForbiddenBody) UnmarshalBinary(b []byte) error {
	var res KubernetesGetNetworkPoliciesForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesGetNetworkPoliciesNotFoundBody kubernetes get network policies not found body
swagger:model KubernetesGetNetworkPoliciesNotFoundBody
*/
type KubernetesGetNetworkPoliciesNotFoundBody struct {

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

// Validate validates this kubernetes get network policies not found body
func (o *KubernetesGetNetworkPoliciesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes get network policies not found body based on context it is used
func (o *KubernetesGetNetworkPoliciesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res KubernetesGetNetworkPoliciesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesGetNetworkPoliciesOKBody kubernetes get network policies o k body
swagger:model KubernetesGetNetworkPoliciesOKBody
*/
type KubernetesGetNetworkPoliciesOKBody struct {

	// data
	Data []*KubernetesGetNetworkPoliciesOKBodyDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this kubernetes get network policies o k body
func (o *KubernetesGetNetworkPoliciesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KubernetesGetNetworkPoliciesOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kubernetesGetNetworkPoliciesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kubernetesGetNetworkPoliciesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this kubernetes get network policies o k body based on the context it is used
func (o *KubernetesGetNetworkPoliciesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KubernetesGetNetworkPoliciesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kubernetesGetNetworkPoliciesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kubernetesGetNetworkPoliciesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesOKBody) UnmarshalBinary(b []byte) error {
	var res KubernetesGetNetworkPoliciesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesGetNetworkPoliciesOKBodyDataItems0 kubernetes get network policies o k body data items0
swagger:model KubernetesGetNetworkPoliciesOKBodyDataItems0
*/
type KubernetesGetNetworkPoliciesOKBodyDataItems0 struct {

	// age
	Age string `json:"age,omitempty"`

	// metadata name
	MetadataName string `json:"metadataName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// pod selector
	PodSelector map[string]string `json:"podSelector,omitempty"`
}

// Validate validates this kubernetes get network policies o k body data items0
func (o *KubernetesGetNetworkPoliciesOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes get network policies o k body data items0 based on context it is used
func (o *KubernetesGetNetworkPoliciesOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesGetNetworkPoliciesOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KubernetesGetNetworkPoliciesUnauthorizedBody kubernetes get network policies unauthorized body
swagger:model KubernetesGetNetworkPoliciesUnauthorizedBody
*/
type KubernetesGetNetworkPoliciesUnauthorizedBody struct {

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

// Validate validates this kubernetes get network policies unauthorized body
func (o *KubernetesGetNetworkPoliciesUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes get network policies unauthorized body based on context it is used
func (o *KubernetesGetNetworkPoliciesUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesGetNetworkPoliciesUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res KubernetesGetNetworkPoliciesUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
