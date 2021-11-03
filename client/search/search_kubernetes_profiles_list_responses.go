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

// SearchKubernetesProfilesListReader is a Reader for the SearchKubernetesProfilesList structure.
type SearchKubernetesProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchKubernetesProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchKubernetesProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchKubernetesProfilesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchKubernetesProfilesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchKubernetesProfilesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchKubernetesProfilesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSearchKubernetesProfilesListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchKubernetesProfilesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchKubernetesProfilesListOK creates a SearchKubernetesProfilesListOK with default headers values
func NewSearchKubernetesProfilesListOK() *SearchKubernetesProfilesListOK {
	return &SearchKubernetesProfilesListOK{}
}

/* SearchKubernetesProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type SearchKubernetesProfilesListOK struct {
	Payload *models.KubernetesProfilesSearchList
}

func (o *SearchKubernetesProfilesListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListOK  %+v", 200, o.Payload)
}
func (o *SearchKubernetesProfilesListOK) GetPayload() *models.KubernetesProfilesSearchList {
	return o.Payload
}

func (o *SearchKubernetesProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesProfilesSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKubernetesProfilesListBadRequest creates a SearchKubernetesProfilesListBadRequest with default headers values
func NewSearchKubernetesProfilesListBadRequest() *SearchKubernetesProfilesListBadRequest {
	return &SearchKubernetesProfilesListBadRequest{}
}

/* SearchKubernetesProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchKubernetesProfilesListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SearchKubernetesProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListBadRequest  %+v", 400, o.Payload)
}
func (o *SearchKubernetesProfilesListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchKubernetesProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKubernetesProfilesListUnauthorized creates a SearchKubernetesProfilesListUnauthorized with default headers values
func NewSearchKubernetesProfilesListUnauthorized() *SearchKubernetesProfilesListUnauthorized {
	return &SearchKubernetesProfilesListUnauthorized{}
}

/* SearchKubernetesProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchKubernetesProfilesListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SearchKubernetesProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchKubernetesProfilesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchKubernetesProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKubernetesProfilesListForbidden creates a SearchKubernetesProfilesListForbidden with default headers values
func NewSearchKubernetesProfilesListForbidden() *SearchKubernetesProfilesListForbidden {
	return &SearchKubernetesProfilesListForbidden{}
}

/* SearchKubernetesProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchKubernetesProfilesListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SearchKubernetesProfilesListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListForbidden  %+v", 403, o.Payload)
}
func (o *SearchKubernetesProfilesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchKubernetesProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKubernetesProfilesListNotFound creates a SearchKubernetesProfilesListNotFound with default headers values
func NewSearchKubernetesProfilesListNotFound() *SearchKubernetesProfilesListNotFound {
	return &SearchKubernetesProfilesListNotFound{}
}

/* SearchKubernetesProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchKubernetesProfilesListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SearchKubernetesProfilesListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListNotFound  %+v", 404, o.Payload)
}
func (o *SearchKubernetesProfilesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchKubernetesProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKubernetesProfilesListTooManyRequests creates a SearchKubernetesProfilesListTooManyRequests with default headers values
func NewSearchKubernetesProfilesListTooManyRequests() *SearchKubernetesProfilesListTooManyRequests {
	return &SearchKubernetesProfilesListTooManyRequests{}
}

/* SearchKubernetesProfilesListTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type SearchKubernetesProfilesListTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *SearchKubernetesProfilesListTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListTooManyRequests  %+v", 429, o.Payload)
}
func (o *SearchKubernetesProfilesListTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchKubernetesProfilesListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKubernetesProfilesListInternalServerError creates a SearchKubernetesProfilesListInternalServerError with default headers values
func NewSearchKubernetesProfilesListInternalServerError() *SearchKubernetesProfilesListInternalServerError {
	return &SearchKubernetesProfilesListInternalServerError{}
}

/* SearchKubernetesProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchKubernetesProfilesListInternalServerError struct {
}

func (o *SearchKubernetesProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListInternalServerError ", 500)
}

func (o *SearchKubernetesProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
