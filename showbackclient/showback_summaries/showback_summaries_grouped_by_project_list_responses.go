// Code generated by go-swagger; DO NOT EDIT.

package showback_summaries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackSummariesGroupedByProjectListReader is a Reader for the ShowbackSummariesGroupedByProjectList structure.
type ShowbackSummariesGroupedByProjectListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackSummariesGroupedByProjectListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackSummariesGroupedByProjectListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackSummariesGroupedByProjectListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackSummariesGroupedByProjectListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackSummariesGroupedByProjectListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackSummariesGroupedByProjectListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackSummariesGroupedByProjectListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackSummariesGroupedByProjectListOK creates a ShowbackSummariesGroupedByProjectListOK with default headers values
func NewShowbackSummariesGroupedByProjectListOK() *ShowbackSummariesGroupedByProjectListOK {
	return &ShowbackSummariesGroupedByProjectListOK{}
}

/*
ShowbackSummariesGroupedByProjectListOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackSummariesGroupedByProjectListOK struct {
	Payload *models.GroupedShowbackByProjectList
}

// IsSuccess returns true when this showback summaries grouped by project list o k response has a 2xx status code
func (o *ShowbackSummariesGroupedByProjectListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback summaries grouped by project list o k response has a 3xx status code
func (o *ShowbackSummariesGroupedByProjectListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped by project list o k response has a 4xx status code
func (o *ShowbackSummariesGroupedByProjectListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries grouped by project list o k response has a 5xx status code
func (o *ShowbackSummariesGroupedByProjectListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped by project list o k response a status code equal to that given
func (o *ShowbackSummariesGroupedByProjectListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ShowbackSummariesGroupedByProjectListOK) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListOK  %+v", 200, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListOK) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListOK  %+v", 200, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListOK) GetPayload() *models.GroupedShowbackByProjectList {
	return o.Payload
}

func (o *ShowbackSummariesGroupedByProjectListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupedShowbackByProjectList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedByProjectListBadRequest creates a ShowbackSummariesGroupedByProjectListBadRequest with default headers values
func NewShowbackSummariesGroupedByProjectListBadRequest() *ShowbackSummariesGroupedByProjectListBadRequest {
	return &ShowbackSummariesGroupedByProjectListBadRequest{}
}

/*
ShowbackSummariesGroupedByProjectListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackSummariesGroupedByProjectListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this showback summaries grouped by project list bad request response has a 2xx status code
func (o *ShowbackSummariesGroupedByProjectListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped by project list bad request response has a 3xx status code
func (o *ShowbackSummariesGroupedByProjectListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped by project list bad request response has a 4xx status code
func (o *ShowbackSummariesGroupedByProjectListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped by project list bad request response has a 5xx status code
func (o *ShowbackSummariesGroupedByProjectListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped by project list bad request response a status code equal to that given
func (o *ShowbackSummariesGroupedByProjectListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ShowbackSummariesGroupedByProjectListBadRequest) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListBadRequest) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackSummariesGroupedByProjectListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedByProjectListUnauthorized creates a ShowbackSummariesGroupedByProjectListUnauthorized with default headers values
func NewShowbackSummariesGroupedByProjectListUnauthorized() *ShowbackSummariesGroupedByProjectListUnauthorized {
	return &ShowbackSummariesGroupedByProjectListUnauthorized{}
}

/*
ShowbackSummariesGroupedByProjectListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackSummariesGroupedByProjectListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this showback summaries grouped by project list unauthorized response has a 2xx status code
func (o *ShowbackSummariesGroupedByProjectListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped by project list unauthorized response has a 3xx status code
func (o *ShowbackSummariesGroupedByProjectListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped by project list unauthorized response has a 4xx status code
func (o *ShowbackSummariesGroupedByProjectListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped by project list unauthorized response has a 5xx status code
func (o *ShowbackSummariesGroupedByProjectListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped by project list unauthorized response a status code equal to that given
func (o *ShowbackSummariesGroupedByProjectListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ShowbackSummariesGroupedByProjectListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListUnauthorized) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ShowbackSummariesGroupedByProjectListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedByProjectListForbidden creates a ShowbackSummariesGroupedByProjectListForbidden with default headers values
func NewShowbackSummariesGroupedByProjectListForbidden() *ShowbackSummariesGroupedByProjectListForbidden {
	return &ShowbackSummariesGroupedByProjectListForbidden{}
}

/*
ShowbackSummariesGroupedByProjectListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackSummariesGroupedByProjectListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this showback summaries grouped by project list forbidden response has a 2xx status code
func (o *ShowbackSummariesGroupedByProjectListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped by project list forbidden response has a 3xx status code
func (o *ShowbackSummariesGroupedByProjectListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped by project list forbidden response has a 4xx status code
func (o *ShowbackSummariesGroupedByProjectListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped by project list forbidden response has a 5xx status code
func (o *ShowbackSummariesGroupedByProjectListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped by project list forbidden response a status code equal to that given
func (o *ShowbackSummariesGroupedByProjectListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ShowbackSummariesGroupedByProjectListForbidden) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListForbidden) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ShowbackSummariesGroupedByProjectListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedByProjectListNotFound creates a ShowbackSummariesGroupedByProjectListNotFound with default headers values
func NewShowbackSummariesGroupedByProjectListNotFound() *ShowbackSummariesGroupedByProjectListNotFound {
	return &ShowbackSummariesGroupedByProjectListNotFound{}
}

/*
ShowbackSummariesGroupedByProjectListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackSummariesGroupedByProjectListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this showback summaries grouped by project list not found response has a 2xx status code
func (o *ShowbackSummariesGroupedByProjectListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped by project list not found response has a 3xx status code
func (o *ShowbackSummariesGroupedByProjectListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped by project list not found response has a 4xx status code
func (o *ShowbackSummariesGroupedByProjectListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped by project list not found response has a 5xx status code
func (o *ShowbackSummariesGroupedByProjectListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped by project list not found response a status code equal to that given
func (o *ShowbackSummariesGroupedByProjectListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ShowbackSummariesGroupedByProjectListNotFound) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListNotFound) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesGroupedByProjectListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ShowbackSummariesGroupedByProjectListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedByProjectListInternalServerError creates a ShowbackSummariesGroupedByProjectListInternalServerError with default headers values
func NewShowbackSummariesGroupedByProjectListInternalServerError() *ShowbackSummariesGroupedByProjectListInternalServerError {
	return &ShowbackSummariesGroupedByProjectListInternalServerError{}
}

/*
ShowbackSummariesGroupedByProjectListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackSummariesGroupedByProjectListInternalServerError struct {
}

// IsSuccess returns true when this showback summaries grouped by project list internal server error response has a 2xx status code
func (o *ShowbackSummariesGroupedByProjectListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped by project list internal server error response has a 3xx status code
func (o *ShowbackSummariesGroupedByProjectListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped by project list internal server error response has a 4xx status code
func (o *ShowbackSummariesGroupedByProjectListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries grouped by project list internal server error response has a 5xx status code
func (o *ShowbackSummariesGroupedByProjectListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback summaries grouped by project list internal server error response a status code equal to that given
func (o *ShowbackSummariesGroupedByProjectListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ShowbackSummariesGroupedByProjectListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListInternalServerError ", 500)
}

func (o *ShowbackSummariesGroupedByProjectListInternalServerError) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/byProject][%d] showbackSummariesGroupedByProjectListInternalServerError ", 500)
}

func (o *ShowbackSummariesGroupedByProjectListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
