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

// KubernetesProfilesListReader is a Reader for the KubernetesProfilesList structure.
type KubernetesProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesProfilesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesProfilesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesProfilesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesProfilesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesProfilesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesProfilesListOK creates a KubernetesProfilesListOK with default headers values
func NewKubernetesProfilesListOK() *KubernetesProfilesListOK {
	return &KubernetesProfilesListOK{}
}

/*
KubernetesProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesProfilesListOK struct {
	Payload *models.KubernetesProfilesList
}

// IsSuccess returns true when this kubernetes profiles list o k response has a 2xx status code
func (o *KubernetesProfilesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes profiles list o k response has a 3xx status code
func (o *KubernetesProfilesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles list o k response has a 4xx status code
func (o *KubernetesProfilesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes profiles list o k response has a 5xx status code
func (o *KubernetesProfilesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles list o k response a status code equal to that given
func (o *KubernetesProfilesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesProfilesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListOK  %+v", 200, o.Payload)
}

func (o *KubernetesProfilesListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListOK  %+v", 200, o.Payload)
}

func (o *KubernetesProfilesListOK) GetPayload() *models.KubernetesProfilesList {
	return o.Payload
}

func (o *KubernetesProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesProfilesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListBadRequest creates a KubernetesProfilesListBadRequest with default headers values
func NewKubernetesProfilesListBadRequest() *KubernetesProfilesListBadRequest {
	return &KubernetesProfilesListBadRequest{}
}

/*
KubernetesProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesProfilesListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes profiles list bad request response has a 2xx status code
func (o *KubernetesProfilesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles list bad request response has a 3xx status code
func (o *KubernetesProfilesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles list bad request response has a 4xx status code
func (o *KubernetesProfilesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles list bad request response has a 5xx status code
func (o *KubernetesProfilesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles list bad request response a status code equal to that given
func (o *KubernetesProfilesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesProfilesListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesProfilesListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListUnauthorized creates a KubernetesProfilesListUnauthorized with default headers values
func NewKubernetesProfilesListUnauthorized() *KubernetesProfilesListUnauthorized {
	return &KubernetesProfilesListUnauthorized{}
}

/*
KubernetesProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesProfilesListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes profiles list unauthorized response has a 2xx status code
func (o *KubernetesProfilesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles list unauthorized response has a 3xx status code
func (o *KubernetesProfilesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles list unauthorized response has a 4xx status code
func (o *KubernetesProfilesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles list unauthorized response has a 5xx status code
func (o *KubernetesProfilesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles list unauthorized response a status code equal to that given
func (o *KubernetesProfilesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesProfilesListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesProfilesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListForbidden creates a KubernetesProfilesListForbidden with default headers values
func NewKubernetesProfilesListForbidden() *KubernetesProfilesListForbidden {
	return &KubernetesProfilesListForbidden{}
}

/*
KubernetesProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesProfilesListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes profiles list forbidden response has a 2xx status code
func (o *KubernetesProfilesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles list forbidden response has a 3xx status code
func (o *KubernetesProfilesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles list forbidden response has a 4xx status code
func (o *KubernetesProfilesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles list forbidden response has a 5xx status code
func (o *KubernetesProfilesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles list forbidden response a status code equal to that given
func (o *KubernetesProfilesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesProfilesListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesProfilesListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesProfilesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListNotFound creates a KubernetesProfilesListNotFound with default headers values
func NewKubernetesProfilesListNotFound() *KubernetesProfilesListNotFound {
	return &KubernetesProfilesListNotFound{}
}

/*
KubernetesProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesProfilesListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes profiles list not found response has a 2xx status code
func (o *KubernetesProfilesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles list not found response has a 3xx status code
func (o *KubernetesProfilesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles list not found response has a 4xx status code
func (o *KubernetesProfilesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes profiles list not found response has a 5xx status code
func (o *KubernetesProfilesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes profiles list not found response a status code equal to that given
func (o *KubernetesProfilesListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesProfilesListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesProfilesListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesProfilesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListInternalServerError creates a KubernetesProfilesListInternalServerError with default headers values
func NewKubernetesProfilesListInternalServerError() *KubernetesProfilesListInternalServerError {
	return &KubernetesProfilesListInternalServerError{}
}

/*
KubernetesProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesProfilesListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes profiles list internal server error response has a 2xx status code
func (o *KubernetesProfilesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes profiles list internal server error response has a 3xx status code
func (o *KubernetesProfilesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes profiles list internal server error response has a 4xx status code
func (o *KubernetesProfilesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes profiles list internal server error response has a 5xx status code
func (o *KubernetesProfilesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes profiles list internal server error response a status code equal to that given
func (o *KubernetesProfilesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListInternalServerError ", 500)
}

func (o *KubernetesProfilesListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListInternalServerError ", 500)
}

func (o *KubernetesProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
