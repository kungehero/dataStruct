package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

}

type Test struct{}

func (test *Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}
func hello(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	fmt.Fprintf(w, "%s", param.ByName("name"))
}
