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

// KubernetesGetHelmReleaseListReader is a Reader for the KubernetesGetHelmReleaseList structure.
type KubernetesGetHelmReleaseListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetHelmReleaseListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetHelmReleaseListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetHelmReleaseListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetHelmReleaseListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetHelmReleaseListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetHelmReleaseListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetHelmReleaseListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetHelmReleaseListOK creates a KubernetesGetHelmReleaseListOK with default headers values
func NewKubernetesGetHelmReleaseListOK() *KubernetesGetHelmReleaseListOK {
	return &KubernetesGetHelmReleaseListOK{}
}

/*
KubernetesGetHelmReleaseListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetHelmReleaseListOK struct {
	Payload *models.HelmReleasesList
}

// IsSuccess returns true when this kubernetes get helm release list o k response has a 2xx status code
func (o *KubernetesGetHelmReleaseListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get helm release list o k response has a 3xx status code
func (o *KubernetesGetHelmReleaseListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get helm release list o k response has a 4xx status code
func (o *KubernetesGetHelmReleaseListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get helm release list o k response has a 5xx status code
func (o *KubernetesGetHelmReleaseListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get helm release list o k response a status code equal to that given
func (o *KubernetesGetHelmReleaseListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetHelmReleaseListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetHelmReleaseListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetHelmReleaseListOK) GetPayload() *models.HelmReleasesList {
	return o.Payload
}

func (o *KubernetesGetHelmReleaseListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HelmReleasesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetHelmReleaseListBadRequest creates a KubernetesGetHelmReleaseListBadRequest with default headers values
func NewKubernetesGetHelmReleaseListBadRequest() *KubernetesGetHelmReleaseListBadRequest {
	return &KubernetesGetHelmReleaseListBadRequest{}
}

/*
KubernetesGetHelmReleaseListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetHelmReleaseListBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get helm release list bad request response has a 2xx status code
func (o *KubernetesGetHelmReleaseListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get helm release list bad request response has a 3xx status code
func (o *KubernetesGetHelmReleaseListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get helm release list bad request response has a 4xx status code
func (o *KubernetesGetHelmReleaseListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get helm release list bad request response has a 5xx status code
func (o *KubernetesGetHelmReleaseListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get helm release list bad request response a status code equal to that given
func (o *KubernetesGetHelmReleaseListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetHelmReleaseListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetHelmReleaseListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetHelmReleaseListBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetHelmReleaseListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetHelmReleaseListUnauthorized creates a KubernetesGetHelmReleaseListUnauthorized with default headers values
func NewKubernetesGetHelmReleaseListUnauthorized() *KubernetesGetHelmReleaseListUnauthorized {
	return &KubernetesGetHelmReleaseListUnauthorized{}
}

/*
KubernetesGetHelmReleaseListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetHelmReleaseListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get helm release list unauthorized response has a 2xx status code
func (o *KubernetesGetHelmReleaseListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get helm release list unauthorized response has a 3xx status code
func (o *KubernetesGetHelmReleaseListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get helm release list unauthorized response has a 4xx status code
func (o *KubernetesGetHelmReleaseListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get helm release list unauthorized response has a 5xx status code
func (o *KubernetesGetHelmReleaseListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get helm release list unauthorized response a status code equal to that given
func (o *KubernetesGetHelmReleaseListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetHelmReleaseListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetHelmReleaseListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetHelmReleaseListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetHelmReleaseListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetHelmReleaseListForbidden creates a KubernetesGetHelmReleaseListForbidden with default headers values
func NewKubernetesGetHelmReleaseListForbidden() *KubernetesGetHelmReleaseListForbidden {
	return &KubernetesGetHelmReleaseListForbidden{}
}

/*
KubernetesGetHelmReleaseListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetHelmReleaseListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get helm release list forbidden response has a 2xx status code
func (o *KubernetesGetHelmReleaseListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get helm release list forbidden response has a 3xx status code
func (o *KubernetesGetHelmReleaseListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get helm release list forbidden response has a 4xx status code
func (o *KubernetesGetHelmReleaseListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get helm release list forbidden response has a 5xx status code
func (o *KubernetesGetHelmReleaseListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get helm release list forbidden response a status code equal to that given
func (o *KubernetesGetHelmReleaseListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetHelmReleaseListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetHelmReleaseListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetHelmReleaseListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetHelmReleaseListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetHelmReleaseListNotFound creates a KubernetesGetHelmReleaseListNotFound with default headers values
func NewKubernetesGetHelmReleaseListNotFound() *KubernetesGetHelmReleaseListNotFound {
	return &KubernetesGetHelmReleaseListNotFound{}
}

/*
KubernetesGetHelmReleaseListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetHelmReleaseListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get helm release list not found response has a 2xx status code
func (o *KubernetesGetHelmReleaseListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get helm release list not found response has a 3xx status code
func (o *KubernetesGetHelmReleaseListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get helm release list not found response has a 4xx status code
func (o *KubernetesGetHelmReleaseListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get helm release list not found response has a 5xx status code
func (o *KubernetesGetHelmReleaseListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get helm release list not found response a status code equal to that given
func (o *KubernetesGetHelmReleaseListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetHelmReleaseListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetHelmReleaseListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetHelmReleaseListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetHelmReleaseListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetHelmReleaseListInternalServerError creates a KubernetesGetHelmReleaseListInternalServerError with default headers values
func NewKubernetesGetHelmReleaseListInternalServerError() *KubernetesGetHelmReleaseListInternalServerError {
	return &KubernetesGetHelmReleaseListInternalServerError{}
}

/*
KubernetesGetHelmReleaseListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetHelmReleaseListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get helm release list internal server error response has a 2xx status code
func (o *KubernetesGetHelmReleaseListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get helm release list internal server error response has a 3xx status code
func (o *KubernetesGetHelmReleaseListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get helm release list internal server error response has a 4xx status code
func (o *KubernetesGetHelmReleaseListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get helm release list internal server error response has a 5xx status code
func (o *KubernetesGetHelmReleaseListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get helm release list internal server error response a status code equal to that given
func (o *KubernetesGetHelmReleaseListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetHelmReleaseListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListInternalServerError ", 500)
}

func (o *KubernetesGetHelmReleaseListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/helmreleases][%d] kubernetesGetHelmReleaseListInternalServerError ", 500)
}

func (o *KubernetesGetHelmReleaseListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
