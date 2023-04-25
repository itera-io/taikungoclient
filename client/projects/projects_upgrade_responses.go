// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectsUpgradeReader is a Reader for the ProjectsUpgrade structure.
type ProjectsUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsUpgradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsUpgradeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsUpgradeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsUpgradeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsUpgradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsUpgradeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsUpgradeOK creates a ProjectsUpgradeOK with default headers values
func NewProjectsUpgradeOK() *ProjectsUpgradeOK {
	return &ProjectsUpgradeOK{}
}

/*
ProjectsUpgradeOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsUpgradeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects upgrade o k response has a 2xx status code
func (o *ProjectsUpgradeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects upgrade o k response has a 3xx status code
func (o *ProjectsUpgradeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects upgrade o k response has a 4xx status code
func (o *ProjectsUpgradeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects upgrade o k response has a 5xx status code
func (o *ProjectsUpgradeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects upgrade o k response a status code equal to that given
func (o *ProjectsUpgradeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the projects upgrade o k response
func (o *ProjectsUpgradeOK) Code() int {
	return 200
}

func (o *ProjectsUpgradeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeOK  %+v", 200, o.Payload)
}

func (o *ProjectsUpgradeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeOK  %+v", 200, o.Payload)
}

func (o *ProjectsUpgradeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsUpgradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeBadRequest creates a ProjectsUpgradeBadRequest with default headers values
func NewProjectsUpgradeBadRequest() *ProjectsUpgradeBadRequest {
	return &ProjectsUpgradeBadRequest{}
}

/*
ProjectsUpgradeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsUpgradeBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this projects upgrade bad request response has a 2xx status code
func (o *ProjectsUpgradeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects upgrade bad request response has a 3xx status code
func (o *ProjectsUpgradeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects upgrade bad request response has a 4xx status code
func (o *ProjectsUpgradeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects upgrade bad request response has a 5xx status code
func (o *ProjectsUpgradeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects upgrade bad request response a status code equal to that given
func (o *ProjectsUpgradeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the projects upgrade bad request response
func (o *ProjectsUpgradeBadRequest) Code() int {
	return 400
}

func (o *ProjectsUpgradeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsUpgradeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsUpgradeBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsUpgradeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeUnauthorized creates a ProjectsUpgradeUnauthorized with default headers values
func NewProjectsUpgradeUnauthorized() *ProjectsUpgradeUnauthorized {
	return &ProjectsUpgradeUnauthorized{}
}

/*
ProjectsUpgradeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsUpgradeUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this projects upgrade unauthorized response has a 2xx status code
func (o *ProjectsUpgradeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects upgrade unauthorized response has a 3xx status code
func (o *ProjectsUpgradeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects upgrade unauthorized response has a 4xx status code
func (o *ProjectsUpgradeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects upgrade unauthorized response has a 5xx status code
func (o *ProjectsUpgradeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects upgrade unauthorized response a status code equal to that given
func (o *ProjectsUpgradeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the projects upgrade unauthorized response
func (o *ProjectsUpgradeUnauthorized) Code() int {
	return 401
}

func (o *ProjectsUpgradeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsUpgradeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsUpgradeUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsUpgradeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeForbidden creates a ProjectsUpgradeForbidden with default headers values
func NewProjectsUpgradeForbidden() *ProjectsUpgradeForbidden {
	return &ProjectsUpgradeForbidden{}
}

/*
ProjectsUpgradeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsUpgradeForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this projects upgrade forbidden response has a 2xx status code
func (o *ProjectsUpgradeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects upgrade forbidden response has a 3xx status code
func (o *ProjectsUpgradeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects upgrade forbidden response has a 4xx status code
func (o *ProjectsUpgradeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects upgrade forbidden response has a 5xx status code
func (o *ProjectsUpgradeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects upgrade forbidden response a status code equal to that given
func (o *ProjectsUpgradeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the projects upgrade forbidden response
func (o *ProjectsUpgradeForbidden) Code() int {
	return 403
}

func (o *ProjectsUpgradeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsUpgradeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsUpgradeForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsUpgradeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeNotFound creates a ProjectsUpgradeNotFound with default headers values
func NewProjectsUpgradeNotFound() *ProjectsUpgradeNotFound {
	return &ProjectsUpgradeNotFound{}
}

/*
ProjectsUpgradeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsUpgradeNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this projects upgrade not found response has a 2xx status code
func (o *ProjectsUpgradeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects upgrade not found response has a 3xx status code
func (o *ProjectsUpgradeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects upgrade not found response has a 4xx status code
func (o *ProjectsUpgradeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects upgrade not found response has a 5xx status code
func (o *ProjectsUpgradeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects upgrade not found response a status code equal to that given
func (o *ProjectsUpgradeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the projects upgrade not found response
func (o *ProjectsUpgradeNotFound) Code() int {
	return 404
}

func (o *ProjectsUpgradeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsUpgradeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsUpgradeNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsUpgradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeInternalServerError creates a ProjectsUpgradeInternalServerError with default headers values
func NewProjectsUpgradeInternalServerError() *ProjectsUpgradeInternalServerError {
	return &ProjectsUpgradeInternalServerError{}
}

/*
ProjectsUpgradeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsUpgradeInternalServerError struct {
}

// IsSuccess returns true when this projects upgrade internal server error response has a 2xx status code
func (o *ProjectsUpgradeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects upgrade internal server error response has a 3xx status code
func (o *ProjectsUpgradeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects upgrade internal server error response has a 4xx status code
func (o *ProjectsUpgradeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects upgrade internal server error response has a 5xx status code
func (o *ProjectsUpgradeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects upgrade internal server error response a status code equal to that given
func (o *ProjectsUpgradeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the projects upgrade internal server error response
func (o *ProjectsUpgradeInternalServerError) Code() int {
	return 500
}

func (o *ProjectsUpgradeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeInternalServerError ", 500)
}

func (o *ProjectsUpgradeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeInternalServerError ", 500)
}

func (o *ProjectsUpgradeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
