// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImagesCommonAzureImagesReader is a Reader for the ImagesCommonAzureImages structure.
type ImagesCommonAzureImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesCommonAzureImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesCommonAzureImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesCommonAzureImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesCommonAzureImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesCommonAzureImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesCommonAzureImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesCommonAzureImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesCommonAzureImagesOK creates a ImagesCommonAzureImagesOK with default headers values
func NewImagesCommonAzureImagesOK() *ImagesCommonAzureImagesOK {
	return &ImagesCommonAzureImagesOK{}
}

/*
ImagesCommonAzureImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesCommonAzureImagesOK struct {
	Payload []*ImagesCommonAzureImagesOKBodyItems0
}

// IsSuccess returns true when this images common azure images o k response has a 2xx status code
func (o *ImagesCommonAzureImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images common azure images o k response has a 3xx status code
func (o *ImagesCommonAzureImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images o k response has a 4xx status code
func (o *ImagesCommonAzureImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common azure images o k response has a 5xx status code
func (o *ImagesCommonAzureImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images o k response a status code equal to that given
func (o *ImagesCommonAzureImagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesCommonAzureImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesCommonAzureImagesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesCommonAzureImagesOK) GetPayload() []*ImagesCommonAzureImagesOKBodyItems0 {
	return o.Payload
}

func (o *ImagesCommonAzureImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesBadRequest creates a ImagesCommonAzureImagesBadRequest with default headers values
func NewImagesCommonAzureImagesBadRequest() *ImagesCommonAzureImagesBadRequest {
	return &ImagesCommonAzureImagesBadRequest{}
}

/*
ImagesCommonAzureImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesCommonAzureImagesBadRequest struct {
	Payload []*ImagesCommonAzureImagesBadRequestBodyItems0
}

// IsSuccess returns true when this images common azure images bad request response has a 2xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images bad request response has a 3xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images bad request response has a 4xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images bad request response has a 5xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images bad request response a status code equal to that given
func (o *ImagesCommonAzureImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesCommonAzureImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonAzureImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonAzureImagesBadRequest) GetPayload() []*ImagesCommonAzureImagesBadRequestBodyItems0 {
	return o.Payload
}

func (o *ImagesCommonAzureImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesUnauthorized creates a ImagesCommonAzureImagesUnauthorized with default headers values
func NewImagesCommonAzureImagesUnauthorized() *ImagesCommonAzureImagesUnauthorized {
	return &ImagesCommonAzureImagesUnauthorized{}
}

/*
ImagesCommonAzureImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesCommonAzureImagesUnauthorized struct {
	Payload *ImagesCommonAzureImagesUnauthorizedBody
}

// IsSuccess returns true when this images common azure images unauthorized response has a 2xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images unauthorized response has a 3xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images unauthorized response has a 4xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images unauthorized response has a 5xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images unauthorized response a status code equal to that given
func (o *ImagesCommonAzureImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesCommonAzureImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonAzureImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonAzureImagesUnauthorized) GetPayload() *ImagesCommonAzureImagesUnauthorizedBody {
	return o.Payload
}

func (o *ImagesCommonAzureImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagesCommonAzureImagesUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesForbidden creates a ImagesCommonAzureImagesForbidden with default headers values
func NewImagesCommonAzureImagesForbidden() *ImagesCommonAzureImagesForbidden {
	return &ImagesCommonAzureImagesForbidden{}
}

/*
ImagesCommonAzureImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesCommonAzureImagesForbidden struct {
	Payload *ImagesCommonAzureImagesForbiddenBody
}

// IsSuccess returns true when this images common azure images forbidden response has a 2xx status code
func (o *ImagesCommonAzureImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images forbidden response has a 3xx status code
func (o *ImagesCommonAzureImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images forbidden response has a 4xx status code
func (o *ImagesCommonAzureImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images forbidden response has a 5xx status code
func (o *ImagesCommonAzureImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images forbidden response a status code equal to that given
func (o *ImagesCommonAzureImagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesCommonAzureImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonAzureImagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonAzureImagesForbidden) GetPayload() *ImagesCommonAzureImagesForbiddenBody {
	return o.Payload
}

func (o *ImagesCommonAzureImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagesCommonAzureImagesForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesNotFound creates a ImagesCommonAzureImagesNotFound with default headers values
func NewImagesCommonAzureImagesNotFound() *ImagesCommonAzureImagesNotFound {
	return &ImagesCommonAzureImagesNotFound{}
}

/*
ImagesCommonAzureImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesCommonAzureImagesNotFound struct {
	Payload *ImagesCommonAzureImagesNotFoundBody
}

// IsSuccess returns true when this images common azure images not found response has a 2xx status code
func (o *ImagesCommonAzureImagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images not found response has a 3xx status code
func (o *ImagesCommonAzureImagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images not found response has a 4xx status code
func (o *ImagesCommonAzureImagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images not found response has a 5xx status code
func (o *ImagesCommonAzureImagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images not found response a status code equal to that given
func (o *ImagesCommonAzureImagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesCommonAzureImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonAzureImagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonAzureImagesNotFound) GetPayload() *ImagesCommonAzureImagesNotFoundBody {
	return o.Payload
}

func (o *ImagesCommonAzureImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagesCommonAzureImagesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesInternalServerError creates a ImagesCommonAzureImagesInternalServerError with default headers values
func NewImagesCommonAzureImagesInternalServerError() *ImagesCommonAzureImagesInternalServerError {
	return &ImagesCommonAzureImagesInternalServerError{}
}

/*
ImagesCommonAzureImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesCommonAzureImagesInternalServerError struct {
}

// IsSuccess returns true when this images common azure images internal server error response has a 2xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images internal server error response has a 3xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images internal server error response has a 4xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common azure images internal server error response has a 5xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images common azure images internal server error response a status code equal to that given
func (o *ImagesCommonAzureImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesCommonAzureImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesInternalServerError ", 500)
}

func (o *ImagesCommonAzureImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesInternalServerError ", 500)
}

func (o *ImagesCommonAzureImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ImagesCommonAzureImagesBadRequestBodyItems0 images common azure images bad request body items0
swagger:model ImagesCommonAzureImagesBadRequestBodyItems0
*/
type ImagesCommonAzureImagesBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this images common azure images bad request body items0
func (o *ImagesCommonAzureImagesBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images common azure images bad request body items0 based on context it is used
func (o *ImagesCommonAzureImagesBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesCommonAzureImagesBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesCommonAzureImagesBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res ImagesCommonAzureImagesBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesCommonAzureImagesForbiddenBody images common azure images forbidden body
swagger:model ImagesCommonAzureImagesForbiddenBody
*/
type ImagesCommonAzureImagesForbiddenBody struct {

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

// Validate validates this images common azure images forbidden body
func (o *ImagesCommonAzureImagesForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images common azure images forbidden body based on context it is used
func (o *ImagesCommonAzureImagesForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesCommonAzureImagesForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesCommonAzureImagesForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ImagesCommonAzureImagesForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesCommonAzureImagesNotFoundBody images common azure images not found body
swagger:model ImagesCommonAzureImagesNotFoundBody
*/
type ImagesCommonAzureImagesNotFoundBody struct {

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

// Validate validates this images common azure images not found body
func (o *ImagesCommonAzureImagesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images common azure images not found body based on context it is used
func (o *ImagesCommonAzureImagesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesCommonAzureImagesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesCommonAzureImagesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImagesCommonAzureImagesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesCommonAzureImagesOKBodyItems0 images common azure images o k body items0
swagger:model ImagesCommonAzureImagesOKBodyItems0
*/
type ImagesCommonAzureImagesOKBodyItems0 struct {

	// image
	Image *ImagesCommonAzureImagesOKBodyItems0Image `json:"image,omitempty"`

	// publisher
	Publisher string `json:"publisher,omitempty"`
}

// Validate validates this images common azure images o k body items0
func (o *ImagesCommonAzureImagesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImagesCommonAzureImagesOKBodyItems0) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(o.Image) { // not required
		return nil
	}

	if o.Image != nil {
		if err := o.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this images common azure images o k body items0 based on the context it is used
func (o *ImagesCommonAzureImagesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImagesCommonAzureImagesOKBodyItems0) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if o.Image != nil {
		if err := o.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ImagesCommonAzureImagesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesCommonAzureImagesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ImagesCommonAzureImagesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesCommonAzureImagesOKBodyItems0Image images common azure images o k body items0 image
swagger:model ImagesCommonAzureImagesOKBodyItems0Image
*/
type ImagesCommonAzureImagesOKBodyItems0Image struct {

	// display name
	DisplayName string `json:"displayName"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this images common azure images o k body items0 image
func (o *ImagesCommonAzureImagesOKBodyItems0Image) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images common azure images o k body items0 image based on context it is used
func (o *ImagesCommonAzureImagesOKBodyItems0Image) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesCommonAzureImagesOKBodyItems0Image) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesCommonAzureImagesOKBodyItems0Image) UnmarshalBinary(b []byte) error {
	var res ImagesCommonAzureImagesOKBodyItems0Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagesCommonAzureImagesUnauthorizedBody images common azure images unauthorized body
swagger:model ImagesCommonAzureImagesUnauthorizedBody
*/
type ImagesCommonAzureImagesUnauthorizedBody struct {

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

// Validate validates this images common azure images unauthorized body
func (o *ImagesCommonAzureImagesUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this images common azure images unauthorized body based on context it is used
func (o *ImagesCommonAzureImagesUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagesCommonAzureImagesUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagesCommonAzureImagesUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ImagesCommonAzureImagesUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
