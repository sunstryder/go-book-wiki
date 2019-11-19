package main

import (
	"fmt"
	"io/ioutil"
)

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
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{
		Title: title,
		Body:  body,
	}, nil
}

func main() {
	// create a new Page struct with the given data
	p1 := &Page{Title: "The Name of the Wind", Body: []byte("It was a silence of three parts")}
	//save it to a file
	p1.save()
	// load the contents of the file
	p2, _ := loadPage("The Name of the Wind")
	// spit out the body into console
	fmt.Println(string(p2.Body))
}
