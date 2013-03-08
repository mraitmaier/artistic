
package main

import (
    "fmt"
    "net/http"
    "bitbucket.org/miranr/artistic/common"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Artistic Test Web Page")
}

func main () {
    //testing import for local code
    p := artistic.CreatePainter()
    p.Name = artistic.CreateName("Vincent", "", "Van Gogh")
    fmt.Println(p.String())

    fmt.Println("Serving application on 'localhost:8080'...")

    http.HandleFunc("/", testHandler)
    http.ListenAndServe(":8080", nil)
}
