// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdatePolicyOKCode is the HTTP code returned for type UpdatePolicyOK
const UpdatePolicyOKCode int = 200

/*UpdatePolicyOK OK

swagger:response updatePolicyOK
*/
type UpdatePolicyOK struct {

	/*
	  In: Body
	*/
	Payload *UpdatePolicyOKBody `json:"body,omitempty"`
}

// NewUpdatePolicyOK creates UpdatePolicyOK with default headers values
func NewUpdatePolicyOK() *UpdatePolicyOK {

	return &UpdatePolicyOK{}
}

// WithPayload adds the payload to the update policy o k response
func (o *UpdatePolicyOK) WithPayload(payload *UpdatePolicyOKBody) *UpdatePolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy o k response
func (o *UpdatePolicyOK) SetPayload(payload *UpdatePolicyOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyBadRequestCode is the HTTP code returned for type UpdatePolicyBadRequest
const UpdatePolicyBadRequestCode int = 400

/*UpdatePolicyBadRequest bad request

swagger:response updatePolicyBadRequest
*/
type UpdatePolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *UpdatePolicyBadRequestBody `json:"body,omitempty"`
}

// NewUpdatePolicyBadRequest creates UpdatePolicyBadRequest with default headers values
func NewUpdatePolicyBadRequest() *UpdatePolicyBadRequest {

	return &UpdatePolicyBadRequest{}
}

// WithPayload adds the payload to the update policy bad request response
func (o *UpdatePolicyBadRequest) WithPayload(payload *UpdatePolicyBadRequestBody) *UpdatePolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy bad request response
func (o *UpdatePolicyBadRequest) SetPayload(payload *UpdatePolicyBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyUnauthorizedCode is the HTTP code returned for type UpdatePolicyUnauthorized
const UpdatePolicyUnauthorizedCode int = 401

/*UpdatePolicyUnauthorized Unauthorized

swagger:response updatePolicyUnauthorized
*/
type UpdatePolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *UpdatePolicyUnauthorizedBody `json:"body,omitempty"`
}

// NewUpdatePolicyUnauthorized creates UpdatePolicyUnauthorized with default headers values
func NewUpdatePolicyUnauthorized() *UpdatePolicyUnauthorized {

	return &UpdatePolicyUnauthorized{}
}

// WithPayload adds the payload to the update policy unauthorized response
func (o *UpdatePolicyUnauthorized) WithPayload(payload *UpdatePolicyUnauthorizedBody) *UpdatePolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy unauthorized response
func (o *UpdatePolicyUnauthorized) SetPayload(payload *UpdatePolicyUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyForbiddenCode is the HTTP code returned for type UpdatePolicyForbidden
const UpdatePolicyForbiddenCode int = 403

/*UpdatePolicyForbidden Forbidden

swagger:response updatePolicyForbidden
*/
type UpdatePolicyForbidden struct {
}

// NewUpdatePolicyForbidden creates UpdatePolicyForbidden with default headers values
func NewUpdatePolicyForbidden() *UpdatePolicyForbidden {

	return &UpdatePolicyForbidden{}
}

// WriteResponse to the client
func (o *UpdatePolicyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// UpdatePolicyNotFoundCode is the HTTP code returned for type UpdatePolicyNotFound
const UpdatePolicyNotFoundCode int = 404

/*UpdatePolicyNotFound resource not found

swagger:response updatePolicyNotFound
*/
type UpdatePolicyNotFound struct {
}

// NewUpdatePolicyNotFound creates UpdatePolicyNotFound with default headers values
func NewUpdatePolicyNotFound() *UpdatePolicyNotFound {

	return &UpdatePolicyNotFound{}
}

// WriteResponse to the client
func (o *UpdatePolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
