// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skunkie/grpc-gateway-example/pkg/httpclient/models"
)

// APIGetUserReader is a Reader for the APIGetUser structure.
type APIGetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIGetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAPIGetUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAPIGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAPIGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIGetUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAPIGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIGetUserOK creates a APIGetUserOK with default headers values
func NewAPIGetUserOK() *APIGetUserOK {
	return &APIGetUserOK{}
}

/*
APIGetUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type APIGetUserOK struct {
	Payload *models.PbUser
}

// IsSuccess returns true when this api get user o k response has a 2xx status code
func (o *APIGetUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this api get user o k response has a 3xx status code
func (o *APIGetUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api get user o k response has a 4xx status code
func (o *APIGetUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this api get user o k response has a 5xx status code
func (o *APIGetUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this api get user o k response a status code equal to that given
func (o *APIGetUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the api get user o k response
func (o *APIGetUserOK) Code() int {
	return 200
}

func (o *APIGetUserOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserOK  %+v", 200, o.Payload)
}

func (o *APIGetUserOK) String() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserOK  %+v", 200, o.Payload)
}

func (o *APIGetUserOK) GetPayload() *models.PbUser {
	return o.Payload
}

func (o *APIGetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetUserBadRequest creates a APIGetUserBadRequest with default headers values
func NewAPIGetUserBadRequest() *APIGetUserBadRequest {
	return &APIGetUserBadRequest{}
}

/*
APIGetUserBadRequest describes a response with status code 400, with default header values.

Bad Request.
*/
type APIGetUserBadRequest struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api get user bad request response has a 2xx status code
func (o *APIGetUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api get user bad request response has a 3xx status code
func (o *APIGetUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api get user bad request response has a 4xx status code
func (o *APIGetUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this api get user bad request response has a 5xx status code
func (o *APIGetUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this api get user bad request response a status code equal to that given
func (o *APIGetUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the api get user bad request response
func (o *APIGetUserBadRequest) Code() int {
	return 400
}

func (o *APIGetUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserBadRequest  %+v", 400, o.Payload)
}

func (o *APIGetUserBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserBadRequest  %+v", 400, o.Payload)
}

func (o *APIGetUserBadRequest) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIGetUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetUserForbidden creates a APIGetUserForbidden with default headers values
func NewAPIGetUserForbidden() *APIGetUserForbidden {
	return &APIGetUserForbidden{}
}

/*
APIGetUserForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type APIGetUserForbidden struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api get user forbidden response has a 2xx status code
func (o *APIGetUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api get user forbidden response has a 3xx status code
func (o *APIGetUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api get user forbidden response has a 4xx status code
func (o *APIGetUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this api get user forbidden response has a 5xx status code
func (o *APIGetUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this api get user forbidden response a status code equal to that given
func (o *APIGetUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the api get user forbidden response
func (o *APIGetUserForbidden) Code() int {
	return 403
}

func (o *APIGetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserForbidden  %+v", 403, o.Payload)
}

func (o *APIGetUserForbidden) String() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserForbidden  %+v", 403, o.Payload)
}

func (o *APIGetUserForbidden) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIGetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetUserNotFound creates a APIGetUserNotFound with default headers values
func NewAPIGetUserNotFound() *APIGetUserNotFound {
	return &APIGetUserNotFound{}
}

/*
APIGetUserNotFound describes a response with status code 404, with default header values.

Not Found.
*/
type APIGetUserNotFound struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api get user not found response has a 2xx status code
func (o *APIGetUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api get user not found response has a 3xx status code
func (o *APIGetUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api get user not found response has a 4xx status code
func (o *APIGetUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this api get user not found response has a 5xx status code
func (o *APIGetUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this api get user not found response a status code equal to that given
func (o *APIGetUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the api get user not found response
func (o *APIGetUserNotFound) Code() int {
	return 404
}

func (o *APIGetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserNotFound  %+v", 404, o.Payload)
}

func (o *APIGetUserNotFound) String() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserNotFound  %+v", 404, o.Payload)
}

func (o *APIGetUserNotFound) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIGetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetUserInternalServerError creates a APIGetUserInternalServerError with default headers values
func NewAPIGetUserInternalServerError() *APIGetUserInternalServerError {
	return &APIGetUserInternalServerError{}
}

/*
APIGetUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error.
*/
type APIGetUserInternalServerError struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api get user internal server error response has a 2xx status code
func (o *APIGetUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api get user internal server error response has a 3xx status code
func (o *APIGetUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api get user internal server error response has a 4xx status code
func (o *APIGetUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this api get user internal server error response has a 5xx status code
func (o *APIGetUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this api get user internal server error response a status code equal to that given
func (o *APIGetUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the api get user internal server error response
func (o *APIGetUserInternalServerError) Code() int {
	return 500
}

func (o *APIGetUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserInternalServerError  %+v", 500, o.Payload)
}

func (o *APIGetUserInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] apiGetUserInternalServerError  %+v", 500, o.Payload)
}

func (o *APIGetUserInternalServerError) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIGetUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIGetUserDefault creates a APIGetUserDefault with default headers values
func NewAPIGetUserDefault(code int) *APIGetUserDefault {
	return &APIGetUserDefault{
		_statusCode: code,
	}
}

/*
APIGetUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type APIGetUserDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this Api get user default response has a 2xx status code
func (o *APIGetUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this Api get user default response has a 3xx status code
func (o *APIGetUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this Api get user default response has a 4xx status code
func (o *APIGetUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this Api get user default response has a 5xx status code
func (o *APIGetUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this Api get user default response a status code equal to that given
func (o *APIGetUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the Api get user default response
func (o *APIGetUserDefault) Code() int {
	return o._statusCode
}

func (o *APIGetUserDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] Api_GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *APIGetUserDefault) String() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] Api_GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *APIGetUserDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIGetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}