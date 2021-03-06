// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SearchBillingCredentialsListReader is a Reader for the SearchBillingCredentialsList structure.
type SearchBillingCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchBillingCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchBillingCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchBillingCredentialsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchBillingCredentialsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchBillingCredentialsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchBillingCredentialsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchBillingCredentialsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchBillingCredentialsListOK creates a SearchBillingCredentialsListOK with default headers values
func NewSearchBillingCredentialsListOK() *SearchBillingCredentialsListOK {
	return &SearchBillingCredentialsListOK{}
}

/* SearchBillingCredentialsListOK describes a response with status code 200, with default header values.

Success
*/
type SearchBillingCredentialsListOK struct {
	Payload *models.BillingCredentialsSearchList
}

func (o *SearchBillingCredentialsListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/billing-credentials][%d] searchBillingCredentialsListOK  %+v", 200, o.Payload)
}
func (o *SearchBillingCredentialsListOK) GetPayload() *models.BillingCredentialsSearchList {
	return o.Payload
}

func (o *SearchBillingCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingCredentialsSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBillingCredentialsListBadRequest creates a SearchBillingCredentialsListBadRequest with default headers values
func NewSearchBillingCredentialsListBadRequest() *SearchBillingCredentialsListBadRequest {
	return &SearchBillingCredentialsListBadRequest{}
}

/* SearchBillingCredentialsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchBillingCredentialsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SearchBillingCredentialsListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/billing-credentials][%d] searchBillingCredentialsListBadRequest  %+v", 400, o.Payload)
}
func (o *SearchBillingCredentialsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchBillingCredentialsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBillingCredentialsListUnauthorized creates a SearchBillingCredentialsListUnauthorized with default headers values
func NewSearchBillingCredentialsListUnauthorized() *SearchBillingCredentialsListUnauthorized {
	return &SearchBillingCredentialsListUnauthorized{}
}

/* SearchBillingCredentialsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchBillingCredentialsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SearchBillingCredentialsListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/billing-credentials][%d] searchBillingCredentialsListUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchBillingCredentialsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchBillingCredentialsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBillingCredentialsListForbidden creates a SearchBillingCredentialsListForbidden with default headers values
func NewSearchBillingCredentialsListForbidden() *SearchBillingCredentialsListForbidden {
	return &SearchBillingCredentialsListForbidden{}
}

/* SearchBillingCredentialsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchBillingCredentialsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SearchBillingCredentialsListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/billing-credentials][%d] searchBillingCredentialsListForbidden  %+v", 403, o.Payload)
}
func (o *SearchBillingCredentialsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchBillingCredentialsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBillingCredentialsListNotFound creates a SearchBillingCredentialsListNotFound with default headers values
func NewSearchBillingCredentialsListNotFound() *SearchBillingCredentialsListNotFound {
	return &SearchBillingCredentialsListNotFound{}
}

/* SearchBillingCredentialsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchBillingCredentialsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SearchBillingCredentialsListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/billing-credentials][%d] searchBillingCredentialsListNotFound  %+v", 404, o.Payload)
}
func (o *SearchBillingCredentialsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchBillingCredentialsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBillingCredentialsListInternalServerError creates a SearchBillingCredentialsListInternalServerError with default headers values
func NewSearchBillingCredentialsListInternalServerError() *SearchBillingCredentialsListInternalServerError {
	return &SearchBillingCredentialsListInternalServerError{}
}

/* SearchBillingCredentialsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchBillingCredentialsListInternalServerError struct {
}

func (o *SearchBillingCredentialsListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/billing-credentials][%d] searchBillingCredentialsListInternalServerError ", 500)
}

func (o *SearchBillingCredentialsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
