package response

import "github.com/hifat/goroger-core/rules"

var CodeCreated string = "CREATED"
var CodeOK string = "OK"

type Response struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`

	Status int `json:"-"`
}

type ResponseMeta struct {
	Total int64 `json:"total"`
}

type ResponseSuccess struct {
	Response

	Item  any           `json:"item,omitempty"`
	Items any           `json:"items,omitempty"`
	Meta  *ResponseMeta `json:"meta,omitempty"`
}

type ResponseErr struct {
	Response
	Attribute rules.ValidateErrs `json:"attribute"`
}

func (e ResponseErr) Error() string {
	return e.Message
}
