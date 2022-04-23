// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// RepositoryListOfRulesReader is a Reader for the RepositoryListOfRules structure.
type RepositoryListOfRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepositoryListOfRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepositoryListOfRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepositoryListOfRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRepositoryListOfRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRepositoryListOfRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepositoryListOfRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRepositoryListOfRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRepositoryListOfRulesOK creates a RepositoryListOfRulesOK with default headers values
func NewRepositoryListOfRulesOK() *RepositoryListOfRulesOK {
	return &RepositoryListOfRulesOK{}
}

/* RepositoryListOfRulesOK describes a response with status code 200, with default header values.

Success
*/
type RepositoryListOfRulesOK struct {
	Payload *models.ArtifactRepositoryList
}

func (o *RepositoryListOfRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Repository/available][%d] repositoryListOfRulesOK  %+v", 200, o.Payload)
}
func (o *RepositoryListOfRulesOK) GetPayload() *models.ArtifactRepositoryList {
	return o.Payload
}

func (o *RepositoryListOfRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArtifactRepositoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryListOfRulesBadRequest creates a RepositoryListOfRulesBadRequest with default headers values
func NewRepositoryListOfRulesBadRequest() *RepositoryListOfRulesBadRequest {
	return &RepositoryListOfRulesBadRequest{}
}

/* RepositoryListOfRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RepositoryListOfRulesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *RepositoryListOfRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Repository/available][%d] repositoryListOfRulesBadRequest  %+v", 400, o.Payload)
}
func (o *RepositoryListOfRulesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *RepositoryListOfRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryListOfRulesUnauthorized creates a RepositoryListOfRulesUnauthorized with default headers values
func NewRepositoryListOfRulesUnauthorized() *RepositoryListOfRulesUnauthorized {
	return &RepositoryListOfRulesUnauthorized{}
}

/* RepositoryListOfRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RepositoryListOfRulesUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *RepositoryListOfRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Repository/available][%d] repositoryListOfRulesUnauthorized  %+v", 401, o.Payload)
}
func (o *RepositoryListOfRulesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RepositoryListOfRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryListOfRulesForbidden creates a RepositoryListOfRulesForbidden with default headers values
func NewRepositoryListOfRulesForbidden() *RepositoryListOfRulesForbidden {
	return &RepositoryListOfRulesForbidden{}
}

/* RepositoryListOfRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RepositoryListOfRulesForbidden struct {
	Payload *models.ProblemDetails
}

func (o *RepositoryListOfRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Repository/available][%d] repositoryListOfRulesForbidden  %+v", 403, o.Payload)
}
func (o *RepositoryListOfRulesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RepositoryListOfRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryListOfRulesNotFound creates a RepositoryListOfRulesNotFound with default headers values
func NewRepositoryListOfRulesNotFound() *RepositoryListOfRulesNotFound {
	return &RepositoryListOfRulesNotFound{}
}

/* RepositoryListOfRulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RepositoryListOfRulesNotFound struct {
	Payload *models.ProblemDetails
}

func (o *RepositoryListOfRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Repository/available][%d] repositoryListOfRulesNotFound  %+v", 404, o.Payload)
}
func (o *RepositoryListOfRulesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RepositoryListOfRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryListOfRulesInternalServerError creates a RepositoryListOfRulesInternalServerError with default headers values
func NewRepositoryListOfRulesInternalServerError() *RepositoryListOfRulesInternalServerError {
	return &RepositoryListOfRulesInternalServerError{}
}

/* RepositoryListOfRulesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type RepositoryListOfRulesInternalServerError struct {
}

func (o *RepositoryListOfRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Repository/available][%d] repositoryListOfRulesInternalServerError ", 500)
}

func (o *RepositoryListOfRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
