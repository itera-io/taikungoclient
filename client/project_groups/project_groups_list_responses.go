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

// ProjectGroupsListReader is a Reader for the ProjectGroupsList structure.
type ProjectGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectGroupsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectGroupsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectGroupsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectGroupsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectGroupsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectGroupsListOK creates a ProjectGroupsListOK with default headers values
func NewProjectGroupsListOK() *ProjectGroupsListOK {
	return &ProjectGroupsListOK{}
}

/*
ProjectGroupsListOK describes a response with status code 200, with default header values.

Success
*/
type ProjectGroupsListOK struct {
	Payload *models.ProjectGroupList
}

// IsSuccess returns true when this project groups list o k response has a 2xx status code
func (o *ProjectGroupsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project groups list o k response has a 3xx status code
func (o *ProjectGroupsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list o k response has a 4xx status code
func (o *ProjectGroupsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups list o k response has a 5xx status code
func (o *ProjectGroupsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list o k response a status code equal to that given
func (o *ProjectGroupsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project groups list o k response
func (o *ProjectGroupsListOK) Code() int {
	return 200
}

func (o *ProjectGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsListOK) GetPayload() *models.ProjectGroupList {
	return o.Payload
}

func (o *ProjectGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectGroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListBadRequest creates a ProjectGroupsListBadRequest with default headers values
func NewProjectGroupsListBadRequest() *ProjectGroupsListBadRequest {
	return &ProjectGroupsListBadRequest{}
}

/*
ProjectGroupsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectGroupsListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups list bad request response has a 2xx status code
func (o *ProjectGroupsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list bad request response has a 3xx status code
func (o *ProjectGroupsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list bad request response has a 4xx status code
func (o *ProjectGroupsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list bad request response has a 5xx status code
func (o *ProjectGroupsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list bad request response a status code equal to that given
func (o *ProjectGroupsListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the project groups list bad request response
func (o *ProjectGroupsListBadRequest) Code() int {
	return 400
}

func (o *ProjectGroupsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListUnauthorized creates a ProjectGroupsListUnauthorized with default headers values
func NewProjectGroupsListUnauthorized() *ProjectGroupsListUnauthorized {
	return &ProjectGroupsListUnauthorized{}
}

/*
ProjectGroupsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectGroupsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups list unauthorized response has a 2xx status code
func (o *ProjectGroupsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list unauthorized response has a 3xx status code
func (o *ProjectGroupsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list unauthorized response has a 4xx status code
func (o *ProjectGroupsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list unauthorized response has a 5xx status code
func (o *ProjectGroupsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list unauthorized response a status code equal to that given
func (o *ProjectGroupsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the project groups list unauthorized response
func (o *ProjectGroupsListUnauthorized) Code() int {
	return 401
}

func (o *ProjectGroupsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListForbidden creates a ProjectGroupsListForbidden with default headers values
func NewProjectGroupsListForbidden() *ProjectGroupsListForbidden {
	return &ProjectGroupsListForbidden{}
}

/*
ProjectGroupsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectGroupsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups list forbidden response has a 2xx status code
func (o *ProjectGroupsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list forbidden response has a 3xx status code
func (o *ProjectGroupsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list forbidden response has a 4xx status code
func (o *ProjectGroupsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list forbidden response has a 5xx status code
func (o *ProjectGroupsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list forbidden response a status code equal to that given
func (o *ProjectGroupsListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the project groups list forbidden response
func (o *ProjectGroupsListForbidden) Code() int {
	return 403
}

func (o *ProjectGroupsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListNotFound creates a ProjectGroupsListNotFound with default headers values
func NewProjectGroupsListNotFound() *ProjectGroupsListNotFound {
	return &ProjectGroupsListNotFound{}
}

/*
ProjectGroupsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectGroupsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups list not found response has a 2xx status code
func (o *ProjectGroupsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list not found response has a 3xx status code
func (o *ProjectGroupsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list not found response has a 4xx status code
func (o *ProjectGroupsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups list not found response has a 5xx status code
func (o *ProjectGroupsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups list not found response a status code equal to that given
func (o *ProjectGroupsListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the project groups list not found response
func (o *ProjectGroupsListNotFound) Code() int {
	return 404
}

func (o *ProjectGroupsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsListInternalServerError creates a ProjectGroupsListInternalServerError with default headers values
func NewProjectGroupsListInternalServerError() *ProjectGroupsListInternalServerError {
	return &ProjectGroupsListInternalServerError{}
}

/*
ProjectGroupsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectGroupsListInternalServerError struct {
}

// IsSuccess returns true when this project groups list internal server error response has a 2xx status code
func (o *ProjectGroupsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups list internal server error response has a 3xx status code
func (o *ProjectGroupsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups list internal server error response has a 4xx status code
func (o *ProjectGroupsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups list internal server error response has a 5xx status code
func (o *ProjectGroupsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project groups list internal server error response a status code equal to that given
func (o *ProjectGroupsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the project groups list internal server error response
func (o *ProjectGroupsListInternalServerError) Code() int {
	return 500
}

func (o *ProjectGroupsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListInternalServerError ", 500)
}

func (o *ProjectGroupsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectGroups/list][%d] projectGroupsListInternalServerError ", 500)
}

func (o *ProjectGroupsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
