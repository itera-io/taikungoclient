// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// AzurePublishersReader is a Reader for the AzurePublishers structure.
type AzurePublishersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzurePublishersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzurePublishersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAzurePublishersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAzurePublishersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAzurePublishersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAzurePublishersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAzurePublishersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAzurePublishersOK creates a AzurePublishersOK with default headers values
func NewAzurePublishersOK() *AzurePublishersOK {
	return &AzurePublishersOK{}
}

/*
AzurePublishersOK describes a response with status code 200, with default header values.

Success
*/
type AzurePublishersOK struct {
	Payload *AzurePublishersOKBody
}

// IsSuccess returns true when this azure publishers o k response has a 2xx status code
func (o *AzurePublishersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure publishers o k response has a 3xx status code
func (o *AzurePublishersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers o k response has a 4xx status code
func (o *AzurePublishersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure publishers o k response has a 5xx status code
func (o *AzurePublishersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers o k response a status code equal to that given
func (o *AzurePublishersOK) IsCode(code int) bool {
	return code == 200
}

func (o *AzurePublishersOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersOK  %+v", 200, o.Payload)
}

func (o *AzurePublishersOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersOK  %+v", 200, o.Payload)
}

func (o *AzurePublishersOK) GetPayload() *AzurePublishersOKBody {
	return o.Payload
}

func (o *AzurePublishersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AzurePublishersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersBadRequest creates a AzurePublishersBadRequest with default headers values
func NewAzurePublishersBadRequest() *AzurePublishersBadRequest {
	return &AzurePublishersBadRequest{}
}

/*
AzurePublishersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AzurePublishersBadRequest struct {
	Payload []*AzurePublishersBadRequestBodyItems0
}

// IsSuccess returns true when this azure publishers bad request response has a 2xx status code
func (o *AzurePublishersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers bad request response has a 3xx status code
func (o *AzurePublishersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers bad request response has a 4xx status code
func (o *AzurePublishersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers bad request response has a 5xx status code
func (o *AzurePublishersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers bad request response a status code equal to that given
func (o *AzurePublishersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AzurePublishersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersBadRequest  %+v", 400, o.Payload)
}

func (o *AzurePublishersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersBadRequest  %+v", 400, o.Payload)
}

func (o *AzurePublishersBadRequest) GetPayload() []*AzurePublishersBadRequestBodyItems0 {
	return o.Payload
}

func (o *AzurePublishersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersUnauthorized creates a AzurePublishersUnauthorized with default headers values
func NewAzurePublishersUnauthorized() *AzurePublishersUnauthorized {
	return &AzurePublishersUnauthorized{}
}

/*
AzurePublishersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AzurePublishersUnauthorized struct {
	Payload *AzurePublishersUnauthorizedBody
}

// IsSuccess returns true when this azure publishers unauthorized response has a 2xx status code
func (o *AzurePublishersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers unauthorized response has a 3xx status code
func (o *AzurePublishersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers unauthorized response has a 4xx status code
func (o *AzurePublishersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers unauthorized response has a 5xx status code
func (o *AzurePublishersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers unauthorized response a status code equal to that given
func (o *AzurePublishersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AzurePublishersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersUnauthorized  %+v", 401, o.Payload)
}

func (o *AzurePublishersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersUnauthorized  %+v", 401, o.Payload)
}

func (o *AzurePublishersUnauthorized) GetPayload() *AzurePublishersUnauthorizedBody {
	return o.Payload
}

func (o *AzurePublishersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AzurePublishersUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersForbidden creates a AzurePublishersForbidden with default headers values
func NewAzurePublishersForbidden() *AzurePublishersForbidden {
	return &AzurePublishersForbidden{}
}

/*
AzurePublishersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AzurePublishersForbidden struct {
	Payload *AzurePublishersForbiddenBody
}

// IsSuccess returns true when this azure publishers forbidden response has a 2xx status code
func (o *AzurePublishersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers forbidden response has a 3xx status code
func (o *AzurePublishersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers forbidden response has a 4xx status code
func (o *AzurePublishersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers forbidden response has a 5xx status code
func (o *AzurePublishersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers forbidden response a status code equal to that given
func (o *AzurePublishersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AzurePublishersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersForbidden  %+v", 403, o.Payload)
}

func (o *AzurePublishersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersForbidden  %+v", 403, o.Payload)
}

func (o *AzurePublishersForbidden) GetPayload() *AzurePublishersForbiddenBody {
	return o.Payload
}

func (o *AzurePublishersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AzurePublishersForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersNotFound creates a AzurePublishersNotFound with default headers values
func NewAzurePublishersNotFound() *AzurePublishersNotFound {
	return &AzurePublishersNotFound{}
}

/*
AzurePublishersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AzurePublishersNotFound struct {
	Payload *AzurePublishersNotFoundBody
}

// IsSuccess returns true when this azure publishers not found response has a 2xx status code
func (o *AzurePublishersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers not found response has a 3xx status code
func (o *AzurePublishersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers not found response has a 4xx status code
func (o *AzurePublishersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this azure publishers not found response has a 5xx status code
func (o *AzurePublishersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this azure publishers not found response a status code equal to that given
func (o *AzurePublishersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AzurePublishersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersNotFound  %+v", 404, o.Payload)
}

func (o *AzurePublishersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersNotFound  %+v", 404, o.Payload)
}

func (o *AzurePublishersNotFound) GetPayload() *AzurePublishersNotFoundBody {
	return o.Payload
}

func (o *AzurePublishersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AzurePublishersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzurePublishersInternalServerError creates a AzurePublishersInternalServerError with default headers values
func NewAzurePublishersInternalServerError() *AzurePublishersInternalServerError {
	return &AzurePublishersInternalServerError{}
}

/*
AzurePublishersInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AzurePublishersInternalServerError struct {
}

// IsSuccess returns true when this azure publishers internal server error response has a 2xx status code
func (o *AzurePublishersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this azure publishers internal server error response has a 3xx status code
func (o *AzurePublishersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure publishers internal server error response has a 4xx status code
func (o *AzurePublishersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure publishers internal server error response has a 5xx status code
func (o *AzurePublishersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this azure publishers internal server error response a status code equal to that given
func (o *AzurePublishersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AzurePublishersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersInternalServerError ", 500)
}

func (o *AzurePublishersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Azure/publishers/{cloudId}][%d] azurePublishersInternalServerError ", 500)
}

func (o *AzurePublishersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
AzurePublishersBadRequestBodyItems0 azure publishers bad request body items0
swagger:model AzurePublishersBadRequestBodyItems0
*/
type AzurePublishersBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this azure publishers bad request body items0
func (o *AzurePublishersBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure publishers bad request body items0 based on context it is used
func (o *AzurePublishersBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AzurePublishersBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AzurePublishersBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res AzurePublishersBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AzurePublishersForbiddenBody azure publishers forbidden body
swagger:model AzurePublishersForbiddenBody
*/
type AzurePublishersForbiddenBody struct {

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

// Validate validates this azure publishers forbidden body
func (o *AzurePublishersForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure publishers forbidden body based on context it is used
func (o *AzurePublishersForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AzurePublishersForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AzurePublishersForbiddenBody) UnmarshalBinary(b []byte) error {
	var res AzurePublishersForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AzurePublishersNotFoundBody azure publishers not found body
swagger:model AzurePublishersNotFoundBody
*/
type AzurePublishersNotFoundBody struct {

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

// Validate validates this azure publishers not found body
func (o *AzurePublishersNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure publishers not found body based on context it is used
func (o *AzurePublishersNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AzurePublishersNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AzurePublishersNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AzurePublishersNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AzurePublishersOKBody azure publishers o k body
swagger:model AzurePublishersOKBody
*/
type AzurePublishersOKBody struct {

	// data
	Data []string `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this azure publishers o k body
func (o *AzurePublishersOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure publishers o k body based on context it is used
func (o *AzurePublishersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AzurePublishersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AzurePublishersOKBody) UnmarshalBinary(b []byte) error {
	var res AzurePublishersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AzurePublishersUnauthorizedBody azure publishers unauthorized body
swagger:model AzurePublishersUnauthorizedBody
*/
type AzurePublishersUnauthorizedBody struct {

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

// Validate validates this azure publishers unauthorized body
func (o *AzurePublishersUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure publishers unauthorized body based on context it is used
func (o *AzurePublishersUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AzurePublishersUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AzurePublishersUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res AzurePublishersUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
