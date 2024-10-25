package web

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func UnMarshalJSON(request *http.Request, out interface{}) error {
	//validation
	return json.NewDecoder(request.Body).Decode(out)

}

type Parser struct {
	Body   io.ReadCloser
	Form   url.Values
	Params map[string]string
}

func NewParser(request *http.Request) *Parser {
	return &Parser{
		Body:   request.Body,
		Form:   request.Form,
		Params: mux.Vars(request),
	}

}
