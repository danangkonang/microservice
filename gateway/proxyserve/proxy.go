package proxyserve

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyServe(p *httputil.ReverseProxy, target *url.URL) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// log.Println(r.URL)
		// log.Println(target.Host)
		// fmt.Println(r)
		r.Host = target.Host
		// w.Header().Set("X-Ben", "Rad")
		p.ServeHTTP(w, r)
	}
}
