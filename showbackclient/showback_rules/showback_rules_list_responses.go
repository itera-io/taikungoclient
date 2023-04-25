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

// ShowbackRulesListReader is a Reader for the ShowbackRulesList structure.
type ShowbackRulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackRulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackRulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackRulesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackRulesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackRulesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackRulesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackRulesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackRulesListOK creates a ShowbackRulesListOK with default headers values
func NewShowbackRulesListOK() *ShowbackRulesListOK {
	return &ShowbackRulesListOK{}
}

/*
ShowbackRulesListOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackRulesListOK struct {
	Payload *models.ShowbackRuleList
}

// IsSuccess returns true when this showback rules list o k response has a 2xx status code
func (o *ShowbackRulesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback rules list o k response has a 3xx status code
func (o *ShowbackRulesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules list o k response has a 4xx status code
func (o *ShowbackRulesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules list o k response has a 5xx status code
func (o *ShowbackRulesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules list o k response a status code equal to that given
func (o *ShowbackRulesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the showback rules list o k response
func (o *ShowbackRulesListOK) Code() int {
	return 200
}

func (o *ShowbackRulesListOK) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesListOK) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListOK  %+v", 200, o.Payload)
}

func (o *ShowbackRulesListOK) GetPayload() *models.ShowbackRuleList {
	return o.Payload
}

func (o *ShowbackRulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShowbackRuleList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListBadRequest creates a ShowbackRulesListBadRequest with default headers values
func NewShowbackRulesListBadRequest() *ShowbackRulesListBadRequest {
	return &ShowbackRulesListBadRequest{}
}

/*
ShowbackRulesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackRulesListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this showback rules list bad request response has a 2xx status code
func (o *ShowbackRulesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules list bad request response has a 3xx status code
func (o *ShowbackRulesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules list bad request response has a 4xx status code
func (o *ShowbackRulesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules list bad request response has a 5xx status code
func (o *ShowbackRulesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules list bad request response a status code equal to that given
func (o *ShowbackRulesListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the showback rules list bad request response
func (o *ShowbackRulesListBadRequest) Code() int {
	return 400
}

func (o *ShowbackRulesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesListBadRequest) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackRulesListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackRulesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListUnauthorized creates a ShowbackRulesListUnauthorized with default headers values
func NewShowbackRulesListUnauthorized() *ShowbackRulesListUnauthorized {
	return &ShowbackRulesListUnauthorized{}
}

/*
ShowbackRulesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackRulesListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this showback rules list unauthorized response has a 2xx status code
func (o *ShowbackRulesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules list unauthorized response has a 3xx status code
func (o *ShowbackRulesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules list unauthorized response has a 4xx status code
func (o *ShowbackRulesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules list unauthorized response has a 5xx status code
func (o *ShowbackRulesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules list unauthorized response a status code equal to that given
func (o *ShowbackRulesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the showback rules list unauthorized response
func (o *ShowbackRulesListUnauthorized) Code() int {
	return 401
}

func (o *ShowbackRulesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesListUnauthorized) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackRulesListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackRulesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListForbidden creates a ShowbackRulesListForbidden with default headers values
func NewShowbackRulesListForbidden() *ShowbackRulesListForbidden {
	return &ShowbackRulesListForbidden{}
}

/*
ShowbackRulesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackRulesListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this showback rules list forbidden response has a 2xx status code
func (o *ShowbackRulesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules list forbidden response has a 3xx status code
func (o *ShowbackRulesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules list forbidden response has a 4xx status code
func (o *ShowbackRulesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules list forbidden response has a 5xx status code
func (o *ShowbackRulesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules list forbidden response a status code equal to that given
func (o *ShowbackRulesListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the showback rules list forbidden response
func (o *ShowbackRulesListForbidden) Code() int {
	return 403
}

func (o *ShowbackRulesListForbidden) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesListForbidden) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackRulesListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackRulesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListNotFound creates a ShowbackRulesListNotFound with default headers values
func NewShowbackRulesListNotFound() *ShowbackRulesListNotFound {
	return &ShowbackRulesListNotFound{}
}

/*
ShowbackRulesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackRulesListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this showback rules list not found response has a 2xx status code
func (o *ShowbackRulesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules list not found response has a 3xx status code
func (o *ShowbackRulesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules list not found response has a 4xx status code
func (o *ShowbackRulesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback rules list not found response has a 5xx status code
func (o *ShowbackRulesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback rules list not found response a status code equal to that given
func (o *ShowbackRulesListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the showback rules list not found response
func (o *ShowbackRulesListNotFound) Code() int {
	return 404
}

func (o *ShowbackRulesListNotFound) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesListNotFound) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackRulesListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackRulesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListInternalServerError creates a ShowbackRulesListInternalServerError with default headers values
func NewShowbackRulesListInternalServerError() *ShowbackRulesListInternalServerError {
	return &ShowbackRulesListInternalServerError{}
}

/*
ShowbackRulesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackRulesListInternalServerError struct {
}

// IsSuccess returns true when this showback rules list internal server error response has a 2xx status code
func (o *ShowbackRulesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback rules list internal server error response has a 3xx status code
func (o *ShowbackRulesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback rules list internal server error response has a 4xx status code
func (o *ShowbackRulesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback rules list internal server error response has a 5xx status code
func (o *ShowbackRulesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback rules list internal server error response a status code equal to that given
func (o *ShowbackRulesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the showback rules list internal server error response
func (o *ShowbackRulesListInternalServerError) Code() int {
	return 500
}

func (o *ShowbackRulesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListInternalServerError ", 500)
}

func (o *ShowbackRulesListInternalServerError) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackRules][%d] showbackRulesListInternalServerError ", 500)
}

func (o *ShowbackRulesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
