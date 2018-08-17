package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
	"math/rand"
	"time"
	"strconv"
	"html/template"
	"log"
)
type Pokemon struct {
	Results	[]struct {
		Name string `json:"name"`
	} `json:"results"`
}

type Specie struct {
	Color struct {
		Name string `json:"name"`
	} `json:"color"`
	Genera       []struct {
		Genus    string `json:"genus"`
	} `json:"genera"`
}

type Text struct {
	Text	string	`json:"flavor_text"`
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
	Error				string
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
	StaticShiny			string
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
	Color				string

	Quiz				[]QuizBox
	Response			string
	Points				int

	AnswerCheck			bool
}

type SearchBox struct {
	PokemonSearch	string
}

type QuizBox struct {
	PokemonQuiz	string
}

var pokemonList []string

func main() {
	pokemonAPI, err := http.Get("https://pokeapi.co/api/v2/pokemon/?limit=802")
	errorCheck(err)

	pokemonData, err := ioutil.ReadAll(pokemonAPI.Body)
	errorCheck(err)

	var pokemonObject Pokemon
	json.Unmarshal(pokemonData, &pokemonObject)

	pokemonList = make([]string, 0)
	for i := 0; i < len(pokemonObject.Results); i++ {
		pokemonList = append(pokemonList, pokemonObject.Results[i].Name)
	}

	fmt.Println(pokemonList)
	fmt.Println("Pokemon liste fullført!", len(pokemonObject.Results), "initialisert!")

	http.HandleFunc("/", pokemon)
	http.HandleFunc("/search", search)
	http.HandleFunc("/gen1", generation)
	http.HandleFunc("/gen2", generation)
	http.HandleFunc("/gen3", generation)
	http.HandleFunc("/gen4", generation)
	http.HandleFunc("/gen5", generation)
	http.HandleFunc("/gen6", generation)
	http.HandleFunc("/gen7", generation)
	http.HandleFunc("/quiz", quiz)
	http.HandleFunc("/response", response)

	http.ListenAndServe(":80", nil)
}

func pokemon(w http.ResponseWriter, r *http.Request) {
	searchBox := []SearchBox {
		SearchBox{"pokemonSearch"},
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(pokemonList))
	randomPokemon := strconv.Itoa(random)

	previousImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(random-1)+".png"
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
	}

	information, err := http.Get("https://pokeapi.co/api/v2/pokemon/"+randomPokemon+"/")
	errorCheck(err)

	informationData, err := ioutil.ReadAll(information.Body)
	errorCheck(err)

	var informationObject Information
	json.Unmarshal(informationData, &informationObject)

	staticSprite := informationObject.Sprite.FrontDefault
	sprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	staticShiny := informationObject.Sprite.FrontShiny
	shinySprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados-shiny/"+informationObject.Name+".gif"
	if random <= 721 {
		sprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
		staticShiny = informationObject.Sprite.FrontShiny
		shinySprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados-shiny/"+informationObject.Name+".gif"
	} else {
		sprite = informationObject.Sprite.FrontDefault
		shinySprite = informationObject.Sprite.FrontShiny
	}

	specie, err := http.Get("https://pokeapi.co/api/v2/pokemon-species/"+randomPokemon+"/")
	errorCheck(err)

	specieData, err := ioutil.ReadAll(specie.Body)
	errorCheck(err)

	var specieObject Specie
	json.Unmarshal(specieData, &specieObject)

	specieColor := specieObject.Color.Name

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
		Search: searchBox,
		PreviousPokemonName: strings.Title(pokemonList[random-2]),
		PreviousPokemon: previousImage,
		Image: image,
		NextPokemonName: strings.Title(pokemonList[random]),
		NextPokemon: nextImage,
		StaticSprite: staticSprite,
		Sprite: sprite,
		StaticShiny: staticShiny,
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

		Color: specieColor,
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
				break
			}
		}
		if searchID == "0" {
			errorMessage = "Fant ingen med navn '" + searchResult + "'!"
			searchID = "1"
		}
	}

	searchINT, _ := strconv.Atoi(searchID)
	previousImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(searchINT-1)+".png"
	previousName := strings.Title(pokemonList[searchINT-1])
	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+searchID+".png"
	nextImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(searchINT+1)+".png"
	nextName := strings.Title(pokemonList[searchINT])
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
		previousName = "Marshadow"
	}

	information, err := http.Get("https://pokeapi.co/api/v2/pokemon/"+searchID+"/")
	errorCheck(err)

	informationData, err := ioutil.ReadAll(information.Body)
	errorCheck(err)

	var informationObject Information
	json.Unmarshal(informationData, &informationObject)

	staticSprite := informationObject.Sprite.FrontDefault
	sprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	staticShiny := informationObject.Sprite.FrontShiny
	shinySprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados-shiny/"+informationObject.Name+".gif"
	if searchINT <= 721 {
		sprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
		staticShiny = informationObject.Sprite.FrontShiny
		shinySprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados-shiny/"+informationObject.Name+".gif"
	} else {
		sprite = informationObject.Sprite.FrontDefault
		shinySprite = informationObject.Sprite.FrontShiny
	}

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
		Search: searchBox,
		Error: errorMessage,
		PreviousPokemonName: previousName,
		PreviousPokemon: previousImage,
		Image: image,
		NextPokemonName: nextName,
		NextPokemon: nextImage,
		StaticSprite: staticSprite,
		Sprite: sprite,
		StaticShiny: staticShiny,
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

var answerCheck bool
var points int
var randomQuiz string

func quiz(w http.ResponseWriter, r *http.Request) {
	if answerCheck == false {
		points--
	}

	answerCheck = false

	QuizBox := []QuizBox {
		QuizBox{"pokemonQuiz"},
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(pokemonList))
	randomQuiz = strconv.Itoa(random)

	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+randomQuiz+".png"
	if random <= 9 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+randomQuiz+".png"
	} else if random <= 99 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+randomQuiz+".png"
	}

	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/quiz.html"))

	pokemonTemplateData := PageData {
		PageTitle: "Quiz who's that Pokemon!",
		Quiz: QuizBox,
		Image: image,
		Points: points,
	}

	tmpl.Execute(w, pokemonTemplateData)
}

func response(w http.ResponseWriter, r *http.Request) {
	QuizBox := []QuizBox {
		QuizBox{"pokemonQuiz"},
	}

	var answerResponse string

	r.ParseForm()
	QuizResult := r.Form.Get("pokemonQuiz")
	lowerQuizResult := strings.ToLower(QuizResult)
	randomQuizINT, _ := strconv.Atoi(randomQuiz)

	errorMessage := "En ukjent feil har oppstått!"
	if _, err := strconv.Atoi(QuizResult); err == nil {
		errorMessage = "Du kan ikke bruke sifre!"
	} else if QuizResult == "" {
		errorMessage = "Du må skrive inn noe!"
	} else {
		if lowerQuizResult == pokemonList[randomQuizINT-1] {
			errorMessage = ""
			answerResponse = "Korrekt!"
			points++
			answerCheck = true
		}
		if answerCheck == false {
			errorMessage = "Feil svar! Prøv igjen!"
			points--
		}
	}

	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+randomQuiz+".png"
	if randomQuizINT <= 9 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/00"+randomQuiz+".png"
	} else if randomQuizINT <= 99 {
		image = "https://assets.pokemon.com/assets/cms2/img/pokedex/full/0"+randomQuiz+".png"
	}

	tmpl := template.Must(template.ParseFiles("C:/GitHub/Team-Lenny/obligatorisk-oppgave-4/src/Quiz.html"))

	pokemonTemplateData := PageData {
		PageTitle: "Quiz who's that Pokemon!",
		Quiz: QuizBox,
		Image: image,
		Response: answerResponse,
		Points: points,
		Error: errorMessage,
		AnswerCheck: answerCheck,
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
	random := rand.Intn(limit - offset) + offset
	randomPokemon := strconv.Itoa(random)

	previousImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(random-1)+".png"
	//previousName := strings.Title(pokemonList[random])
	image := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+randomPokemon+".png"
	nextImage := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/"+strconv.Itoa(random+1)+".png"
	//nextName := strings.Title(pokemonList[random])
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
	}

	information, err := http.Get("https://pokeapi.co/api/v2/pokemon/"+randomPokemon+"/")
	errorCheck(err)

	informationData, err := ioutil.ReadAll(information.Body)
	errorCheck(err)

	var informationObject Information
	json.Unmarshal(informationData, &informationObject)

	staticSprite := informationObject.Sprite.FrontDefault
	sprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
	staticShiny := informationObject.Sprite.FrontShiny
	shinySprite := "http://www.pkparaiso.com/imagenes/xy/sprites/animados-shiny/"+informationObject.Name+".gif"
	if random <= 721 {
		sprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados/"+informationObject.Name+".gif"
		staticShiny = informationObject.Sprite.FrontShiny
		shinySprite = "http://www.pkparaiso.com/imagenes/xy/sprites/animados-shiny/"+informationObject.Name+".gif"
	} else {
		sprite = informationObject.Sprite.FrontDefault
		shinySprite = informationObject.Sprite.FrontShiny
	}

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
		PokemonAmount: limit,
		PokemonGen: chosenGeneration,
		PokemonList: pokemonList,
		//PreviousPokemonName: previousName,
		PreviousPokemon: previousImage,
		Image: image,
		//NextPokemonName: nextName,
		NextPokemon: nextImage,
		StaticSprite: staticSprite,
		Sprite: sprite,
		StaticShiny: staticShiny,
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

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}