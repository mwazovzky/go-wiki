package page

import (
	"io/ioutil"
)

// Page represents web page
type Page struct {
	Title string
	Body []byte
}

// Save page to a file
func (p *Page) Save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage from a file
func LoadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
