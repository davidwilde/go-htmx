package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Results struct {
	People     []Person
	TotalCount int
}

type Event struct {
	Title     string
	EventType int
}

type Person struct {
	Name        string
	LatestEvent Event
}

type SavedSearch struct {
	Name string
}

type Details struct {
	SmartSearches []SavedSearch
	Action        string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		// use a template to render the search results
		t := template.Must(template.ParseFiles("base.html", "index.html"))

		data := []Person{
			{Name: "Peter", LatestEvent: Event{Title: "Peter's Event", EventType: 1}},
			{
				Name: "John",
				LatestEvent: Event{
					Title:     "John's Event",
					EventType: 1,
				},
			},
			{
				Name: "Jane",
				LatestEvent: Event{
					Title:     "Jane's Event",
					EventType: 2,
				},
			},
		}

		eventTypeString := r.URL.Query().Get("eventType")
		action := r.URL.Query().Get("action")

		if action == "save-as-smart-search" {
			// redirect to /smart-search but with the query string
			http.Redirect(w, r, "/smart-search?eventType="+eventTypeString, 302)
			return
		}

		subset := []Person{}
		eventType, _ := strconv.Atoi(eventTypeString)

		for _, p := range data {
			if p.LatestEvent.EventType == eventType {
				subset = append(subset, p)
			}
		}

		results := Results{
			People:     subset,
			TotalCount: len(subset),
		}

		if r.Header.Get("Hx-Target") == "results" {
			t.ExecuteTemplate(w, "results", results)
			return
		}

		t.Execute(w, results)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		// use a template to render the search results
		t := template.Must(template.ParseFiles("person.html"))
		t.Execute(w, nil)
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {
		// use a template to render the smart search form
		action := r.URL.Query().Get("action")

		data := []SavedSearch{
			{
				Name: "All people who worked at Initech",
			},
			{
				Name: "New SDRs from US and Canada",
			},
		}

		details := Details{
			SmartSearches: data,
			Action:        action,
		}

		t := template.Must(template.ParseFiles("base.html", "smart-search.html"))
		if r.Header.Get("Hx-Target") == "fullscreen-overlay" {
			t.ExecuteTemplate(w, "fullscreen-overlay", details)
			return
		}
		t.Execute(w, details)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/person", h2)
	http.HandleFunc("/smart-search", h3)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
