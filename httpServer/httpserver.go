package httpServer

import (
	"errors"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

const SuffixFile = ".txt"
const PrefixFile = "./txt/"
const PermissinFile = 0600

var templates = template.Must(template.ParseFiles("./html/edit.html", "./html/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {

	result := validPath.FindStringSubmatch(r.URL.Path)

	if result == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return result[2], nil
}

func (page *Page) savePage() error {
	fileName := PrefixFile + page.Title + SuffixFile
	return os.WriteFile(fileName, page.Body, PermissinFile)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title, er := getTitle(w, r)
	if er != nil {
		return
	}

	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}

	err := page.savePage()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func viewPage(title string) (*Page, error) {
	fileName := PrefixFile + title + SuffixFile
	body, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	page, err := viewPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view.html", page)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	page, err := viewPage(title)

	if err != nil {
		page = &Page{Title: title}
		page.savePage()
	}
	renderTemplate(w, "edit.html", page)
}

func renderTemplate(w http.ResponseWriter, htmlTemplatePath string, page *Page) {
	err := templates.ExecuteTemplate(w, htmlTemplatePath, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
