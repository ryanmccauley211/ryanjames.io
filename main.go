package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"net/http"
)

var projectTodos = map[string]TodoList{}
var smtpCred = SmtpCred{}

func handleIndex(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/index.html")
	checkCriticalErr(err)
	checkCriticalErr(tmpl.Execute(w, nil))
}

func handleExplore(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/explore.html")

	checkCriticalErr(err)
	_, err = tmpl.ParseGlob("static/templates/explore-templates/modals/*.html")
	checkCriticalErr(err)

	_, err = tmpl.ParseGlob("static/templates/explore-templates/summaries/*.html")
	checkCriticalErr(err)

	err = tmpl.ExecuteTemplate(w, "explore.html", getExploreTemplateAliases(tmpl))
	checkCriticalErr(err)
}

func handlePortfolio(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/portfolio.html")
	checkCriticalErr(err)

	_, err = tmpl.ParseGlob("static/templates/portfolio-templates/modals/*.html")
	checkCriticalErr(err)

	_, err = tmpl.ParseGlob("static/templates/portfolio-templates/summaries/*.html")
	checkCriticalErr(err)

	checkCriticalErr(tmpl.Execute(w, projectTodos))
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/contact.html")
	checkCriticalErr(err)

	formSubmitted := false
	if r.Method == "POST" {
		handleSend(r)
		formSubmitted = true
	}
	checkCriticalErr(tmpl.Execute(w, formSubmitted))
}

func handleAbout(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/about.html")
	checkCriticalErr(err)
	checkCriticalErr(tmpl.Execute(w, nil))
}

func handleSend(response *http.Request) {
	checkCriticalErr(response.ParseForm())
	formType := response.PostFormValue("form-type")

	if formType == "hire" {
		name := response.PostFormValue("first-name") + " " + response.PostFormValue("last-name") +
			" contact: " + response.PostFormValue("contact-num")
		SendMail("Hire", name, response.PostFormValue("email"), response.PostFormValue("message"), smtpCred.email, smtpCred.pass, response.PostFormValue("contact-num"), "")
	} else if formType == "contribute" {
		name := response.PostFormValue("first-name") + " " + response.PostFormValue("last-name")
		SendMail("Contribute to "+response.PostFormValue("project-name"), name,
			response.PostFormValue("email"), response.PostFormValue("message"), smtpCred.email, smtpCred.pass, "", response.PostFormValue("project-name"))
	} else if formType == "other" {
		name := response.PostFormValue("first-name") + " " + response.PostFormValue("last-name")
		SendMail("General", name, response.PostFormValue("email"), response.PostFormValue("message"), smtpCred.email, smtpCred.pass, "", "")
	}
}

func main() {
	Init()

	smtpCred.email = viper.GetString("email")
	smtpCred.pass = viper.GetString("pass")

	fmt.Printf("Initializing...\n")
	checkCriticalErr(getTodoList("oak", "https://github.com/ryanmccauley211/Oak/blob/master/README.md"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/explore", handleExplore)
	http.HandleFunc("/portfolio", handlePortfolio)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/about", handleAbout)

	host := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	fmt.Printf("Serving at port %s\n", host)
	log.Fatal(http.ListenAndServe(host, nil))
}

func getExploreTemplateAliases(templates *template.Template) map[string]string {
	tmplAliases := map[string]string{}
	htmlExtensionLen := 5
	for _, tmpl := range templates.Templates()[1:] {
		nameKey := tmpl.Name()[:len(tmpl.Name())-htmlExtensionLen]
		tmplAliases[nameKey] = tmpl.Name()
	}
	return tmplAliases
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
		var checkedTasks []string
		var uncheckedTasks []string
		var delTasks []string
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

type SmtpCred struct {
	email string
	pass  string
}

func checkCriticalErr(err error) {
	if err != nil {
		panic(err)
	}
}
