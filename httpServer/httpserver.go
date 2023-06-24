package httpServer

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}

const SuffixFile = ".txt"
const PrefixFile = "./txt/"
const PermissinFile = 0600

func (page *Page) save() error {
	fileName := PrefixFile + page.Title + SuffixFile
	return os.WriteFile(fileName, page.Body, PermissinFile)
}

func loadPage(title string) (*Page, error) {
	fileName := PrefixFile + title + SuffixFile
	body, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")

	page := &Page{Title: title, Body: []byte(body)}
	err := page.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)

	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w, page, "./html/edit.html")
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, page, "./html/view.html")
}

func renderTemplate(w http.ResponseWriter, page *Page, templateFileName string) {
	temp, err := template.ParseFiles(templateFileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, page)

}

// -----------------------------------------------------

func SayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello user...\n")
	fmt.Printf("[%s]; status: %d\n", time.Now(), 200)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi there, I love %s!", r.URL.Path[1:])
}
