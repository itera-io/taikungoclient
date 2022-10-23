// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentGetFinalPriceReader is a Reader for the PaymentGetFinalPrice structure.
type PaymentGetFinalPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGetFinalPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGetFinalPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentGetFinalPriceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentGetFinalPriceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentGetFinalPriceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentGetFinalPriceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentGetFinalPriceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentGetFinalPriceOK creates a PaymentGetFinalPriceOK with default headers values
func NewPaymentGetFinalPriceOK() *PaymentGetFinalPriceOK {
	return &PaymentGetFinalPriceOK{}
}

/*
PaymentGetFinalPriceOK describes a response with status code 200, with default header values.

Success
*/
type PaymentGetFinalPriceOK struct {
	Payload *PaymentGetFinalPriceOKBody
}

// IsSuccess returns true when this payment get final price o k response has a 2xx status code
func (o *PaymentGetFinalPriceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment get final price o k response has a 3xx status code
func (o *PaymentGetFinalPriceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price o k response has a 4xx status code
func (o *PaymentGetFinalPriceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get final price o k response has a 5xx status code
func (o *PaymentGetFinalPriceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price o k response a status code equal to that given
func (o *PaymentGetFinalPriceOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentGetFinalPriceOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceOK  %+v", 200, o.Payload)
}

func (o *PaymentGetFinalPriceOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceOK  %+v", 200, o.Payload)
}

func (o *PaymentGetFinalPriceOK) GetPayload() *PaymentGetFinalPriceOKBody {
	return o.Payload
}

func (o *PaymentGetFinalPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PaymentGetFinalPriceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceBadRequest creates a PaymentGetFinalPriceBadRequest with default headers values
func NewPaymentGetFinalPriceBadRequest() *PaymentGetFinalPriceBadRequest {
	return &PaymentGetFinalPriceBadRequest{}
}

/*
PaymentGetFinalPriceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentGetFinalPriceBadRequest struct {
	Payload []*PaymentGetFinalPriceBadRequestBodyItems0
}

// IsSuccess returns true when this payment get final price bad request response has a 2xx status code
func (o *PaymentGetFinalPriceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price bad request response has a 3xx status code
func (o *PaymentGetFinalPriceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price bad request response has a 4xx status code
func (o *PaymentGetFinalPriceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price bad request response has a 5xx status code
func (o *PaymentGetFinalPriceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price bad request response a status code equal to that given
func (o *PaymentGetFinalPriceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PaymentGetFinalPriceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetFinalPriceBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetFinalPriceBadRequest) GetPayload() []*PaymentGetFinalPriceBadRequestBodyItems0 {
	return o.Payload
}

func (o *PaymentGetFinalPriceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceUnauthorized creates a PaymentGetFinalPriceUnauthorized with default headers values
func NewPaymentGetFinalPriceUnauthorized() *PaymentGetFinalPriceUnauthorized {
	return &PaymentGetFinalPriceUnauthorized{}
}

/*
PaymentGetFinalPriceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentGetFinalPriceUnauthorized struct {
	Payload *PaymentGetFinalPriceUnauthorizedBody
}

// IsSuccess returns true when this payment get final price unauthorized response has a 2xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price unauthorized response has a 3xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price unauthorized response has a 4xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price unauthorized response has a 5xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price unauthorized response a status code equal to that given
func (o *PaymentGetFinalPriceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PaymentGetFinalPriceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetFinalPriceUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetFinalPriceUnauthorized) GetPayload() *PaymentGetFinalPriceUnauthorizedBody {
	return o.Payload
}

func (o *PaymentGetFinalPriceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PaymentGetFinalPriceUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceForbidden creates a PaymentGetFinalPriceForbidden with default headers values
func NewPaymentGetFinalPriceForbidden() *PaymentGetFinalPriceForbidden {
	return &PaymentGetFinalPriceForbidden{}
}

/*
PaymentGetFinalPriceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentGetFinalPriceForbidden struct {
	Payload *PaymentGetFinalPriceForbiddenBody
}

// IsSuccess returns true when this payment get final price forbidden response has a 2xx status code
func (o *PaymentGetFinalPriceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price forbidden response has a 3xx status code
func (o *PaymentGetFinalPriceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price forbidden response has a 4xx status code
func (o *PaymentGetFinalPriceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price forbidden response has a 5xx status code
func (o *PaymentGetFinalPriceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price forbidden response a status code equal to that given
func (o *PaymentGetFinalPriceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PaymentGetFinalPriceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetFinalPriceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetFinalPriceForbidden) GetPayload() *PaymentGetFinalPriceForbiddenBody {
	return o.Payload
}

func (o *PaymentGetFinalPriceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PaymentGetFinalPriceForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceNotFound creates a PaymentGetFinalPriceNotFound with default headers values
func NewPaymentGetFinalPriceNotFound() *PaymentGetFinalPriceNotFound {
	return &PaymentGetFinalPriceNotFound{}
}

/*
PaymentGetFinalPriceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentGetFinalPriceNotFound struct {
	Payload *PaymentGetFinalPriceNotFoundBody
}

// IsSuccess returns true when this payment get final price not found response has a 2xx status code
func (o *PaymentGetFinalPriceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price not found response has a 3xx status code
func (o *PaymentGetFinalPriceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price not found response has a 4xx status code
func (o *PaymentGetFinalPriceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price not found response has a 5xx status code
func (o *PaymentGetFinalPriceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price not found response a status code equal to that given
func (o *PaymentGetFinalPriceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PaymentGetFinalPriceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetFinalPriceNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetFinalPriceNotFound) GetPayload() *PaymentGetFinalPriceNotFoundBody {
	return o.Payload
}

func (o *PaymentGetFinalPriceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PaymentGetFinalPriceNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceInternalServerError creates a PaymentGetFinalPriceInternalServerError with default headers values
func NewPaymentGetFinalPriceInternalServerError() *PaymentGetFinalPriceInternalServerError {
	return &PaymentGetFinalPriceInternalServerError{}
}

/*
PaymentGetFinalPriceInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentGetFinalPriceInternalServerError struct {
}

// IsSuccess returns true when this payment get final price internal server error response has a 2xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price internal server error response has a 3xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price internal server error response has a 4xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get final price internal server error response has a 5xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment get final price internal server error response a status code equal to that given
func (o *PaymentGetFinalPriceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PaymentGetFinalPriceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceInternalServerError ", 500)
}

func (o *PaymentGetFinalPriceInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceInternalServerError ", 500)
}

func (o *PaymentGetFinalPriceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
PaymentGetFinalPriceBadRequestBodyItems0 payment get final price bad request body items0
swagger:model PaymentGetFinalPriceBadRequestBodyItems0
*/
type PaymentGetFinalPriceBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this payment get final price bad request body items0
func (o *PaymentGetFinalPriceBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment get final price bad request body items0 based on context it is used
func (o *PaymentGetFinalPriceBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PaymentGetFinalPriceBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PaymentGetFinalPriceBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res PaymentGetFinalPriceBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PaymentGetFinalPriceBody payment get final price body
swagger:model PaymentGetFinalPriceBody
*/
type PaymentGetFinalPriceBody struct {

	// organization Uuid
	OrganizationUUID string `json:"organizationUuid,omitempty"`

	// subscription Id
	SubscriptionID int32 `json:"subscriptionId,omitempty"`
}

// Validate validates this payment get final price body
func (o *PaymentGetFinalPriceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment get final price body based on context it is used
func (o *PaymentGetFinalPriceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PaymentGetFinalPriceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PaymentGetFinalPriceBody) UnmarshalBinary(b []byte) error {
	var res PaymentGetFinalPriceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PaymentGetFinalPriceForbiddenBody payment get final price forbidden body
swagger:model PaymentGetFinalPriceForbiddenBody
*/
type PaymentGetFinalPriceForbiddenBody struct {

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

// Validate validates this payment get final price forbidden body
func (o *PaymentGetFinalPriceForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment get final price forbidden body based on context it is used
func (o *PaymentGetFinalPriceForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PaymentGetFinalPriceForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PaymentGetFinalPriceForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PaymentGetFinalPriceForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PaymentGetFinalPriceNotFoundBody payment get final price not found body
swagger:model PaymentGetFinalPriceNotFoundBody
*/
type PaymentGetFinalPriceNotFoundBody struct {

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

// Validate validates this payment get final price not found body
func (o *PaymentGetFinalPriceNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment get final price not found body based on context it is used
func (o *PaymentGetFinalPriceNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PaymentGetFinalPriceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PaymentGetFinalPriceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PaymentGetFinalPriceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PaymentGetFinalPriceOKBody payment get final price o k body
swagger:model PaymentGetFinalPriceOKBody
*/
type PaymentGetFinalPriceOKBody struct {

	// monthly final price
	MonthlyFinalPrice float64 `json:"monthlyFinalPrice,omitempty"`

	// yearly final price
	YearlyFinalPrice float64 `json:"yearlyFinalPrice,omitempty"`
}

// Validate validates this payment get final price o k body
func (o *PaymentGetFinalPriceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment get final price o k body based on context it is used
func (o *PaymentGetFinalPriceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PaymentGetFinalPriceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PaymentGetFinalPriceOKBody) UnmarshalBinary(b []byte) error {
	var res PaymentGetFinalPriceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PaymentGetFinalPriceUnauthorizedBody payment get final price unauthorized body
swagger:model PaymentGetFinalPriceUnauthorizedBody
*/
type PaymentGetFinalPriceUnauthorizedBody struct {

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

// Validate validates this payment get final price unauthorized body
func (o *PaymentGetFinalPriceUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment get final price unauthorized body based on context it is used
func (o *PaymentGetFinalPriceUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PaymentGetFinalPriceUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PaymentGetFinalPriceUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PaymentGetFinalPriceUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
