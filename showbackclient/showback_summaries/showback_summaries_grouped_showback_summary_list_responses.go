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

// ShowbackSummariesGroupedShowbackSummaryListReader is a Reader for the ShowbackSummariesGroupedShowbackSummaryList structure.
type ShowbackSummariesGroupedShowbackSummaryListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackSummariesGroupedShowbackSummaryListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackSummariesGroupedShowbackSummaryListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackSummariesGroupedShowbackSummaryListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackSummariesGroupedShowbackSummaryListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackSummariesGroupedShowbackSummaryListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackSummariesGroupedShowbackSummaryListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackSummariesGroupedShowbackSummaryListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackSummariesGroupedShowbackSummaryListOK creates a ShowbackSummariesGroupedShowbackSummaryListOK with default headers values
func NewShowbackSummariesGroupedShowbackSummaryListOK() *ShowbackSummariesGroupedShowbackSummaryListOK {
	return &ShowbackSummariesGroupedShowbackSummaryListOK{}
}

/*
ShowbackSummariesGroupedShowbackSummaryListOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackSummariesGroupedShowbackSummaryListOK struct {
	Payload []*models.GroupedShowbackSummaryListDto
}

// IsSuccess returns true when this showback summaries grouped showback summary list o k response has a 2xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback summaries grouped showback summary list o k response has a 3xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped showback summary list o k response has a 4xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries grouped showback summary list o k response has a 5xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped showback summary list o k response a status code equal to that given
func (o *ShowbackSummariesGroupedShowbackSummaryListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the showback summaries grouped showback summary list o k response
func (o *ShowbackSummariesGroupedShowbackSummaryListOK) Code() int {
	return 200
}

func (o *ShowbackSummariesGroupedShowbackSummaryListOK) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListOK  %+v", 200, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListOK) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListOK  %+v", 200, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListOK) GetPayload() []*models.GroupedShowbackSummaryListDto {
	return o.Payload
}

func (o *ShowbackSummariesGroupedShowbackSummaryListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedShowbackSummaryListBadRequest creates a ShowbackSummariesGroupedShowbackSummaryListBadRequest with default headers values
func NewShowbackSummariesGroupedShowbackSummaryListBadRequest() *ShowbackSummariesGroupedShowbackSummaryListBadRequest {
	return &ShowbackSummariesGroupedShowbackSummaryListBadRequest{}
}

/*
ShowbackSummariesGroupedShowbackSummaryListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackSummariesGroupedShowbackSummaryListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this showback summaries grouped showback summary list bad request response has a 2xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped showback summary list bad request response has a 3xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped showback summary list bad request response has a 4xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped showback summary list bad request response has a 5xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped showback summary list bad request response a status code equal to that given
func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the showback summaries grouped showback summary list bad request response
func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) Code() int {
	return 400
}

func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackSummariesGroupedShowbackSummaryListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedShowbackSummaryListUnauthorized creates a ShowbackSummariesGroupedShowbackSummaryListUnauthorized with default headers values
func NewShowbackSummariesGroupedShowbackSummaryListUnauthorized() *ShowbackSummariesGroupedShowbackSummaryListUnauthorized {
	return &ShowbackSummariesGroupedShowbackSummaryListUnauthorized{}
}

/*
ShowbackSummariesGroupedShowbackSummaryListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackSummariesGroupedShowbackSummaryListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this showback summaries grouped showback summary list unauthorized response has a 2xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped showback summary list unauthorized response has a 3xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped showback summary list unauthorized response has a 4xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped showback summary list unauthorized response has a 5xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped showback summary list unauthorized response a status code equal to that given
func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the showback summaries grouped showback summary list unauthorized response
func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) Code() int {
	return 401
}

func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackSummariesGroupedShowbackSummaryListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedShowbackSummaryListForbidden creates a ShowbackSummariesGroupedShowbackSummaryListForbidden with default headers values
func NewShowbackSummariesGroupedShowbackSummaryListForbidden() *ShowbackSummariesGroupedShowbackSummaryListForbidden {
	return &ShowbackSummariesGroupedShowbackSummaryListForbidden{}
}

/*
ShowbackSummariesGroupedShowbackSummaryListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackSummariesGroupedShowbackSummaryListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this showback summaries grouped showback summary list forbidden response has a 2xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped showback summary list forbidden response has a 3xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped showback summary list forbidden response has a 4xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped showback summary list forbidden response has a 5xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped showback summary list forbidden response a status code equal to that given
func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the showback summaries grouped showback summary list forbidden response
func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) Code() int {
	return 403
}

func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackSummariesGroupedShowbackSummaryListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedShowbackSummaryListNotFound creates a ShowbackSummariesGroupedShowbackSummaryListNotFound with default headers values
func NewShowbackSummariesGroupedShowbackSummaryListNotFound() *ShowbackSummariesGroupedShowbackSummaryListNotFound {
	return &ShowbackSummariesGroupedShowbackSummaryListNotFound{}
}

/*
ShowbackSummariesGroupedShowbackSummaryListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackSummariesGroupedShowbackSummaryListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this showback summaries grouped showback summary list not found response has a 2xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped showback summary list not found response has a 3xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped showback summary list not found response has a 4xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries grouped showback summary list not found response has a 5xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries grouped showback summary list not found response a status code equal to that given
func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the showback summaries grouped showback summary list not found response
func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) Code() int {
	return 404
}

func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackSummariesGroupedShowbackSummaryListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesGroupedShowbackSummaryListInternalServerError creates a ShowbackSummariesGroupedShowbackSummaryListInternalServerError with default headers values
func NewShowbackSummariesGroupedShowbackSummaryListInternalServerError() *ShowbackSummariesGroupedShowbackSummaryListInternalServerError {
	return &ShowbackSummariesGroupedShowbackSummaryListInternalServerError{}
}

/*
ShowbackSummariesGroupedShowbackSummaryListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackSummariesGroupedShowbackSummaryListInternalServerError struct {
}

// IsSuccess returns true when this showback summaries grouped showback summary list internal server error response has a 2xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries grouped showback summary list internal server error response has a 3xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries grouped showback summary list internal server error response has a 4xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries grouped showback summary list internal server error response has a 5xx status code
func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback summaries grouped showback summary list internal server error response a status code equal to that given
func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the showback summaries grouped showback summary list internal server error response
func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) Code() int {
	return 500
}

func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListInternalServerError ", 500)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/grouped/list][%d] showbackSummariesGroupedShowbackSummaryListInternalServerError ", 500)
}

func (o *ShowbackSummariesGroupedShowbackSummaryListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
