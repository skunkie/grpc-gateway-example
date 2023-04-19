// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/skunkie/grpc-gateway-example/pkg/httpclient/models"
)

// NewAPICreateUserParams creates a new APICreateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPICreateUserParams() *APICreateUserParams {
	return &APICreateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPICreateUserParamsWithTimeout creates a new APICreateUserParams object
// with the ability to set a timeout on a request.
func NewAPICreateUserParamsWithTimeout(timeout time.Duration) *APICreateUserParams {
	return &APICreateUserParams{
		timeout: timeout,
	}
}

// NewAPICreateUserParamsWithContext creates a new APICreateUserParams object
// with the ability to set a context for a request.
func NewAPICreateUserParamsWithContext(ctx context.Context) *APICreateUserParams {
	return &APICreateUserParams{
		Context: ctx,
	}
}

// NewAPICreateUserParamsWithHTTPClient creates a new APICreateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPICreateUserParamsWithHTTPClient(client *http.Client) *APICreateUserParams {
	return &APICreateUserParams{
		HTTPClient: client,
	}
}

/*
APICreateUserParams contains all the parameters to send to the API endpoint

	for the Api create user operation.

	Typically these are written to a http.Request.
*/
type APICreateUserParams struct {

	// Body.
	Body *models.PbCreateUserReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APICreateUserParams) WithDefaults() *APICreateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APICreateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api create user params
func (o *APICreateUserParams) WithTimeout(timeout time.Duration) *APICreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api create user params
func (o *APICreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api create user params
func (o *APICreateUserParams) WithContext(ctx context.Context) *APICreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api create user params
func (o *APICreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api create user params
func (o *APICreateUserParams) WithHTTPClient(client *http.Client) *APICreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api create user params
func (o *APICreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the Api create user params
func (o *APICreateUserParams) WithBody(body *models.PbCreateUserReq) *APICreateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the Api create user params
func (o *APICreateUserParams) SetBody(body *models.PbCreateUserReq) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *APICreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
