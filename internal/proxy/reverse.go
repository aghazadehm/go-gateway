package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy() http.Handler {
	target, _ := url.Parse("http://localhost:9000")

	return httputil.NewSingleHostReverseProxy(target)
}