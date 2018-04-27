package app

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/", handler_index)
	http.HandleFunc("/view", handler_index)
	http.HandleFunc("/ds", dsput)

}

// for Template.go

func handler_index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<body>")
	pageMaker(w, r)
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}

type FEEDITEM struct {
	Title   string
	Url     string
	GetDate time.Time
}

func dsput(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	e1 := FEEDITEM{Title: "ahoabc", Url: "https://bbb.blogs.example.com", GetDate: time.Now()}

	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Geeed", nil), &e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var e2 FEEDITEM
	if err = datastore.Get(ctx, key, &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<body>")

	fmt.Fprintf(w, "Stored retrieved %q ::: %q", e2.Title, e2.Url)

	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")

}
