// Code generated by go-swagger; DO NOT EDIT.

package project_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectAppEditParamsReader is a Reader for the ProjectAppEditParams structure.
type ProjectAppEditParamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppEditParamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppEditParamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppEditParamsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppEditParamsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppEditParamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppEditParamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppEditParamsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppEditParamsOK creates a ProjectAppEditParamsOK with default headers values
func NewProjectAppEditParamsOK() *ProjectAppEditParamsOK {
	return &ProjectAppEditParamsOK{}
}

/* ProjectAppEditParamsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppEditParamsOK struct {
	Payload models.Unit
}

func (o *ProjectAppEditParamsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/edit-params][%d] projectAppEditParamsOK  %+v", 200, o.Payload)
}
func (o *ProjectAppEditParamsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectAppEditParamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppEditParamsBadRequest creates a ProjectAppEditParamsBadRequest with default headers values
func NewProjectAppEditParamsBadRequest() *ProjectAppEditParamsBadRequest {
	return &ProjectAppEditParamsBadRequest{}
}

/* ProjectAppEditParamsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppEditParamsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectAppEditParamsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/edit-params][%d] projectAppEditParamsBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectAppEditParamsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectAppEditParamsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppEditParamsUnauthorized creates a ProjectAppEditParamsUnauthorized with default headers values
func NewProjectAppEditParamsUnauthorized() *ProjectAppEditParamsUnauthorized {
	return &ProjectAppEditParamsUnauthorized{}
}

/* ProjectAppEditParamsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppEditParamsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppEditParamsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/edit-params][%d] projectAppEditParamsUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectAppEditParamsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppEditParamsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppEditParamsForbidden creates a ProjectAppEditParamsForbidden with default headers values
func NewProjectAppEditParamsForbidden() *ProjectAppEditParamsForbidden {
	return &ProjectAppEditParamsForbidden{}
}

/* ProjectAppEditParamsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppEditParamsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppEditParamsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/edit-params][%d] projectAppEditParamsForbidden  %+v", 403, o.Payload)
}
func (o *ProjectAppEditParamsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppEditParamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppEditParamsNotFound creates a ProjectAppEditParamsNotFound with default headers values
func NewProjectAppEditParamsNotFound() *ProjectAppEditParamsNotFound {
	return &ProjectAppEditParamsNotFound{}
}

/* ProjectAppEditParamsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppEditParamsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppEditParamsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/edit-params][%d] projectAppEditParamsNotFound  %+v", 404, o.Payload)
}
func (o *ProjectAppEditParamsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppEditParamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppEditParamsInternalServerError creates a ProjectAppEditParamsInternalServerError with default headers values
func NewProjectAppEditParamsInternalServerError() *ProjectAppEditParamsInternalServerError {
	return &ProjectAppEditParamsInternalServerError{}
}

/* ProjectAppEditParamsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppEditParamsInternalServerError struct {
}

func (o *ProjectAppEditParamsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/edit-params][%d] projectAppEditParamsInternalServerError ", 500)
}

func (o *ProjectAppEditParamsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}