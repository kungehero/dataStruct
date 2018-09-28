package main

import (
	"dataStruct/stackzk"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	stackzk.Push(1)
	stackzk.Push(2)
	stackzk.Push(3)
	stackzk.Push(4)
	stackzk.Push(5)
	stackzk.Push(6)
	println(len(stackzk.Data))
}

type Test struct{}

func (test *Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}
func hello(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	fmt.Fprintln(w, "%s", param.ByName("name"))
}
