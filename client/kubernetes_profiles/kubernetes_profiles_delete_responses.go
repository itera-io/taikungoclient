// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubernetesProfilesDeleteReader is a Reader for the KubernetesProfilesDelete structure.
type KubernetesProfilesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesProfilesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesProfilesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewKubernetesProfilesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesProfilesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesProfilesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesProfilesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesProfilesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesProfilesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesProfilesDeleteOK creates a KubernetesProfilesDeleteOK with default headers values
func NewKubernetesProfilesDeleteOK() *KubernetesProfilesDeleteOK {
	return &KubernetesProfilesDeleteOK{}
}

/*
KubernetesProfilesDeleteOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesProfilesDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes profiles delete o k response has a 2xx status code
func (o *KubernetesProfilesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes profiles delete o k response has a 3xx status code
func (o *KubernetesProfilesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles delete o k response has a 4xx status code
func (o *KubernetesProfilesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes profiles delete o k response has a 5xx status code
func (o *KubernetesProfilesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles delete o k response a status code equal to that given
func (o *KubernetesProfilesDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes profiles delete o k response
func (o *KubernetesProfilesDeleteOK) Code() int {
	return 200
}

func (o *KubernetesProfilesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteOK  %+v", 200, o.Payload)
}

func (o *KubernetesProfilesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteOK  %+v", 200, o.Payload)
}

func (o *KubernetesProfilesDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesProfilesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesDeleteNoContent creates a KubernetesProfilesDeleteNoContent with default headers values
func NewKubernetesProfilesDeleteNoContent() *KubernetesProfilesDeleteNoContent {
	return &KubernetesProfilesDeleteNoContent{}
}

/*
KubernetesProfilesDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type KubernetesProfilesDeleteNoContent struct {
}

// IsSuccess returns true when this kubernetes profiles delete no content response has a 2xx status code
func (o *KubernetesProfilesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes profiles delete no content response has a 3xx status code
func (o *KubernetesProfilesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles delete no content response has a 4xx status code
func (o *KubernetesProfilesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes profiles delete no content response has a 5xx status code
func (o *KubernetesProfilesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles delete no content response a status code equal to that given
func (o *KubernetesProfilesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the kubernetes profiles delete no content response
func (o *KubernetesProfilesDeleteNoContent) Code() int {
	return 204
}

func (o *KubernetesProfilesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteNoContent ", 204)
}

func (o *KubernetesProfilesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteNoContent ", 204)
}

func (o *KubernetesProfilesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKubernetesProfilesDeleteBadRequest creates a KubernetesProfilesDeleteBadRequest with default headers values
func NewKubernetesProfilesDeleteBadRequest() *KubernetesProfilesDeleteBadRequest {
	return &KubernetesProfilesDeleteBadRequest{}
}

/*
KubernetesProfilesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesProfilesDeleteBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes profiles delete bad request response has a 2xx status code
func (o *KubernetesProfilesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles delete bad request response has a 3xx status code
func (o *KubernetesProfilesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles delete bad request response has a 4xx status code
func (o *KubernetesProfilesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles delete bad request response has a 5xx status code
func (o *KubernetesProfilesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles delete bad request response a status code equal to that given
func (o *KubernetesProfilesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes profiles delete bad request response
func (o *KubernetesProfilesDeleteBadRequest) Code() int {
	return 400
}

func (o *KubernetesProfilesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesProfilesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesProfilesDeleteBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesDeleteUnauthorized creates a KubernetesProfilesDeleteUnauthorized with default headers values
func NewKubernetesProfilesDeleteUnauthorized() *KubernetesProfilesDeleteUnauthorized {
	return &KubernetesProfilesDeleteUnauthorized{}
}

/*
KubernetesProfilesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesProfilesDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes profiles delete unauthorized response has a 2xx status code
func (o *KubernetesProfilesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles delete unauthorized response has a 3xx status code
func (o *KubernetesProfilesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles delete unauthorized response has a 4xx status code
func (o *KubernetesProfilesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles delete unauthorized response has a 5xx status code
func (o *KubernetesProfilesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles delete unauthorized response a status code equal to that given
func (o *KubernetesProfilesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes profiles delete unauthorized response
func (o *KubernetesProfilesDeleteUnauthorized) Code() int {
	return 401
}

func (o *KubernetesProfilesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesProfilesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesProfilesDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesDeleteForbidden creates a KubernetesProfilesDeleteForbidden with default headers values
func NewKubernetesProfilesDeleteForbidden() *KubernetesProfilesDeleteForbidden {
	return &KubernetesProfilesDeleteForbidden{}
}

/*
KubernetesProfilesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesProfilesDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes profiles delete forbidden response has a 2xx status code
func (o *KubernetesProfilesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles delete forbidden response has a 3xx status code
func (o *KubernetesProfilesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles delete forbidden response has a 4xx status code
func (o *KubernetesProfilesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles delete forbidden response has a 5xx status code
func (o *KubernetesProfilesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles delete forbidden response a status code equal to that given
func (o *KubernetesProfilesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes profiles delete forbidden response
func (o *KubernetesProfilesDeleteForbidden) Code() int {
	return 403
}

func (o *KubernetesProfilesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesProfilesDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesProfilesDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesDeleteNotFound creates a KubernetesProfilesDeleteNotFound with default headers values
func NewKubernetesProfilesDeleteNotFound() *KubernetesProfilesDeleteNotFound {
	return &KubernetesProfilesDeleteNotFound{}
}

/*
KubernetesProfilesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesProfilesDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes profiles delete not found response has a 2xx status code
func (o *KubernetesProfilesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles delete not found response has a 3xx status code
func (o *KubernetesProfilesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles delete not found response has a 4xx status code
func (o *KubernetesProfilesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles delete not found response has a 5xx status code
func (o *KubernetesProfilesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles delete not found response a status code equal to that given
func (o *KubernetesProfilesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes profiles delete not found response
func (o *KubernetesProfilesDeleteNotFound) Code() int {
	return 404
}

func (o *KubernetesProfilesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesProfilesDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesProfilesDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesDeleteInternalServerError creates a KubernetesProfilesDeleteInternalServerError with default headers values
func NewKubernetesProfilesDeleteInternalServerError() *KubernetesProfilesDeleteInternalServerError {
	return &KubernetesProfilesDeleteInternalServerError{}
}

/*
KubernetesProfilesDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesProfilesDeleteInternalServerError struct {
}

// IsSuccess returns true when this kubernetes profiles delete internal server error response has a 2xx status code
func (o *KubernetesProfilesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles delete internal server error response has a 3xx status code
func (o *KubernetesProfilesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles delete internal server error response has a 4xx status code
func (o *KubernetesProfilesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes profiles delete internal server error response has a 5xx status code
func (o *KubernetesProfilesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes profiles delete internal server error response a status code equal to that given
func (o *KubernetesProfilesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes profiles delete internal server error response
func (o *KubernetesProfilesDeleteInternalServerError) Code() int {
	return 500
}

func (o *KubernetesProfilesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteInternalServerError ", 500)
}

func (o *KubernetesProfilesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/KubernetesProfiles/{id}][%d] kubernetesProfilesDeleteInternalServerError ", 500)
}

func (o *KubernetesProfilesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
