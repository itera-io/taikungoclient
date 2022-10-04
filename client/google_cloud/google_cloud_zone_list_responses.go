// Code generated by go-swagger; DO NOT EDIT.

package google_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// GoogleCloudZoneListReader is a Reader for the GoogleCloudZoneList structure.
type GoogleCloudZoneListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GoogleCloudZoneListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGoogleCloudZoneListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGoogleCloudZoneListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGoogleCloudZoneListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGoogleCloudZoneListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGoogleCloudZoneListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGoogleCloudZoneListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGoogleCloudZoneListOK creates a GoogleCloudZoneListOK with default headers values
func NewGoogleCloudZoneListOK() *GoogleCloudZoneListOK {
	return &GoogleCloudZoneListOK{}
}

/*
GoogleCloudZoneListOK describes a response with status code 200, with default header values.

Success
*/
type GoogleCloudZoneListOK struct {
	Payload []string
}

// IsSuccess returns true when this google cloud zone list o k response has a 2xx status code
func (o *GoogleCloudZoneListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this google cloud zone list o k response has a 3xx status code
func (o *GoogleCloudZoneListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud zone list o k response has a 4xx status code
func (o *GoogleCloudZoneListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this google cloud zone list o k response has a 5xx status code
func (o *GoogleCloudZoneListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud zone list o k response a status code equal to that given
func (o *GoogleCloudZoneListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GoogleCloudZoneListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListOK  %+v", 200, o.Payload)
}

func (o *GoogleCloudZoneListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListOK  %+v", 200, o.Payload)
}

func (o *GoogleCloudZoneListOK) GetPayload() []string {
	return o.Payload
}

func (o *GoogleCloudZoneListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListBadRequest creates a GoogleCloudZoneListBadRequest with default headers values
func NewGoogleCloudZoneListBadRequest() *GoogleCloudZoneListBadRequest {
	return &GoogleCloudZoneListBadRequest{}
}

/*
GoogleCloudZoneListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GoogleCloudZoneListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this google cloud zone list bad request response has a 2xx status code
func (o *GoogleCloudZoneListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud zone list bad request response has a 3xx status code
func (o *GoogleCloudZoneListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud zone list bad request response has a 4xx status code
func (o *GoogleCloudZoneListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud zone list bad request response has a 5xx status code
func (o *GoogleCloudZoneListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud zone list bad request response a status code equal to that given
func (o *GoogleCloudZoneListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GoogleCloudZoneListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListBadRequest  %+v", 400, o.Payload)
}

func (o *GoogleCloudZoneListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListBadRequest  %+v", 400, o.Payload)
}

func (o *GoogleCloudZoneListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *GoogleCloudZoneListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListUnauthorized creates a GoogleCloudZoneListUnauthorized with default headers values
func NewGoogleCloudZoneListUnauthorized() *GoogleCloudZoneListUnauthorized {
	return &GoogleCloudZoneListUnauthorized{}
}

/*
GoogleCloudZoneListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GoogleCloudZoneListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this google cloud zone list unauthorized response has a 2xx status code
func (o *GoogleCloudZoneListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud zone list unauthorized response has a 3xx status code
func (o *GoogleCloudZoneListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud zone list unauthorized response has a 4xx status code
func (o *GoogleCloudZoneListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud zone list unauthorized response has a 5xx status code
func (o *GoogleCloudZoneListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud zone list unauthorized response a status code equal to that given
func (o *GoogleCloudZoneListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GoogleCloudZoneListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListUnauthorized  %+v", 401, o.Payload)
}

func (o *GoogleCloudZoneListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListUnauthorized  %+v", 401, o.Payload)
}

func (o *GoogleCloudZoneListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *GoogleCloudZoneListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListForbidden creates a GoogleCloudZoneListForbidden with default headers values
func NewGoogleCloudZoneListForbidden() *GoogleCloudZoneListForbidden {
	return &GoogleCloudZoneListForbidden{}
}

/*
GoogleCloudZoneListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GoogleCloudZoneListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this google cloud zone list forbidden response has a 2xx status code
func (o *GoogleCloudZoneListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud zone list forbidden response has a 3xx status code
func (o *GoogleCloudZoneListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud zone list forbidden response has a 4xx status code
func (o *GoogleCloudZoneListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud zone list forbidden response has a 5xx status code
func (o *GoogleCloudZoneListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud zone list forbidden response a status code equal to that given
func (o *GoogleCloudZoneListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GoogleCloudZoneListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListForbidden  %+v", 403, o.Payload)
}

func (o *GoogleCloudZoneListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListForbidden  %+v", 403, o.Payload)
}

func (o *GoogleCloudZoneListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *GoogleCloudZoneListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListNotFound creates a GoogleCloudZoneListNotFound with default headers values
func NewGoogleCloudZoneListNotFound() *GoogleCloudZoneListNotFound {
	return &GoogleCloudZoneListNotFound{}
}

/*
GoogleCloudZoneListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GoogleCloudZoneListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this google cloud zone list not found response has a 2xx status code
func (o *GoogleCloudZoneListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud zone list not found response has a 3xx status code
func (o *GoogleCloudZoneListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud zone list not found response has a 4xx status code
func (o *GoogleCloudZoneListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this google cloud zone list not found response has a 5xx status code
func (o *GoogleCloudZoneListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this google cloud zone list not found response a status code equal to that given
func (o *GoogleCloudZoneListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GoogleCloudZoneListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListNotFound  %+v", 404, o.Payload)
}

func (o *GoogleCloudZoneListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListNotFound  %+v", 404, o.Payload)
}

func (o *GoogleCloudZoneListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *GoogleCloudZoneListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListInternalServerError creates a GoogleCloudZoneListInternalServerError with default headers values
func NewGoogleCloudZoneListInternalServerError() *GoogleCloudZoneListInternalServerError {
	return &GoogleCloudZoneListInternalServerError{}
}

/*
GoogleCloudZoneListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GoogleCloudZoneListInternalServerError struct {
}

// IsSuccess returns true when this google cloud zone list internal server error response has a 2xx status code
func (o *GoogleCloudZoneListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this google cloud zone list internal server error response has a 3xx status code
func (o *GoogleCloudZoneListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this google cloud zone list internal server error response has a 4xx status code
func (o *GoogleCloudZoneListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this google cloud zone list internal server error response has a 5xx status code
func (o *GoogleCloudZoneListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this google cloud zone list internal server error response a status code equal to that given
func (o *GoogleCloudZoneListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GoogleCloudZoneListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListInternalServerError ", 500)
}

func (o *GoogleCloudZoneListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListInternalServerError ", 500)
}

func (o *GoogleCloudZoneListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
