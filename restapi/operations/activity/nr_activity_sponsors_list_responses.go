// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injapi/models"
)

// ActivitySponsorsListOKCode is the HTTP code returned for type ActivitySponsorsListOK
const ActivitySponsorsListOKCode int = 200

/*ActivitySponsorsListOK 获取列表

swagger:response activitySponsorsListOK
*/
type ActivitySponsorsListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2009 `json:"body,omitempty"`
}

// NewActivitySponsorsListOK creates ActivitySponsorsListOK with default headers values
func NewActivitySponsorsListOK() *ActivitySponsorsListOK {
	return &ActivitySponsorsListOK{}
}

// WithPayload adds the payload to the activity sponsors list o k response
func (o *ActivitySponsorsListOK) WithPayload(payload *models.InlineResponse2009) *ActivitySponsorsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity sponsors list o k response
func (o *ActivitySponsorsListOK) SetPayload(payload *models.InlineResponse2009) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivitySponsorsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrActivitySponsorsListDefault error

swagger:response activitySponsorsListDefault
*/
type NrActivitySponsorsListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrActivitySponsorsListDefault creates NrActivitySponsorsListDefault with default headers values
func NewNrActivitySponsorsListDefault(code int) *NrActivitySponsorsListDefault {
	if code <= 0 {
		code = 500
	}

	return &NrActivitySponsorsListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity sponsors list default response
func (o *NrActivitySponsorsListDefault) WithStatusCode(code int) *NrActivitySponsorsListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity sponsors list default response
func (o *NrActivitySponsorsListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity sponsors list default response
func (o *NrActivitySponsorsListDefault) WithPayload(payload *models.Error) *NrActivitySponsorsListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity sponsors list default response
func (o *NrActivitySponsorsListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrActivitySponsorsListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
