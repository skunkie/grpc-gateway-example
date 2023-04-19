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

// APIListUsersReader is a Reader for the APIListUsers structure.
type APIListUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIListUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIListUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAPIListUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAPIListUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAPIListUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAPIListUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAPIListUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIListUsersOK creates a APIListUsersOK with default headers values
func NewAPIListUsersOK() *APIListUsersOK {
	return &APIListUsersOK{}
}

/*
APIListUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type APIListUsersOK struct {
	Payload *models.PbUsers
}

// IsSuccess returns true when this api list users o k response has a 2xx status code
func (o *APIListUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this api list users o k response has a 3xx status code
func (o *APIListUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api list users o k response has a 4xx status code
func (o *APIListUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this api list users o k response has a 5xx status code
func (o *APIListUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this api list users o k response a status code equal to that given
func (o *APIListUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the api list users o k response
func (o *APIListUsersOK) Code() int {
	return 200
}

func (o *APIListUsersOK) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersOK  %+v", 200, o.Payload)
}

func (o *APIListUsersOK) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersOK  %+v", 200, o.Payload)
}

func (o *APIListUsersOK) GetPayload() *models.PbUsers {
	return o.Payload
}

func (o *APIListUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbUsers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIListUsersBadRequest creates a APIListUsersBadRequest with default headers values
func NewAPIListUsersBadRequest() *APIListUsersBadRequest {
	return &APIListUsersBadRequest{}
}

/*
APIListUsersBadRequest describes a response with status code 400, with default header values.

Bad Request.
*/
type APIListUsersBadRequest struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api list users bad request response has a 2xx status code
func (o *APIListUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api list users bad request response has a 3xx status code
func (o *APIListUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api list users bad request response has a 4xx status code
func (o *APIListUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this api list users bad request response has a 5xx status code
func (o *APIListUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this api list users bad request response a status code equal to that given
func (o *APIListUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the api list users bad request response
func (o *APIListUsersBadRequest) Code() int {
	return 400
}

func (o *APIListUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersBadRequest  %+v", 400, o.Payload)
}

func (o *APIListUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersBadRequest  %+v", 400, o.Payload)
}

func (o *APIListUsersBadRequest) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIListUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIListUsersForbidden creates a APIListUsersForbidden with default headers values
func NewAPIListUsersForbidden() *APIListUsersForbidden {
	return &APIListUsersForbidden{}
}

/*
APIListUsersForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type APIListUsersForbidden struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api list users forbidden response has a 2xx status code
func (o *APIListUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api list users forbidden response has a 3xx status code
func (o *APIListUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api list users forbidden response has a 4xx status code
func (o *APIListUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this api list users forbidden response has a 5xx status code
func (o *APIListUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this api list users forbidden response a status code equal to that given
func (o *APIListUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the api list users forbidden response
func (o *APIListUsersForbidden) Code() int {
	return 403
}

func (o *APIListUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersForbidden  %+v", 403, o.Payload)
}

func (o *APIListUsersForbidden) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersForbidden  %+v", 403, o.Payload)
}

func (o *APIListUsersForbidden) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIListUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIListUsersNotFound creates a APIListUsersNotFound with default headers values
func NewAPIListUsersNotFound() *APIListUsersNotFound {
	return &APIListUsersNotFound{}
}

/*
APIListUsersNotFound describes a response with status code 404, with default header values.

Not Found.
*/
type APIListUsersNotFound struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api list users not found response has a 2xx status code
func (o *APIListUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api list users not found response has a 3xx status code
func (o *APIListUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api list users not found response has a 4xx status code
func (o *APIListUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this api list users not found response has a 5xx status code
func (o *APIListUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this api list users not found response a status code equal to that given
func (o *APIListUsersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the api list users not found response
func (o *APIListUsersNotFound) Code() int {
	return 404
}

func (o *APIListUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersNotFound  %+v", 404, o.Payload)
}

func (o *APIListUsersNotFound) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersNotFound  %+v", 404, o.Payload)
}

func (o *APIListUsersNotFound) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIListUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIListUsersInternalServerError creates a APIListUsersInternalServerError with default headers values
func NewAPIListUsersInternalServerError() *APIListUsersInternalServerError {
	return &APIListUsersInternalServerError{}
}

/*
APIListUsersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error.
*/
type APIListUsersInternalServerError struct {
	Payload *models.RPCStatus
}

// IsSuccess returns true when this api list users internal server error response has a 2xx status code
func (o *APIListUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api list users internal server error response has a 3xx status code
func (o *APIListUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api list users internal server error response has a 4xx status code
func (o *APIListUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this api list users internal server error response has a 5xx status code
func (o *APIListUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this api list users internal server error response a status code equal to that given
func (o *APIListUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the api list users internal server error response
func (o *APIListUsersInternalServerError) Code() int {
	return 500
}

func (o *APIListUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *APIListUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] apiListUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *APIListUsersInternalServerError) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIListUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIListUsersDefault creates a APIListUsersDefault with default headers values
func NewAPIListUsersDefault(code int) *APIListUsersDefault {
	return &APIListUsersDefault{
		_statusCode: code,
	}
}

/*
APIListUsersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type APIListUsersDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this Api list users default response has a 2xx status code
func (o *APIListUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this Api list users default response has a 3xx status code
func (o *APIListUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this Api list users default response has a 4xx status code
func (o *APIListUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this Api list users default response has a 5xx status code
func (o *APIListUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this Api list users default response a status code equal to that given
func (o *APIListUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the Api list users default response
func (o *APIListUsersDefault) Code() int {
	return o._statusCode
}

func (o *APIListUsersDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] Api_ListUsers default  %+v", o._statusCode, o.Payload)
}

func (o *APIListUsersDefault) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] Api_ListUsers default  %+v", o._statusCode, o.Payload)
}

func (o *APIListUsersDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *APIListUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
