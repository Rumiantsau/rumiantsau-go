package main
import (
    "fmt"
    "net/http"
)
type msg string
func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
   fmt.Fprint(resp, m) 
}
func main() {
    msgHandler := msg("Rumiantsau golang-app. Hello World 2")
    http.ListenAndServe("0.0.0.0:8080", msgHandler)
}
