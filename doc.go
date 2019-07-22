/*
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
*/
package healthcheck
