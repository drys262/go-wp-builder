package main


import (
	"fmt"
	"net/http"
)



func main() {
    
    http.HandleFunc("/", home)

    http.ListenAndServe(":80", nil)
    
}


func home(w http.ResponseWriter, r *http.Request) {

    fmt.Fprint(w, "<h1>WP BUILDER</h2>")
}
