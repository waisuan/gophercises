package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func main() {
	f, err := ioutil.ReadFile("story.json")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to open file: %s", err))
	}

	var adventure map[string]Chapter
	if err = json.Unmarshal(f, &adventure); err != nil {
		log.Fatalf("Unable to parse JSON: %s", err)
	}

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Unrecognized page! Please turn back.")
	})

	tmpl := template.Must(template.ParseFiles("story.html"))
	for a := range adventure {
		path := fmt.Sprintf("/%s", a)
		r.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
			chapter := request.URL.String()[1:]
			tmpl.Execute(writer, adventure[chapter])
		})
	}

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Unable to start server: %s", err)
	}
}

func (chapter Chapter) FullStory() string {
	return strings.Join(chapter.Story, " ")
}
