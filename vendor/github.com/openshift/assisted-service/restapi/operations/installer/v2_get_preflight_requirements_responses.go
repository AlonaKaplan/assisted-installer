// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2GetPreflightRequirementsOKCode is the HTTP code returned for type V2GetPreflightRequirementsOK
const V2GetPreflightRequirementsOKCode int = 200

/*V2GetPreflightRequirementsOK Success.

swagger:response v2GetPreflightRequirementsOK
*/
type V2GetPreflightRequirementsOK struct {

	/*
	  In: Body
	*/
	Payload *models.PreflightHardwareRequirements `json:"body,omitempty"`
}

// NewV2GetPreflightRequirementsOK creates V2GetPreflightRequirementsOK with default headers values
func NewV2GetPreflightRequirementsOK() *V2GetPreflightRequirementsOK {

	return &V2GetPreflightRequirementsOK{}
}

// WithPayload adds the payload to the v2 get preflight requirements o k response
func (o *V2GetPreflightRequirementsOK) WithPayload(payload *models.PreflightHardwareRequirements) *V2GetPreflightRequirementsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get preflight requirements o k response
func (o *V2GetPreflightRequirementsOK) SetPayload(payload *models.PreflightHardwareRequirements) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPreflightRequirementsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPreflightRequirementsUnauthorizedCode is the HTTP code returned for type V2GetPreflightRequirementsUnauthorized
const V2GetPreflightRequirementsUnauthorizedCode int = 401

/*V2GetPreflightRequirementsUnauthorized Unauthorized.

swagger:response v2GetPreflightRequirementsUnauthorized
*/
type V2GetPreflightRequirementsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetPreflightRequirementsUnauthorized creates V2GetPreflightRequirementsUnauthorized with default headers values
func NewV2GetPreflightRequirementsUnauthorized() *V2GetPreflightRequirementsUnauthorized {

	return &V2GetPreflightRequirementsUnauthorized{}
}

// WithPayload adds the payload to the v2 get preflight requirements unauthorized response
func (o *V2GetPreflightRequirementsUnauthorized) WithPayload(payload *models.InfraError) *V2GetPreflightRequirementsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get preflight requirements unauthorized response
func (o *V2GetPreflightRequirementsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPreflightRequirementsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPreflightRequirementsForbiddenCode is the HTTP code returned for type V2GetPreflightRequirementsForbidden
const V2GetPreflightRequirementsForbiddenCode int = 403

/*V2GetPreflightRequirementsForbidden Forbidden.

swagger:response v2GetPreflightRequirementsForbidden
*/
type V2GetPreflightRequirementsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetPreflightRequirementsForbidden creates V2GetPreflightRequirementsForbidden with default headers values
func NewV2GetPreflightRequirementsForbidden() *V2GetPreflightRequirementsForbidden {

	return &V2GetPreflightRequirementsForbidden{}
}

// WithPayload adds the payload to the v2 get preflight requirements forbidden response
func (o *V2GetPreflightRequirementsForbidden) WithPayload(payload *models.InfraError) *V2GetPreflightRequirementsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get preflight requirements forbidden response
func (o *V2GetPreflightRequirementsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPreflightRequirementsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPreflightRequirementsNotFoundCode is the HTTP code returned for type V2GetPreflightRequirementsNotFound
const V2GetPreflightRequirementsNotFoundCode int = 404

/*V2GetPreflightRequirementsNotFound Error.

swagger:response v2GetPreflightRequirementsNotFound
*/
type V2GetPreflightRequirementsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPreflightRequirementsNotFound creates V2GetPreflightRequirementsNotFound with default headers values
func NewV2GetPreflightRequirementsNotFound() *V2GetPreflightRequirementsNotFound {

	return &V2GetPreflightRequirementsNotFound{}
}

// WithPayload adds the payload to the v2 get preflight requirements not found response
func (o *V2GetPreflightRequirementsNotFound) WithPayload(payload *models.Error) *V2GetPreflightRequirementsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get preflight requirements not found response
func (o *V2GetPreflightRequirementsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPreflightRequirementsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPreflightRequirementsMethodNotAllowedCode is the HTTP code returned for type V2GetPreflightRequirementsMethodNotAllowed
const V2GetPreflightRequirementsMethodNotAllowedCode int = 405

/*V2GetPreflightRequirementsMethodNotAllowed Method Not Allowed.

swagger:response v2GetPreflightRequirementsMethodNotAllowed
*/
type V2GetPreflightRequirementsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPreflightRequirementsMethodNotAllowed creates V2GetPreflightRequirementsMethodNotAllowed with default headers values
func NewV2GetPreflightRequirementsMethodNotAllowed() *V2GetPreflightRequirementsMethodNotAllowed {

	return &V2GetPreflightRequirementsMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 get preflight requirements method not allowed response
func (o *V2GetPreflightRequirementsMethodNotAllowed) WithPayload(payload *models.Error) *V2GetPreflightRequirementsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get preflight requirements method not allowed response
func (o *V2GetPreflightRequirementsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPreflightRequirementsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPreflightRequirementsInternalServerErrorCode is the HTTP code returned for type V2GetPreflightRequirementsInternalServerError
const V2GetPreflightRequirementsInternalServerErrorCode int = 500

/*V2GetPreflightRequirementsInternalServerError Error.

swagger:response v2GetPreflightRequirementsInternalServerError
*/
type V2GetPreflightRequirementsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPreflightRequirementsInternalServerError creates V2GetPreflightRequirementsInternalServerError with default headers values
func NewV2GetPreflightRequirementsInternalServerError() *V2GetPreflightRequirementsInternalServerError {

	return &V2GetPreflightRequirementsInternalServerError{}
}

// WithPayload adds the payload to the v2 get preflight requirements internal server error response
func (o *V2GetPreflightRequirementsInternalServerError) WithPayload(payload *models.Error) *V2GetPreflightRequirementsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get preflight requirements internal server error response
func (o *V2GetPreflightRequirementsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPreflightRequirementsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
