package main

import (
	"fmt"
    "log"
    "net/http"
    "html/template"
    "strconv"
)

type Results struct {
    People []Person
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

    h1 := func(w http.ResponseWriter, r *http.Request) {
		// use a template to render the search results
		t  := template.Must(template.ParseFiles("index.html"))

        data := []Person{
                Person{
                    Name: "John",
                    LatestEvent: Event{
                        Title: "John's Event",
                        EventType: 1,
                    },
                },
                Person{
                    Name: "Jane",
                    LatestEvent: Event{
                        Title: "Jane's Event",
                        EventType: 2,
                    },
                },
            }

        eventTypeString := r.URL.Query().Get("eventType")
        subset := []Person{}
        eventType, err := strconv.Atoi(eventTypeString)

        if (err != nil) {
            subset = data
        } else {
            for _, p := range data {
                if (p.LatestEvent.EventType == eventType) {
                    subset = append(subset, p)
                }
            }
        }

        results := Results{
            People: subset,
            TotalCount: len(subset),
        }

        if (r.Header.Get("Hx-Target") == "results") {
            t.ExecuteTemplate(w, "results", results)
            return
        }

        t.Execute(w, results)
	}
    http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
