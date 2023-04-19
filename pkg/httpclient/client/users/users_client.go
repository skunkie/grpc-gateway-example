// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	APICreateUser(params *APICreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APICreateUserCreated, error)

	APIDeleteUser(params *APIDeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIDeleteUserNoContent, error)

	APIGetUser(params *APIGetUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIGetUserOK, error)

	APIListUsers(params *APIListUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIListUsersOK, error)

	APIUpdateUser(params *APIUpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIUpdateUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
APICreateUser creates user
*/
func (a *Client) APICreateUser(params *APICreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APICreateUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPICreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Api_CreateUser",
		Method:             "POST",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APICreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APICreateUserCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APICreateUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
APIDeleteUser deletes user
*/
func (a *Client) APIDeleteUser(params *APIDeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIDeleteUserNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIDeleteUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Api_DeleteUser",
		Method:             "DELETE",
		PathPattern:        "/v1/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIDeleteUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIDeleteUserNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIDeleteUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
APIGetUser gets user
*/
func (a *Client) APIGetUser(params *APIGetUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIGetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIGetUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Api_GetUser",
		Method:             "GET",
		PathPattern:        "/v1/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIGetUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIGetUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIGetUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
APIListUsers lists users
*/
func (a *Client) APIListUsers(params *APIListUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIListUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIListUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Api_ListUsers",
		Method:             "GET",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIListUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIListUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIListUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
APIUpdateUser updates user info
*/
func (a *Client) APIUpdateUser(params *APIUpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIUpdateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIUpdateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Api_UpdateUser",
		Method:             "PATCH",
		PathPattern:        "/v1/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIUpdateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIUpdateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIUpdateUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
