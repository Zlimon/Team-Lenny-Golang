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

type PageData struct {
	PageTitle	string
	PokemonList	[]string
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

	var pokemonList = make([]string, 0)
	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
		pokemonList = append(pokemonList, responseObject.Pokemon[i].Species.Name)
	}

	fmt.Println("File List: ", pokemonList)

	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData {
			PageTitle: "Pokemon Kanto Region",
			PokemonList: pokemonList,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}