package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"html/template"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

type TodoPageData struct {
	PageTitle string
	PokemonListe     string
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	var pokeListe string

	for i := 0; i < len(responseObject.Pokemon); i++ {
		pokeListe = responseObject.Pokemon[i].Species.Name
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}

	fmt.Println(pokeListe)

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/1/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData {
			PageTitle: "Pokemon Kanto Region",
			PokemonListe: pokeListe,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}