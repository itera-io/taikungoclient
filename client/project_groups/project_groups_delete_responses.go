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

// ProjectGroupsDeleteReader is a Reader for the ProjectGroupsDelete structure.
type ProjectGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGroupsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectGroupsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectGroupsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectGroupsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectGroupsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectGroupsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectGroupsDeleteOK creates a ProjectGroupsDeleteOK with default headers values
func NewProjectGroupsDeleteOK() *ProjectGroupsDeleteOK {
	return &ProjectGroupsDeleteOK{}
}

/*
ProjectGroupsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type ProjectGroupsDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this project groups delete o k response has a 2xx status code
func (o *ProjectGroupsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project groups delete o k response has a 3xx status code
func (o *ProjectGroupsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups delete o k response has a 4xx status code
func (o *ProjectGroupsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups delete o k response has a 5xx status code
func (o *ProjectGroupsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups delete o k response a status code equal to that given
func (o *ProjectGroupsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project groups delete o k response
func (o *ProjectGroupsDeleteOK) Code() int {
	return 200
}

func (o *ProjectGroupsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectGroupsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsDeleteNoContent creates a ProjectGroupsDeleteNoContent with default headers values
func NewProjectGroupsDeleteNoContent() *ProjectGroupsDeleteNoContent {
	return &ProjectGroupsDeleteNoContent{}
}

/*
ProjectGroupsDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type ProjectGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this project groups delete no content response has a 2xx status code
func (o *ProjectGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project groups delete no content response has a 3xx status code
func (o *ProjectGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups delete no content response has a 4xx status code
func (o *ProjectGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups delete no content response has a 5xx status code
func (o *ProjectGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups delete no content response a status code equal to that given
func (o *ProjectGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the project groups delete no content response
func (o *ProjectGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *ProjectGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteNoContent ", 204)
}

func (o *ProjectGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteNoContent ", 204)
}

func (o *ProjectGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectGroupsDeleteBadRequest creates a ProjectGroupsDeleteBadRequest with default headers values
func NewProjectGroupsDeleteBadRequest() *ProjectGroupsDeleteBadRequest {
	return &ProjectGroupsDeleteBadRequest{}
}

/*
ProjectGroupsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectGroupsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this project groups delete bad request response has a 2xx status code
func (o *ProjectGroupsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups delete bad request response has a 3xx status code
func (o *ProjectGroupsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups delete bad request response has a 4xx status code
func (o *ProjectGroupsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups delete bad request response has a 5xx status code
func (o *ProjectGroupsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups delete bad request response a status code equal to that given
func (o *ProjectGroupsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the project groups delete bad request response
func (o *ProjectGroupsDeleteBadRequest) Code() int {
	return 400
}

func (o *ProjectGroupsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectGroupsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsDeleteUnauthorized creates a ProjectGroupsDeleteUnauthorized with default headers values
func NewProjectGroupsDeleteUnauthorized() *ProjectGroupsDeleteUnauthorized {
	return &ProjectGroupsDeleteUnauthorized{}
}

/*
ProjectGroupsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectGroupsDeleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this project groups delete unauthorized response has a 2xx status code
func (o *ProjectGroupsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups delete unauthorized response has a 3xx status code
func (o *ProjectGroupsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups delete unauthorized response has a 4xx status code
func (o *ProjectGroupsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups delete unauthorized response has a 5xx status code
func (o *ProjectGroupsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups delete unauthorized response a status code equal to that given
func (o *ProjectGroupsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the project groups delete unauthorized response
func (o *ProjectGroupsDeleteUnauthorized) Code() int {
	return 401
}

func (o *ProjectGroupsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsDeleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectGroupsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsDeleteForbidden creates a ProjectGroupsDeleteForbidden with default headers values
func NewProjectGroupsDeleteForbidden() *ProjectGroupsDeleteForbidden {
	return &ProjectGroupsDeleteForbidden{}
}

/*
ProjectGroupsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectGroupsDeleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this project groups delete forbidden response has a 2xx status code
func (o *ProjectGroupsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups delete forbidden response has a 3xx status code
func (o *ProjectGroupsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups delete forbidden response has a 4xx status code
func (o *ProjectGroupsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups delete forbidden response has a 5xx status code
func (o *ProjectGroupsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups delete forbidden response a status code equal to that given
func (o *ProjectGroupsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the project groups delete forbidden response
func (o *ProjectGroupsDeleteForbidden) Code() int {
	return 403
}

func (o *ProjectGroupsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsDeleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectGroupsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsDeleteNotFound creates a ProjectGroupsDeleteNotFound with default headers values
func NewProjectGroupsDeleteNotFound() *ProjectGroupsDeleteNotFound {
	return &ProjectGroupsDeleteNotFound{}
}

/*
ProjectGroupsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectGroupsDeleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this project groups delete not found response has a 2xx status code
func (o *ProjectGroupsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups delete not found response has a 3xx status code
func (o *ProjectGroupsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups delete not found response has a 4xx status code
func (o *ProjectGroupsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups delete not found response has a 5xx status code
func (o *ProjectGroupsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups delete not found response a status code equal to that given
func (o *ProjectGroupsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the project groups delete not found response
func (o *ProjectGroupsDeleteNotFound) Code() int {
	return 404
}

func (o *ProjectGroupsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectGroupsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsDeleteInternalServerError creates a ProjectGroupsDeleteInternalServerError with default headers values
func NewProjectGroupsDeleteInternalServerError() *ProjectGroupsDeleteInternalServerError {
	return &ProjectGroupsDeleteInternalServerError{}
}

/*
ProjectGroupsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectGroupsDeleteInternalServerError struct {
}

// IsSuccess returns true when this project groups delete internal server error response has a 2xx status code
func (o *ProjectGroupsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups delete internal server error response has a 3xx status code
func (o *ProjectGroupsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups delete internal server error response has a 4xx status code
func (o *ProjectGroupsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups delete internal server error response has a 5xx status code
func (o *ProjectGroupsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project groups delete internal server error response a status code equal to that given
func (o *ProjectGroupsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the project groups delete internal server error response
func (o *ProjectGroupsDeleteInternalServerError) Code() int {
	return 500
}

func (o *ProjectGroupsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteInternalServerError ", 500)
}

func (o *ProjectGroupsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectGroups][%d] projectGroupsDeleteInternalServerError ", 500)
}

func (o *ProjectGroupsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
