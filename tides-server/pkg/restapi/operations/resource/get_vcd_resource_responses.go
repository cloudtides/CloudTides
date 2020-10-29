// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetVcdResourceOKCode is the HTTP code returned for type GetVcdResourceOK
const GetVcdResourceOKCode int = 200

/*GetVcdResourceOK return vcd info

swagger:response getVcdResourceOK
*/
type GetVcdResourceOK struct {

	/*
	  In: Body
	*/
	Payload *GetVcdResourceOKBody `json:"body,omitempty"`
}

// NewGetVcdResourceOK creates GetVcdResourceOK with default headers values
func NewGetVcdResourceOK() *GetVcdResourceOK {

	return &GetVcdResourceOK{}
}

// WithPayload adds the payload to the get vcd resource o k response
func (o *GetVcdResourceOK) WithPayload(payload *GetVcdResourceOKBody) *GetVcdResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vcd resource o k response
func (o *GetVcdResourceOK) SetPayload(payload *GetVcdResourceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVcdResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVcdResourceUnauthorizedCode is the HTTP code returned for type GetVcdResourceUnauthorized
const GetVcdResourceUnauthorizedCode int = 401

/*GetVcdResourceUnauthorized Unauthorized

swagger:response getVcdResourceUnauthorized
*/
type GetVcdResourceUnauthorized struct {
}

// NewGetVcdResourceUnauthorized creates GetVcdResourceUnauthorized with default headers values
func NewGetVcdResourceUnauthorized() *GetVcdResourceUnauthorized {

	return &GetVcdResourceUnauthorized{}
}

// WriteResponse to the client
func (o *GetVcdResourceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}