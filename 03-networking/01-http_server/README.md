# HTTP Servers in Go

## Overview

This lesson introduces HTTP server development using Go's standard `net/http`
package. The examples build from a basic server to request cancellation,
routing, HTTP methods, request bodies, and JSON APIs.

## Important Points

- `http.HandleFunc` maps URL paths to handler functions.
- A handler receives an `http.ResponseWriter` for the response and
  `*http.Request` for request data.
- `http.ListenAndServe(":8000", nil)` starts a server using the default
  `ServeMux`.
- Use `r.Method` and the `http.Method...` constants to handle HTTP methods.
- Use `r.URL.Query()` to read query parameters.
- Decode JSON request bodies with `json.Decoder` and encode JSON responses with
  `json.Encoder`.
- Return `http.StatusMethodNotAllowed` for unsupported HTTP methods.
- Use `r.Context().Done()` to stop work when the client cancels a request.

## Key Takeaway

**Go's standard `net/http` package provides the core tools needed to build
HTTP servers and JSON APIs without a web framework.**