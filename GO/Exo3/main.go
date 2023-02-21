package main

import (
    "log"
    "net/http"
	"math/rand"
	"time"
	"text/template"

)
type Case struct {
	Cases  int
	Numbers  []int
}





func main() {
	var c Case
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
    http.HandleFunc("/",c.InitPage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func (c *Case) InitPage(w http.ResponseWriter, r *http.Request) {
	a := template.Must(template.ParseFiles("index.html"))
	details := Case{
		Numbers: c.Generate(r, w),
	}
	a.Execute(w, details)
}
func (c *Case) Generate(r *http.Request, w http.ResponseWriter) []int {
	max := 25
	min := 3
	var tab []int
	c.Numbers = tab
	rand.Seed(time.Now().UnixNano())
	c.Cases = rand.Intn(max-min) + min
	for i := 0; i < c.Cases; i++ {
		c.Numbers = append(c.Numbers, rand.Int())
	}
	return c.Numbers
}

