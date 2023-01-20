// Code generated by go-swagger; DO NOT EDIT.

package project_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
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
	Payload []*models.CommonDropdownDto
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

// Code gets the status code for the project groups list by user group Id o k response
func (o *ProjectGroupsListByUserGroupIDOK) Code() int {
	return 200
}

func (o *ProjectGroupsListByUserGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDOK) GetPayload() []*models.CommonDropdownDto {
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
	Payload interface{}
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

// Code gets the status code for the project groups list by user group Id bad request response
func (o *ProjectGroupsListByUserGroupIDBadRequest) Code() int {
	return 400
}

func (o *ProjectGroupsListByUserGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDBadRequest) GetPayload() interface{} {
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
	Payload interface{}
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

// Code gets the status code for the project groups list by user group Id unauthorized response
func (o *ProjectGroupsListByUserGroupIDUnauthorized) Code() int {
	return 401
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
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
	Payload interface{}
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

// Code gets the status code for the project groups list by user group Id forbidden response
func (o *ProjectGroupsListByUserGroupIDForbidden) Code() int {
	return 403
}

func (o *ProjectGroupsListByUserGroupIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
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
	Payload interface{}
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

// Code gets the status code for the project groups list by user group Id not found response
func (o *ProjectGroupsListByUserGroupIDNotFound) Code() int {
	return 404
}

func (o *ProjectGroupsListByUserGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list-by-user-group-id][%d] projectGroupsListByUserGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsListByUserGroupIDNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectGroupsListByUserGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
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

// Code gets the status code for the project groups list by user group Id internal server error response
func (o *ProjectGroupsListByUserGroupIDInternalServerError) Code() int {
	return 500
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
