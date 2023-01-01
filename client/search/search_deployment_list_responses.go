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

// SearchDeploymentListReader is a Reader for the SearchDeploymentList structure.
type SearchDeploymentListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchDeploymentListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchDeploymentListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchDeploymentListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchDeploymentListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchDeploymentListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchDeploymentListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchDeploymentListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchDeploymentListOK creates a SearchDeploymentListOK with default headers values
func NewSearchDeploymentListOK() *SearchDeploymentListOK {
	return &SearchDeploymentListOK{}
}

/*
SearchDeploymentListOK describes a response with status code 200, with default header values.

Success
*/
type SearchDeploymentListOK struct {
	Payload *models.DeploymentSearchList
}

// IsSuccess returns true when this search deployment list o k response has a 2xx status code
func (o *SearchDeploymentListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search deployment list o k response has a 3xx status code
func (o *SearchDeploymentListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search deployment list o k response has a 4xx status code
func (o *SearchDeploymentListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search deployment list o k response has a 5xx status code
func (o *SearchDeploymentListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search deployment list o k response a status code equal to that given
func (o *SearchDeploymentListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchDeploymentListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListOK  %+v", 200, o.Payload)
}

func (o *SearchDeploymentListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListOK  %+v", 200, o.Payload)
}

func (o *SearchDeploymentListOK) GetPayload() *models.DeploymentSearchList {
	return o.Payload
}

func (o *SearchDeploymentListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDeploymentListBadRequest creates a SearchDeploymentListBadRequest with default headers values
func NewSearchDeploymentListBadRequest() *SearchDeploymentListBadRequest {
	return &SearchDeploymentListBadRequest{}
}

/*
SearchDeploymentListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchDeploymentListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search deployment list bad request response has a 2xx status code
func (o *SearchDeploymentListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search deployment list bad request response has a 3xx status code
func (o *SearchDeploymentListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search deployment list bad request response has a 4xx status code
func (o *SearchDeploymentListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search deployment list bad request response has a 5xx status code
func (o *SearchDeploymentListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search deployment list bad request response a status code equal to that given
func (o *SearchDeploymentListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchDeploymentListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchDeploymentListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchDeploymentListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchDeploymentListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDeploymentListUnauthorized creates a SearchDeploymentListUnauthorized with default headers values
func NewSearchDeploymentListUnauthorized() *SearchDeploymentListUnauthorized {
	return &SearchDeploymentListUnauthorized{}
}

/*
SearchDeploymentListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchDeploymentListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search deployment list unauthorized response has a 2xx status code
func (o *SearchDeploymentListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search deployment list unauthorized response has a 3xx status code
func (o *SearchDeploymentListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search deployment list unauthorized response has a 4xx status code
func (o *SearchDeploymentListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search deployment list unauthorized response has a 5xx status code
func (o *SearchDeploymentListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search deployment list unauthorized response a status code equal to that given
func (o *SearchDeploymentListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchDeploymentListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchDeploymentListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchDeploymentListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchDeploymentListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDeploymentListForbidden creates a SearchDeploymentListForbidden with default headers values
func NewSearchDeploymentListForbidden() *SearchDeploymentListForbidden {
	return &SearchDeploymentListForbidden{}
}

/*
SearchDeploymentListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchDeploymentListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search deployment list forbidden response has a 2xx status code
func (o *SearchDeploymentListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search deployment list forbidden response has a 3xx status code
func (o *SearchDeploymentListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search deployment list forbidden response has a 4xx status code
func (o *SearchDeploymentListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search deployment list forbidden response has a 5xx status code
func (o *SearchDeploymentListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search deployment list forbidden response a status code equal to that given
func (o *SearchDeploymentListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchDeploymentListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListForbidden  %+v", 403, o.Payload)
}

func (o *SearchDeploymentListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListForbidden  %+v", 403, o.Payload)
}

func (o *SearchDeploymentListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchDeploymentListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDeploymentListNotFound creates a SearchDeploymentListNotFound with default headers values
func NewSearchDeploymentListNotFound() *SearchDeploymentListNotFound {
	return &SearchDeploymentListNotFound{}
}

/*
SearchDeploymentListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchDeploymentListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search deployment list not found response has a 2xx status code
func (o *SearchDeploymentListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search deployment list not found response has a 3xx status code
func (o *SearchDeploymentListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search deployment list not found response has a 4xx status code
func (o *SearchDeploymentListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search deployment list not found response has a 5xx status code
func (o *SearchDeploymentListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search deployment list not found response a status code equal to that given
func (o *SearchDeploymentListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchDeploymentListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListNotFound  %+v", 404, o.Payload)
}

func (o *SearchDeploymentListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListNotFound  %+v", 404, o.Payload)
}

func (o *SearchDeploymentListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchDeploymentListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDeploymentListInternalServerError creates a SearchDeploymentListInternalServerError with default headers values
func NewSearchDeploymentListInternalServerError() *SearchDeploymentListInternalServerError {
	return &SearchDeploymentListInternalServerError{}
}

/*
SearchDeploymentListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchDeploymentListInternalServerError struct {
}

// IsSuccess returns true when this search deployment list internal server error response has a 2xx status code
func (o *SearchDeploymentListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search deployment list internal server error response has a 3xx status code
func (o *SearchDeploymentListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search deployment list internal server error response has a 4xx status code
func (o *SearchDeploymentListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search deployment list internal server error response has a 5xx status code
func (o *SearchDeploymentListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search deployment list internal server error response a status code equal to that given
func (o *SearchDeploymentListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchDeploymentListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListInternalServerError ", 500)
}

func (o *SearchDeploymentListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/deployments][%d] searchDeploymentListInternalServerError ", 500)
}

func (o *SearchDeploymentListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
