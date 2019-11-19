package main

import "io/ioutil"

// The Page struct describes how page data is stored in memory
type Page struct {
	Title string
	Body  []byte // this means "byte slice". Body is of this type because our io libraries use this type, not string
}

// saves Page body to a text file named "title.txt" This is a method and not a func because it ˙has a receiver
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // the 0600 is a octo literall indicating this file has read-write permission for current user only.
}

// function local to this package, lowercase first letter to indicate that it's not exported.
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{
		Title: title,
		Body:  body,
	}
}
