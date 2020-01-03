package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := "pages/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "pages/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main()  {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a test page body")}
	p1.save()
	p2, err := loadPage("TestPage")
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(string(p2.Body))
}
