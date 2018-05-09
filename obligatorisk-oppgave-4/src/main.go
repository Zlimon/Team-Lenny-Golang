package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
	"math/rand"
	"time"
	"strconv"
	"html/template"
)

type Pokemon struct {
	Results	[]struct {
		Name string `json:"name"`
	} `json:"results"`
}

type Information struct {
	ID		int	`json:"id"`
	Name	string `json:"name"`
	Height	int	`json:"height"`
	Type 	[]Type `json:"types"`
	Weight	int	`json:"weight"`
	Ability	[]Ability `json:"abilities"`
}

type Type struct {
	Type PokemonType `json:"type"`
}

type PokemonType struct {
	TypeName string `json:"name"`
}

type Ability struct {
	Ability PokemonAbility `json:"ability"`
}

type PokemonAbility struct {
	AbilityName string `json:"name"`
}

type PageData struct {
	PageTitle				string
	PageBackground			string
	PokemonList				[]string
	PokemonID				int
	PokemonName				string
	PokemonHeight			float64
	PokemonType1			string
	PokemonType2			string
	PokemonTypesColor		string
	PokemonSecondTypesColor	string
	PokemonWeight			float64
	PokemonAbilities		[]string
	PokemonSearch 			[]SearchBox
	PokemonResult			string
	SearchError				string
	PokemonImage			string
}

type SearchBox struct {
	PokemonNameSearch	string
}

var region = "kanto"
var firstTypeColor = "3564AE" // Standard farge
var secondTypeColor = "3564AE"
var DittoBackground = "" // Standard bakgrunn
var pokemonList = make([]string, 0)
//var pokemonListKanto = make([]string, 0)
//var pokemonListHoenn = make([]string, 0)
var pokemonTypeList = make([]string, 0)
var pokemonTypeList1 string
var pokemonTypeList2 string
var pokemonAbilityList = make([]string, 0)
var randomPokemon string
var errorMessage string

func pokemon(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/pokemon.html"))

	pokemon, err := http.Get("https://pokeapi.co/api/v2/pokemon/?limit=949") // Henter informasjon om Pokemon
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	pokemonData, err := ioutil.ReadAll(pokemon.Body)
	if err != nil {
		log.Fatal(err)
	}

	var pokemonObject Pokemon
	json.Unmarshal(pokemonData, &pokemonObject)

	pokemonList = make([]string, 0)
	for i := 0; i < len(pokemonObject.Results); i++ {
		pokemonList = append(pokemonList, pokemonObject.Results[i].Name)
	}

	searchBox := []SearchBox {
		SearchBox{"pokemonSearch"},
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(pokemonList))
	//random := rand.Intn(len(regionObject.Pokemon)-151)+151 // If Hoenn region
	randomPokemon = strconv.Itoa(random)

	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+randomPokemon+".png"
	searchResultINT, _ := strconv.Atoi(randomPokemon)
	if searchResultINT <= 9 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+randomPokemon+".png"
	} else if searchResultINT <= 99 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+randomPokemon+".png"
	}

	information, err := http.Get("https://pokeapi.co/api/v2/pokemon/"+randomPokemon+"/") // Henter informasjon om Pokemon en tilfeldig Pokemon
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	informationData, err := ioutil.ReadAll(information.Body)
	if err != nil {
		log.Fatal(err)
	}

	var informationObject Information
	json.Unmarshal(informationData, &informationObject)

	correctHeight := float64(informationObject.Height) / 10
	correctWeight := float64(informationObject.Weight) / 10
	uppercaseName := strings.Title(informationObject.Name)

	pokemonTypeList := make([]string, 0)
	for i := 0; i < len(informationObject.Type); i++ {
		pokemonTypeList = append(pokemonTypeList, "[" +strings.Title(informationObject.Type[i].Type.TypeName)+ "]")
	}

	pokemonAbilityList := make([]string, 0)
	for i := 0; i < len(informationObject.Ability); i++ {
		pokemonAbilityList = append(pokemonAbilityList, "[" +strings.Title(informationObject.Ability[i].Ability.AbilityName)+ "]")
	}

	pokemonTemplateData := PageData {
		// Globale variabler
		PageTitle: "Pokemon " + strings.Title(region) + " Region",
		PageBackground: DittoBackground,
		PokemonType1: pokemonTypeList1,
		PokemonType2: pokemonTypeList2,
		PokemonTypesColor: firstTypeColor,
		PokemonSecondTypesColor: secondTypeColor,

		PokemonID: informationObject.ID,
		PokemonName: uppercaseName,
		PokemonHeight: correctHeight,
		PokemonWeight: correctWeight,
		PokemonAbilities: pokemonAbilityList,
		PokemonSearch: searchBox,
		PokemonImage: image,
	}

	tmpl.Execute(w, pokemonTemplateData)
}

func kanto(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/pokemon.html"))

	searchBox := []SearchBox {
		SearchBox{"pokemonSearch"},
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(151) // If Kanto region
	randomPokemon = strconv.Itoa(random)

	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+randomPokemon+".png"
	searchResultINT, _ := strconv.Atoi(randomPokemon)
	if searchResultINT <= 9 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+randomPokemon+".png"
	} else if searchResultINT <= 99 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+randomPokemon+".png"
	}

	information, err := http.Get("https://pokeapi.co/api/v2/pokemon/"+randomPokemon+"/") // Henter informasjon om en tilfeldig Pokemon
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	informationData, err := ioutil.ReadAll(information.Body)
	if err != nil {
		log.Fatal(err)
	}

	var informationObject Information
	json.Unmarshal(informationData, &informationObject)

	correctHeight := float64(informationObject.Height) / 10
	correctWeight := float64(informationObject.Weight) / 10
	uppercaseName := strings.Title(informationObject.Name)

	pokemonTypeList := make([]string, 0)
	for i := 0; i < len(informationObject.Type); i++ {
		pokemonTypeList = append(pokemonTypeList, "[" +strings.Title(informationObject.Type[i].Type.TypeName)+ "]")
	}

	pokemonAbilityList := make([]string, 0)
	for i := 0; i < len(informationObject.Ability); i++ {
		pokemonAbilityList = append(pokemonAbilityList, "[" +strings.Title(informationObject.Ability[i].Ability.AbilityName)+ "]")
	}

	var pokemonTemplateData = PageData {
		// Globale variabler
		PageTitle: "Pokemon " + strings.Title(region) + " Region",
		PageBackground: DittoBackground,
		PokemonType1: pokemonTypeList1,
		PokemonType2: pokemonTypeList2,
		PokemonTypesColor: firstTypeColor,
		PokemonSecondTypesColor: secondTypeColor,

		PokemonID: informationObject.ID,
		PokemonName: uppercaseName,
		PokemonHeight: correctHeight,
		PokemonWeight: correctWeight,
		PokemonAbilities: pokemonAbilityList,
		PokemonSearch: searchBox,
		PokemonImage: image,
	}

	tmpl.Execute(w, pokemonTemplateData)
}

func search(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/pokemon.html"))

	pokemon, err := http.Get("https://pokeapi.co/api/v2/pokemon/?limit=949") // Henter informasjon om Pokemon en tilfeldig Pokemon
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	pokemonData, err := ioutil.ReadAll(pokemon.Body)
	if err != nil {
		log.Fatal(err)
	}

	var pokemonObject Pokemon
	json.Unmarshal(pokemonData, &pokemonObject)

	pokemonList := make([]string, 0)
	for i := 0; i < len(pokemonObject.Results); i++ {
		pokemonList = append(pokemonList, pokemonObject.Results[i].Name)
	}

	searchBox := []SearchBox {
		SearchBox{"pokemonSearch"},
	}

	r.ParseForm()
	searchResult := r.Form.Get("pokemonSearch")
	lowerSearchResult := strings.ToLower(searchResult)

	searchResultID := "0"
	searchMaxAmount, _ := strconv.Atoi(lowerSearchResult)
	if searchMaxAmount >= 803 {
		errorMessage = "Du kan ikke skrive inn mer enn 802!" //  + strconv.Itoa(len(regionObject.Pokemon)) +
		searchResultID = "1"
	} else {
		for i := range pokemonList {
			if lowerSearchResult == pokemonList[i] {
				errorMessage = ""
				searchResultID = strconv.Itoa(i + 1)
				break
			} else if _, err := strconv.Atoi(lowerSearchResult); err == nil {
				errorMessage = ""
				searchResultID = lowerSearchResult
			}
		}
		if searchResultID == "0" { //  || pokemonAmount > len(regionObject.Pokemon)
			errorMessage = "Fant ingen med navn '" + searchResult + "'!"
			searchResultID = "1"
		}
	}

	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+searchResultID+".png"
	searchResultINT, _ := strconv.Atoi(searchResultID)
	if searchResultINT <= 9 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+searchResultID+".png"
	} else if searchResultINT <= 99 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+searchResultID+".png"
	}

	information, err := http.Get("https://pokeapi.co/api/v2/pokemon/"+searchResultID+"/") // Henter informasjon om Pokemon
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	informationData, err := ioutil.ReadAll(information.Body)
	if err != nil {
		log.Fatal(err)
	}

	var informationObject Information
	json.Unmarshal(informationData, &informationObject)

	correctHeight := float64(informationObject.Height) / 10
	correctWeight := float64(informationObject.Weight) / 10
	uppercaseName := strings.Title(informationObject.Name)

	pokemonTypeList := make([]string, 0)
	for i := 0; i < len(informationObject.Type); i++ {
		pokemonTypeList = append(pokemonTypeList, "[" +strings.Title(informationObject.Type[i].Type.TypeName)+ "]")
	}

	pokemonAbilityList := make([]string, 0)
	for i := 0; i < len(informationObject.Ability); i++ {
		pokemonAbilityList = append(pokemonAbilityList, "[" +strings.Title(informationObject.Ability[i].Ability.AbilityName)+ "]")
	}

	pokemonTemplateData := PageData {
		// Globale variabler
		PageTitle: "Pokemon " + strings.Title(region) + " Region",
		PageBackground: DittoBackground,
		PokemonType1: pokemonTypeList1,
		PokemonType2: pokemonTypeList2,
		PokemonTypesColor: firstTypeColor,
		PokemonSecondTypesColor: secondTypeColor,

		PokemonID: informationObject.ID,
		PokemonName: uppercaseName,
		PokemonHeight: correctHeight,
		PokemonWeight: correctWeight,
		PokemonAbilities: pokemonAbilityList,
		PokemonSearch: searchBox,
		PokemonImage: image,

		SearchError: errorMessage,
	}

	tmpl.Execute(w, pokemonTemplateData)
}

func main() {
	http.HandleFunc("/", pokemon)
	http.HandleFunc("/kanto", kanto)
	//http.HandleFunc("/selected", UserSelected)
	http.HandleFunc("/search", search)
	//http.HandleFunc("/hoenn", hoenn)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}