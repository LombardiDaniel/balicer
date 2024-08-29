package handlers

import (
	"net/textproto"
	"strings"
	"time"

	"github.com/patos-ufscar/balicer/models"
)

type HandlerProxyL7Impl struct {
	path       string
	addHeaders map[string]string
	ttl        time.Duration
}

func NewHandlerProxyL7Impl(path string, ret models.ReturnProxy) Handler {
	return &HandlerProxyL7Impl{
		path:       path,
		addHeaders: ret.AddHeaders,
		ttl:        ret.TTL,
	}
}

func (h *HandlerProxyL7Impl) ValidPath(path string) bool {
	return strings.HasPrefix(path, h.path)
}

func (h *HandlerProxyL7Impl) Handle(req models.HttpRequest) (models.HttpResponse, error) {

	resp := models.NewHttpResponse()

	for k, v := range h.addHeaders {
		resp.Headers[textproto.CanonicalMIMEHeaderKey(k)] = v
	}

	// resp.StatusText = http.StatusText(h.StatusCode)
	// resp.HTTPVersion = "HTTP/1.1"
	// resp.Body = h.Body

	return resp, nil
}
