// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injapi/models"
)

// ActivityFavOKCode is the HTTP code returned for type ActivityFavOK
const ActivityFavOKCode int = 200

/*ActivityFavOK 获取列表

swagger:response activityFavOK
*/
type ActivityFavOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewActivityFavOK creates ActivityFavOK with default headers values
func NewActivityFavOK() *ActivityFavOK {
	return &ActivityFavOK{}
}

// WithPayload adds the payload to the activity fav o k response
func (o *ActivityFavOK) WithPayload(payload *models.InlineResponse2003) *ActivityFavOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity fav o k response
func (o *ActivityFavOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityFavOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrActivityFavDefault error

swagger:response activityFavDefault
*/
type NrActivityFavDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrActivityFavDefault creates NrActivityFavDefault with default headers values
func NewNrActivityFavDefault(code int) *NrActivityFavDefault {
	if code <= 0 {
		code = 500
	}

	return &NrActivityFavDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity fav default response
func (o *NrActivityFavDefault) WithStatusCode(code int) *NrActivityFavDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity fav default response
func (o *NrActivityFavDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity fav default response
func (o *NrActivityFavDefault) WithPayload(payload *models.Error) *NrActivityFavDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity fav default response
func (o *NrActivityFavDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrActivityFavDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
