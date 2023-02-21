package main

import (
    "log"
    "net/http"
	"text/template"
	"fmt"

)
type Input struct {
	button int
}





func main() {
	var i Input
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
    http.HandleFunc("/",i.InitPage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func (i *Input) InitPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Page not found")
		return
	}
	a := template.Must(template.ParseFiles("index.html"))
	details := Input{
		button: i.redirect(w,r),
	}
	a.Execute(w, details)
}

func (i *Input)redirect(w http.ResponseWriter, r *http.Request) int{
	println(r.PostFormValue("bt"))
	if r.PostFormValue("bt") == "Github" {
		http.Redirect(w, r, "https://github.com/NoaYnov", http.StatusSeeOther)
	} else if r.PostFormValue("bt") == "Instagram" {
		http.Redirect(w, r, "https://www.instagram.com/boosted_esport/", http.StatusSeeOther)
	} else if r.PostFormValue("bt") == "Twitter" {
		http.Redirect(w, r, "https://twitter.com/boosted_esport", http.StatusSeeOther)
	}
	return 0
}