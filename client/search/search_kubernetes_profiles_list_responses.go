// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
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

/*
SearchKubernetesProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type SearchKubernetesProfilesListOK struct {
	Payload *SearchKubernetesProfilesListOKBody
}

// IsSuccess returns true when this search kubernetes profiles list o k response has a 2xx status code
func (o *SearchKubernetesProfilesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search kubernetes profiles list o k response has a 3xx status code
func (o *SearchKubernetesProfilesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search kubernetes profiles list o k response has a 4xx status code
func (o *SearchKubernetesProfilesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search kubernetes profiles list o k response has a 5xx status code
func (o *SearchKubernetesProfilesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search kubernetes profiles list o k response a status code equal to that given
func (o *SearchKubernetesProfilesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchKubernetesProfilesListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListOK  %+v", 200, o.Payload)
}

func (o *SearchKubernetesProfilesListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListOK  %+v", 200, o.Payload)
}

func (o *SearchKubernetesProfilesListOK) GetPayload() *SearchKubernetesProfilesListOKBody {
	return o.Payload
}

func (o *SearchKubernetesProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchKubernetesProfilesListOKBody)

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

/*
SearchKubernetesProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchKubernetesProfilesListBadRequest struct {
	Payload []*SearchKubernetesProfilesListBadRequestBodyItems0
}

// IsSuccess returns true when this search kubernetes profiles list bad request response has a 2xx status code
func (o *SearchKubernetesProfilesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search kubernetes profiles list bad request response has a 3xx status code
func (o *SearchKubernetesProfilesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search kubernetes profiles list bad request response has a 4xx status code
func (o *SearchKubernetesProfilesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search kubernetes profiles list bad request response has a 5xx status code
func (o *SearchKubernetesProfilesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search kubernetes profiles list bad request response a status code equal to that given
func (o *SearchKubernetesProfilesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchKubernetesProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchKubernetesProfilesListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchKubernetesProfilesListBadRequest) GetPayload() []*SearchKubernetesProfilesListBadRequestBodyItems0 {
	return o.Payload
}

func (o *SearchKubernetesProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKubernetesProfilesListUnauthorized creates a SearchKubernetesProfilesListUnauthorized with default headers values
func NewSearchKubernetesProfilesListUnauthorized() *SearchKubernetesProfilesListUnauthorized {
	return &SearchKubernetesProfilesListUnauthorized{}
}

/*
SearchKubernetesProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchKubernetesProfilesListUnauthorized struct {
	Payload *SearchKubernetesProfilesListUnauthorizedBody
}

// IsSuccess returns true when this search kubernetes profiles list unauthorized response has a 2xx status code
func (o *SearchKubernetesProfilesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search kubernetes profiles list unauthorized response has a 3xx status code
func (o *SearchKubernetesProfilesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search kubernetes profiles list unauthorized response has a 4xx status code
func (o *SearchKubernetesProfilesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search kubernetes profiles list unauthorized response has a 5xx status code
func (o *SearchKubernetesProfilesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search kubernetes profiles list unauthorized response a status code equal to that given
func (o *SearchKubernetesProfilesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchKubernetesProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchKubernetesProfilesListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchKubernetesProfilesListUnauthorized) GetPayload() *SearchKubernetesProfilesListUnauthorizedBody {
	return o.Payload
}

func (o *SearchKubernetesProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchKubernetesProfilesListUnauthorizedBody)

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

/*
SearchKubernetesProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchKubernetesProfilesListForbidden struct {
	Payload *SearchKubernetesProfilesListForbiddenBody
}

// IsSuccess returns true when this search kubernetes profiles list forbidden response has a 2xx status code
func (o *SearchKubernetesProfilesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search kubernetes profiles list forbidden response has a 3xx status code
func (o *SearchKubernetesProfilesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search kubernetes profiles list forbidden response has a 4xx status code
func (o *SearchKubernetesProfilesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search kubernetes profiles list forbidden response has a 5xx status code
func (o *SearchKubernetesProfilesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search kubernetes profiles list forbidden response a status code equal to that given
func (o *SearchKubernetesProfilesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchKubernetesProfilesListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchKubernetesProfilesListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchKubernetesProfilesListForbidden) GetPayload() *SearchKubernetesProfilesListForbiddenBody {
	return o.Payload
}

func (o *SearchKubernetesProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchKubernetesProfilesListForbiddenBody)

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

/*
SearchKubernetesProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchKubernetesProfilesListNotFound struct {
	Payload *SearchKubernetesProfilesListNotFoundBody
}

// IsSuccess returns true when this search kubernetes profiles list not found response has a 2xx status code
func (o *SearchKubernetesProfilesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search kubernetes profiles list not found response has a 3xx status code
func (o *SearchKubernetesProfilesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search kubernetes profiles list not found response has a 4xx status code
func (o *SearchKubernetesProfilesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search kubernetes profiles list not found response has a 5xx status code
func (o *SearchKubernetesProfilesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search kubernetes profiles list not found response a status code equal to that given
func (o *SearchKubernetesProfilesListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchKubernetesProfilesListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchKubernetesProfilesListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchKubernetesProfilesListNotFound) GetPayload() *SearchKubernetesProfilesListNotFoundBody {
	return o.Payload
}

func (o *SearchKubernetesProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchKubernetesProfilesListNotFoundBody)

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

/*
SearchKubernetesProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchKubernetesProfilesListInternalServerError struct {
}

// IsSuccess returns true when this search kubernetes profiles list internal server error response has a 2xx status code
func (o *SearchKubernetesProfilesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search kubernetes profiles list internal server error response has a 3xx status code
func (o *SearchKubernetesProfilesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search kubernetes profiles list internal server error response has a 4xx status code
func (o *SearchKubernetesProfilesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search kubernetes profiles list internal server error response has a 5xx status code
func (o *SearchKubernetesProfilesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search kubernetes profiles list internal server error response a status code equal to that given
func (o *SearchKubernetesProfilesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchKubernetesProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListInternalServerError ", 500)
}

func (o *SearchKubernetesProfilesListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/kubernetes-profiles][%d] searchKubernetesProfilesListInternalServerError ", 500)
}

func (o *SearchKubernetesProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
SearchKubernetesProfilesListBadRequestBodyItems0 search kubernetes profiles list bad request body items0
swagger:model SearchKubernetesProfilesListBadRequestBodyItems0
*/
type SearchKubernetesProfilesListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this search kubernetes profiles list bad request body items0
func (o *SearchKubernetesProfilesListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search kubernetes profiles list bad request body items0 based on context it is used
func (o *SearchKubernetesProfilesListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchKubernetesProfilesListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchKubernetesProfilesListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res SearchKubernetesProfilesListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchKubernetesProfilesListBody search kubernetes profiles list body
swagger:model SearchKubernetesProfilesListBody
*/
type SearchKubernetesProfilesListBody struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// search term
	SearchTerm string `json:"searchTerm,omitempty"`
}

// Validate validates this search kubernetes profiles list body
func (o *SearchKubernetesProfilesListBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search kubernetes profiles list body based on context it is used
func (o *SearchKubernetesProfilesListBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchKubernetesProfilesListBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchKubernetesProfilesListBody) UnmarshalBinary(b []byte) error {
	var res SearchKubernetesProfilesListBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchKubernetesProfilesListForbiddenBody search kubernetes profiles list forbidden body
swagger:model SearchKubernetesProfilesListForbiddenBody
*/
type SearchKubernetesProfilesListForbiddenBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search kubernetes profiles list forbidden body
func (o *SearchKubernetesProfilesListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search kubernetes profiles list forbidden body based on context it is used
func (o *SearchKubernetesProfilesListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchKubernetesProfilesListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchKubernetesProfilesListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SearchKubernetesProfilesListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchKubernetesProfilesListNotFoundBody search kubernetes profiles list not found body
swagger:model SearchKubernetesProfilesListNotFoundBody
*/
type SearchKubernetesProfilesListNotFoundBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search kubernetes profiles list not found body
func (o *SearchKubernetesProfilesListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search kubernetes profiles list not found body based on context it is used
func (o *SearchKubernetesProfilesListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchKubernetesProfilesListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchKubernetesProfilesListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SearchKubernetesProfilesListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchKubernetesProfilesListOKBody search kubernetes profiles list o k body
swagger:model SearchKubernetesProfilesListOKBody
*/
type SearchKubernetesProfilesListOKBody struct {

	// data
	Data []*SearchKubernetesProfilesListOKBodyDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this search kubernetes profiles list o k body
func (o *SearchKubernetesProfilesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchKubernetesProfilesListOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchKubernetesProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchKubernetesProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search kubernetes profiles list o k body based on the context it is used
func (o *SearchKubernetesProfilesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchKubernetesProfilesListOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchKubernetesProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchKubernetesProfilesListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchKubernetesProfilesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchKubernetesProfilesListOKBody) UnmarshalBinary(b []byte) error {
	var res SearchKubernetesProfilesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchKubernetesProfilesListOKBodyDataItems0 search kubernetes profiles list o k body data items0
swagger:model SearchKubernetesProfilesListOKBodyDataItems0
*/
type SearchKubernetesProfilesListOKBodyDataItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`
}

// Validate validates this search kubernetes profiles list o k body data items0
func (o *SearchKubernetesProfilesListOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search kubernetes profiles list o k body data items0 based on context it is used
func (o *SearchKubernetesProfilesListOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchKubernetesProfilesListOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchKubernetesProfilesListOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res SearchKubernetesProfilesListOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchKubernetesProfilesListUnauthorizedBody search kubernetes profiles list unauthorized body
swagger:model SearchKubernetesProfilesListUnauthorizedBody
*/
type SearchKubernetesProfilesListUnauthorizedBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search kubernetes profiles list unauthorized body
func (o *SearchKubernetesProfilesListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search kubernetes profiles list unauthorized body based on context it is used
func (o *SearchKubernetesProfilesListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchKubernetesProfilesListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchKubernetesProfilesListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res SearchKubernetesProfilesListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
