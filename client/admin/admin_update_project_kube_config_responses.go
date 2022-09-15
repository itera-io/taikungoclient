// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminUpdateProjectKubeConfigReader is a Reader for the AdminUpdateProjectKubeConfig structure.
type AdminUpdateProjectKubeConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateProjectKubeConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateProjectKubeConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateProjectKubeConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateProjectKubeConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateProjectKubeConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminUpdateProjectKubeConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateProjectKubeConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUpdateProjectKubeConfigOK creates a AdminUpdateProjectKubeConfigOK with default headers values
func NewAdminUpdateProjectKubeConfigOK() *AdminUpdateProjectKubeConfigOK {
	return &AdminUpdateProjectKubeConfigOK{}
}

/*
AdminUpdateProjectKubeConfigOK describes a response with status code 200, with default header values.

Success
*/
type AdminUpdateProjectKubeConfigOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this admin update project kube config o k response has a 2xx status code
func (o *AdminUpdateProjectKubeConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin update project kube config o k response has a 3xx status code
func (o *AdminUpdateProjectKubeConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update project kube config o k response has a 4xx status code
func (o *AdminUpdateProjectKubeConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update project kube config o k response has a 5xx status code
func (o *AdminUpdateProjectKubeConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update project kube config o k response a status code equal to that given
func (o *AdminUpdateProjectKubeConfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminUpdateProjectKubeConfigOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigOK  %+v", 200, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigOK  %+v", 200, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigBadRequest creates a AdminUpdateProjectKubeConfigBadRequest with default headers values
func NewAdminUpdateProjectKubeConfigBadRequest() *AdminUpdateProjectKubeConfigBadRequest {
	return &AdminUpdateProjectKubeConfigBadRequest{}
}

/*
AdminUpdateProjectKubeConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminUpdateProjectKubeConfigBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this admin update project kube config bad request response has a 2xx status code
func (o *AdminUpdateProjectKubeConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update project kube config bad request response has a 3xx status code
func (o *AdminUpdateProjectKubeConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update project kube config bad request response has a 4xx status code
func (o *AdminUpdateProjectKubeConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update project kube config bad request response has a 5xx status code
func (o *AdminUpdateProjectKubeConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update project kube config bad request response a status code equal to that given
func (o *AdminUpdateProjectKubeConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AdminUpdateProjectKubeConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigUnauthorized creates a AdminUpdateProjectKubeConfigUnauthorized with default headers values
func NewAdminUpdateProjectKubeConfigUnauthorized() *AdminUpdateProjectKubeConfigUnauthorized {
	return &AdminUpdateProjectKubeConfigUnauthorized{}
}

/*
AdminUpdateProjectKubeConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminUpdateProjectKubeConfigUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin update project kube config unauthorized response has a 2xx status code
func (o *AdminUpdateProjectKubeConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update project kube config unauthorized response has a 3xx status code
func (o *AdminUpdateProjectKubeConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update project kube config unauthorized response has a 4xx status code
func (o *AdminUpdateProjectKubeConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update project kube config unauthorized response has a 5xx status code
func (o *AdminUpdateProjectKubeConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update project kube config unauthorized response a status code equal to that given
func (o *AdminUpdateProjectKubeConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AdminUpdateProjectKubeConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigForbidden creates a AdminUpdateProjectKubeConfigForbidden with default headers values
func NewAdminUpdateProjectKubeConfigForbidden() *AdminUpdateProjectKubeConfigForbidden {
	return &AdminUpdateProjectKubeConfigForbidden{}
}

/*
AdminUpdateProjectKubeConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminUpdateProjectKubeConfigForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin update project kube config forbidden response has a 2xx status code
func (o *AdminUpdateProjectKubeConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update project kube config forbidden response has a 3xx status code
func (o *AdminUpdateProjectKubeConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update project kube config forbidden response has a 4xx status code
func (o *AdminUpdateProjectKubeConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update project kube config forbidden response has a 5xx status code
func (o *AdminUpdateProjectKubeConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update project kube config forbidden response a status code equal to that given
func (o *AdminUpdateProjectKubeConfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AdminUpdateProjectKubeConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigForbidden  %+v", 403, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigForbidden  %+v", 403, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigNotFound creates a AdminUpdateProjectKubeConfigNotFound with default headers values
func NewAdminUpdateProjectKubeConfigNotFound() *AdminUpdateProjectKubeConfigNotFound {
	return &AdminUpdateProjectKubeConfigNotFound{}
}

/*
AdminUpdateProjectKubeConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminUpdateProjectKubeConfigNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin update project kube config not found response has a 2xx status code
func (o *AdminUpdateProjectKubeConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update project kube config not found response has a 3xx status code
func (o *AdminUpdateProjectKubeConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update project kube config not found response has a 4xx status code
func (o *AdminUpdateProjectKubeConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update project kube config not found response has a 5xx status code
func (o *AdminUpdateProjectKubeConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update project kube config not found response a status code equal to that given
func (o *AdminUpdateProjectKubeConfigNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AdminUpdateProjectKubeConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigNotFound  %+v", 404, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigNotFound  %+v", 404, o.Payload)
}

func (o *AdminUpdateProjectKubeConfigNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigInternalServerError creates a AdminUpdateProjectKubeConfigInternalServerError with default headers values
func NewAdminUpdateProjectKubeConfigInternalServerError() *AdminUpdateProjectKubeConfigInternalServerError {
	return &AdminUpdateProjectKubeConfigInternalServerError{}
}

/*
AdminUpdateProjectKubeConfigInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminUpdateProjectKubeConfigInternalServerError struct {
}

// IsSuccess returns true when this admin update project kube config internal server error response has a 2xx status code
func (o *AdminUpdateProjectKubeConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update project kube config internal server error response has a 3xx status code
func (o *AdminUpdateProjectKubeConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update project kube config internal server error response has a 4xx status code
func (o *AdminUpdateProjectKubeConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update project kube config internal server error response has a 5xx status code
func (o *AdminUpdateProjectKubeConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin update project kube config internal server error response a status code equal to that given
func (o *AdminUpdateProjectKubeConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AdminUpdateProjectKubeConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigInternalServerError ", 500)
}

func (o *AdminUpdateProjectKubeConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigInternalServerError ", 500)
}

func (o *AdminUpdateProjectKubeConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
