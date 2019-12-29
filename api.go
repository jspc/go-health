package healthcheck

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// Handle provides an API endpoint/ router for healthcheck endpoints
// which can be mounted into ct fasthttp apps.
//
// It is compliant with github.com/beamly/go-http-middleware
func (h *Healthchecks) Handle(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	if !ctx.IsGet() {
		sendError(ctx, "method not supported", fasthttp.StatusMethodNotAllowed)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	routes := map[string]func(*fasthttp.RequestCtx){
		"version":   h.handleVersion,
		"readiness": h.handleReadiness,
		"liveness":  h.handleLiveness,
		"":          h.handleAll,
	}

	path := string(ctx.URI().LastPathSegment())
	if r, ok := routes[path]; ok {
		r(ctx)
	} else {
		sendError(ctx, "not found", fasthttp.StatusNotFound)

		return
	}
}

func (h *Healthchecks) handleVersion(ctx *fasthttp.RequestCtx) {
	sendJson(ctx, h.Version)
}

func (h *Healthchecks) handleAll(ctx *fasthttp.RequestCtx) {
	sendJson(ctx, h)
}

func (h *Healthchecks) handleReadiness(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(h.statusFromType("Readiness"))
}

func (h *Healthchecks) handleLiveness(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(h.statusFromType("Liveness"))
}

func (h Healthchecks) statusFromType(t string) int {
	for _, hc := range h.byType(t) {
		if !hc.Success {
			return fasthttp.StatusServiceUnavailable
		}
	}

	return fasthttp.StatusOK
}

func sendJson(ctx *fasthttp.RequestCtx, i interface{}) {
	o, err := json.Marshal(i)
	if err != nil {
		sendError(ctx, "unable to marshal json", 500)
	}

	fmt.Fprintf(ctx, string(o))
}

func sendError(ctx *fasthttp.RequestCtx, msg string, code int) {
	fmt.Fprintf(ctx, fmt.Sprintf(`{"status": %q}`, msg))

	ctx.SetStatusCode(code)
}
