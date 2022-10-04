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

// KubernetesRestartDeploymentReader is a Reader for the KubernetesRestartDeployment structure.
type KubernetesRestartDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesRestartDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesRestartDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesRestartDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesRestartDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesRestartDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesRestartDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesRestartDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesRestartDeploymentOK creates a KubernetesRestartDeploymentOK with default headers values
func NewKubernetesRestartDeploymentOK() *KubernetesRestartDeploymentOK {
	return &KubernetesRestartDeploymentOK{}
}

/*
KubernetesRestartDeploymentOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesRestartDeploymentOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes restart deployment o k response has a 2xx status code
func (o *KubernetesRestartDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes restart deployment o k response has a 3xx status code
func (o *KubernetesRestartDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart deployment o k response has a 4xx status code
func (o *KubernetesRestartDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes restart deployment o k response has a 5xx status code
func (o *KubernetesRestartDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart deployment o k response a status code equal to that given
func (o *KubernetesRestartDeploymentOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesRestartDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentOK  %+v", 200, o.Payload)
}

func (o *KubernetesRestartDeploymentOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentOK  %+v", 200, o.Payload)
}

func (o *KubernetesRestartDeploymentOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesRestartDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDeploymentBadRequest creates a KubernetesRestartDeploymentBadRequest with default headers values
func NewKubernetesRestartDeploymentBadRequest() *KubernetesRestartDeploymentBadRequest {
	return &KubernetesRestartDeploymentBadRequest{}
}

/*
KubernetesRestartDeploymentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesRestartDeploymentBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes restart deployment bad request response has a 2xx status code
func (o *KubernetesRestartDeploymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart deployment bad request response has a 3xx status code
func (o *KubernetesRestartDeploymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart deployment bad request response has a 4xx status code
func (o *KubernetesRestartDeploymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart deployment bad request response has a 5xx status code
func (o *KubernetesRestartDeploymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart deployment bad request response a status code equal to that given
func (o *KubernetesRestartDeploymentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesRestartDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesRestartDeploymentBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesRestartDeploymentBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesRestartDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDeploymentUnauthorized creates a KubernetesRestartDeploymentUnauthorized with default headers values
func NewKubernetesRestartDeploymentUnauthorized() *KubernetesRestartDeploymentUnauthorized {
	return &KubernetesRestartDeploymentUnauthorized{}
}

/*
KubernetesRestartDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesRestartDeploymentUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes restart deployment unauthorized response has a 2xx status code
func (o *KubernetesRestartDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart deployment unauthorized response has a 3xx status code
func (o *KubernetesRestartDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart deployment unauthorized response has a 4xx status code
func (o *KubernetesRestartDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart deployment unauthorized response has a 5xx status code
func (o *KubernetesRestartDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart deployment unauthorized response a status code equal to that given
func (o *KubernetesRestartDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesRestartDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesRestartDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesRestartDeploymentUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesRestartDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDeploymentForbidden creates a KubernetesRestartDeploymentForbidden with default headers values
func NewKubernetesRestartDeploymentForbidden() *KubernetesRestartDeploymentForbidden {
	return &KubernetesRestartDeploymentForbidden{}
}

/*
KubernetesRestartDeploymentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesRestartDeploymentForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes restart deployment forbidden response has a 2xx status code
func (o *KubernetesRestartDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart deployment forbidden response has a 3xx status code
func (o *KubernetesRestartDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart deployment forbidden response has a 4xx status code
func (o *KubernetesRestartDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart deployment forbidden response has a 5xx status code
func (o *KubernetesRestartDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart deployment forbidden response a status code equal to that given
func (o *KubernetesRestartDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesRestartDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesRestartDeploymentForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesRestartDeploymentForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesRestartDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDeploymentNotFound creates a KubernetesRestartDeploymentNotFound with default headers values
func NewKubernetesRestartDeploymentNotFound() *KubernetesRestartDeploymentNotFound {
	return &KubernetesRestartDeploymentNotFound{}
}

/*
KubernetesRestartDeploymentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesRestartDeploymentNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes restart deployment not found response has a 2xx status code
func (o *KubernetesRestartDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart deployment not found response has a 3xx status code
func (o *KubernetesRestartDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart deployment not found response has a 4xx status code
func (o *KubernetesRestartDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart deployment not found response has a 5xx status code
func (o *KubernetesRestartDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart deployment not found response a status code equal to that given
func (o *KubernetesRestartDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesRestartDeploymentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesRestartDeploymentNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesRestartDeploymentNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesRestartDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDeploymentInternalServerError creates a KubernetesRestartDeploymentInternalServerError with default headers values
func NewKubernetesRestartDeploymentInternalServerError() *KubernetesRestartDeploymentInternalServerError {
	return &KubernetesRestartDeploymentInternalServerError{}
}

/*
KubernetesRestartDeploymentInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesRestartDeploymentInternalServerError struct {
}

// IsSuccess returns true when this kubernetes restart deployment internal server error response has a 2xx status code
func (o *KubernetesRestartDeploymentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart deployment internal server error response has a 3xx status code
func (o *KubernetesRestartDeploymentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart deployment internal server error response has a 4xx status code
func (o *KubernetesRestartDeploymentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes restart deployment internal server error response has a 5xx status code
func (o *KubernetesRestartDeploymentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes restart deployment internal server error response a status code equal to that given
func (o *KubernetesRestartDeploymentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesRestartDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentInternalServerError ", 500)
}

func (o *KubernetesRestartDeploymentInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/deployment][%d] kubernetesRestartDeploymentInternalServerError ", 500)
}

func (o *KubernetesRestartDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
