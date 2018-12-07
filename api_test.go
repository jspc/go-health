package healthcheck

import (
	"testing"

	"github.com/valyala/fasthttp"
)

func TestHealthchecks_Serve(t *testing.T) {
	for _, test := range []struct {
		name   string
		method string
		path   string
		expect int
	}{
		{"PUT on /", "PUT", "/", 405},
		{"DELETE on /", "DELETE", "/", 405},
		{"PATCH on /", "PATCH", "/", 405},
		{"HEAD on /", "HEAD", "/", 405},
		{"TRACE on /", "TRACE", "/", 405},
		{"CONNECT on /", "CONNECT", "/", 405},

		{"GET on /poop", "GET", "/poop", 404},

		{"GET on /a/b/c/version", "GET", "/a/b/c/version", 200},
		{"GET on //aaaaa/readiness", "GET", "//aaaaa/readiness", 200},
		{"GET on /liveness", "GET", "/liveness", 200},
		{"GET on /", "GET", "/", 200},
	} {
		t.Run(test.name, func(t *testing.T) {
			req := fasthttp.AcquireRequest()
			req.SetRequestURI(test.path)
			req.Header.SetMethod(test.method)

			resp := fasthttp.AcquireResponse()

			c := &fasthttp.RequestCtx{
				Request:  *req,
				Response: *resp,
			}

			Healthchecks{}.Handle(c)
			s := c.Response.StatusCode()
			if test.expect != s {
				t.Errorf("expected status %d, received %d", test.expect, s)
			}
		})
	}
}
