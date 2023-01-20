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

// KubernetesProfilesCreateReader is a Reader for the KubernetesProfilesCreate structure.
type KubernetesProfilesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesProfilesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesProfilesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesProfilesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesProfilesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesProfilesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesProfilesCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesProfilesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesProfilesCreateOK creates a KubernetesProfilesCreateOK with default headers values
func NewKubernetesProfilesCreateOK() *KubernetesProfilesCreateOK {
	return &KubernetesProfilesCreateOK{}
}

/*
KubernetesProfilesCreateOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesProfilesCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this kubernetes profiles create o k response has a 2xx status code
func (o *KubernetesProfilesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes profiles create o k response has a 3xx status code
func (o *KubernetesProfilesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles create o k response has a 4xx status code
func (o *KubernetesProfilesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes profiles create o k response has a 5xx status code
func (o *KubernetesProfilesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles create o k response a status code equal to that given
func (o *KubernetesProfilesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes profiles create o k response
func (o *KubernetesProfilesCreateOK) Code() int {
	return 200
}

func (o *KubernetesProfilesCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *KubernetesProfilesCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *KubernetesProfilesCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *KubernetesProfilesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesCreateBadRequest creates a KubernetesProfilesCreateBadRequest with default headers values
func NewKubernetesProfilesCreateBadRequest() *KubernetesProfilesCreateBadRequest {
	return &KubernetesProfilesCreateBadRequest{}
}

/*
KubernetesProfilesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesProfilesCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes profiles create bad request response has a 2xx status code
func (o *KubernetesProfilesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles create bad request response has a 3xx status code
func (o *KubernetesProfilesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles create bad request response has a 4xx status code
func (o *KubernetesProfilesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles create bad request response has a 5xx status code
func (o *KubernetesProfilesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles create bad request response a status code equal to that given
func (o *KubernetesProfilesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes profiles create bad request response
func (o *KubernetesProfilesCreateBadRequest) Code() int {
	return 400
}

func (o *KubernetesProfilesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesProfilesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesProfilesCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesProfilesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesCreateUnauthorized creates a KubernetesProfilesCreateUnauthorized with default headers values
func NewKubernetesProfilesCreateUnauthorized() *KubernetesProfilesCreateUnauthorized {
	return &KubernetesProfilesCreateUnauthorized{}
}

/*
KubernetesProfilesCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesProfilesCreateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes profiles create unauthorized response has a 2xx status code
func (o *KubernetesProfilesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles create unauthorized response has a 3xx status code
func (o *KubernetesProfilesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles create unauthorized response has a 4xx status code
func (o *KubernetesProfilesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles create unauthorized response has a 5xx status code
func (o *KubernetesProfilesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles create unauthorized response a status code equal to that given
func (o *KubernetesProfilesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes profiles create unauthorized response
func (o *KubernetesProfilesCreateUnauthorized) Code() int {
	return 401
}

func (o *KubernetesProfilesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesProfilesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesProfilesCreateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesProfilesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesCreateForbidden creates a KubernetesProfilesCreateForbidden with default headers values
func NewKubernetesProfilesCreateForbidden() *KubernetesProfilesCreateForbidden {
	return &KubernetesProfilesCreateForbidden{}
}

/*
KubernetesProfilesCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesProfilesCreateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes profiles create forbidden response has a 2xx status code
func (o *KubernetesProfilesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles create forbidden response has a 3xx status code
func (o *KubernetesProfilesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles create forbidden response has a 4xx status code
func (o *KubernetesProfilesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles create forbidden response has a 5xx status code
func (o *KubernetesProfilesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles create forbidden response a status code equal to that given
func (o *KubernetesProfilesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes profiles create forbidden response
func (o *KubernetesProfilesCreateForbidden) Code() int {
	return 403
}

func (o *KubernetesProfilesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesProfilesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesProfilesCreateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesProfilesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesCreateNotFound creates a KubernetesProfilesCreateNotFound with default headers values
func NewKubernetesProfilesCreateNotFound() *KubernetesProfilesCreateNotFound {
	return &KubernetesProfilesCreateNotFound{}
}

/*
KubernetesProfilesCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesProfilesCreateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes profiles create not found response has a 2xx status code
func (o *KubernetesProfilesCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles create not found response has a 3xx status code
func (o *KubernetesProfilesCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles create not found response has a 4xx status code
func (o *KubernetesProfilesCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles create not found response has a 5xx status code
func (o *KubernetesProfilesCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles create not found response a status code equal to that given
func (o *KubernetesProfilesCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes profiles create not found response
func (o *KubernetesProfilesCreateNotFound) Code() int {
	return 404
}

func (o *KubernetesProfilesCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesProfilesCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesProfilesCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesProfilesCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesCreateInternalServerError creates a KubernetesProfilesCreateInternalServerError with default headers values
func NewKubernetesProfilesCreateInternalServerError() *KubernetesProfilesCreateInternalServerError {
	return &KubernetesProfilesCreateInternalServerError{}
}

/*
KubernetesProfilesCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesProfilesCreateInternalServerError struct {
}

// IsSuccess returns true when this kubernetes profiles create internal server error response has a 2xx status code
func (o *KubernetesProfilesCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles create internal server error response has a 3xx status code
func (o *KubernetesProfilesCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles create internal server error response has a 4xx status code
func (o *KubernetesProfilesCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes profiles create internal server error response has a 5xx status code
func (o *KubernetesProfilesCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes profiles create internal server error response a status code equal to that given
func (o *KubernetesProfilesCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes profiles create internal server error response
func (o *KubernetesProfilesCreateInternalServerError) Code() int {
	return 500
}

func (o *KubernetesProfilesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateInternalServerError ", 500)
}

func (o *KubernetesProfilesCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubernetesProfiles][%d] kubernetesProfilesCreateInternalServerError ", 500)
}

func (o *KubernetesProfilesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
