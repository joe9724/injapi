// Code generated by go-swagger; DO NOT EDIT.

package praise

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "injapi/models"
)

// PraiseClickOKCode is the HTTP code returned for type PraiseClickOK
const PraiseClickOKCode int = 200

/*PraiseClickOK 详情

swagger:response praiseClickOK
*/
type PraiseClickOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewPraiseClickOK creates PraiseClickOK with default headers values
func NewPraiseClickOK() *PraiseClickOK {
	return &PraiseClickOK{}
}

// WithPayload adds the payload to the praise click o k response
func (o *PraiseClickOK) WithPayload(payload *models.InlineResponse2003) *PraiseClickOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the praise click o k response
func (o *PraiseClickOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PraiseClickOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrPraiseClickDefault error

swagger:response praiseClickDefault
*/
type NrPraiseClickDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrPraiseClickDefault creates NrPraiseClickDefault with default headers values
func NewNrPraiseClickDefault(code int) *NrPraiseClickDefault {
	if code <= 0 {
		code = 500
	}

	return &NrPraiseClickDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the praise click default response
func (o *NrPraiseClickDefault) WithStatusCode(code int) *NrPraiseClickDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the praise click default response
func (o *NrPraiseClickDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the praise click default response
func (o *NrPraiseClickDefault) WithPayload(payload *models.Error) *NrPraiseClickDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the praise click default response
func (o *NrPraiseClickDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrPraiseClickDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
