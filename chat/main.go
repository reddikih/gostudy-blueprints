package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// temp1 represents a template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	// route
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// start a chat room
	go r.run()

	// start a web servrr
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
