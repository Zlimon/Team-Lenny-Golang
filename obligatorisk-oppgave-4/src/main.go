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

type Specie struct {
	Color struct {
		Name string `json:"name"`
	} `json:"color"`
	ID        int `json:"id"`
	FlavorTextEntries	[]Text	`json:"flavor_text_entries"`
	Genera       []struct {
		Genus    string `json:"genus"`
	} `json:"genera"`
	EvolvesFromSpecies struct {
		Name string `json:"name"`
	} `json:"evolves_from_species"`
	Name           string `json:"name"`
	EvolutionChain struct {
	} `json:"evolution_chain"`
}

type Text struct {
	Text	string	`json:"flavor_text"`
}

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
	Sprite	Sprite `json:"sprites"`
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

type Sprite struct {
	FrontDefault	string      `json:"front_default"`
	FrontShiny		string      `json:"front_shiny"`
}

type PageData struct {
	PageTitle			string
	PageBackground		string

	Search 				[]SearchBox
	SearchError			string
	PokemonAmount		int
	PokemonGen			string
	PokemonList			[]string

	PreviousPokemonName	string
	PreviousPokemon		string
	Image				string
	NextPokemonName		string
	NextPokemon			string
	StaticSprite		string
	Sprite				string
	ShinySprite			string

	ID					int
	Name				string
	Height				float64
	Weight				float64
	PrimaryType			string
	SecondaryType		string
	PrimaryTypeColor	string
	SecondaryTypeColor	string
	Abilities			[]string

	Text				string
}

type SearchBox struct {
	PokemonSearch	string
}

var pokemonList []string

func main() {
	pokemonAPI, err := http.Get("https://pokeapi.co/api/v2/pokemon/?limit=802")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	pokemonData, err := ioutil.ReadAll(pokemonAPI.Body)
	if err != nil {
		log.Fatal(err)
	}

	var pokemonObject Pokemon
	json.Unmarshal(pokemonData, &pokemonObject)

	pokemonList = make([]string, 0)
	for i := 0; i < len(pokemonObject.Results); i++ {
		pokemonList = append(pokemonList, pokemonObject.Results[i].Name)
	}

	fmt.Println(pokemonList)
	fmt.Println("Pokemon liste fullført!", len(pokemonObject.Results), "initialisert!")

	http.HandleFunc("/", pokemon)
	http.HandleFunc("/gen1", generation)
	http.HandleFunc("/gen2", generation)
	http.HandleFunc("/gen3", generation)
	http.HandleFunc("/gen4", generation)
	http.HandleFunc("/gen5", generation)
	http.HandleFunc("/gen6", generation)
	http.HandleFunc("/gen7", generation)
	//http.HandleFunc("/selected", UserSelected)
	http.HandleFunc("/search", search)
	//http.HandleFunc("/hoenn", hoenn)
	http.HandleFunc("/test", test)

	errr := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(errr)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	/*
	fmt.Println("let go")

	specie, err := http.Get("https://pokeapi.co/api/v2/pokemon-species/54/") // Henter informasjon om en tilfeldig Pokemon
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	specieData, err := ioutil.ReadAll(specie.Body)
	if err != nil {
		log.Fatal(err)
	}

	var specieObject Specie
	json.Unmarshal(specieData, &specieObject)

	testColor := specieObject.Color.Name
	testID := specieObject.ID
	testText := specieObject.FlavorTextEntries[2]
	testGenera := specieObject.Genera[2]
	testEvolveFrom := specieObject.EvolvesFromSpecies.Name
	testName := specieObject.Name


	fmt.Println(testColor)
	fmt.Println(testID)
	fmt.Println(testText)
	fmt.Println(testGenera)
	fmt.Println(testEvolveFrom)
	fmt.Println(testName)

	fmt.Println("ok")
	*/
}

func pokemon(w http.ResponseWriter, r *http.Request) {
	searchBox := []SearchBox {
		SearchBox{"pokemonSearch"},
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(pokemonList))
	randomPokemon := strconv.Itoa(random)

	previousImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(random-1)+".png"
	previousPokemon := strconv.Itoa(random-1)
	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+randomPokemon+".png"
	nextImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(random+1)+".png"
	if random <= 9 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+strconv.Itoa(random-1)+".png"
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+randomPokemon+".png"
		nextImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+strconv.Itoa(random+1)+".png"
	} else if random <= 99 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+strconv.Itoa(random-1)+".png"
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+randomPokemon+".png"
		nextImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+strconv.Itoa(random+1)+".png"
	}
	if random == 1 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(len(pokemonList))+".png"
		previousPokemon = strconv.Itoa(len(pokemonList))
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

	staticSprite := informationObject.Sprite.FrontDefault
	sprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	if random <= 721 {
		sprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	} else {
		sprite = informationObject.Sprite.FrontDefault
	}
	shinySprite := informationObject.Sprite.FrontShiny

	specie, err := http.Get("https://pokeapi.co/api/v2/pokemon-species/"+randomPokemon+"/") // Henter informasjon om en tilfeldig Pokemon
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	specieData, err := ioutil.ReadAll(specie.Body)
	if err != nil {
		log.Fatal(err)
	}

	var specieObject Specie
	json.Unmarshal(specieData, &specieObject)

	//testColor := specieObject.Color.Name
	//testID := specieObject.ID
	pokemonTextList := make([]string, 0)
	for i := 0; i < len(specieObject.FlavorTextEntries); i++ {
		pokemonTextList = append(pokemonTextList, strings.Title(specieObject.FlavorTextEntries[i].Text))
	}
	//testGenera := specieObject.Genera[2]
	//testEvolveFrom := specieObject.EvolvesFromSpecies.Name
	//testName := specieObject.Name

	calcHeight := float64(informationObject.Height) / 10
	calcWeight := float64(informationObject.Weight) / 10
	titleName := strings.Title(informationObject.Name)

	pokemonTypeList := make([]string, 0)
	for i := 0; i < len(informationObject.Type); i++ {
		pokemonTypeList = append(pokemonTypeList, strings.Title(informationObject.Type[i].Type.TypeName))
	}

	var pokemonPrimaryType string
	var pokemonSecondaryType string
	var pokemonPrimaryTypeColor string
	var pokemonSecondaryTypeColor string

	pokemonPrimaryType = pokemonTypeList[0]
	if pokemonTypeList[0] == "Normal" { //|| pokemonTypeList[1] == "[Normal]"
		pokemonPrimaryTypeColor = "A8A878"
	} else if pokemonTypeList[0] == "Fighting" { //|| pokemonTypeList[1] == "[Fighting]"
		pokemonPrimaryTypeColor = "C03028"
	} else if pokemonTypeList[0] == "Flying" { //|| pokemonTypeList[1] == "[Flying]"
		pokemonPrimaryTypeColor = "A890F0"
	} else if pokemonTypeList[0] == "Poison" { //|| pokemonTypeList[1] == "[Poison]"
		pokemonPrimaryTypeColor = "A040A0"
	} else if pokemonTypeList[0] == "Ground" { //|| pokemonTypeList[1] == "[Ground]"
		pokemonPrimaryTypeColor = "E0C068"
	} else if pokemonTypeList[0] == "Rock" { //|| pokemonTypeList[1] == "[Rock]"
		pokemonPrimaryTypeColor = "B8A038"
	} else if pokemonTypeList[0] == "Bug" { //|| pokemonTypeList[1] == "[Bug]"
		pokemonPrimaryTypeColor = "A8B820"
	} else if pokemonTypeList[0] == "Ghost" { //|| pokemonTypeList[1] == "[Ghost]"
		pokemonPrimaryTypeColor = "705898"
	} else if pokemonTypeList[0] == "Steel" { //|| pokemonTypeList[1] == "[Steel]"
		pokemonPrimaryTypeColor = "B8B8D0"
	} else if pokemonTypeList[0] == "Fire" { //|| pokemonTypeList[1] == "[Fire]"
		pokemonPrimaryTypeColor = "F08030"
	} else if pokemonTypeList[0] == "Water" { //|| pokemonTypeList[1] == "[Water]"
		pokemonPrimaryTypeColor = "6890F0"
	} else if pokemonTypeList[0] == "Grass" { //|| pokemonTypeList[1] == "[Grass]"
		pokemonPrimaryTypeColor = "78C850"
	} else if pokemonTypeList[0] == "Electric" { //|| pokemonTypeList[1] == "[Electric]"
		pokemonPrimaryTypeColor = "F8D030"
	} else if pokemonTypeList[0] == "Psychic" { //|| pokemonTypeList[1] == "[Psychic]"
		pokemonPrimaryTypeColor = "F85888"
	} else if pokemonTypeList[0] == "Ice" { //|| pokemonTypeList[1] == "[Ice]"
		pokemonPrimaryTypeColor = "98D8D8"
	} else if pokemonTypeList[0] == "Dragon" { //|| pokemonTypeList[1] == "[Dragon]"
		pokemonPrimaryTypeColor = "7038F8"
	} else if pokemonTypeList[0] == "Dark" { //|| pokemonTypeList[1] == "[Dark]"
		pokemonPrimaryTypeColor = "705848"
	} else if pokemonTypeList[0] == "Fairy" { //|| pokemonTypeList[1] == "[Fairy]"
		pokemonPrimaryTypeColor = "EE99AC"
	} else if pokemonTypeList[0] == "Unknown" { //|| pokemonTypeList[1] == "[Unknown]"
		pokemonPrimaryTypeColor = "68A090"
	} else if pokemonTypeList[0] == "Shadow" { //|| pokemonTypeList[1] == "[Shadow]"
		pokemonPrimaryTypeColor = "000000"
	} else {
		pokemonPrimaryTypeColor = "3564AE"
	}

	if len(pokemonTypeList) > 1 {
		pokemonSecondaryType = pokemonTypeList[1]
		if pokemonTypeList[1] == "Normal" { //|| pokemonTypeList[1] == "[Normal]"
			pokemonSecondaryTypeColor = "A8A878"
		} else if pokemonTypeList[1] == "Fighting" { //|| pokemonTypeList[1] == "[Fighting]"
			pokemonSecondaryTypeColor = "C03028"
		} else if pokemonTypeList[1] == "Flying" { //|| pokemonTypeList[1] == "[Flying]"
			pokemonSecondaryTypeColor = "A890F0"
		} else if pokemonTypeList[1] == "Poison" { //|| pokemonTypeList[1] == "[Poison]"
			pokemonSecondaryTypeColor = "A040A0"
		} else if pokemonTypeList[1] == "Ground" { //|| pokemonTypeList[1] == "[Ground]"
			pokemonSecondaryTypeColor = "E0C068"
		} else if pokemonTypeList[1] == "Rock" { //|| pokemonTypeList[1] == "[Rock]"
			pokemonSecondaryTypeColor = "B8A038"
		} else if pokemonTypeList[1] == "Bug" { //|| pokemonTypeList[1] == "[Bug]"
			pokemonSecondaryTypeColor = "A8B820"
		} else if pokemonTypeList[1] == "Ghost" { //|| pokemonTypeList[1] == "[Ghost]"
			pokemonSecondaryTypeColor = "705898"
		} else if pokemonTypeList[1] == "Steel" { //|| pokemonTypeList[1] == "[Steel]"
			pokemonSecondaryTypeColor = "B8B8D0"
		} else if pokemonTypeList[1] == "Fire" { //|| pokemonTypeList[1] == "[Fire]"
			pokemonSecondaryTypeColor = "F08030"
		} else if pokemonTypeList[1] == "Water" { //|| pokemonTypeList[1] == "[Water]"
			pokemonSecondaryTypeColor = "6890F0"
		} else if pokemonTypeList[1] == "Grass" { //|| pokemonTypeList[1] == "[Grass]"
			pokemonSecondaryTypeColor = "78C850"
		} else if pokemonTypeList[1] == "Electric" { //|| pokemonTypeList[1] == "[Electric]"
			pokemonSecondaryTypeColor = "F8D030"
		} else if pokemonTypeList[1] == "Psychic" { //|| pokemonTypeList[1] == "[Psychic]"
			pokemonSecondaryTypeColor = "F85888"
		} else if pokemonTypeList[1] == "Ice" { //|| pokemonTypeList[1] == "[Ice]"
			pokemonSecondaryTypeColor = "98D8D8"
		} else if pokemonTypeList[1] == "Dragon" { //|| pokemonTypeList[1] == "[Dragon]"
			pokemonSecondaryTypeColor = "7038F8"
		} else if pokemonTypeList[1] == "Dark" { //|| pokemonTypeList[1] == "[Dark]"
			pokemonSecondaryTypeColor = "705848"
		} else if pokemonTypeList[1] == "Fairy" { //|| pokemonTypeList[1] == "[Fairy]"
			pokemonSecondaryTypeColor = "EE99AC"
		} else if pokemonTypeList[1] == "Unknown" { //|| pokemonTypeList[1] == "[Unknown]"
			pokemonSecondaryTypeColor = "68A090"
		} else if pokemonTypeList[1] == "Shadow" { //|| pokemonTypeList[1] == "[Shadow]"
			pokemonSecondaryTypeColor = "000000"
		} else {
			pokemonSecondaryTypeColor = "3564AE"
		}
	}

	pokemonAbilityList := make([]string, 0)
	for i := 0; i < len(informationObject.Ability); i++ {
		pokemonAbilityList = append(pokemonAbilityList, "[" +strings.Title(informationObject.Ability[i].Ability.AbilityName)+ "]")
	}

	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/index.html"))

	pokemonTemplateData := PageData {
		PageTitle: "Pokemon GO(lang)",
		//PageBackground: DittoBackground,
		Search: searchBox,
		PreviousPokemonName: previousPokemon,
		PreviousPokemon: previousImage,
		Image: image,
		NextPokemonName: strconv.Itoa(random+1),
		NextPokemon: nextImage,
		StaticSprite: staticSprite,
		Sprite: sprite,
		ShinySprite: shinySprite,

		ID: informationObject.ID,
		Name: titleName,
		Height: calcHeight,
		Weight: calcWeight,
		PrimaryType: pokemonPrimaryType,
		SecondaryType: pokemonSecondaryType,
		PrimaryTypeColor: pokemonPrimaryTypeColor,
		SecondaryTypeColor: pokemonSecondaryTypeColor,
		Abilities: pokemonAbilityList,

		//Text: pokemonTextList[1],
	}

	tmpl.Execute(w, pokemonTemplateData)
}

func generation(w http.ResponseWriter, r *http.Request) {
	var chosenGeneration string

	limit := len(pokemonList)
	offset := 0

	if r.URL.Path == "/gen1" {
		chosenGeneration = "Generasjon 1"
		limit = 151
	} else if r.URL.Path == "/gen2" {
		chosenGeneration = "Generasjon 2"
		offset = 151
		limit = 251
	} else if r.URL.Path == "/gen3" {
		chosenGeneration = "Generasjon 3"
		offset = 251
		limit = 386
	} else if r.URL.Path == "/gen4" {
		chosenGeneration = "Generasjon 4"
		offset = 386
		limit = 493
	} else if r.URL.Path == "/gen5" {
		chosenGeneration = "Generasjon 5"
		offset = 493
		limit = 649
	} else if r.URL.Path == "/gen6" {
		chosenGeneration = "Generasjon 6"
		offset = 649
		limit = 721
	} else if r.URL.Path == "/gen7" {
		chosenGeneration = "Generasjon 7"
		offset = 721
		limit = 802
	}

	pokemonList := pokemonList[offset:limit]

	rand.Seed(time.Now().UnixNano())
	//random := rand.Intn(((offset)+limit)-offset+1)+offset
	random := rand.Intn(limit - offset) + offset
	randomPokemon := strconv.Itoa(random)

	previousImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(random-1)+".png"
	previousPokemon := strconv.Itoa(random-1)
	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+randomPokemon+".png"
	nextImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(random+1)+".png"
	if random <= 9 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+strconv.Itoa(random-1)+".png"
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+randomPokemon+".png"
		nextImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+strconv.Itoa(random+1)+".png"
	} else if random <= 99 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+strconv.Itoa(random-1)+".png"
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+randomPokemon+".png"
		nextImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+strconv.Itoa(random+1)+".png"
	}
	if random == 1 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(len(pokemonList))+".png"
		previousPokemon = strconv.Itoa(len(pokemonList))
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

	staticSprite := informationObject.Sprite.FrontDefault
	sprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	if r.URL.Path != "/gen7" {
		sprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	} else {
		sprite = informationObject.Sprite.FrontDefault
	}
	shinySprite := informationObject.Sprite.FrontShiny

	calcHeight := float64(informationObject.Height) / 10
	calcWeight := float64(informationObject.Weight) / 10
	titleName := strings.Title(informationObject.Name)

	pokemonTypeList := make([]string, 0)
	for i := 0; i < len(informationObject.Type); i++ {
		pokemonTypeList = append(pokemonTypeList, strings.Title(informationObject.Type[i].Type.TypeName))
	}

	var pokemonPrimaryType string
	var pokemonSecondaryType string
	var pokemonPrimaryTypeColor string
	var pokemonSecondaryTypeColor string

	pokemonPrimaryType = pokemonTypeList[0]
	if pokemonTypeList[0] == "Normal" { //|| pokemonTypeList[1] == "[Normal]"
		pokemonPrimaryTypeColor = "A8A878"
	} else if pokemonTypeList[0] == "Fighting" { //|| pokemonTypeList[1] == "[Fighting]"
		pokemonPrimaryTypeColor = "C03028"
	} else if pokemonTypeList[0] == "Flying" { //|| pokemonTypeList[1] == "[Flying]"
		pokemonPrimaryTypeColor = "A890F0"
	} else if pokemonTypeList[0] == "Poison" { //|| pokemonTypeList[1] == "[Poison]"
		pokemonPrimaryTypeColor = "A040A0"
	} else if pokemonTypeList[0] == "Ground" { //|| pokemonTypeList[1] == "[Ground]"
		pokemonPrimaryTypeColor = "E0C068"
	} else if pokemonTypeList[0] == "Rock" { //|| pokemonTypeList[1] == "[Rock]"
		pokemonPrimaryTypeColor = "B8A038"
	} else if pokemonTypeList[0] == "Bug" { //|| pokemonTypeList[1] == "[Bug]"
		pokemonPrimaryTypeColor = "A8B820"
	} else if pokemonTypeList[0] == "Ghost" { //|| pokemonTypeList[1] == "[Ghost]"
		pokemonPrimaryTypeColor = "705898"
	} else if pokemonTypeList[0] == "Steel" { //|| pokemonTypeList[1] == "[Steel]"
		pokemonPrimaryTypeColor = "B8B8D0"
	} else if pokemonTypeList[0] == "Fire" { //|| pokemonTypeList[1] == "[Fire]"
		pokemonPrimaryTypeColor = "F08030"
	} else if pokemonTypeList[0] == "Water" { //|| pokemonTypeList[1] == "[Water]"
		pokemonPrimaryTypeColor = "6890F0"
	} else if pokemonTypeList[0] == "Grass" { //|| pokemonTypeList[1] == "[Grass]"
		pokemonPrimaryTypeColor = "78C850"
	} else if pokemonTypeList[0] == "Electric" { //|| pokemonTypeList[1] == "[Electric]"
		pokemonPrimaryTypeColor = "F8D030"
	} else if pokemonTypeList[0] == "Psychic" { //|| pokemonTypeList[1] == "[Psychic]"
		pokemonPrimaryTypeColor = "F85888"
	} else if pokemonTypeList[0] == "Ice" { //|| pokemonTypeList[1] == "[Ice]"
		pokemonPrimaryTypeColor = "98D8D8"
	} else if pokemonTypeList[0] == "Dragon" { //|| pokemonTypeList[1] == "[Dragon]"
		pokemonPrimaryTypeColor = "7038F8"
	} else if pokemonTypeList[0] == "Dark" { //|| pokemonTypeList[1] == "[Dark]"
		pokemonPrimaryTypeColor = "705848"
	} else if pokemonTypeList[0] == "Fairy" { //|| pokemonTypeList[1] == "[Fairy]"
		pokemonPrimaryTypeColor = "EE99AC"
	} else if pokemonTypeList[0] == "Unknown" { //|| pokemonTypeList[1] == "[Unknown]"
		pokemonPrimaryTypeColor = "68A090"
	} else if pokemonTypeList[0] == "Shadow" { //|| pokemonTypeList[1] == "[Shadow]"
		pokemonPrimaryTypeColor = "000000"
	} else {
		pokemonPrimaryTypeColor = "3564AE"
	}

	if len(pokemonTypeList) > 1 {
		pokemonSecondaryType = pokemonTypeList[1]
		if pokemonTypeList[1] == "Normal" { //|| pokemonTypeList[1] == "[Normal]"
			pokemonSecondaryTypeColor = "A8A878"
		} else if pokemonTypeList[1] == "Fighting" { //|| pokemonTypeList[1] == "[Fighting]"
			pokemonSecondaryTypeColor = "C03028"
		} else if pokemonTypeList[1] == "Flying" { //|| pokemonTypeList[1] == "[Flying]"
			pokemonSecondaryTypeColor = "A890F0"
		} else if pokemonTypeList[1] == "Poison" { //|| pokemonTypeList[1] == "[Poison]"
			pokemonSecondaryTypeColor = "A040A0"
		} else if pokemonTypeList[1] == "Ground" { //|| pokemonTypeList[1] == "[Ground]"
			pokemonSecondaryTypeColor = "E0C068"
		} else if pokemonTypeList[1] == "Rock" { //|| pokemonTypeList[1] == "[Rock]"
			pokemonSecondaryTypeColor = "B8A038"
		} else if pokemonTypeList[1] == "Bug" { //|| pokemonTypeList[1] == "[Bug]"
			pokemonSecondaryTypeColor = "A8B820"
		} else if pokemonTypeList[1] == "Ghost" { //|| pokemonTypeList[1] == "[Ghost]"
			pokemonSecondaryTypeColor = "705898"
		} else if pokemonTypeList[1] == "Steel" { //|| pokemonTypeList[1] == "[Steel]"
			pokemonSecondaryTypeColor = "B8B8D0"
		} else if pokemonTypeList[1] == "Fire" { //|| pokemonTypeList[1] == "[Fire]"
			pokemonSecondaryTypeColor = "F08030"
		} else if pokemonTypeList[1] == "Water" { //|| pokemonTypeList[1] == "[Water]"
			pokemonSecondaryTypeColor = "6890F0"
		} else if pokemonTypeList[1] == "Grass" { //|| pokemonTypeList[1] == "[Grass]"
			pokemonSecondaryTypeColor = "78C850"
		} else if pokemonTypeList[1] == "Electric" { //|| pokemonTypeList[1] == "[Electric]"
			pokemonSecondaryTypeColor = "F8D030"
		} else if pokemonTypeList[1] == "Psychic" { //|| pokemonTypeList[1] == "[Psychic]"
			pokemonSecondaryTypeColor = "F85888"
		} else if pokemonTypeList[1] == "Ice" { //|| pokemonTypeList[1] == "[Ice]"
			pokemonSecondaryTypeColor = "98D8D8"
		} else if pokemonTypeList[1] == "Dragon" { //|| pokemonTypeList[1] == "[Dragon]"
			pokemonSecondaryTypeColor = "7038F8"
		} else if pokemonTypeList[1] == "Dark" { //|| pokemonTypeList[1] == "[Dark]"
			pokemonSecondaryTypeColor = "705848"
		} else if pokemonTypeList[1] == "Fairy" { //|| pokemonTypeList[1] == "[Fairy]"
			pokemonSecondaryTypeColor = "EE99AC"
		} else if pokemonTypeList[1] == "Unknown" { //|| pokemonTypeList[1] == "[Unknown]"
			pokemonSecondaryTypeColor = "68A090"
		} else if pokemonTypeList[1] == "Shadow" { //|| pokemonTypeList[1] == "[Shadow]"
			pokemonSecondaryTypeColor = "000000"
		} else {
			pokemonSecondaryTypeColor = "3564AE"
		}
	}

	pokemonAbilityList := make([]string, 0)
	for i := 0; i < len(informationObject.Ability); i++ {
		pokemonAbilityList = append(pokemonAbilityList, "[" +strings.Title(informationObject.Ability[i].Ability.AbilityName)+ "]")
	}

	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/generation.html"))

	pokemonTemplateData := PageData {
		PageTitle: chosenGeneration,
		//PageBackground: DittoBackground,
		PokemonAmount: limit,
		PokemonGen: chosenGeneration,
		PokemonList: pokemonList,
		PreviousPokemonName: previousPokemon,
		PreviousPokemon: previousImage,
		Image: image,
		NextPokemonName: strconv.Itoa(random+1),
		NextPokemon: nextImage,
		StaticSprite: staticSprite,
		Sprite: sprite,
		ShinySprite: shinySprite,

		ID: informationObject.ID,
		Name: titleName,
		Height: calcHeight,
		Weight: calcWeight,
		PrimaryType: pokemonPrimaryType,
		SecondaryType: pokemonSecondaryType,
		PrimaryTypeColor: pokemonPrimaryTypeColor,
		SecondaryTypeColor: pokemonSecondaryTypeColor,
		Abilities: pokemonAbilityList,
	}

	tmpl.Execute(w, pokemonTemplateData)
}

func search(w http.ResponseWriter, r *http.Request) {
	searchBox := []SearchBox {
		SearchBox{"pokemonSearch"},
	}

	r.ParseForm()
	searchResult := r.Form.Get("pokemonSearch")
	lowerSearchResult := strings.ToLower(searchResult)

	errorMessage := "En ukjent feil har oppstått!"
	searchID := "0"
	maxSearch, _ := strconv.Atoi(lowerSearchResult)
	if maxSearch > len(pokemonList) {
		errorMessage = "Du kan ikke skrive inn mer enn " + strconv.Itoa(len(pokemonList)) + "!"
		searchID = "1"
	} else if searchResult == "" {
		errorMessage = "Du må skrive inn noe!"
		searchID = "1"
	} else {
		for i := range pokemonList {
			if lowerSearchResult == pokemonList[i] {
				errorMessage = ""
				searchID = strconv.Itoa(i + 1)
				break
			} else if _, err := strconv.Atoi(lowerSearchResult); err == nil {
				errorMessage = ""
				searchID = lowerSearchResult
			}
		}
		if searchID == "0" {
			errorMessage = "Fant ingen med navn '" + searchResult + "'!"
			searchID = "1"
		}
	}

	searchINT, _ := strconv.Atoi(searchID)
	previousImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(searchINT-1)+".png"
	previousPokemon := strconv.Itoa(searchINT-1)
	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+searchID+".png"
	nextImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(searchINT+1)+".png"
	if searchINT <= 9 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+strconv.Itoa(searchINT-1)+".png"
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+searchID+".png"
		nextImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+strconv.Itoa(searchINT+1)+".png"
	} else if searchINT <= 99 {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+strconv.Itoa(searchINT-1)+".png"
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+searchID+".png"
		nextImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+strconv.Itoa(searchINT+1)+".png"
	}
	if searchID == "1" {
		previousImage = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(len(pokemonList))+".png"
		previousPokemon = strconv.Itoa(len(pokemonList))
	}

	information, err := http.Get("https://pokeapi.co/api/v2/pokemon/"+searchID+"/") // Henter informasjon om Pokemon
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

	staticSprite := informationObject.Sprite.FrontDefault
	sprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	if searchINT <= 721 {
		sprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	} else {
		sprite = informationObject.Sprite.FrontDefault
	}
	shinySprite := informationObject.Sprite.FrontShiny

	calcHeight := float64(informationObject.Height) / 10
	calcWeight := float64(informationObject.Weight) / 10
	titleName := strings.Title(informationObject.Name)

	pokemonTypeList := make([]string, 0)
	for i := 0; i < len(informationObject.Type); i++ {
		pokemonTypeList = append(pokemonTypeList, strings.Title(informationObject.Type[i].Type.TypeName))
	}

	var pokemonPrimaryType string
	var pokemonSecondaryType string
	var pokemonPrimaryTypeColor string
	var pokemonSecondaryTypeColor string

	pokemonPrimaryType = pokemonTypeList[0]
	if pokemonTypeList[0] == "Normal" { //|| pokemonTypeList[1] == "[Normal]"
		pokemonPrimaryTypeColor = "A8A878"
	} else if pokemonTypeList[0] == "Fighting" { //|| pokemonTypeList[1] == "[Fighting]"
		pokemonPrimaryTypeColor = "C03028"
	} else if pokemonTypeList[0] == "Flying" { //|| pokemonTypeList[1] == "[Flying]"
		pokemonPrimaryTypeColor = "A890F0"
	} else if pokemonTypeList[0] == "Poison" { //|| pokemonTypeList[1] == "[Poison]"
		pokemonPrimaryTypeColor = "A040A0"
	} else if pokemonTypeList[0] == "Ground" { //|| pokemonTypeList[1] == "[Ground]"
		pokemonPrimaryTypeColor = "E0C068"
	} else if pokemonTypeList[0] == "Rock" { //|| pokemonTypeList[1] == "[Rock]"
		pokemonPrimaryTypeColor = "B8A038"
	} else if pokemonTypeList[0] == "Bug" { //|| pokemonTypeList[1] == "[Bug]"
		pokemonPrimaryTypeColor = "A8B820"
	} else if pokemonTypeList[0] == "Ghost" { //|| pokemonTypeList[1] == "[Ghost]"
		pokemonPrimaryTypeColor = "705898"
	} else if pokemonTypeList[0] == "Steel" { //|| pokemonTypeList[1] == "[Steel]"
		pokemonPrimaryTypeColor = "B8B8D0"
	} else if pokemonTypeList[0] == "Fire" { //|| pokemonTypeList[1] == "[Fire]"
		pokemonPrimaryTypeColor = "F08030"
	} else if pokemonTypeList[0] == "Water" { //|| pokemonTypeList[1] == "[Water]"
		pokemonPrimaryTypeColor = "6890F0"
	} else if pokemonTypeList[0] == "Grass" { //|| pokemonTypeList[1] == "[Grass]"
		pokemonPrimaryTypeColor = "78C850"
	} else if pokemonTypeList[0] == "Electric" { //|| pokemonTypeList[1] == "[Electric]"
		pokemonPrimaryTypeColor = "F8D030"
	} else if pokemonTypeList[0] == "Psychic" { //|| pokemonTypeList[1] == "[Psychic]"
		pokemonPrimaryTypeColor = "F85888"
	} else if pokemonTypeList[0] == "Ice" { //|| pokemonTypeList[1] == "[Ice]"
		pokemonPrimaryTypeColor = "98D8D8"
	} else if pokemonTypeList[0] == "Dragon" { //|| pokemonTypeList[1] == "[Dragon]"
		pokemonPrimaryTypeColor = "7038F8"
	} else if pokemonTypeList[0] == "Dark" { //|| pokemonTypeList[1] == "[Dark]"
		pokemonPrimaryTypeColor = "705848"
	} else if pokemonTypeList[0] == "Fairy" { //|| pokemonTypeList[1] == "[Fairy]"
		pokemonPrimaryTypeColor = "EE99AC"
	} else if pokemonTypeList[0] == "Unknown" { //|| pokemonTypeList[1] == "[Unknown]"
		pokemonPrimaryTypeColor = "68A090"
	} else if pokemonTypeList[0] == "Shadow" { //|| pokemonTypeList[1] == "[Shadow]"
		pokemonPrimaryTypeColor = "000000"
	} else {
		pokemonPrimaryTypeColor = "3564AE"
	}

	if len(pokemonTypeList) > 1 {
		pokemonSecondaryType = pokemonTypeList[1]
		if pokemonTypeList[1] == "Normal" { //|| pokemonTypeList[1] == "[Normal]"
			pokemonSecondaryTypeColor = "A8A878"
		} else if pokemonTypeList[1] == "Fighting" { //|| pokemonTypeList[1] == "[Fighting]"
			pokemonSecondaryTypeColor = "C03028"
		} else if pokemonTypeList[1] == "Flying" { //|| pokemonTypeList[1] == "[Flying]"
			pokemonSecondaryTypeColor = "A890F0"
		} else if pokemonTypeList[1] == "Poison" { //|| pokemonTypeList[1] == "[Poison]"
			pokemonSecondaryTypeColor = "A040A0"
		} else if pokemonTypeList[1] == "Ground" { //|| pokemonTypeList[1] == "[Ground]"
			pokemonSecondaryTypeColor = "E0C068"
		} else if pokemonTypeList[1] == "Rock" { //|| pokemonTypeList[1] == "[Rock]"
			pokemonSecondaryTypeColor = "B8A038"
		} else if pokemonTypeList[1] == "Bug" { //|| pokemonTypeList[1] == "[Bug]"
			pokemonSecondaryTypeColor = "A8B820"
		} else if pokemonTypeList[1] == "Ghost" { //|| pokemonTypeList[1] == "[Ghost]"
			pokemonSecondaryTypeColor = "705898"
		} else if pokemonTypeList[1] == "Steel" { //|| pokemonTypeList[1] == "[Steel]"
			pokemonSecondaryTypeColor = "B8B8D0"
		} else if pokemonTypeList[1] == "Fire" { //|| pokemonTypeList[1] == "[Fire]"
			pokemonSecondaryTypeColor = "F08030"
		} else if pokemonTypeList[1] == "Water" { //|| pokemonTypeList[1] == "[Water]"
			pokemonSecondaryTypeColor = "6890F0"
		} else if pokemonTypeList[1] == "Grass" { //|| pokemonTypeList[1] == "[Grass]"
			pokemonSecondaryTypeColor = "78C850"
		} else if pokemonTypeList[1] == "Electric" { //|| pokemonTypeList[1] == "[Electric]"
			pokemonSecondaryTypeColor = "F8D030"
		} else if pokemonTypeList[1] == "Psychic" { //|| pokemonTypeList[1] == "[Psychic]"
			pokemonSecondaryTypeColor = "F85888"
		} else if pokemonTypeList[1] == "Ice" { //|| pokemonTypeList[1] == "[Ice]"
			pokemonSecondaryTypeColor = "98D8D8"
		} else if pokemonTypeList[1] == "Dragon" { //|| pokemonTypeList[1] == "[Dragon]"
			pokemonSecondaryTypeColor = "7038F8"
		} else if pokemonTypeList[1] == "Dark" { //|| pokemonTypeList[1] == "[Dark]"
			pokemonSecondaryTypeColor = "705848"
		} else if pokemonTypeList[1] == "Fairy" { //|| pokemonTypeList[1] == "[Fairy]"
			pokemonSecondaryTypeColor = "EE99AC"
		} else if pokemonTypeList[1] == "Unknown" { //|| pokemonTypeList[1] == "[Unknown]"
			pokemonSecondaryTypeColor = "68A090"
		} else if pokemonTypeList[1] == "Shadow" { //|| pokemonTypeList[1] == "[Shadow]"
			pokemonSecondaryTypeColor = "000000"
		} else {
			pokemonSecondaryTypeColor = "3564AE"
		}
	}

	pokemonAbilityList := make([]string, 0)
	for i := 0; i < len(informationObject.Ability); i++ {
		pokemonAbilityList = append(pokemonAbilityList, "[" +strings.Title(informationObject.Ability[i].Ability.AbilityName)+ "]")
	}

	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/index.html"))

	pokemonTemplateData := PageData {
		PageTitle: "Søk",
		//PageBackground: DittoBackground,
		Search: searchBox,
		SearchError: errorMessage,
		PreviousPokemonName: previousPokemon,
		PreviousPokemon: previousImage,
		Image: image,
		NextPokemonName: strconv.Itoa(searchINT+1),
		NextPokemon: nextImage,
		StaticSprite: staticSprite,
		Sprite: sprite,
		ShinySprite: shinySprite,

		ID: informationObject.ID,
		Name: titleName,
		Height: calcHeight,
		Weight: calcWeight,
		PrimaryType: pokemonPrimaryType,
		SecondaryType: pokemonSecondaryType,
		PrimaryTypeColor: pokemonPrimaryTypeColor,
		SecondaryTypeColor: pokemonSecondaryTypeColor,
		Abilities: pokemonAbilityList,
	}

	tmpl.Execute(w, pokemonTemplateData)
}