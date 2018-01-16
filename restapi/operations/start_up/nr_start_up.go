// Code generated by go-swagger; DO NOT EDIT.

package start_up

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrStartUpHandlerFunc turns a function with the right signature into a start up handler
type NrStartUpHandlerFunc func(NrStartUpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrStartUpHandlerFunc) Handle(params NrStartUpParams) middleware.Responder {
	return fn(params)
}

// NrStartUpHandler interface for that can handle valid start up params
type NrStartUpHandler interface {
	Handle(NrStartUpParams) middleware.Responder
}

// NewNrStartUp creates a new http.Handler for the start up operation
func NewNrStartUp(ctx *middleware.Context, handler NrStartUpHandler) *NrStartUp {
	return &NrStartUp{Context: ctx, Handler: handler}
}

/*NrStartUp swagger:route GET /startUp startUp startUp

启动页接口

启动页接口

*/
type NrStartUp struct {
	Context *middleware.Context
	Handler NrStartUpHandler
}

func (o *NrStartUp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrStartUpParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
