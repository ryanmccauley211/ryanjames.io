package main

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

var email string
var pass string
var projectTodos = map[string]TodoList{}


func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/index.html")
	if (err != nil) {
		log.Fatalf("%v", err)
	}
	tmpl.Execute(w, nil)
}

func handleExplore(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/explore.html", "static/templates/explore-templates/pytorch-classification-iris.html")
	if (err != nil) {
		log.Fatalf("%v", err)
	}
	tmpl.Execute(w, nil)
}

func handlePortfolio(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/portfolio.html", "static/templates/project-templates/oak.html")
	if (err != nil) {
		log.Fatalf("%v", err)
	}
	tmpl.Execute(w, projectTodos)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/contact.html")
	if (err != nil) {
		log.Fatalf("%v", err)
	}

	formSubmitted := false
	if r.Method == "POST" {
		handleSend(r)
		formSubmitted = true
	}

	tmpl.Execute(w, formSubmitted)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/about.html")
	if (err != nil) {
		log.Fatalf("%v", err)
	}
	tmpl.Execute(w, nil)
}

func handleSend(response *http.Request) {
	response.ParseForm()
	formType := response.PostFormValue("form-type")
	if (formType == "hire") {
		name := response.PostFormValue("first-name") + " " + response.PostFormValue("last-name") +
			" contact: " + response.PostFormValue("contact-num");
		SendMail("Hire", name, response.PostFormValue("email"), response.PostFormValue("message"), email, pass, response.PostFormValue("contact-num"), "")
	} else if (formType == "contribute") {
		name := response.PostFormValue("first-name") + " " + response.PostFormValue("last-name");
		SendMail("Contribute to " + response.PostFormValue("project-name"), name,
			response.PostFormValue("email"), response.PostFormValue("message"), email, pass, "", response.PostFormValue("project-name"))
	} else if (formType == "other") {
		name := response.PostFormValue("first-name") + " " + response.PostFormValue("last-name");
		SendMail("General", name, response.PostFormValue("email"), response.PostFormValue("message"), email, pass, "", "")
	}
}

func main() {

	email = os.Args[1]
	pass = os.Args[2]

	fmt.Printf("Initializing...\n")
	getTodoList("oak", "https://github.com/ryanmccauley211/Oak/blob/master/README.md")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/explore", handleExplore)
	http.HandleFunc("/portfolio", handlePortfolio)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/about", handleAbout)

	fmt.Println("Serving...\n")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func getTodoList(projectName string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		checkedTasks := []string{}
		uncheckedTasks := []string{}
		delTasks := []string{}
		doc.Find(".contains-task-list li").Each(func(i int, s *goquery.Selection) {
			input := s.Find("input")
			_, checked := input.Attr("checked")
			del := s.Find("del")

			if checked {
				checkedTasks = append(checkedTasks, s.Text())
			} else if del.Size() > 0 {
				delTasks = append(delTasks, del.Text())
			} else {
				uncheckedTasks = append(uncheckedTasks, s.Text())
			}
			projectTodos[projectName] = TodoList{checkedTasks, uncheckedTasks, delTasks, true}
		})
	}
	return nil
}

type TodoList struct {
	Checked   []string
	Unchecked []string
	Deleted   []string
	Loaded    bool
}