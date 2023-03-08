package main

import (
	"log"
	"net/http"
	"text/template"
)

//	func StaticHandler(w http.ResponseWriter, r *http.Request) {
//		f, err := os.Open("static" + r.URL.Path)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		if strings.HasSuffix(r.URL.Path, ".css") {
//			w.Header().Add("Content-Type", "text/css")
//		}
//		io.Copy(w, f)
//
// }

func RenderTemplate(w http.ResponseWriter, page string) {
	tp, err := template.ParseFiles("templates/" + page + ".page.tmpl")
	if err != nil {
		log.Println(err)
	}

	tp.Execute(w, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")
}

func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/contact", ContactHandler)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "about")
	})

	http.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", nil)
}
