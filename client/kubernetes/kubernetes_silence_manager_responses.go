// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubernetesSilenceManagerReader is a Reader for the KubernetesSilenceManager structure.
type KubernetesSilenceManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesSilenceManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesSilenceManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesSilenceManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesSilenceManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesSilenceManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesSilenceManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesSilenceManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesSilenceManagerOK creates a KubernetesSilenceManagerOK with default headers values
func NewKubernetesSilenceManagerOK() *KubernetesSilenceManagerOK {
	return &KubernetesSilenceManagerOK{}
}

/*
KubernetesSilenceManagerOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesSilenceManagerOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes silence manager o k response has a 2xx status code
func (o *KubernetesSilenceManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes silence manager o k response has a 3xx status code
func (o *KubernetesSilenceManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes silence manager o k response has a 4xx status code
func (o *KubernetesSilenceManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes silence manager o k response has a 5xx status code
func (o *KubernetesSilenceManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes silence manager o k response a status code equal to that given
func (o *KubernetesSilenceManagerOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesSilenceManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerOK  %+v", 200, o.Payload)
}

func (o *KubernetesSilenceManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerOK  %+v", 200, o.Payload)
}

func (o *KubernetesSilenceManagerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesSilenceManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesSilenceManagerBadRequest creates a KubernetesSilenceManagerBadRequest with default headers values
func NewKubernetesSilenceManagerBadRequest() *KubernetesSilenceManagerBadRequest {
	return &KubernetesSilenceManagerBadRequest{}
}

/*
KubernetesSilenceManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesSilenceManagerBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes silence manager bad request response has a 2xx status code
func (o *KubernetesSilenceManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes silence manager bad request response has a 3xx status code
func (o *KubernetesSilenceManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes silence manager bad request response has a 4xx status code
func (o *KubernetesSilenceManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes silence manager bad request response has a 5xx status code
func (o *KubernetesSilenceManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes silence manager bad request response a status code equal to that given
func (o *KubernetesSilenceManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesSilenceManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesSilenceManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesSilenceManagerBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesSilenceManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesSilenceManagerUnauthorized creates a KubernetesSilenceManagerUnauthorized with default headers values
func NewKubernetesSilenceManagerUnauthorized() *KubernetesSilenceManagerUnauthorized {
	return &KubernetesSilenceManagerUnauthorized{}
}

/*
KubernetesSilenceManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesSilenceManagerUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes silence manager unauthorized response has a 2xx status code
func (o *KubernetesSilenceManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes silence manager unauthorized response has a 3xx status code
func (o *KubernetesSilenceManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes silence manager unauthorized response has a 4xx status code
func (o *KubernetesSilenceManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes silence manager unauthorized response has a 5xx status code
func (o *KubernetesSilenceManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes silence manager unauthorized response a status code equal to that given
func (o *KubernetesSilenceManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesSilenceManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesSilenceManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesSilenceManagerUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesSilenceManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesSilenceManagerForbidden creates a KubernetesSilenceManagerForbidden with default headers values
func NewKubernetesSilenceManagerForbidden() *KubernetesSilenceManagerForbidden {
	return &KubernetesSilenceManagerForbidden{}
}

/*
KubernetesSilenceManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesSilenceManagerForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes silence manager forbidden response has a 2xx status code
func (o *KubernetesSilenceManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes silence manager forbidden response has a 3xx status code
func (o *KubernetesSilenceManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes silence manager forbidden response has a 4xx status code
func (o *KubernetesSilenceManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes silence manager forbidden response has a 5xx status code
func (o *KubernetesSilenceManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes silence manager forbidden response a status code equal to that given
func (o *KubernetesSilenceManagerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesSilenceManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesSilenceManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesSilenceManagerForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesSilenceManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesSilenceManagerNotFound creates a KubernetesSilenceManagerNotFound with default headers values
func NewKubernetesSilenceManagerNotFound() *KubernetesSilenceManagerNotFound {
	return &KubernetesSilenceManagerNotFound{}
}

/*
KubernetesSilenceManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesSilenceManagerNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes silence manager not found response has a 2xx status code
func (o *KubernetesSilenceManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes silence manager not found response has a 3xx status code
func (o *KubernetesSilenceManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes silence manager not found response has a 4xx status code
func (o *KubernetesSilenceManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes silence manager not found response has a 5xx status code
func (o *KubernetesSilenceManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes silence manager not found response a status code equal to that given
func (o *KubernetesSilenceManagerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesSilenceManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesSilenceManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesSilenceManagerNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesSilenceManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesSilenceManagerInternalServerError creates a KubernetesSilenceManagerInternalServerError with default headers values
func NewKubernetesSilenceManagerInternalServerError() *KubernetesSilenceManagerInternalServerError {
	return &KubernetesSilenceManagerInternalServerError{}
}

/*
KubernetesSilenceManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesSilenceManagerInternalServerError struct {
}

// IsSuccess returns true when this kubernetes silence manager internal server error response has a 2xx status code
func (o *KubernetesSilenceManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes silence manager internal server error response has a 3xx status code
func (o *KubernetesSilenceManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes silence manager internal server error response has a 4xx status code
func (o *KubernetesSilenceManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes silence manager internal server error response has a 5xx status code
func (o *KubernetesSilenceManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes silence manager internal server error response a status code equal to that given
func (o *KubernetesSilenceManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesSilenceManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerInternalServerError ", 500)
}

func (o *KubernetesSilenceManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/silencemanager][%d] kubernetesSilenceManagerInternalServerError ", 500)
}

func (o *KubernetesSilenceManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
