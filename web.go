/*
    web.go -
 */
package main

import (
    "fmt"
    "net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Artistic Test Web Page")
}


