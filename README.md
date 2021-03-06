

# healthcheck
`import "github.com/jspc/go-health"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
healthcheck is the reference implementation of the Culture Trip Healthcheck API.

It can be instantiated as per:


	package main
	
	import (
	    "github.com/jspc/go-health"
	)
	
	func main() {
	    h := healthcheck.New()
	}

It exposes a fasthttp handler which can be mounted to do stuff




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Healthcheck](#Healthcheck)
  * [func (h *Healthcheck) Run()](#Healthcheck.Run)
* [type Healthchecks](#Healthchecks)
  * [func New(v Version) Healthchecks](#New)
  * [func (h *Healthchecks) Add(hc Healthcheck)](#Healthchecks.Add)
  * [func (h Healthchecks) Handle(ctx *fasthttp.RequestCtx)](#Healthchecks.Handle)
  * [func (h *Healthchecks) Start()](#Healthchecks.Start)
* [type Version](#Version)


#### <a name="pkg-files">Package files</a>
[api.go](/src/github.com/jspc/go-health/api.go) [doc.go](/src/github.com/jspc/go-health/doc.go) [healthcheck.go](/src/github.com/jspc/go-health/healthcheck.go) [healthchecks.go](/src/github.com/jspc/go-health/healthchecks.go) [version.go](/src/github.com/jspc/go-health/version.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    // Tick is the time to wait between healthcheck tests
    Tick = 10 * time.Second
)
```



## <a name="Healthcheck">type</a> [Healthcheck](/src/target/healthcheck.go?s=93:1300#L8)
``` go
type Healthcheck struct {
    // Name is a simple string to help explain the point of the healthcheck
    // It doesn't have to even be unique- it's just to make output more
    // readable
    Name string `json:"name"`

    // Readiness and Liveness determines what check endpoints are affected by
    // what healthchecks- it's also possible to set these as both false- this
    // will, essentially, make a healthcheck advisory only
    Readiness bool `json:"readiness"`
    Liveness  bool `json:"liveness"`

    // RunbookItem should point to the url and anchor, where possible, of
    // the healthcheck/ dependency. It may be in the runbook of the service
    // or it may be another runbook
    RunbookItem string `json:"runbook_item"`

    // F is the function that a healthcheck runs. It returns true for a
    // successful healthcheck, and false for a failing healthcheck.
    // It also returns optional output.
    F func() (bool, interface{}) `json:"-"`

    // The following are overwritten when a healthcheck is run
    State    string      `json:"state"` // enum: not_run, running, run
    LastRun  time.Time   `json:"last_run"`
    Duration float64     `json:"duration_ms"`
    Success  bool        `json:"success"`
    Output   interface{} `json:"output"`
}

```
Healthcheck represents an individual healthcheck










### <a name="Healthcheck.Run">func</a> (\*Healthcheck) [Run](/src/target/healthcheck.go?s=1338:1365#L39)
``` go
func (h *Healthcheck) Run()
```
Run will.... run the healthcheck




## <a name="Healthchecks">type</a> [Healthchecks](/src/target/healthchecks.go?s=213:382#L15)
``` go
type Healthchecks struct {
    ReportTime   time.Time     `json:"report_as_of"`
    Healthchecks []Healthcheck `json:"healthchecks"`
    // contains filtered or unexported fields
}

```
Healthchecks represent the state of a run set of healthchecks
and







### <a name="New">func</a> [New](/src/target/healthchecks.go?s=500:532#L25)
``` go
func New(v Version) Healthchecks
```
New returns a Healthchecks object which exposes an API and contains
timers and logic for running healthchecks





### <a name="Healthchecks.Add">func</a> (\*Healthchecks) [Add](/src/target/healthchecks.go?s=705:747#L35)
``` go
func (h *Healthchecks) Add(hc Healthcheck)
```
Add takes a Healthcheck and enrolls it into the Healthchecks thing




### <a name="Healthchecks.Handle">func</a> (Healthchecks) [Handle](/src/target/api.go?s=269:323#L14)
``` go
func (h Healthchecks) Handle(ctx *fasthttp.RequestCtx)
```
Handle provides an API endpoint/ router for healthcheck endpoints
which can be mounted into ct fasthttp apps.

It is compliant with github.com/beamly/go-http-middleware




### <a name="Healthchecks.Start">func</a> (\*Healthchecks) [Start](/src/target/healthchecks.go?s=853:883#L40)
``` go
func (h *Healthchecks) Start()
```
Start will run a timer doing healthchecks and stuff




## <a name="Version">type</a> [Version](/src/target/version.go?s=114:347#L5)
``` go
type Version struct {
    Name      string `json:"release_name"`
    Built     int64  `json:"built"`
    CircleSha string `json:"version"`
    Oracle    string `json:"oracle"`
    Runbook   string `json:"runbook"`
    Squad     string `json:"squad"`
}

```
Version exposes the version, version config, and basic
links to bits and bobs it needs














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
