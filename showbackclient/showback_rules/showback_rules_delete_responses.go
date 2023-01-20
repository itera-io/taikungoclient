// Code generated by go-swagger; DO NOT EDIT.

package showback_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackRulesDeleteReader is a Reader for the ShowbackRulesDelete structure.
type ShowbackRulesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackRulesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackRulesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewShowbackRulesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackRulesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackRulesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackRulesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackRulesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackRulesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackRulesDeleteOK creates a ShowbackRulesDeleteOK with default headers values
func NewShowbackRulesDeleteOK() *ShowbackRulesDeleteOK {
	return &ShowbackRulesDeleteOK{}
}

/*
ShowbackRulesDeleteOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackRulesDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this showback rules delete o k response has a 2xx status code
func (o *ShowbackRulesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback rules delete o k response has a 3xx status code
func (o *ShowbackRulesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules delete o k response has a 4xx status code
func (o *ShowbackRulesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules delete o k response has a 5xx status code
func (o *ShowbackRulesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules delete o k response a status code equal to that given
func (o *ShowbackRulesDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the showback rules delete o k response
func (o *ShowbackRulesDeleteOK) Code() int {
	return 200
}

func (o *ShowbackRulesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ShowbackRulesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesDeleteNoContent creates a ShowbackRulesDeleteNoContent with default headers values
func NewShowbackRulesDeleteNoContent() *ShowbackRulesDeleteNoContent {
	return &ShowbackRulesDeleteNoContent{}
}

/*
ShowbackRulesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ShowbackRulesDeleteNoContent struct {
}

// IsSuccess returns true when this showback rules delete no content response has a 2xx status code
func (o *ShowbackRulesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback rules delete no content response has a 3xx status code
func (o *ShowbackRulesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules delete no content response has a 4xx status code
func (o *ShowbackRulesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules delete no content response has a 5xx status code
func (o *ShowbackRulesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules delete no content response a status code equal to that given
func (o *ShowbackRulesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the showback rules delete no content response
func (o *ShowbackRulesDeleteNoContent) Code() int {
	return 204
}

func (o *ShowbackRulesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteNoContent ", 204)
}

func (o *ShowbackRulesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteNoContent ", 204)
}

func (o *ShowbackRulesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowbackRulesDeleteBadRequest creates a ShowbackRulesDeleteBadRequest with default headers values
func NewShowbackRulesDeleteBadRequest() *ShowbackRulesDeleteBadRequest {
	return &ShowbackRulesDeleteBadRequest{}
}

/*
ShowbackRulesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackRulesDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this showback rules delete bad request response has a 2xx status code
func (o *ShowbackRulesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules delete bad request response has a 3xx status code
func (o *ShowbackRulesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules delete bad request response has a 4xx status code
func (o *ShowbackRulesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules delete bad request response has a 5xx status code
func (o *ShowbackRulesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules delete bad request response a status code equal to that given
func (o *ShowbackRulesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the showback rules delete bad request response
func (o *ShowbackRulesDeleteBadRequest) Code() int {
	return 400
}

func (o *ShowbackRulesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackRulesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesDeleteUnauthorized creates a ShowbackRulesDeleteUnauthorized with default headers values
func NewShowbackRulesDeleteUnauthorized() *ShowbackRulesDeleteUnauthorized {
	return &ShowbackRulesDeleteUnauthorized{}
}

/*
ShowbackRulesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackRulesDeleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this showback rules delete unauthorized response has a 2xx status code
func (o *ShowbackRulesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules delete unauthorized response has a 3xx status code
func (o *ShowbackRulesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules delete unauthorized response has a 4xx status code
func (o *ShowbackRulesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules delete unauthorized response has a 5xx status code
func (o *ShowbackRulesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules delete unauthorized response a status code equal to that given
func (o *ShowbackRulesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the showback rules delete unauthorized response
func (o *ShowbackRulesDeleteUnauthorized) Code() int {
	return 401
}

func (o *ShowbackRulesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesDeleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackRulesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesDeleteForbidden creates a ShowbackRulesDeleteForbidden with default headers values
func NewShowbackRulesDeleteForbidden() *ShowbackRulesDeleteForbidden {
	return &ShowbackRulesDeleteForbidden{}
}

/*
ShowbackRulesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackRulesDeleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this showback rules delete forbidden response has a 2xx status code
func (o *ShowbackRulesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules delete forbidden response has a 3xx status code
func (o *ShowbackRulesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules delete forbidden response has a 4xx status code
func (o *ShowbackRulesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules delete forbidden response has a 5xx status code
func (o *ShowbackRulesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules delete forbidden response a status code equal to that given
func (o *ShowbackRulesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the showback rules delete forbidden response
func (o *ShowbackRulesDeleteForbidden) Code() int {
	return 403
}

func (o *ShowbackRulesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesDeleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackRulesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesDeleteNotFound creates a ShowbackRulesDeleteNotFound with default headers values
func NewShowbackRulesDeleteNotFound() *ShowbackRulesDeleteNotFound {
	return &ShowbackRulesDeleteNotFound{}
}

/*
ShowbackRulesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackRulesDeleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this showback rules delete not found response has a 2xx status code
func (o *ShowbackRulesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules delete not found response has a 3xx status code
func (o *ShowbackRulesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules delete not found response has a 4xx status code
func (o *ShowbackRulesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules delete not found response has a 5xx status code
func (o *ShowbackRulesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules delete not found response a status code equal to that given
func (o *ShowbackRulesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the showback rules delete not found response
func (o *ShowbackRulesDeleteNotFound) Code() int {
	return 404
}

func (o *ShowbackRulesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackRulesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesDeleteInternalServerError creates a ShowbackRulesDeleteInternalServerError with default headers values
func NewShowbackRulesDeleteInternalServerError() *ShowbackRulesDeleteInternalServerError {
	return &ShowbackRulesDeleteInternalServerError{}
}

/*
ShowbackRulesDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackRulesDeleteInternalServerError struct {
}

// IsSuccess returns true when this showback rules delete internal server error response has a 2xx status code
func (o *ShowbackRulesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules delete internal server error response has a 3xx status code
func (o *ShowbackRulesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules delete internal server error response has a 4xx status code
func (o *ShowbackRulesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules delete internal server error response has a 5xx status code
func (o *ShowbackRulesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback rules delete internal server error response a status code equal to that given
func (o *ShowbackRulesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the showback rules delete internal server error response
func (o *ShowbackRulesDeleteInternalServerError) Code() int {
	return 500
}

func (o *ShowbackRulesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteInternalServerError ", 500)
}

func (o *ShowbackRulesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /showback/v{v}/ShowbackRules/{id}][%d] showbackRulesDeleteInternalServerError ", 500)
}

func (o *ShowbackRulesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
