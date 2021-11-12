package main

type Response struct {
	code int         `json:"code"`
	msg  string      `json:"msg"`
	data interface{} `json:"data"`
}

type Builder interface {
	Code(int) ResponseBuilder
	Msg(string) ResponseBuilder
	Data(interface{}) ResponseBuilder
	Build() Response
}

func (rb ResponseBuilder) Code(code int) ResponseBuilder {
	rb.code = code
	return rb
}

func (rb ResponseBuilder) Msg(msg string) ResponseBuilder {
	rb.msg = msg
	return rb
}

func (rb ResponseBuilder) Data(data interface{}) ResponseBuilder {
	rb.data = data
	return rb
}

func (rb ResponseBuilder) Build() Response {
	return Response{
		code: rb.code,
		msg:  rb.msg,
		data: rb.data,
	}
}

type ResponseBuilder struct {
	Response
}

func NewResponseBuilder() ResponseBuilder {
	return ResponseBuilder{}
}
