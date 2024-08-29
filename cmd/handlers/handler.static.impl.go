package handlers

import (
	"net/http"
	"net/textproto"
	"strings"

	"github.com/patos-ufscar/balicer/models"
)

type HandlerStaticImpl struct {
	path       string
	statusCode int
	headers    map[string]string
	body       []byte
}

func NewHandlerStaticImpl(path string, ret models.ReturnStatic) Handler {
	// headers := make(map[string]string)
	return &HandlerStaticImpl{
		path:       path,
		statusCode: ret.Code,
		headers:    ret.Headers,
		body:       ret.Body,
	}
}

func (h *HandlerStaticImpl) ValidPath(path string) bool {
	return strings.HasPrefix(path, h.path)
}

func (h *HandlerStaticImpl) Handle(req models.HttpRequest) (models.HttpResponse, error) {
	resp := models.NewHttpResponse()

	for k, v := range h.headers {
		resp.Headers[textproto.CanonicalMIMEHeaderKey(k)] = v
	}

	resp.StatusCode = h.statusCode
	resp.StatusText = http.StatusText(h.statusCode)
	resp.HTTPVersion = "HTTP/1.1"
	resp.Body = h.body

	return resp, nil
}
