# CRUD API in Go

## Overview

This example creates an in-memory REST-style CRUD API for student records using
Go's standard `net/http` package. The server stores records in a package-level
slice and protects shared access with a `sync.Mutex`.

## Important Points

- `http.HandleFunc` connects the `/students` path to `studentHandler`.
- `r.Method` selects the CRUD operation.
- `r.URL.Query().Get("id")` reads the student ID from the query string.
- `json.Decoder` decodes request bodies.
- `json.Encoder` returns JSON responses.
- `sync.Mutex` prevents concurrent requests from corrupting the in-memory data.
- Data is not persistent: restarting the server clears all records.