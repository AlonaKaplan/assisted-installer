// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2RegisterClusterCreatedCode is the HTTP code returned for type V2RegisterClusterCreated
const V2RegisterClusterCreatedCode int = 201

/*V2RegisterClusterCreated Success.

swagger:response v2RegisterClusterCreated
*/
type V2RegisterClusterCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewV2RegisterClusterCreated creates V2RegisterClusterCreated with default headers values
func NewV2RegisterClusterCreated() *V2RegisterClusterCreated {

	return &V2RegisterClusterCreated{}
}

// WithPayload adds the payload to the v2 register cluster created response
func (o *V2RegisterClusterCreated) WithPayload(payload *models.Cluster) *V2RegisterClusterCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 register cluster created response
func (o *V2RegisterClusterCreated) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2RegisterClusterCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2RegisterClusterBadRequestCode is the HTTP code returned for type V2RegisterClusterBadRequest
const V2RegisterClusterBadRequestCode int = 400

/*V2RegisterClusterBadRequest Error.

swagger:response v2RegisterClusterBadRequest
*/
type V2RegisterClusterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2RegisterClusterBadRequest creates V2RegisterClusterBadRequest with default headers values
func NewV2RegisterClusterBadRequest() *V2RegisterClusterBadRequest {

	return &V2RegisterClusterBadRequest{}
}

// WithPayload adds the payload to the v2 register cluster bad request response
func (o *V2RegisterClusterBadRequest) WithPayload(payload *models.Error) *V2RegisterClusterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 register cluster bad request response
func (o *V2RegisterClusterBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2RegisterClusterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2RegisterClusterUnauthorizedCode is the HTTP code returned for type V2RegisterClusterUnauthorized
const V2RegisterClusterUnauthorizedCode int = 401

/*V2RegisterClusterUnauthorized Unauthorized.

swagger:response v2RegisterClusterUnauthorized
*/
type V2RegisterClusterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2RegisterClusterUnauthorized creates V2RegisterClusterUnauthorized with default headers values
func NewV2RegisterClusterUnauthorized() *V2RegisterClusterUnauthorized {

	return &V2RegisterClusterUnauthorized{}
}

// WithPayload adds the payload to the v2 register cluster unauthorized response
func (o *V2RegisterClusterUnauthorized) WithPayload(payload *models.InfraError) *V2RegisterClusterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 register cluster unauthorized response
func (o *V2RegisterClusterUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2RegisterClusterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2RegisterClusterForbiddenCode is the HTTP code returned for type V2RegisterClusterForbidden
const V2RegisterClusterForbiddenCode int = 403

/*V2RegisterClusterForbidden Forbidden.

swagger:response v2RegisterClusterForbidden
*/
type V2RegisterClusterForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2RegisterClusterForbidden creates V2RegisterClusterForbidden with default headers values
func NewV2RegisterClusterForbidden() *V2RegisterClusterForbidden {

	return &V2RegisterClusterForbidden{}
}

// WithPayload adds the payload to the v2 register cluster forbidden response
func (o *V2RegisterClusterForbidden) WithPayload(payload *models.InfraError) *V2RegisterClusterForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 register cluster forbidden response
func (o *V2RegisterClusterForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2RegisterClusterForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2RegisterClusterMethodNotAllowedCode is the HTTP code returned for type V2RegisterClusterMethodNotAllowed
const V2RegisterClusterMethodNotAllowedCode int = 405

/*V2RegisterClusterMethodNotAllowed Method Not Allowed.

swagger:response v2RegisterClusterMethodNotAllowed
*/
type V2RegisterClusterMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2RegisterClusterMethodNotAllowed creates V2RegisterClusterMethodNotAllowed with default headers values
func NewV2RegisterClusterMethodNotAllowed() *V2RegisterClusterMethodNotAllowed {

	return &V2RegisterClusterMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 register cluster method not allowed response
func (o *V2RegisterClusterMethodNotAllowed) WithPayload(payload *models.Error) *V2RegisterClusterMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 register cluster method not allowed response
func (o *V2RegisterClusterMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2RegisterClusterMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2RegisterClusterInternalServerErrorCode is the HTTP code returned for type V2RegisterClusterInternalServerError
const V2RegisterClusterInternalServerErrorCode int = 500

/*V2RegisterClusterInternalServerError Error.

swagger:response v2RegisterClusterInternalServerError
*/
type V2RegisterClusterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2RegisterClusterInternalServerError creates V2RegisterClusterInternalServerError with default headers values
func NewV2RegisterClusterInternalServerError() *V2RegisterClusterInternalServerError {

	return &V2RegisterClusterInternalServerError{}
}

// WithPayload adds the payload to the v2 register cluster internal server error response
func (o *V2RegisterClusterInternalServerError) WithPayload(payload *models.Error) *V2RegisterClusterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 register cluster internal server error response
func (o *V2RegisterClusterInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2RegisterClusterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
