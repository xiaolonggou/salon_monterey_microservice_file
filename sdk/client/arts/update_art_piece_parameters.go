// Code generated by go-swagger; DO NOT EDIT.

package arts

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
)

// NewUpdateArtPieceParams creates a new UpdateArtPieceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateArtPieceParams() *UpdateArtPieceParams {
	return &UpdateArtPieceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateArtPieceParamsWithTimeout creates a new UpdateArtPieceParams object
// with the ability to set a timeout on a request.
func NewUpdateArtPieceParamsWithTimeout(timeout time.Duration) *UpdateArtPieceParams {
	return &UpdateArtPieceParams{
		timeout: timeout,
	}
}

// NewUpdateArtPieceParamsWithContext creates a new UpdateArtPieceParams object
// with the ability to set a context for a request.
func NewUpdateArtPieceParamsWithContext(ctx context.Context) *UpdateArtPieceParams {
	return &UpdateArtPieceParams{
		Context: ctx,
	}
}

// NewUpdateArtPieceParamsWithHTTPClient creates a new UpdateArtPieceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateArtPieceParamsWithHTTPClient(client *http.Client) *UpdateArtPieceParams {
	return &UpdateArtPieceParams{
		HTTPClient: client,
	}
}

/* UpdateArtPieceParams contains all the parameters to send to the API endpoint
   for the update art piece operation.

   Typically these are written to a http.Request.
*/
type UpdateArtPieceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update art piece params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateArtPieceParams) WithDefaults() *UpdateArtPieceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update art piece params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateArtPieceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update art piece params
func (o *UpdateArtPieceParams) WithTimeout(timeout time.Duration) *UpdateArtPieceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update art piece params
func (o *UpdateArtPieceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update art piece params
func (o *UpdateArtPieceParams) WithContext(ctx context.Context) *UpdateArtPieceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update art piece params
func (o *UpdateArtPieceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update art piece params
func (o *UpdateArtPieceParams) WithHTTPClient(client *http.Client) *UpdateArtPieceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update art piece params
func (o *UpdateArtPieceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateArtPieceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
