// Code generated by go-swagger; DO NOT EDIT.

package project_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectActionsDeleteReader is a Reader for the ProjectActionsDelete structure.
type ProjectActionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectActionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectActionsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectActionsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectActionsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectActionsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectActionsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectActionsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectActionsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectActionsDeleteOK creates a ProjectActionsDeleteOK with default headers values
func NewProjectActionsDeleteOK() *ProjectActionsDeleteOK {
	return &ProjectActionsDeleteOK{}
}

/*
ProjectActionsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type ProjectActionsDeleteOK struct {
}

// IsSuccess returns true when this project actions delete o k response has a 2xx status code
func (o *ProjectActionsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project actions delete o k response has a 3xx status code
func (o *ProjectActionsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions delete o k response has a 4xx status code
func (o *ProjectActionsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project actions delete o k response has a 5xx status code
func (o *ProjectActionsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions delete o k response a status code equal to that given
func (o *ProjectActionsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project actions delete o k response
func (o *ProjectActionsDeleteOK) Code() int {
	return 200
}

func (o *ProjectActionsDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteOK ", 200)
}

func (o *ProjectActionsDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteOK ", 200)
}

func (o *ProjectActionsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectActionsDeleteNoContent creates a ProjectActionsDeleteNoContent with default headers values
func NewProjectActionsDeleteNoContent() *ProjectActionsDeleteNoContent {
	return &ProjectActionsDeleteNoContent{}
}

/*
ProjectActionsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ProjectActionsDeleteNoContent struct {
}

// IsSuccess returns true when this project actions delete no content response has a 2xx status code
func (o *ProjectActionsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project actions delete no content response has a 3xx status code
func (o *ProjectActionsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions delete no content response has a 4xx status code
func (o *ProjectActionsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this project actions delete no content response has a 5xx status code
func (o *ProjectActionsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions delete no content response a status code equal to that given
func (o *ProjectActionsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the project actions delete no content response
func (o *ProjectActionsDeleteNoContent) Code() int {
	return 204
}

func (o *ProjectActionsDeleteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteNoContent ", 204)
}

func (o *ProjectActionsDeleteNoContent) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteNoContent ", 204)
}

func (o *ProjectActionsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectActionsDeleteBadRequest creates a ProjectActionsDeleteBadRequest with default headers values
func NewProjectActionsDeleteBadRequest() *ProjectActionsDeleteBadRequest {
	return &ProjectActionsDeleteBadRequest{}
}

/*
ProjectActionsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectActionsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this project actions delete bad request response has a 2xx status code
func (o *ProjectActionsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions delete bad request response has a 3xx status code
func (o *ProjectActionsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions delete bad request response has a 4xx status code
func (o *ProjectActionsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions delete bad request response has a 5xx status code
func (o *ProjectActionsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions delete bad request response a status code equal to that given
func (o *ProjectActionsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the project actions delete bad request response
func (o *ProjectActionsDeleteBadRequest) Code() int {
	return 400
}

func (o *ProjectActionsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectActionsDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectActionsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectActionsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsDeleteUnauthorized creates a ProjectActionsDeleteUnauthorized with default headers values
func NewProjectActionsDeleteUnauthorized() *ProjectActionsDeleteUnauthorized {
	return &ProjectActionsDeleteUnauthorized{}
}

/*
ProjectActionsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectActionsDeleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this project actions delete unauthorized response has a 2xx status code
func (o *ProjectActionsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions delete unauthorized response has a 3xx status code
func (o *ProjectActionsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions delete unauthorized response has a 4xx status code
func (o *ProjectActionsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions delete unauthorized response has a 5xx status code
func (o *ProjectActionsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions delete unauthorized response a status code equal to that given
func (o *ProjectActionsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the project actions delete unauthorized response
func (o *ProjectActionsDeleteUnauthorized) Code() int {
	return 401
}

func (o *ProjectActionsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectActionsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectActionsDeleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectActionsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsDeleteForbidden creates a ProjectActionsDeleteForbidden with default headers values
func NewProjectActionsDeleteForbidden() *ProjectActionsDeleteForbidden {
	return &ProjectActionsDeleteForbidden{}
}

/*
ProjectActionsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectActionsDeleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this project actions delete forbidden response has a 2xx status code
func (o *ProjectActionsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions delete forbidden response has a 3xx status code
func (o *ProjectActionsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions delete forbidden response has a 4xx status code
func (o *ProjectActionsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions delete forbidden response has a 5xx status code
func (o *ProjectActionsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions delete forbidden response a status code equal to that given
func (o *ProjectActionsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the project actions delete forbidden response
func (o *ProjectActionsDeleteForbidden) Code() int {
	return 403
}

func (o *ProjectActionsDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectActionsDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectActionsDeleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectActionsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsDeleteNotFound creates a ProjectActionsDeleteNotFound with default headers values
func NewProjectActionsDeleteNotFound() *ProjectActionsDeleteNotFound {
	return &ProjectActionsDeleteNotFound{}
}

/*
ProjectActionsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectActionsDeleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this project actions delete not found response has a 2xx status code
func (o *ProjectActionsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions delete not found response has a 3xx status code
func (o *ProjectActionsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions delete not found response has a 4xx status code
func (o *ProjectActionsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions delete not found response has a 5xx status code
func (o *ProjectActionsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions delete not found response a status code equal to that given
func (o *ProjectActionsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the project actions delete not found response
func (o *ProjectActionsDeleteNotFound) Code() int {
	return 404
}

func (o *ProjectActionsDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectActionsDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectActionsDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectActionsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsDeleteInternalServerError creates a ProjectActionsDeleteInternalServerError with default headers values
func NewProjectActionsDeleteInternalServerError() *ProjectActionsDeleteInternalServerError {
	return &ProjectActionsDeleteInternalServerError{}
}

/*
ProjectActionsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectActionsDeleteInternalServerError struct {
}

// IsSuccess returns true when this project actions delete internal server error response has a 2xx status code
func (o *ProjectActionsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions delete internal server error response has a 3xx status code
func (o *ProjectActionsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions delete internal server error response has a 4xx status code
func (o *ProjectActionsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project actions delete internal server error response has a 5xx status code
func (o *ProjectActionsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project actions delete internal server error response a status code equal to that given
func (o *ProjectActionsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the project actions delete internal server error response
func (o *ProjectActionsDeleteInternalServerError) Code() int {
	return 500
}

func (o *ProjectActionsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteInternalServerError ", 500)
}

func (o *ProjectActionsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/delete][%d] projectActionsDeleteInternalServerError ", 500)
}

func (o *ProjectActionsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
