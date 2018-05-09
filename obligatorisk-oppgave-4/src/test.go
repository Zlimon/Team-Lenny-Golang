package main

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
)

type RadioButton struct {
	Name       string
	Value      string
}

type PageVariables struct {
	PageTitle        string
	PageRadioButtons []RadioButton
	Answer           string
}

func main() {
	http.HandleFunc("/", DisplayRadioButtons)
	http.HandleFunc("/selected", UserSelected)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request){
	MyRadioButtons := []RadioButton {
		RadioButton{"animalselect", "cats"},
	}

	MyPageVariables := PageVariables {
		PageRadioButtons : MyRadioButtons,
	}

	t, err := template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/select.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func UserSelected(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	youranimal := r.Form.Get("animalselect")

	if youranimal == "wow" {
		fmt.Println("kult")
	}

	MyPageVariables := PageVariables{
		Answer : youranimal,
	}

	t, err := template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/select.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}