package api

import (
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func start() {
	http.HandleFunc("/tasks/", makeHandler(viewHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
