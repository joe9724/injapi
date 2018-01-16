// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrActivityFavHandlerFunc turns a function with the right signature into a activity fav handler
type NrActivityFavHandlerFunc func(NrActivityFavParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrActivityFavHandlerFunc) Handle(params NrActivityFavParams) middleware.Responder {
	return fn(params)
}

// NrActivityFavHandler interface for that can handle valid activity fav params
type NrActivityFavHandler interface {
	Handle(NrActivityFavParams) middleware.Responder
}

// NewNrActivityFav creates a new http.Handler for the activity fav operation
func NewNrActivityFav(ctx *middleware.Context, handler NrActivityFavHandler) *NrActivityFav {
	return &NrActivityFav{Context: ctx, Handler: handler}
}

/*NrActivityFav swagger:route GET /activity/fav Activity activityFav

我收藏的活动列表

我收藏的活动列表

*/
type NrActivityFav struct {
	Context *middleware.Context
	Handler NrActivityFavHandler
}

func (o *NrActivityFav) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrActivityFavParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
