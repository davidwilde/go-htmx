package main

import (
	"fmt"
    "log"
    "net/http"
    "html/template"
)

type Results struct {
    Events []Event
    TotalCount int
}

type Event struct {
    Title string
    EventType int
}

type Person struct {
    Name string
    LatestEvent Event
}

func main() {
    fmt.Println("Hello, World!")

    h1 := func(w http.ResponseWriter, r *http.Request) {
		// use a template to render the search results
		t  := template.Must(template.ParseFiles("index.html"))
        people := map[string][]Person{
            "Results": {
                {Name: "John", LatestEvent: Event{Title: "Birthday", EventType: 1}},
                {Name: "Jane", LatestEvent: Event{Title: "Wedding", EventType: 2}},
            },
            "Count": 
        }
        t.Execute(w, people)
	}
    http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
