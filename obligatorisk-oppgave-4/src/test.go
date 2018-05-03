package main

import (
	"net/http"
	"log"
	"html/template"
)

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
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
	// Display some radio buttons to the user
	Title := "Which do you prefer?"
	MyRadioButtons := []RadioButton {
		RadioButton{"animalselect", "cats", false, false, "Cats"},
		RadioButton{"animalselect", "dogs", false, false, "Dogs"},
	}

	MyPageVariables := PageVariables {
		PageTitle: Title,
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
	// r.Form is now either
	// map[animalselect:[cats]] OR
	// map[animalselect:[dogs]]
	// so get the animal which has been selected
	youranimal := r.Form.Get("animalselect")

	Title := "Your preferred animal"
	MyPageVariables := PageVariables{
		PageTitle: Title,
		Answer : youranimal,
	}

	// generate page by passing page variables into template
	t, err := template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/select.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}