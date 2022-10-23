// Code generated by go-swagger; DO NOT EDIT.

package kubespray

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubesprayDeleteReader is a Reader for the KubesprayDelete structure.
type KubesprayDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubesprayDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubesprayDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewKubesprayDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubesprayDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubesprayDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubesprayDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubesprayDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubesprayDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubesprayDeleteOK creates a KubesprayDeleteOK with default headers values
func NewKubesprayDeleteOK() *KubesprayDeleteOK {
	return &KubesprayDeleteOK{}
}

/*
KubesprayDeleteOK describes a response with status code 200, with default header values.

Success
*/
type KubesprayDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubespray delete o k response has a 2xx status code
func (o *KubesprayDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubespray delete o k response has a 3xx status code
func (o *KubesprayDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray delete o k response has a 4xx status code
func (o *KubesprayDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubespray delete o k response has a 5xx status code
func (o *KubesprayDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray delete o k response a status code equal to that given
func (o *KubesprayDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubesprayDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteOK  %+v", 200, o.Payload)
}

func (o *KubesprayDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteOK  %+v", 200, o.Payload)
}

func (o *KubesprayDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubesprayDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayDeleteNoContent creates a KubesprayDeleteNoContent with default headers values
func NewKubesprayDeleteNoContent() *KubesprayDeleteNoContent {
	return &KubesprayDeleteNoContent{}
}

/*
KubesprayDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type KubesprayDeleteNoContent struct {
}

// IsSuccess returns true when this kubespray delete no content response has a 2xx status code
func (o *KubesprayDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubespray delete no content response has a 3xx status code
func (o *KubesprayDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray delete no content response has a 4xx status code
func (o *KubesprayDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubespray delete no content response has a 5xx status code
func (o *KubesprayDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray delete no content response a status code equal to that given
func (o *KubesprayDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *KubesprayDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteNoContent ", 204)
}

func (o *KubesprayDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteNoContent ", 204)
}

func (o *KubesprayDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKubesprayDeleteBadRequest creates a KubesprayDeleteBadRequest with default headers values
func NewKubesprayDeleteBadRequest() *KubesprayDeleteBadRequest {
	return &KubesprayDeleteBadRequest{}
}

/*
KubesprayDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubesprayDeleteBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubespray delete bad request response has a 2xx status code
func (o *KubesprayDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray delete bad request response has a 3xx status code
func (o *KubesprayDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray delete bad request response has a 4xx status code
func (o *KubesprayDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray delete bad request response has a 5xx status code
func (o *KubesprayDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray delete bad request response a status code equal to that given
func (o *KubesprayDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubesprayDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KubesprayDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KubesprayDeleteBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubesprayDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayDeleteUnauthorized creates a KubesprayDeleteUnauthorized with default headers values
func NewKubesprayDeleteUnauthorized() *KubesprayDeleteUnauthorized {
	return &KubesprayDeleteUnauthorized{}
}

/*
KubesprayDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubesprayDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubespray delete unauthorized response has a 2xx status code
func (o *KubesprayDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray delete unauthorized response has a 3xx status code
func (o *KubesprayDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray delete unauthorized response has a 4xx status code
func (o *KubesprayDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray delete unauthorized response has a 5xx status code
func (o *KubesprayDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray delete unauthorized response a status code equal to that given
func (o *KubesprayDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubesprayDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KubesprayDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KubesprayDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubesprayDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayDeleteForbidden creates a KubesprayDeleteForbidden with default headers values
func NewKubesprayDeleteForbidden() *KubesprayDeleteForbidden {
	return &KubesprayDeleteForbidden{}
}

/*
KubesprayDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubesprayDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubespray delete forbidden response has a 2xx status code
func (o *KubesprayDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray delete forbidden response has a 3xx status code
func (o *KubesprayDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray delete forbidden response has a 4xx status code
func (o *KubesprayDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray delete forbidden response has a 5xx status code
func (o *KubesprayDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray delete forbidden response a status code equal to that given
func (o *KubesprayDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubesprayDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KubesprayDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KubesprayDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubesprayDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayDeleteNotFound creates a KubesprayDeleteNotFound with default headers values
func NewKubesprayDeleteNotFound() *KubesprayDeleteNotFound {
	return &KubesprayDeleteNotFound{}
}

/*
KubesprayDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubesprayDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubespray delete not found response has a 2xx status code
func (o *KubesprayDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray delete not found response has a 3xx status code
func (o *KubesprayDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray delete not found response has a 4xx status code
func (o *KubesprayDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubespray delete not found response has a 5xx status code
func (o *KubesprayDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubespray delete not found response a status code equal to that given
func (o *KubesprayDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubesprayDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KubesprayDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KubesprayDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubesprayDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubesprayDeleteInternalServerError creates a KubesprayDeleteInternalServerError with default headers values
func NewKubesprayDeleteInternalServerError() *KubesprayDeleteInternalServerError {
	return &KubesprayDeleteInternalServerError{}
}

/*
KubesprayDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubesprayDeleteInternalServerError struct {
}

// IsSuccess returns true when this kubespray delete internal server error response has a 2xx status code
func (o *KubesprayDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubespray delete internal server error response has a 3xx status code
func (o *KubesprayDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubespray delete internal server error response has a 4xx status code
func (o *KubesprayDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubespray delete internal server error response has a 5xx status code
func (o *KubesprayDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubespray delete internal server error response a status code equal to that given
func (o *KubesprayDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubesprayDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteInternalServerError ", 500)
}

func (o *KubesprayDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Kubespray/{id}][%d] kubesprayDeleteInternalServerError ", 500)
}

func (o *KubesprayDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
