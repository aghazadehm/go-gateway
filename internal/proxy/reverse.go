package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy() http.Handler {
	target, _ := url.Parse("http://localhost:9000")

	proxy := httputil.NewSingleHostReverseProxy(target)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServerHTTP(w, r)
	})
}