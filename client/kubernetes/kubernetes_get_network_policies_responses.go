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

// KubernetesGetNetworkPoliciesReader is a Reader for the KubernetesGetNetworkPolicies structure.
type KubernetesGetNetworkPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetNetworkPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetNetworkPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetNetworkPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetNetworkPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetNetworkPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetNetworkPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetNetworkPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetNetworkPoliciesOK creates a KubernetesGetNetworkPoliciesOK with default headers values
func NewKubernetesGetNetworkPoliciesOK() *KubernetesGetNetworkPoliciesOK {
	return &KubernetesGetNetworkPoliciesOK{}
}

/*
KubernetesGetNetworkPoliciesOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetNetworkPoliciesOK struct {
	Payload *models.NetworkPolicies
}

// IsSuccess returns true when this kubernetes get network policies o k response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get network policies o k response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies o k response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get network policies o k response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies o k response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetNetworkPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesOK) GetPayload() *models.NetworkPolicies {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPolicies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesBadRequest creates a KubernetesGetNetworkPoliciesBadRequest with default headers values
func NewKubernetesGetNetworkPoliciesBadRequest() *KubernetesGetNetworkPoliciesBadRequest {
	return &KubernetesGetNetworkPoliciesBadRequest{}
}

/*
KubernetesGetNetworkPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetNetworkPoliciesBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get network policies bad request response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies bad request response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies bad request response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies bad request response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies bad request response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetNetworkPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesUnauthorized creates a KubernetesGetNetworkPoliciesUnauthorized with default headers values
func NewKubernetesGetNetworkPoliciesUnauthorized() *KubernetesGetNetworkPoliciesUnauthorized {
	return &KubernetesGetNetworkPoliciesUnauthorized{}
}

/*
KubernetesGetNetworkPoliciesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetNetworkPoliciesUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get network policies unauthorized response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies unauthorized response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies unauthorized response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies unauthorized response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies unauthorized response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesForbidden creates a KubernetesGetNetworkPoliciesForbidden with default headers values
func NewKubernetesGetNetworkPoliciesForbidden() *KubernetesGetNetworkPoliciesForbidden {
	return &KubernetesGetNetworkPoliciesForbidden{}
}

/*
KubernetesGetNetworkPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetNetworkPoliciesForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get network policies forbidden response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies forbidden response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies forbidden response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies forbidden response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies forbidden response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetNetworkPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesNotFound creates a KubernetesGetNetworkPoliciesNotFound with default headers values
func NewKubernetesGetNetworkPoliciesNotFound() *KubernetesGetNetworkPoliciesNotFound {
	return &KubernetesGetNetworkPoliciesNotFound{}
}

/*
KubernetesGetNetworkPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetNetworkPoliciesNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get network policies not found response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies not found response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies not found response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get network policies not found response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get network policies not found response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetNetworkPoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetNetworkPoliciesNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetNetworkPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNetworkPoliciesInternalServerError creates a KubernetesGetNetworkPoliciesInternalServerError with default headers values
func NewKubernetesGetNetworkPoliciesInternalServerError() *KubernetesGetNetworkPoliciesInternalServerError {
	return &KubernetesGetNetworkPoliciesInternalServerError{}
}

/*
KubernetesGetNetworkPoliciesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetNetworkPoliciesInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get network policies internal server error response has a 2xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get network policies internal server error response has a 3xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get network policies internal server error response has a 4xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get network policies internal server error response has a 5xx status code
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get network policies internal server error response a status code equal to that given
func (o *KubernetesGetNetworkPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetNetworkPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesInternalServerError ", 500)
}

func (o *KubernetesGetNetworkPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/network-policies][%d] kubernetesGetNetworkPoliciesInternalServerError ", 500)
}

func (o *KubernetesGetNetworkPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
