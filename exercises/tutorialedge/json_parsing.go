package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type users struct {
	Users []user `json:"users"`
}

type user struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age int `json:"age"`
	Social social `json:"social"`
}

// Social struct which contains a
// list of links
type social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func runJsonParsing() {
	/*
	Reading and Parsing a JSON File
	 */
	// Open our jsonFile
	jsonFile, err := os.Open("exercises/tutorialedge/users.json")
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully opened users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// initialize users array
	var users users
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		log.Println("User type: " + users.Users[i].Type)
		log.Println("User age: " + strconv.Itoa(users.Users[i].Age))
		log.Println("User name: " + users.Users[i].Name)
		log.Println("Facebook url: " + users.Users[i].Social.Facebook)
	}

	/*
	Working with Unstructured Data
	Sometimes, going through the process of creating structs for everything can be somewhat time consuming and overly
	verbose for the problems you are trying to solve. In this instance, we can use standard interfaces{} in order to
	read in any JSON data.
	*/
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	fmt.Println(result["users"])
	defer jsonFile.Close()
}