// Code generated by go-swagger; DO NOT EDIT.

package project_app

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

// ProjectAppUninstallReader is a Reader for the ProjectAppUninstall structure.
type ProjectAppUninstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppUninstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppUninstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppUninstallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppUninstallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppUninstallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppUninstallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppUninstallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppUninstallOK creates a ProjectAppUninstallOK with default headers values
func NewProjectAppUninstallOK() *ProjectAppUninstallOK {
	return &ProjectAppUninstallOK{}
}

/*
ProjectAppUninstallOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppUninstallOK struct {
	Payload interface{}
}

// IsSuccess returns true when this project app uninstall o k response has a 2xx status code
func (o *ProjectAppUninstallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project app uninstall o k response has a 3xx status code
func (o *ProjectAppUninstallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app uninstall o k response has a 4xx status code
func (o *ProjectAppUninstallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app uninstall o k response has a 5xx status code
func (o *ProjectAppUninstallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project app uninstall o k response a status code equal to that given
func (o *ProjectAppUninstallOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectAppUninstallOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallOK  %+v", 200, o.Payload)
}

func (o *ProjectAppUninstallOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallOK  %+v", 200, o.Payload)
}

func (o *ProjectAppUninstallOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectAppUninstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallBadRequest creates a ProjectAppUninstallBadRequest with default headers values
func NewProjectAppUninstallBadRequest() *ProjectAppUninstallBadRequest {
	return &ProjectAppUninstallBadRequest{}
}

/*
ProjectAppUninstallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppUninstallBadRequest struct {
	Payload []*ProjectAppUninstallBadRequestBodyItems0
}

// IsSuccess returns true when this project app uninstall bad request response has a 2xx status code
func (o *ProjectAppUninstallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app uninstall bad request response has a 3xx status code
func (o *ProjectAppUninstallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app uninstall bad request response has a 4xx status code
func (o *ProjectAppUninstallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app uninstall bad request response has a 5xx status code
func (o *ProjectAppUninstallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project app uninstall bad request response a status code equal to that given
func (o *ProjectAppUninstallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectAppUninstallBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppUninstallBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppUninstallBadRequest) GetPayload() []*ProjectAppUninstallBadRequestBodyItems0 {
	return o.Payload
}

func (o *ProjectAppUninstallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallUnauthorized creates a ProjectAppUninstallUnauthorized with default headers values
func NewProjectAppUninstallUnauthorized() *ProjectAppUninstallUnauthorized {
	return &ProjectAppUninstallUnauthorized{}
}

/*
ProjectAppUninstallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppUninstallUnauthorized struct {
	Payload *ProjectAppUninstallUnauthorizedBody
}

// IsSuccess returns true when this project app uninstall unauthorized response has a 2xx status code
func (o *ProjectAppUninstallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app uninstall unauthorized response has a 3xx status code
func (o *ProjectAppUninstallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app uninstall unauthorized response has a 4xx status code
func (o *ProjectAppUninstallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app uninstall unauthorized response has a 5xx status code
func (o *ProjectAppUninstallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project app uninstall unauthorized response a status code equal to that given
func (o *ProjectAppUninstallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectAppUninstallUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppUninstallUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppUninstallUnauthorized) GetPayload() *ProjectAppUninstallUnauthorizedBody {
	return o.Payload
}

func (o *ProjectAppUninstallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectAppUninstallUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallForbidden creates a ProjectAppUninstallForbidden with default headers values
func NewProjectAppUninstallForbidden() *ProjectAppUninstallForbidden {
	return &ProjectAppUninstallForbidden{}
}

/*
ProjectAppUninstallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppUninstallForbidden struct {
	Payload *ProjectAppUninstallForbiddenBody
}

// IsSuccess returns true when this project app uninstall forbidden response has a 2xx status code
func (o *ProjectAppUninstallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app uninstall forbidden response has a 3xx status code
func (o *ProjectAppUninstallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app uninstall forbidden response has a 4xx status code
func (o *ProjectAppUninstallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app uninstall forbidden response has a 5xx status code
func (o *ProjectAppUninstallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project app uninstall forbidden response a status code equal to that given
func (o *ProjectAppUninstallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectAppUninstallForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppUninstallForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppUninstallForbidden) GetPayload() *ProjectAppUninstallForbiddenBody {
	return o.Payload
}

func (o *ProjectAppUninstallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectAppUninstallForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallNotFound creates a ProjectAppUninstallNotFound with default headers values
func NewProjectAppUninstallNotFound() *ProjectAppUninstallNotFound {
	return &ProjectAppUninstallNotFound{}
}

/*
ProjectAppUninstallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppUninstallNotFound struct {
	Payload *ProjectAppUninstallNotFoundBody
}

// IsSuccess returns true when this project app uninstall not found response has a 2xx status code
func (o *ProjectAppUninstallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app uninstall not found response has a 3xx status code
func (o *ProjectAppUninstallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app uninstall not found response has a 4xx status code
func (o *ProjectAppUninstallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app uninstall not found response has a 5xx status code
func (o *ProjectAppUninstallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project app uninstall not found response a status code equal to that given
func (o *ProjectAppUninstallNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectAppUninstallNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppUninstallNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppUninstallNotFound) GetPayload() *ProjectAppUninstallNotFoundBody {
	return o.Payload
}

func (o *ProjectAppUninstallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectAppUninstallNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallInternalServerError creates a ProjectAppUninstallInternalServerError with default headers values
func NewProjectAppUninstallInternalServerError() *ProjectAppUninstallInternalServerError {
	return &ProjectAppUninstallInternalServerError{}
}

/*
ProjectAppUninstallInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppUninstallInternalServerError struct {
}

// IsSuccess returns true when this project app uninstall internal server error response has a 2xx status code
func (o *ProjectAppUninstallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app uninstall internal server error response has a 3xx status code
func (o *ProjectAppUninstallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app uninstall internal server error response has a 4xx status code
func (o *ProjectAppUninstallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app uninstall internal server error response has a 5xx status code
func (o *ProjectAppUninstallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project app uninstall internal server error response a status code equal to that given
func (o *ProjectAppUninstallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectAppUninstallInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallInternalServerError ", 500)
}

func (o *ProjectAppUninstallInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallInternalServerError ", 500)
}

func (o *ProjectAppUninstallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ProjectAppUninstallBadRequestBodyItems0 project app uninstall bad request body items0
swagger:model ProjectAppUninstallBadRequestBodyItems0
*/
type ProjectAppUninstallBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this project app uninstall bad request body items0
func (o *ProjectAppUninstallBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project app uninstall bad request body items0 based on context it is used
func (o *ProjectAppUninstallBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectAppUninstallBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectAppUninstallBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res ProjectAppUninstallBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectAppUninstallForbiddenBody project app uninstall forbidden body
swagger:model ProjectAppUninstallForbiddenBody
*/
type ProjectAppUninstallForbiddenBody struct {

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

// Validate validates this project app uninstall forbidden body
func (o *ProjectAppUninstallForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project app uninstall forbidden body based on context it is used
func (o *ProjectAppUninstallForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectAppUninstallForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectAppUninstallForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ProjectAppUninstallForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectAppUninstallNotFoundBody project app uninstall not found body
swagger:model ProjectAppUninstallNotFoundBody
*/
type ProjectAppUninstallNotFoundBody struct {

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

// Validate validates this project app uninstall not found body
func (o *ProjectAppUninstallNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project app uninstall not found body based on context it is used
func (o *ProjectAppUninstallNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectAppUninstallNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectAppUninstallNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ProjectAppUninstallNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectAppUninstallUnauthorizedBody project app uninstall unauthorized body
swagger:model ProjectAppUninstallUnauthorizedBody
*/
type ProjectAppUninstallUnauthorizedBody struct {

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

// Validate validates this project app uninstall unauthorized body
func (o *ProjectAppUninstallUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project app uninstall unauthorized body based on context it is used
func (o *ProjectAppUninstallUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectAppUninstallUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectAppUninstallUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ProjectAppUninstallUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
