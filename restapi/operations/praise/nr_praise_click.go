// Code generated by go-swagger; DO NOT EDIT.

package praise

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrPraiseClickHandlerFunc turns a function with the right signature into a praise click handler
type NrPraiseClickHandlerFunc func(NrPraiseClickParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrPraiseClickHandlerFunc) Handle(params NrPraiseClickParams) middleware.Responder {
	return fn(params)
}

// NrPraiseClickHandler interface for that can handle valid praise click params
type NrPraiseClickHandler interface {
	Handle(NrPraiseClickParams) middleware.Responder
}

// NewNrPraiseClick creates a new http.Handler for the praise click operation
func NewNrPraiseClick(ctx *middleware.Context, handler NrPraiseClickHandler) *NrPraiseClick {
	return &NrPraiseClick{Context: ctx, Handler: handler}
}

/*NrPraiseClick swagger:route GET /praise/click Praise praiseClick

点赞

点赞

*/
type NrPraiseClick struct {
	Context *middleware.Context
	Handler NrPraiseClickHandler
}

func (o *NrPraiseClick) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrPraiseClickParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}