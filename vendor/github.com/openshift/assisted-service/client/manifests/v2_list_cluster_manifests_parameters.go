// Code generated by go-swagger; DO NOT EDIT.

package manifests

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

// NewV2ListClusterManifestsParams creates a new V2ListClusterManifestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2ListClusterManifestsParams() *V2ListClusterManifestsParams {
	return &V2ListClusterManifestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2ListClusterManifestsParamsWithTimeout creates a new V2ListClusterManifestsParams object
// with the ability to set a timeout on a request.
func NewV2ListClusterManifestsParamsWithTimeout(timeout time.Duration) *V2ListClusterManifestsParams {
	return &V2ListClusterManifestsParams{
		timeout: timeout,
	}
}

// NewV2ListClusterManifestsParamsWithContext creates a new V2ListClusterManifestsParams object
// with the ability to set a context for a request.
func NewV2ListClusterManifestsParamsWithContext(ctx context.Context) *V2ListClusterManifestsParams {
	return &V2ListClusterManifestsParams{
		Context: ctx,
	}
}

// NewV2ListClusterManifestsParamsWithHTTPClient creates a new V2ListClusterManifestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2ListClusterManifestsParamsWithHTTPClient(client *http.Client) *V2ListClusterManifestsParams {
	return &V2ListClusterManifestsParams{
		HTTPClient: client,
	}
}

/* V2ListClusterManifestsParams contains all the parameters to send to the API endpoint
   for the v2 list cluster manifests operation.

   Typically these are written to a http.Request.
*/
type V2ListClusterManifestsParams struct {

	/* ClusterID.

	   The cluster for which the manifests should be listed.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 list cluster manifests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2ListClusterManifestsParams) WithDefaults() *V2ListClusterManifestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 list cluster manifests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2ListClusterManifestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) WithTimeout(timeout time.Duration) *V2ListClusterManifestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) WithContext(ctx context.Context) *V2ListClusterManifestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) WithHTTPClient(client *http.Client) *V2ListClusterManifestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) WithClusterID(clusterID strfmt.UUID) *V2ListClusterManifestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 list cluster manifests params
func (o *V2ListClusterManifestsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *V2ListClusterManifestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
