// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// ProjectGroupsListByUserGroupIDReader is a Reader for the ProjectGroupsListByUserGroupID structure.
type ProjectGroupsListByUserGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGroupsListByUserGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGroupsListByUserGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectGroupsListByUserGroupIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectGroupsListByUserGroupIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectGroupsListByUserGroupIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectGroupsListByUserGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectGroupsListByUserGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectGroupsListByUserGroupIDOK creates a ProjectGroupsListByUserGroupIDOK with default headers values
func NewProjectGroupsListByUserGroupIDOK() *ProjectGroupsListByUserGroupIDOK {
	return &ProjectGroupsListByUserGroupIDOK{}
}

/*
ProjectGroupsListByUserGroupIDOK describes a response with status code 200, with default header values.

Success
*/
type ProjectGroupsListByUserGroupIDOK struct {
	Payload []*ProjectGroupsListByUserGroupIDOKBodyItems0
}

// IsSuccess returns true when this project groups list by user group Id o k response has a 2xx status code
func (o *ProjectGroupsListByUserGroupIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project groups list by user group Id o k response has a 3xx status code
func (o *ProjectGroupsListByUserGroupIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list by user group Id o k response has a 4xx status code
func (o *ProjectGroupsListByUserGroupIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups list by user group Id o k response has a 5xx status code
func (o *ProjectGroupsListByUserGroupIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list by user group Id o k response a status code equal to that given
func (o *ProjectGroupsListByUserGroupIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectGroupsListByUserGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDOK) GetPayload() []*ProjectGroupsListByUserGroupIDOKBodyItems0 {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListByUserGroupIDBadRequest creates a ProjectGroupsListByUserGroupIDBadRequest with default headers values
func NewProjectGroupsListByUserGroupIDBadRequest() *ProjectGroupsListByUserGroupIDBadRequest {
	return &ProjectGroupsListByUserGroupIDBadRequest{}
}

/*
ProjectGroupsListByUserGroupIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectGroupsListByUserGroupIDBadRequest struct {
	Payload []*ProjectGroupsListByUserGroupIDBadRequestBodyItems0
}

// IsSuccess returns true when this project groups list by user group Id bad request response has a 2xx status code
func (o *ProjectGroupsListByUserGroupIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list by user group Id bad request response has a 3xx status code
func (o *ProjectGroupsListByUserGroupIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list by user group Id bad request response has a 4xx status code
func (o *ProjectGroupsListByUserGroupIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list by user group Id bad request response has a 5xx status code
func (o *ProjectGroupsListByUserGroupIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list by user group Id bad request response a status code equal to that given
func (o *ProjectGroupsListByUserGroupIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectGroupsListByUserGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDBadRequest) GetPayload() []*ProjectGroupsListByUserGroupIDBadRequestBodyItems0 {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListByUserGroupIDUnauthorized creates a ProjectGroupsListByUserGroupIDUnauthorized with default headers values
func NewProjectGroupsListByUserGroupIDUnauthorized() *ProjectGroupsListByUserGroupIDUnauthorized {
	return &ProjectGroupsListByUserGroupIDUnauthorized{}
}

/*
ProjectGroupsListByUserGroupIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectGroupsListByUserGroupIDUnauthorized struct {
	Payload *ProjectGroupsListByUserGroupIDUnauthorizedBody
}

// IsSuccess returns true when this project groups list by user group Id unauthorized response has a 2xx status code
func (o *ProjectGroupsListByUserGroupIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list by user group Id unauthorized response has a 3xx status code
func (o *ProjectGroupsListByUserGroupIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list by user group Id unauthorized response has a 4xx status code
func (o *ProjectGroupsListByUserGroupIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list by user group Id unauthorized response has a 5xx status code
func (o *ProjectGroupsListByUserGroupIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list by user group Id unauthorized response a status code equal to that given
func (o *ProjectGroupsListByUserGroupIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) GetPayload() *ProjectGroupsListByUserGroupIDUnauthorizedBody {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectGroupsListByUserGroupIDUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListByUserGroupIDForbidden creates a ProjectGroupsListByUserGroupIDForbidden with default headers values
func NewProjectGroupsListByUserGroupIDForbidden() *ProjectGroupsListByUserGroupIDForbidden {
	return &ProjectGroupsListByUserGroupIDForbidden{}
}

/*
ProjectGroupsListByUserGroupIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectGroupsListByUserGroupIDForbidden struct {
	Payload *ProjectGroupsListByUserGroupIDForbiddenBody
}

// IsSuccess returns true when this project groups list by user group Id forbidden response has a 2xx status code
func (o *ProjectGroupsListByUserGroupIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list by user group Id forbidden response has a 3xx status code
func (o *ProjectGroupsListByUserGroupIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list by user group Id forbidden response has a 4xx status code
func (o *ProjectGroupsListByUserGroupIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list by user group Id forbidden response has a 5xx status code
func (o *ProjectGroupsListByUserGroupIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list by user group Id forbidden response a status code equal to that given
func (o *ProjectGroupsListByUserGroupIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectGroupsListByUserGroupIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDForbidden) GetPayload() *ProjectGroupsListByUserGroupIDForbiddenBody {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectGroupsListByUserGroupIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListByUserGroupIDNotFound creates a ProjectGroupsListByUserGroupIDNotFound with default headers values
func NewProjectGroupsListByUserGroupIDNotFound() *ProjectGroupsListByUserGroupIDNotFound {
	return &ProjectGroupsListByUserGroupIDNotFound{}
}

/*
ProjectGroupsListByUserGroupIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectGroupsListByUserGroupIDNotFound struct {
	Payload *ProjectGroupsListByUserGroupIDNotFoundBody
}

// IsSuccess returns true when this project groups list by user group Id not found response has a 2xx status code
func (o *ProjectGroupsListByUserGroupIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list by user group Id not found response has a 3xx status code
func (o *ProjectGroupsListByUserGroupIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list by user group Id not found response has a 4xx status code
func (o *ProjectGroupsListByUserGroupIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list by user group Id not found response has a 5xx status code
func (o *ProjectGroupsListByUserGroupIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list by user group Id not found response a status code equal to that given
func (o *ProjectGroupsListByUserGroupIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectGroupsListByUserGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDNotFound) GetPayload() *ProjectGroupsListByUserGroupIDNotFoundBody {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectGroupsListByUserGroupIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListByUserGroupIDInternalServerError creates a ProjectGroupsListByUserGroupIDInternalServerError with default headers values
func NewProjectGroupsListByUserGroupIDInternalServerError() *ProjectGroupsListByUserGroupIDInternalServerError {
	return &ProjectGroupsListByUserGroupIDInternalServerError{}
}

/*
ProjectGroupsListByUserGroupIDInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectGroupsListByUserGroupIDInternalServerError struct {
}

// IsSuccess returns true when this project groups list by user group Id internal server error response has a 2xx status code
func (o *ProjectGroupsListByUserGroupIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list by user group Id internal server error response has a 3xx status code
func (o *ProjectGroupsListByUserGroupIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list by user group Id internal server error response has a 4xx status code
func (o *ProjectGroupsListByUserGroupIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups list by user group Id internal server error response has a 5xx status code
func (o *ProjectGroupsListByUserGroupIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project groups list by user group Id internal server error response a status code equal to that given
func (o *ProjectGroupsListByUserGroupIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectGroupsListByUserGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdInternalServerError ", 500)
}

func (o *ProjectGroupsListByUserGroupIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdInternalServerError ", 500)
}

func (o *ProjectGroupsListByUserGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ProjectGroupsListByUserGroupIDBadRequestBodyItems0 project groups list by user group ID bad request body items0
swagger:model ProjectGroupsListByUserGroupIDBadRequestBodyItems0
*/
type ProjectGroupsListByUserGroupIDBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this project groups list by user group ID bad request body items0
func (o *ProjectGroupsListByUserGroupIDBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project groups list by user group ID bad request body items0 based on context it is used
func (o *ProjectGroupsListByUserGroupIDBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res ProjectGroupsListByUserGroupIDBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectGroupsListByUserGroupIDForbiddenBody project groups list by user group ID forbidden body
swagger:model ProjectGroupsListByUserGroupIDForbiddenBody
*/
type ProjectGroupsListByUserGroupIDForbiddenBody struct {

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

// Validate validates this project groups list by user group ID forbidden body
func (o *ProjectGroupsListByUserGroupIDForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project groups list by user group ID forbidden body based on context it is used
func (o *ProjectGroupsListByUserGroupIDForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ProjectGroupsListByUserGroupIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectGroupsListByUserGroupIDNotFoundBody project groups list by user group ID not found body
swagger:model ProjectGroupsListByUserGroupIDNotFoundBody
*/
type ProjectGroupsListByUserGroupIDNotFoundBody struct {

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

// Validate validates this project groups list by user group ID not found body
func (o *ProjectGroupsListByUserGroupIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project groups list by user group ID not found body based on context it is used
func (o *ProjectGroupsListByUserGroupIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ProjectGroupsListByUserGroupIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectGroupsListByUserGroupIDOKBodyItems0 project groups list by user group ID o k body items0
swagger:model ProjectGroupsListByUserGroupIDOKBodyItems0
*/
type ProjectGroupsListByUserGroupIDOKBodyItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this project groups list by user group ID o k body items0
func (o *ProjectGroupsListByUserGroupIDOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project groups list by user group ID o k body items0 based on context it is used
func (o *ProjectGroupsListByUserGroupIDOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ProjectGroupsListByUserGroupIDOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectGroupsListByUserGroupIDUnauthorizedBody project groups list by user group ID unauthorized body
swagger:model ProjectGroupsListByUserGroupIDUnauthorizedBody
*/
type ProjectGroupsListByUserGroupIDUnauthorizedBody struct {

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

// Validate validates this project groups list by user group ID unauthorized body
func (o *ProjectGroupsListByUserGroupIDUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project groups list by user group ID unauthorized body based on context it is used
func (o *ProjectGroupsListByUserGroupIDUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectGroupsListByUserGroupIDUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ProjectGroupsListByUserGroupIDUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
