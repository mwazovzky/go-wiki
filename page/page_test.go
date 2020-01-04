package page

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	success = "\u2713"
	failed = "\u2717"
)

func TestMain(m *testing.M) {
	os.Mkdir("data", os.FileMode(0522)) // setup
	m.Run()
	os.RemoveAll("data") // teardown
 }

func TestSave(t *testing.T) {
	body := "Lorem ipsum"
	p := Page{Title: "TestPage", Body: []byte(body)}
	err := p.Save()

	if err != nil {
		t.Fatalf("\t%s Should save file without errors, got error: %v", failed, err)
	}

	content, err := ioutil.ReadFile("data/TestPage.txt")

	if err != nil {
		t.Fatalf("\t%s Should be able to create a file with the name %s, error: %v", failed, p.Title, err)
	}
	t.Logf("\t%s Should be able to create a file with the name %s", success, p.Title)

	if body != string(content) {
		t.Fatalf("\t%s File content should be equal to page body, expected: %s, got: %s", failed, body, string(content))
	}
	t.Logf("\t%s File content should be equal to page body", success)
}

func TestLoadPage(t *testing.T)  {
	p := Page{Title: "TestPage", Body: []byte("Lorem ipsum")}
	p.Save()

	page, err := LoadPage("TestPage")
	
	if err != nil {
		t.Fatalf("\t%s Should load page without errors, got error: %v", failed, err)
	}

	if page.Title != "TestPage" {
		t.Fatalf("\t%s Page title should be equal to %s, got: %s", failed, "TestPage", page.Title)
	}
	t.Logf("\t%s Page title should be equal to %s", success, "TestPage")

	if string(page.Body) != "Lorem ipsum" {
		t.Fatalf("\t%s Page body should be equal to %s, got: %s", failed, "Lorem ipsum", string(page.Body))
	}
	t.Logf("\t%s Page body should be equal to %s", success, "Lorem ipsum")
}
