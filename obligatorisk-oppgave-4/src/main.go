package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"html/template"
	"strconv"
	"math/rand"
	"time"
	"strings"
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
	PokemonSum	int
	RandomPokemon	int
	RandomPokemonID	string
	RandomPokemonName	string

}

func pokemon(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/pokemon.html"))

	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
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

	var pokemonList = make([]string, 1)
	for i := 0; i < len(responseObject.Pokemon); i++ {
		//pokemonList = append(pokemonList, strconv.Itoa(responseObject.Pokemon[i].EntryNo), "-", strings.Title(responseObject.Pokemon[i].Species.Name+", "))
		pokemonList = append(pokemonList, strconv.Itoa(responseObject.Pokemon[i].EntryNo))
		pokemonList = append(pokemonList, strings.Title(responseObject.Pokemon[i].Species.Name))

	}

	for counter, name := range pokemonList {
		fmt.Println(counter, name)
		counter++
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(responseObject.Pokemon))
	randomString := strconv.Itoa(random)
	randomName := strings.Title(responseObject.Pokemon[random-1].Species.Name)

	data := PageData {
		PageTitle: "Pokemon Kanto Region",
		PokemonList: pokemonList,
		PokemonSum: len(responseObject.Pokemon),
		RandomPokemon: random,
		RandomPokemonID: randomString,
		RandomPokemonName: randomName,
	}

	tmpl.Execute(w, data)
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/index.html"))

	data := PageData {
		PageTitle: "Velkommen!",
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", welcomePage)
	http.HandleFunc("/1", pokemon)

	err := http.ListenAndServe(":80", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}