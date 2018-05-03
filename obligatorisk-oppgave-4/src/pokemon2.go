package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Abilities []struct {
		Ability  struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height                 int           `json:"height"`
	ID          int `json:"id"`
	Types          []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/1/")
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

	fmt.Println(responseObject.Abilities)
	fmt.Println(responseObject.Name)
	fmt.Println(responseObject.Weight)
	fmt.Println(responseObject.Height)
	fmt.Println(responseObject.ID)
	fmt.Println(responseObject.Types)
}