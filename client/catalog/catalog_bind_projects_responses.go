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

// CatalogBindProjectsReader is a Reader for the CatalogBindProjects structure.
type CatalogBindProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogBindProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogBindProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogBindProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogBindProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogBindProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogBindProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogBindProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogBindProjectsOK creates a CatalogBindProjectsOK with default headers values
func NewCatalogBindProjectsOK() *CatalogBindProjectsOK {
	return &CatalogBindProjectsOK{}
}

/* CatalogBindProjectsOK describes a response with status code 200, with default header values.

Success
*/
type CatalogBindProjectsOK struct {
	Payload models.Unit
}

func (o *CatalogBindProjectsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/bind-project][%d] catalogBindProjectsOK  %+v", 200, o.Payload)
}
func (o *CatalogBindProjectsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CatalogBindProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogBindProjectsBadRequest creates a CatalogBindProjectsBadRequest with default headers values
func NewCatalogBindProjectsBadRequest() *CatalogBindProjectsBadRequest {
	return &CatalogBindProjectsBadRequest{}
}

/* CatalogBindProjectsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogBindProjectsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CatalogBindProjectsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/bind-project][%d] catalogBindProjectsBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogBindProjectsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CatalogBindProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogBindProjectsUnauthorized creates a CatalogBindProjectsUnauthorized with default headers values
func NewCatalogBindProjectsUnauthorized() *CatalogBindProjectsUnauthorized {
	return &CatalogBindProjectsUnauthorized{}
}

/* CatalogBindProjectsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogBindProjectsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CatalogBindProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/bind-project][%d] catalogBindProjectsUnauthorized  %+v", 401, o.Payload)
}
func (o *CatalogBindProjectsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogBindProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogBindProjectsForbidden creates a CatalogBindProjectsForbidden with default headers values
func NewCatalogBindProjectsForbidden() *CatalogBindProjectsForbidden {
	return &CatalogBindProjectsForbidden{}
}

/* CatalogBindProjectsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogBindProjectsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CatalogBindProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/bind-project][%d] catalogBindProjectsForbidden  %+v", 403, o.Payload)
}
func (o *CatalogBindProjectsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogBindProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogBindProjectsNotFound creates a CatalogBindProjectsNotFound with default headers values
func NewCatalogBindProjectsNotFound() *CatalogBindProjectsNotFound {
	return &CatalogBindProjectsNotFound{}
}

/* CatalogBindProjectsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogBindProjectsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CatalogBindProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/bind-project][%d] catalogBindProjectsNotFound  %+v", 404, o.Payload)
}
func (o *CatalogBindProjectsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogBindProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogBindProjectsInternalServerError creates a CatalogBindProjectsInternalServerError with default headers values
func NewCatalogBindProjectsInternalServerError() *CatalogBindProjectsInternalServerError {
	return &CatalogBindProjectsInternalServerError{}
}

/* CatalogBindProjectsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogBindProjectsInternalServerError struct {
}

func (o *CatalogBindProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Catalog/bind-project][%d] catalogBindProjectsInternalServerError ", 500)
}

func (o *CatalogBindProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
