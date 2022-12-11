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

// CatalogEditCatalogAppParamsReader is a Reader for the CatalogEditCatalogAppParams structure.
type CatalogEditCatalogAppParamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogEditCatalogAppParamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogEditCatalogAppParamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogEditCatalogAppParamsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogEditCatalogAppParamsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogEditCatalogAppParamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogEditCatalogAppParamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCatalogEditCatalogAppParamsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCatalogEditCatalogAppParamsOK creates a CatalogEditCatalogAppParamsOK with default headers values
func NewCatalogEditCatalogAppParamsOK() *CatalogEditCatalogAppParamsOK {
	return &CatalogEditCatalogAppParamsOK{}
}

/*
CatalogEditCatalogAppParamsOK describes a response with status code 200, with default header values.

Success
*/
type CatalogEditCatalogAppParamsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this catalog edit catalog app params o k response has a 2xx status code
func (o *CatalogEditCatalogAppParamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog edit catalog app params o k response has a 3xx status code
func (o *CatalogEditCatalogAppParamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog edit catalog app params o k response has a 4xx status code
func (o *CatalogEditCatalogAppParamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog edit catalog app params o k response has a 5xx status code
func (o *CatalogEditCatalogAppParamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog edit catalog app params o k response a status code equal to that given
func (o *CatalogEditCatalogAppParamsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CatalogEditCatalogAppParamsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsOK  %+v", 200, o.Payload)
}

func (o *CatalogEditCatalogAppParamsOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsOK  %+v", 200, o.Payload)
}

func (o *CatalogEditCatalogAppParamsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CatalogEditCatalogAppParamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogEditCatalogAppParamsBadRequest creates a CatalogEditCatalogAppParamsBadRequest with default headers values
func NewCatalogEditCatalogAppParamsBadRequest() *CatalogEditCatalogAppParamsBadRequest {
	return &CatalogEditCatalogAppParamsBadRequest{}
}

/*
CatalogEditCatalogAppParamsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogEditCatalogAppParamsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this catalog edit catalog app params bad request response has a 2xx status code
func (o *CatalogEditCatalogAppParamsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog edit catalog app params bad request response has a 3xx status code
func (o *CatalogEditCatalogAppParamsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog edit catalog app params bad request response has a 4xx status code
func (o *CatalogEditCatalogAppParamsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog edit catalog app params bad request response has a 5xx status code
func (o *CatalogEditCatalogAppParamsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog edit catalog app params bad request response a status code equal to that given
func (o *CatalogEditCatalogAppParamsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CatalogEditCatalogAppParamsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogEditCatalogAppParamsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogEditCatalogAppParamsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CatalogEditCatalogAppParamsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogEditCatalogAppParamsUnauthorized creates a CatalogEditCatalogAppParamsUnauthorized with default headers values
func NewCatalogEditCatalogAppParamsUnauthorized() *CatalogEditCatalogAppParamsUnauthorized {
	return &CatalogEditCatalogAppParamsUnauthorized{}
}

/*
CatalogEditCatalogAppParamsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogEditCatalogAppParamsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog edit catalog app params unauthorized response has a 2xx status code
func (o *CatalogEditCatalogAppParamsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog edit catalog app params unauthorized response has a 3xx status code
func (o *CatalogEditCatalogAppParamsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog edit catalog app params unauthorized response has a 4xx status code
func (o *CatalogEditCatalogAppParamsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog edit catalog app params unauthorized response has a 5xx status code
func (o *CatalogEditCatalogAppParamsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog edit catalog app params unauthorized response a status code equal to that given
func (o *CatalogEditCatalogAppParamsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CatalogEditCatalogAppParamsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogEditCatalogAppParamsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogEditCatalogAppParamsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogEditCatalogAppParamsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogEditCatalogAppParamsForbidden creates a CatalogEditCatalogAppParamsForbidden with default headers values
func NewCatalogEditCatalogAppParamsForbidden() *CatalogEditCatalogAppParamsForbidden {
	return &CatalogEditCatalogAppParamsForbidden{}
}

/*
CatalogEditCatalogAppParamsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogEditCatalogAppParamsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog edit catalog app params forbidden response has a 2xx status code
func (o *CatalogEditCatalogAppParamsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog edit catalog app params forbidden response has a 3xx status code
func (o *CatalogEditCatalogAppParamsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog edit catalog app params forbidden response has a 4xx status code
func (o *CatalogEditCatalogAppParamsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog edit catalog app params forbidden response has a 5xx status code
func (o *CatalogEditCatalogAppParamsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog edit catalog app params forbidden response a status code equal to that given
func (o *CatalogEditCatalogAppParamsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CatalogEditCatalogAppParamsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsForbidden  %+v", 403, o.Payload)
}

func (o *CatalogEditCatalogAppParamsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsForbidden  %+v", 403, o.Payload)
}

func (o *CatalogEditCatalogAppParamsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogEditCatalogAppParamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogEditCatalogAppParamsNotFound creates a CatalogEditCatalogAppParamsNotFound with default headers values
func NewCatalogEditCatalogAppParamsNotFound() *CatalogEditCatalogAppParamsNotFound {
	return &CatalogEditCatalogAppParamsNotFound{}
}

/*
CatalogEditCatalogAppParamsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogEditCatalogAppParamsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this catalog edit catalog app params not found response has a 2xx status code
func (o *CatalogEditCatalogAppParamsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog edit catalog app params not found response has a 3xx status code
func (o *CatalogEditCatalogAppParamsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog edit catalog app params not found response has a 4xx status code
func (o *CatalogEditCatalogAppParamsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog edit catalog app params not found response has a 5xx status code
func (o *CatalogEditCatalogAppParamsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog edit catalog app params not found response a status code equal to that given
func (o *CatalogEditCatalogAppParamsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CatalogEditCatalogAppParamsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsNotFound  %+v", 404, o.Payload)
}

func (o *CatalogEditCatalogAppParamsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsNotFound  %+v", 404, o.Payload)
}

func (o *CatalogEditCatalogAppParamsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CatalogEditCatalogAppParamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogEditCatalogAppParamsInternalServerError creates a CatalogEditCatalogAppParamsInternalServerError with default headers values
func NewCatalogEditCatalogAppParamsInternalServerError() *CatalogEditCatalogAppParamsInternalServerError {
	return &CatalogEditCatalogAppParamsInternalServerError{}
}

/*
CatalogEditCatalogAppParamsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CatalogEditCatalogAppParamsInternalServerError struct {
}

// IsSuccess returns true when this catalog edit catalog app params internal server error response has a 2xx status code
func (o *CatalogEditCatalogAppParamsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog edit catalog app params internal server error response has a 3xx status code
func (o *CatalogEditCatalogAppParamsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog edit catalog app params internal server error response has a 4xx status code
func (o *CatalogEditCatalogAppParamsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog edit catalog app params internal server error response has a 5xx status code
func (o *CatalogEditCatalogAppParamsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this catalog edit catalog app params internal server error response a status code equal to that given
func (o *CatalogEditCatalogAppParamsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CatalogEditCatalogAppParamsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsInternalServerError ", 500)
}

func (o *CatalogEditCatalogAppParamsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Catalog/edit-catalogapp-params][%d] catalogEditCatalogAppParamsInternalServerError ", 500)
}

func (o *CatalogEditCatalogAppParamsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
