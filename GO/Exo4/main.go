package main

import (
    "log"
    "net/http"
	"text/template"

)
type Form struct {
	formulaire string
}





func main() {
	var f Form
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
    http.HandleFunc("/",f.InitPage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func (f *Form) InitPage(w http.ResponseWriter, r *http.Request) {
	a := template.Must(template.ParseFiles("register.html"))
	details := Form{
		formulaire: f.retour(w,r),
	}
	a.Execute(w, details)
}

func (f *Form)retour(w http.ResponseWriter, r *http.Request) string{
	println(r.FormValue("letter"))
	return r.FormValue("letter")
}