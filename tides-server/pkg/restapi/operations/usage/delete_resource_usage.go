// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteResourceUsageHandlerFunc turns a function with the right signature into a delete resource usage handler
type DeleteResourceUsageHandlerFunc func(DeleteResourceUsageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteResourceUsageHandlerFunc) Handle(params DeleteResourceUsageParams) middleware.Responder {
	return fn(params)
}

// DeleteResourceUsageHandler interface for that can handle valid delete resource usage params
type DeleteResourceUsageHandler interface {
	Handle(DeleteResourceUsageParams) middleware.Responder
}

// NewDeleteResourceUsage creates a new http.Handler for the delete resource usage operation
func NewDeleteResourceUsage(ctx *middleware.Context, handler DeleteResourceUsageHandler) *DeleteResourceUsage {
	return &DeleteResourceUsage{Context: ctx, Handler: handler}
}

/* DeleteResourceUsage swagger:route DELETE /usage/{id} usage deleteResourceUsage

delete resource usage info

*/
type DeleteResourceUsage struct {
	Context *middleware.Context
	Handler DeleteResourceUsageHandler
}

func (o *DeleteResourceUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteResourceUsageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteResourceUsageOKBody delete resource usage o k body
//
// swagger:model DeleteResourceUsageOKBody
type DeleteResourceUsageOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete resource usage o k body
func (o *DeleteResourceUsageOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete resource usage o k body based on context it is used
func (o *DeleteResourceUsageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteResourceUsageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteResourceUsageOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteResourceUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
