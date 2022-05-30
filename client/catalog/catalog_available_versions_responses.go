// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CatalogAvailableVersionsReader is a Reader for the CatalogAvailableVersions structure.
type CatalogAvailableVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogAvailableVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogAvailableVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogAvailableVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogAvailableVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogAvailableVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogAvailableVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogAvailableVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogAvailableVersionsOK creates a CatalogAvailableVersionsOK with default headers values
func NewCatalogAvailableVersionsOK() *CatalogAvailableVersionsOK {
	return &CatalogAvailableVersionsOK{}
}

/* CatalogAvailableVersionsOK describes a response with status code 200, with default header values.

Success
*/
type CatalogAvailableVersionsOK struct {
	Payload []string
}

func (o *CatalogAvailableVersionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/available/versions][%d] catalogAvailableVersionsOK  %+v", 200, o.Payload)
}
func (o *CatalogAvailableVersionsOK) GetPayload() []string {
	return o.Payload
}

func (o *CatalogAvailableVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailableVersionsBadRequest creates a CatalogAvailableVersionsBadRequest with default headers values
func NewCatalogAvailableVersionsBadRequest() *CatalogAvailableVersionsBadRequest {
	return &CatalogAvailableVersionsBadRequest{}
}

/* CatalogAvailableVersionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogAvailableVersionsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CatalogAvailableVersionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/available/versions][%d] catalogAvailableVersionsBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogAvailableVersionsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CatalogAvailableVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailableVersionsUnauthorized creates a CatalogAvailableVersionsUnauthorized with default headers values
func NewCatalogAvailableVersionsUnauthorized() *CatalogAvailableVersionsUnauthorized {
	return &CatalogAvailableVersionsUnauthorized{}
}

/* CatalogAvailableVersionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogAvailableVersionsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CatalogAvailableVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/available/versions][%d] catalogAvailableVersionsUnauthorized  %+v", 401, o.Payload)
}
func (o *CatalogAvailableVersionsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogAvailableVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailableVersionsForbidden creates a CatalogAvailableVersionsForbidden with default headers values
func NewCatalogAvailableVersionsForbidden() *CatalogAvailableVersionsForbidden {
	return &CatalogAvailableVersionsForbidden{}
}

/* CatalogAvailableVersionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogAvailableVersionsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CatalogAvailableVersionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/available/versions][%d] catalogAvailableVersionsForbidden  %+v", 403, o.Payload)
}
func (o *CatalogAvailableVersionsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogAvailableVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailableVersionsNotFound creates a CatalogAvailableVersionsNotFound with default headers values
func NewCatalogAvailableVersionsNotFound() *CatalogAvailableVersionsNotFound {
	return &CatalogAvailableVersionsNotFound{}
}

/* CatalogAvailableVersionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogAvailableVersionsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CatalogAvailableVersionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/available/versions][%d] catalogAvailableVersionsNotFound  %+v", 404, o.Payload)
}
func (o *CatalogAvailableVersionsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogAvailableVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailableVersionsInternalServerError creates a CatalogAvailableVersionsInternalServerError with default headers values
func NewCatalogAvailableVersionsInternalServerError() *CatalogAvailableVersionsInternalServerError {
	return &CatalogAvailableVersionsInternalServerError{}
}

/* CatalogAvailableVersionsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogAvailableVersionsInternalServerError struct {
}

func (o *CatalogAvailableVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/available/versions][%d] catalogAvailableVersionsInternalServerError ", 500)
}

func (o *CatalogAvailableVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}