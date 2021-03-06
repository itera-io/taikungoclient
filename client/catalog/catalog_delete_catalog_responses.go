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

// CatalogDeleteCatalogReader is a Reader for the CatalogDeleteCatalog structure.
type CatalogDeleteCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogDeleteCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogDeleteCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogDeleteCatalogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogDeleteCatalogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogDeleteCatalogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogDeleteCatalogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogDeleteCatalogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogDeleteCatalogOK creates a CatalogDeleteCatalogOK with default headers values
func NewCatalogDeleteCatalogOK() *CatalogDeleteCatalogOK {
	return &CatalogDeleteCatalogOK{}
}

/* CatalogDeleteCatalogOK describes a response with status code 200, with default header values.

Success
*/
type CatalogDeleteCatalogOK struct {
	Payload models.Unit
}

func (o *CatalogDeleteCatalogOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete/{id}][%d] catalogDeleteCatalogOK  %+v", 200, o.Payload)
}
func (o *CatalogDeleteCatalogOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CatalogDeleteCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogBadRequest creates a CatalogDeleteCatalogBadRequest with default headers values
func NewCatalogDeleteCatalogBadRequest() *CatalogDeleteCatalogBadRequest {
	return &CatalogDeleteCatalogBadRequest{}
}

/* CatalogDeleteCatalogBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogDeleteCatalogBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CatalogDeleteCatalogBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete/{id}][%d] catalogDeleteCatalogBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogDeleteCatalogBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CatalogDeleteCatalogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogUnauthorized creates a CatalogDeleteCatalogUnauthorized with default headers values
func NewCatalogDeleteCatalogUnauthorized() *CatalogDeleteCatalogUnauthorized {
	return &CatalogDeleteCatalogUnauthorized{}
}

/* CatalogDeleteCatalogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogDeleteCatalogUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CatalogDeleteCatalogUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete/{id}][%d] catalogDeleteCatalogUnauthorized  %+v", 401, o.Payload)
}
func (o *CatalogDeleteCatalogUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogDeleteCatalogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogForbidden creates a CatalogDeleteCatalogForbidden with default headers values
func NewCatalogDeleteCatalogForbidden() *CatalogDeleteCatalogForbidden {
	return &CatalogDeleteCatalogForbidden{}
}

/* CatalogDeleteCatalogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogDeleteCatalogForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CatalogDeleteCatalogForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete/{id}][%d] catalogDeleteCatalogForbidden  %+v", 403, o.Payload)
}
func (o *CatalogDeleteCatalogForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogDeleteCatalogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogNotFound creates a CatalogDeleteCatalogNotFound with default headers values
func NewCatalogDeleteCatalogNotFound() *CatalogDeleteCatalogNotFound {
	return &CatalogDeleteCatalogNotFound{}
}

/* CatalogDeleteCatalogNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogDeleteCatalogNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CatalogDeleteCatalogNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete/{id}][%d] catalogDeleteCatalogNotFound  %+v", 404, o.Payload)
}
func (o *CatalogDeleteCatalogNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogDeleteCatalogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogDeleteCatalogInternalServerError creates a CatalogDeleteCatalogInternalServerError with default headers values
func NewCatalogDeleteCatalogInternalServerError() *CatalogDeleteCatalogInternalServerError {
	return &CatalogDeleteCatalogInternalServerError{}
}

/* CatalogDeleteCatalogInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogDeleteCatalogInternalServerError struct {
}

func (o *CatalogDeleteCatalogInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Catalog/delete/{id}][%d] catalogDeleteCatalogInternalServerError ", 500)
}

func (o *CatalogDeleteCatalogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
