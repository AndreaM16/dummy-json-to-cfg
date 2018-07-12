package main

import (
	"fmt"
	"github.com/andream16/dummy-json-to-cfg/pkg/config"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

const fileName = "../configs/example-config.json"

func main() {

	var person Person

	_, dirName, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirName), fileName)

	if err := config.UnmarshalJsonFile(filePath, &person); err != nil {
		log.Fatal(err)
	}

	// "Name = Mark, Surname = Lenders, Age = 18"
	fmt.Printf("Name = %s, Surname = %s, Age = %d", person.Name, person.Surname, person.Age)

}
