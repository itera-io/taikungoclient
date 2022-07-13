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

// CatalogAvailablePackageDetailsReader is a Reader for the CatalogAvailablePackageDetails structure.
type CatalogAvailablePackageDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogAvailablePackageDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogAvailablePackageDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogAvailablePackageDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogAvailablePackageDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogAvailablePackageDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogAvailablePackageDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogAvailablePackageDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogAvailablePackageDetailsOK creates a CatalogAvailablePackageDetailsOK with default headers values
func NewCatalogAvailablePackageDetailsOK() *CatalogAvailablePackageDetailsOK {
	return &CatalogAvailablePackageDetailsOK{}
}

/* CatalogAvailablePackageDetailsOK describes a response with status code 200, with default header values.

Success
*/
type CatalogAvailablePackageDetailsOK struct {
	Payload *models.AvailablePackageDetailsDto
}

func (o *CatalogAvailablePackageDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/available/{repoName}/{packageName}][%d] catalogAvailablePackageDetailsOK  %+v", 200, o.Payload)
}
func (o *CatalogAvailablePackageDetailsOK) GetPayload() *models.AvailablePackageDetailsDto {
	return o.Payload
}

func (o *CatalogAvailablePackageDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailablePackageDetailsDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailablePackageDetailsBadRequest creates a CatalogAvailablePackageDetailsBadRequest with default headers values
func NewCatalogAvailablePackageDetailsBadRequest() *CatalogAvailablePackageDetailsBadRequest {
	return &CatalogAvailablePackageDetailsBadRequest{}
}

/* CatalogAvailablePackageDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogAvailablePackageDetailsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CatalogAvailablePackageDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/available/{repoName}/{packageName}][%d] catalogAvailablePackageDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogAvailablePackageDetailsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CatalogAvailablePackageDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailablePackageDetailsUnauthorized creates a CatalogAvailablePackageDetailsUnauthorized with default headers values
func NewCatalogAvailablePackageDetailsUnauthorized() *CatalogAvailablePackageDetailsUnauthorized {
	return &CatalogAvailablePackageDetailsUnauthorized{}
}

/* CatalogAvailablePackageDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogAvailablePackageDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CatalogAvailablePackageDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/available/{repoName}/{packageName}][%d] catalogAvailablePackageDetailsUnauthorized  %+v", 401, o.Payload)
}
func (o *CatalogAvailablePackageDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogAvailablePackageDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailablePackageDetailsForbidden creates a CatalogAvailablePackageDetailsForbidden with default headers values
func NewCatalogAvailablePackageDetailsForbidden() *CatalogAvailablePackageDetailsForbidden {
	return &CatalogAvailablePackageDetailsForbidden{}
}

/* CatalogAvailablePackageDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogAvailablePackageDetailsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CatalogAvailablePackageDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/available/{repoName}/{packageName}][%d] catalogAvailablePackageDetailsForbidden  %+v", 403, o.Payload)
}
func (o *CatalogAvailablePackageDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogAvailablePackageDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailablePackageDetailsNotFound creates a CatalogAvailablePackageDetailsNotFound with default headers values
func NewCatalogAvailablePackageDetailsNotFound() *CatalogAvailablePackageDetailsNotFound {
	return &CatalogAvailablePackageDetailsNotFound{}
}

/* CatalogAvailablePackageDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogAvailablePackageDetailsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CatalogAvailablePackageDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/available/{repoName}/{packageName}][%d] catalogAvailablePackageDetailsNotFound  %+v", 404, o.Payload)
}
func (o *CatalogAvailablePackageDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogAvailablePackageDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogAvailablePackageDetailsInternalServerError creates a CatalogAvailablePackageDetailsInternalServerError with default headers values
func NewCatalogAvailablePackageDetailsInternalServerError() *CatalogAvailablePackageDetailsInternalServerError {
	return &CatalogAvailablePackageDetailsInternalServerError{}
}

/* CatalogAvailablePackageDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogAvailablePackageDetailsInternalServerError struct {
}

func (o *CatalogAvailablePackageDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Catalog/available/{repoName}/{packageName}][%d] catalogAvailablePackageDetailsInternalServerError ", 500)
}

func (o *CatalogAvailablePackageDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
