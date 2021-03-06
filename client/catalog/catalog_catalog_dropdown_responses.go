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

// CatalogCatalogDropdownReader is a Reader for the CatalogCatalogDropdown structure.
type CatalogCatalogDropdownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCatalogDropdownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCatalogDropdownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCatalogDropdownBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogCatalogDropdownUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogCatalogDropdownForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogCatalogDropdownNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogCatalogDropdownInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogCatalogDropdownOK creates a CatalogCatalogDropdownOK with default headers values
func NewCatalogCatalogDropdownOK() *CatalogCatalogDropdownOK {
	return &CatalogCatalogDropdownOK{}
}

/* CatalogCatalogDropdownOK describes a response with status code 200, with default header values.

Success
*/
type CatalogCatalogDropdownOK struct {
	Payload []*models.CatalogDropdownDto
}

func (o *CatalogCatalogDropdownOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/dropdown-list][%d] catalogCatalogDropdownOK  %+v", 200, o.Payload)
}
func (o *CatalogCatalogDropdownOK) GetPayload() []*models.CatalogDropdownDto {
	return o.Payload
}

func (o *CatalogCatalogDropdownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogDropdownBadRequest creates a CatalogCatalogDropdownBadRequest with default headers values
func NewCatalogCatalogDropdownBadRequest() *CatalogCatalogDropdownBadRequest {
	return &CatalogCatalogDropdownBadRequest{}
}

/* CatalogCatalogDropdownBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogCatalogDropdownBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CatalogCatalogDropdownBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/dropdown-list][%d] catalogCatalogDropdownBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogCatalogDropdownBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogDropdownBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogDropdownUnauthorized creates a CatalogCatalogDropdownUnauthorized with default headers values
func NewCatalogCatalogDropdownUnauthorized() *CatalogCatalogDropdownUnauthorized {
	return &CatalogCatalogDropdownUnauthorized{}
}

/* CatalogCatalogDropdownUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogCatalogDropdownUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CatalogCatalogDropdownUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/dropdown-list][%d] catalogCatalogDropdownUnauthorized  %+v", 401, o.Payload)
}
func (o *CatalogCatalogDropdownUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogDropdownUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogDropdownForbidden creates a CatalogCatalogDropdownForbidden with default headers values
func NewCatalogCatalogDropdownForbidden() *CatalogCatalogDropdownForbidden {
	return &CatalogCatalogDropdownForbidden{}
}

/* CatalogCatalogDropdownForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogCatalogDropdownForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CatalogCatalogDropdownForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/dropdown-list][%d] catalogCatalogDropdownForbidden  %+v", 403, o.Payload)
}
func (o *CatalogCatalogDropdownForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogDropdownForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogDropdownNotFound creates a CatalogCatalogDropdownNotFound with default headers values
func NewCatalogCatalogDropdownNotFound() *CatalogCatalogDropdownNotFound {
	return &CatalogCatalogDropdownNotFound{}
}

/* CatalogCatalogDropdownNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogCatalogDropdownNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CatalogCatalogDropdownNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/dropdown-list][%d] catalogCatalogDropdownNotFound  %+v", 404, o.Payload)
}
func (o *CatalogCatalogDropdownNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogCatalogDropdownNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCatalogDropdownInternalServerError creates a CatalogCatalogDropdownInternalServerError with default headers values
func NewCatalogCatalogDropdownInternalServerError() *CatalogCatalogDropdownInternalServerError {
	return &CatalogCatalogDropdownInternalServerError{}
}

/* CatalogCatalogDropdownInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogCatalogDropdownInternalServerError struct {
}

func (o *CatalogCatalogDropdownInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/dropdown-list][%d] catalogCatalogDropdownInternalServerError ", 500)
}

func (o *CatalogCatalogDropdownInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
