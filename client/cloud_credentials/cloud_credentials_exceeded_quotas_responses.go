// Code generated by go-swagger; DO NOT EDIT.

package cloud_credentials

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

// CloudCredentialsExceededQuotasReader is a Reader for the CloudCredentialsExceededQuotas structure.
type CloudCredentialsExceededQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsExceededQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsExceededQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsExceededQuotasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsExceededQuotasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsExceededQuotasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsExceededQuotasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsExceededQuotasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsExceededQuotasOK creates a CloudCredentialsExceededQuotasOK with default headers values
func NewCloudCredentialsExceededQuotasOK() *CloudCredentialsExceededQuotasOK {
	return &CloudCredentialsExceededQuotasOK{}
}

/*
CloudCredentialsExceededQuotasOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsExceededQuotasOK struct {
	Payload *CloudCredentialsExceededQuotasOKBody
}

// IsSuccess returns true when this cloud credentials exceeded quotas o k response has a 2xx status code
func (o *CloudCredentialsExceededQuotasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud credentials exceeded quotas o k response has a 3xx status code
func (o *CloudCredentialsExceededQuotasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials exceeded quotas o k response has a 4xx status code
func (o *CloudCredentialsExceededQuotasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials exceeded quotas o k response has a 5xx status code
func (o *CloudCredentialsExceededQuotasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials exceeded quotas o k response a status code equal to that given
func (o *CloudCredentialsExceededQuotasOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudCredentialsExceededQuotasOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsExceededQuotasOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsExceededQuotasOK) GetPayload() *CloudCredentialsExceededQuotasOKBody {
	return o.Payload
}

func (o *CloudCredentialsExceededQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloudCredentialsExceededQuotasOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsExceededQuotasBadRequest creates a CloudCredentialsExceededQuotasBadRequest with default headers values
func NewCloudCredentialsExceededQuotasBadRequest() *CloudCredentialsExceededQuotasBadRequest {
	return &CloudCredentialsExceededQuotasBadRequest{}
}

/*
CloudCredentialsExceededQuotasBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsExceededQuotasBadRequest struct {
	Payload []*CloudCredentialsExceededQuotasBadRequestBodyItems0
}

// IsSuccess returns true when this cloud credentials exceeded quotas bad request response has a 2xx status code
func (o *CloudCredentialsExceededQuotasBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials exceeded quotas bad request response has a 3xx status code
func (o *CloudCredentialsExceededQuotasBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials exceeded quotas bad request response has a 4xx status code
func (o *CloudCredentialsExceededQuotasBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials exceeded quotas bad request response has a 5xx status code
func (o *CloudCredentialsExceededQuotasBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials exceeded quotas bad request response a status code equal to that given
func (o *CloudCredentialsExceededQuotasBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CloudCredentialsExceededQuotasBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsExceededQuotasBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsExceededQuotasBadRequest) GetPayload() []*CloudCredentialsExceededQuotasBadRequestBodyItems0 {
	return o.Payload
}

func (o *CloudCredentialsExceededQuotasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsExceededQuotasUnauthorized creates a CloudCredentialsExceededQuotasUnauthorized with default headers values
func NewCloudCredentialsExceededQuotasUnauthorized() *CloudCredentialsExceededQuotasUnauthorized {
	return &CloudCredentialsExceededQuotasUnauthorized{}
}

/*
CloudCredentialsExceededQuotasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsExceededQuotasUnauthorized struct {
	Payload *CloudCredentialsExceededQuotasUnauthorizedBody
}

// IsSuccess returns true when this cloud credentials exceeded quotas unauthorized response has a 2xx status code
func (o *CloudCredentialsExceededQuotasUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials exceeded quotas unauthorized response has a 3xx status code
func (o *CloudCredentialsExceededQuotasUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials exceeded quotas unauthorized response has a 4xx status code
func (o *CloudCredentialsExceededQuotasUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials exceeded quotas unauthorized response has a 5xx status code
func (o *CloudCredentialsExceededQuotasUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials exceeded quotas unauthorized response a status code equal to that given
func (o *CloudCredentialsExceededQuotasUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CloudCredentialsExceededQuotasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsExceededQuotasUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsExceededQuotasUnauthorized) GetPayload() *CloudCredentialsExceededQuotasUnauthorizedBody {
	return o.Payload
}

func (o *CloudCredentialsExceededQuotasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloudCredentialsExceededQuotasUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsExceededQuotasForbidden creates a CloudCredentialsExceededQuotasForbidden with default headers values
func NewCloudCredentialsExceededQuotasForbidden() *CloudCredentialsExceededQuotasForbidden {
	return &CloudCredentialsExceededQuotasForbidden{}
}

/*
CloudCredentialsExceededQuotasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsExceededQuotasForbidden struct {
	Payload *CloudCredentialsExceededQuotasForbiddenBody
}

// IsSuccess returns true when this cloud credentials exceeded quotas forbidden response has a 2xx status code
func (o *CloudCredentialsExceededQuotasForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials exceeded quotas forbidden response has a 3xx status code
func (o *CloudCredentialsExceededQuotasForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials exceeded quotas forbidden response has a 4xx status code
func (o *CloudCredentialsExceededQuotasForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials exceeded quotas forbidden response has a 5xx status code
func (o *CloudCredentialsExceededQuotasForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials exceeded quotas forbidden response a status code equal to that given
func (o *CloudCredentialsExceededQuotasForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CloudCredentialsExceededQuotasForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsExceededQuotasForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsExceededQuotasForbidden) GetPayload() *CloudCredentialsExceededQuotasForbiddenBody {
	return o.Payload
}

func (o *CloudCredentialsExceededQuotasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloudCredentialsExceededQuotasForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsExceededQuotasNotFound creates a CloudCredentialsExceededQuotasNotFound with default headers values
func NewCloudCredentialsExceededQuotasNotFound() *CloudCredentialsExceededQuotasNotFound {
	return &CloudCredentialsExceededQuotasNotFound{}
}

/*
CloudCredentialsExceededQuotasNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsExceededQuotasNotFound struct {
	Payload *CloudCredentialsExceededQuotasNotFoundBody
}

// IsSuccess returns true when this cloud credentials exceeded quotas not found response has a 2xx status code
func (o *CloudCredentialsExceededQuotasNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials exceeded quotas not found response has a 3xx status code
func (o *CloudCredentialsExceededQuotasNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials exceeded quotas not found response has a 4xx status code
func (o *CloudCredentialsExceededQuotasNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials exceeded quotas not found response has a 5xx status code
func (o *CloudCredentialsExceededQuotasNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials exceeded quotas not found response a status code equal to that given
func (o *CloudCredentialsExceededQuotasNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloudCredentialsExceededQuotasNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsExceededQuotasNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsExceededQuotasNotFound) GetPayload() *CloudCredentialsExceededQuotasNotFoundBody {
	return o.Payload
}

func (o *CloudCredentialsExceededQuotasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloudCredentialsExceededQuotasNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsExceededQuotasInternalServerError creates a CloudCredentialsExceededQuotasInternalServerError with default headers values
func NewCloudCredentialsExceededQuotasInternalServerError() *CloudCredentialsExceededQuotasInternalServerError {
	return &CloudCredentialsExceededQuotasInternalServerError{}
}

/*
CloudCredentialsExceededQuotasInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsExceededQuotasInternalServerError struct {
}

// IsSuccess returns true when this cloud credentials exceeded quotas internal server error response has a 2xx status code
func (o *CloudCredentialsExceededQuotasInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials exceeded quotas internal server error response has a 3xx status code
func (o *CloudCredentialsExceededQuotasInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials exceeded quotas internal server error response has a 4xx status code
func (o *CloudCredentialsExceededQuotasInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials exceeded quotas internal server error response has a 5xx status code
func (o *CloudCredentialsExceededQuotasInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud credentials exceeded quotas internal server error response a status code equal to that given
func (o *CloudCredentialsExceededQuotasInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudCredentialsExceededQuotasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasInternalServerError ", 500)
}

func (o *CloudCredentialsExceededQuotasInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/exceeded-quotas/{organizationId}][%d] cloudCredentialsExceededQuotasInternalServerError ", 500)
}

func (o *CloudCredentialsExceededQuotasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
CloudCredentialsExceededQuotasBadRequestBodyItems0 cloud credentials exceeded quotas bad request body items0
swagger:model CloudCredentialsExceededQuotasBadRequestBodyItems0
*/
type CloudCredentialsExceededQuotasBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this cloud credentials exceeded quotas bad request body items0
func (o *CloudCredentialsExceededQuotasBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud credentials exceeded quotas bad request body items0 based on context it is used
func (o *CloudCredentialsExceededQuotasBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res CloudCredentialsExceededQuotasBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloudCredentialsExceededQuotasForbiddenBody cloud credentials exceeded quotas forbidden body
swagger:model CloudCredentialsExceededQuotasForbiddenBody
*/
type CloudCredentialsExceededQuotasForbiddenBody struct {

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

// Validate validates this cloud credentials exceeded quotas forbidden body
func (o *CloudCredentialsExceededQuotasForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud credentials exceeded quotas forbidden body based on context it is used
func (o *CloudCredentialsExceededQuotasForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CloudCredentialsExceededQuotasForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloudCredentialsExceededQuotasNotFoundBody cloud credentials exceeded quotas not found body
swagger:model CloudCredentialsExceededQuotasNotFoundBody
*/
type CloudCredentialsExceededQuotasNotFoundBody struct {

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

// Validate validates this cloud credentials exceeded quotas not found body
func (o *CloudCredentialsExceededQuotasNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud credentials exceeded quotas not found body based on context it is used
func (o *CloudCredentialsExceededQuotasNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CloudCredentialsExceededQuotasNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloudCredentialsExceededQuotasOKBody cloud credentials exceeded quotas o k body
swagger:model CloudCredentialsExceededQuotasOKBody
*/
type CloudCredentialsExceededQuotasOKBody struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this cloud credentials exceeded quotas o k body
func (o *CloudCredentialsExceededQuotasOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud credentials exceeded quotas o k body based on context it is used
func (o *CloudCredentialsExceededQuotasOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasOKBody) UnmarshalBinary(b []byte) error {
	var res CloudCredentialsExceededQuotasOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloudCredentialsExceededQuotasUnauthorizedBody cloud credentials exceeded quotas unauthorized body
swagger:model CloudCredentialsExceededQuotasUnauthorizedBody
*/
type CloudCredentialsExceededQuotasUnauthorizedBody struct {

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

// Validate validates this cloud credentials exceeded quotas unauthorized body
func (o *CloudCredentialsExceededQuotasUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud credentials exceeded quotas unauthorized body based on context it is used
func (o *CloudCredentialsExceededQuotasUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloudCredentialsExceededQuotasUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CloudCredentialsExceededQuotasUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
