package response

import "github.com/hifat/goroger-core/rules"

var CodeCreated string = "CREATED"
var CodeOK string = "OK"

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`

	Status int `json:"-"`
}

type ResponseSuccess struct {
	Response

	Item  any `json:"item,omitempty"`
	Items any `json:"items,omitempty"`
}

type ResponseErr struct {
	Response
	Attribute rules.ValidateErrs `json:"attribute"`
}

func (e ResponseErr) Error() string {
	return e.Message
}
