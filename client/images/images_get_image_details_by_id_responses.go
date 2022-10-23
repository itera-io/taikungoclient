// Code generated by go-swagger; DO NOT EDIT.

package images

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

// ImagesGetImageDetailsByIDReader is a Reader for the ImagesGetImageDetailsByID structure.
type ImagesGetImageDetailsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesGetImageDetailsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesGetImageDetailsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesGetImageDetailsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesGetImageDetailsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesGetImageDetailsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesGetImageDetailsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesGetImageDetailsByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesGetImageDetailsByIDOK creates a ImagesGetImageDetailsByIDOK with default headers values
func NewImagesGetImageDetailsByIDOK() *ImagesGetImageDetailsByIDOK {
	return &ImagesGetImageDetailsByIDOK{}
}

/*
ImagesGetImageDetailsByIDOK describes a response with status code 200, with default header values.

Success
*/
type ImagesGetImageDetailsByIDOK struct {
	Payload string
}

// IsSuccess returns true when this images get image details by Id o k response has a 2xx status code
func (o *ImagesGetImageDetailsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images get image details by Id o k response has a 3xx status code
func (o *ImagesGetImageDetailsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get image details by Id o k response has a 4xx status code
func (o *ImagesGetImageDetailsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images get image details by Id o k response has a 5xx status code
func (o *ImagesGetImageDetailsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images get image details by Id o k response a status code equal to that given
func (o *ImagesGetImageDetailsByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesGetImageDetailsByIDOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdOK  %+v", 200, o.Payload)
}

func (o *ImagesGetImageDetailsByIDOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdOK  %+v", 200, o.Payload)
}

func (o *ImagesGetImageDetailsByIDOK) GetPayload() string {
	return o.Payload
}

func (o *ImagesGetImageDetailsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetImageDetailsByIDBadRequest creates a ImagesGetImageDetailsByIDBadRequest with default headers values
func NewImagesGetImageDetailsByIDBadRequest() *ImagesGetImageDetailsByIDBadRequest {
	return &ImagesGetImageDetailsByIDBadRequest{}
}

/*
ImagesGetImageDetailsByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesGetImageDetailsByIDBadRequest struct {
	Payload []*ImagesGetImageDetailsByIDBadRequestBodyItems0
}

// IsSuccess returns true when this images get image details by Id bad request response has a 2xx status code
func (o *ImagesGetImageDetailsByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get image details by Id bad request response has a 3xx status code
func (o *ImagesGetImageDetailsByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get image details by Id bad request response has a 4xx status code
func (o *ImagesGetImageDetailsByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get image details by Id bad request response has a 5xx status code
func (o *ImagesGetImageDetailsByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images get image details by Id bad request response a status code equal to that given
func (o *ImagesGetImageDetailsByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesGetImageDetailsByIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesGetImageDetailsByIDBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesGetImageDetailsByIDBadRequest) GetPayload() []*ImagesGetImageDetailsByIDBadRequestBodyItems0 {
	return o.Payload
}

func (o *ImagesGetImageDetailsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetImageDetailsByIDUnauthorized creates a ImagesGetImageDetailsByIDUnauthorized with default headers values
func NewImagesGetImageDetailsByIDUnauthorized() *ImagesGetImageDetailsByIDUnauthorized {
	return &ImagesGetImageDetailsByIDUnauthorized{}
}

/*
ImagesGetImageDetailsByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesGetImageDetailsByIDUnauthorized struct {
	Payload *ImagesGetImageDetailsByIDUnauthorizedBody
}

// IsSuccess returns true when this images get image details by Id unauthorized response has a 2xx status code
func (o *ImagesGetImageDetailsByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get image details by Id unauthorized response has a 3xx status code
func (o *ImagesGetImageDetailsByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get image details by Id unauthorized response has a 4xx status code
func (o *ImagesGetImageDetailsByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get image details by Id unauthorized response has a 5xx status code
func (o *ImagesGetImageDetailsByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images get image details by Id unauthorized response a status code equal to that given
func (o *ImagesGetImageDetailsByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesGetImageDetailsByIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesGetImageDetailsByIDUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesGetImageDetailsByIDUnauthorized) GetPayload() *ImagesGetImageDetailsByIDUnauthorizedBody {
	return o.Payload
}

func (o *ImagesGetImageDetailsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagesGetImageDetailsByIDUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetImageDetailsByIDForbidden creates a ImagesGetImageDetailsByIDForbidden with default headers values
func NewImagesGetImageDetailsByIDForbidden() *ImagesGetImageDetailsByIDForbidden {
	return &ImagesGetImageDetailsByIDForbidden{}
}

/*
ImagesGetImageDetailsByIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesGetImageDetailsByIDForbidden struct {
	Payload *ImagesGetImageDetailsByIDForbiddenBody
}

// IsSuccess returns true when this images get image details by Id forbidden response has a 2xx status code
func (o *ImagesGetImageDetailsByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get image details by Id forbidden response has a 3xx status code
func (o *ImagesGetImageDetailsByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get image details by Id forbidden response has a 4xx status code
func (o *ImagesGetImageDetailsByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get image details by Id forbidden response has a 5xx status code
func (o *ImagesGetImageDetailsByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images get image details by Id forbidden response a status code equal to that given
func (o *ImagesGetImageDetailsByIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesGetImageDetailsByIDForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdForbidden  %+v", 403, o.Payload)
}

func (o *ImagesGetImageDetailsByIDForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdForbidden  %+v", 403, o.Payload)
}

func (o *ImagesGetImageDetailsByIDForbidden) GetPayload() *ImagesGetImageDetailsByIDForbiddenBody {
	return o.Payload
}

func (o *ImagesGetImageDetailsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagesGetImageDetailsByIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetImageDetailsByIDNotFound creates a ImagesGetImageDetailsByIDNotFound with default headers values
func NewImagesGetImageDetailsByIDNotFound() *ImagesGetImageDetailsByIDNotFound {
	return &ImagesGetImageDetailsByIDNotFound{}
}

/*
ImagesGetImageDetailsByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesGetImageDetailsByIDNotFound struct {
	Payload *ImagesGetImageDetailsByIDNotFoundBody
}

// IsSuccess returns true when this images get image details by Id not found response has a 2xx status code
func (o *ImagesGetImageDetailsByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get image details by Id not found response has a 3xx status code
func (o *ImagesGetImageDetailsByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get image details by Id not found response has a 4xx status code
func (o *ImagesGetImageDetailsByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get image details by Id not found response has a 5xx status code
func (o *ImagesGetImageDetailsByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images get image details by Id not found response a status code equal to that given
func (o *ImagesGetImageDetailsByIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesGetImageDetailsByIDNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdNotFound  %+v", 404, o.Payload)
}

func (o *ImagesGetImageDetailsByIDNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdNotFound  %+v", 404, o.Payload)
}

func (o *ImagesGetImageDetailsByIDNotFound) GetPayload() *ImagesGetImageDetailsByIDNotFoundBody {
	return o.Payload
}

func (o *ImagesGetImageDetailsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagesGetImageDetailsByIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetImageDetailsByIDInternalServerError creates a ImagesGetImageDetailsByIDInternalServerError with default headers values
func NewImagesGetImageDetailsByIDInternalServerError() *ImagesGetImageDetailsByIDInternalServerError {
	return &ImagesGetImageDetailsByIDInternalServerError{}
}

/*
ImagesGetImageDetailsByIDInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesGetImageDetailsByIDInternalServerError struct {
}

// IsSuccess returns true when this images get image details by Id internal server error response has a 2xx status code
func (o *ImagesGetImageDetailsByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get image details by Id internal server error response has a 3xx status code
func (o *ImagesGetImageDetailsByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get image details by Id internal server error response has a 4xx status code
func (o *ImagesGetImageDetailsByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images get image details by Id internal server error response has a 5xx status code
func (o *ImagesGetImageDetailsByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images get image details by Id internal server error response a status code equal to that given
func (o *ImagesGetImageDetailsByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesGetImageDetailsByIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdInternalServerError ", 500)
}

func (o *ImagesGetImageDetailsByIDInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/details][%d] imagesGetImageDetailsByIdInternalServerError ", 500)
}

func (o *ImagesGetImageDetailsByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ImagesGetImageDetailsByIDBadRequestBodyItems0 images get image details by ID bad request body items0
swagger:model ImagesGetImageDetailsByIDBadRequestBodyItems0
*/
type ImagesGetImageDetailsByIDBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this images get image details by ID bad request body items0
func (o *ImagesGetImageDetailsByIDBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images get image details by ID bad request body items0 based on context it is used
func (o *ImagesGetImageDetailsByIDBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res ImagesGetImageDetailsByIDBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesGetImageDetailsByIDBody images get image details by ID body
swagger:model ImagesGetImageDetailsByIDBody
*/
type ImagesGetImageDetailsByIDBody struct {

	// cloud Id
	CloudID int32 `json:"cloudId,omitempty"`

	// image Id
	ImageID string `json:"imageId,omitempty"`
}

// Validate validates this images get image details by ID body
func (o *ImagesGetImageDetailsByIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images get image details by ID body based on context it is used
func (o *ImagesGetImageDetailsByIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDBody) UnmarshalBinary(b []byte) error {
	var res ImagesGetImageDetailsByIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesGetImageDetailsByIDForbiddenBody images get image details by ID forbidden body
swagger:model ImagesGetImageDetailsByIDForbiddenBody
*/
type ImagesGetImageDetailsByIDForbiddenBody struct {

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

// Validate validates this images get image details by ID forbidden body
func (o *ImagesGetImageDetailsByIDForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images get image details by ID forbidden body based on context it is used
func (o *ImagesGetImageDetailsByIDForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ImagesGetImageDetailsByIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesGetImageDetailsByIDNotFoundBody images get image details by ID not found body
swagger:model ImagesGetImageDetailsByIDNotFoundBody
*/
type ImagesGetImageDetailsByIDNotFoundBody struct {

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

// Validate validates this images get image details by ID not found body
func (o *ImagesGetImageDetailsByIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images get image details by ID not found body based on context it is used
func (o *ImagesGetImageDetailsByIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImagesGetImageDetailsByIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesGetImageDetailsByIDUnauthorizedBody images get image details by ID unauthorized body
swagger:model ImagesGetImageDetailsByIDUnauthorizedBody
*/
type ImagesGetImageDetailsByIDUnauthorizedBody struct {

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

// Validate validates this images get image details by ID unauthorized body
func (o *ImagesGetImageDetailsByIDUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images get image details by ID unauthorized body based on context it is used
func (o *ImagesGetImageDetailsByIDUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesGetImageDetailsByIDUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ImagesGetImageDetailsByIDUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
