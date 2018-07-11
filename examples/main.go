package main

import (
	"github.com/andream16/dummy-json-to-cfg/pkg/config"
	"log"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age int `json:"age"`
}

const filePath = "../../configs/example-config.json"

func main() {

	var person Person

	if err := config.UnmarshalJsonFile(filePath, &person); err != nil {
		log.Fatal(err)
	}

	// "Name = Mark, Surname = Lenders, Age = 18"
	fmt.Printf("Name = %s, Surname = %s, Age = %d", person.Name, person.Surname, person.Age)

}
