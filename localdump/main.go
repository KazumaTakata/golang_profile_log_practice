package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "<!--\n"+html.EscapeString(spew.Sdump(w))+"\n-->")
}

type dog struct {
	kind string
	age  int
	next *dog
}

func main() {
	printFunc()
}

func printFunc() {
	d := dog{kind: "siba", age: 2}
	d2 := dog{kind: "akita", age: 6, next: &d}

	fmt.Printf((spew.Sdump(d2)))
}
