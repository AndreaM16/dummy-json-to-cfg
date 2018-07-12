package config

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strings"
)

const jsonFileExtension = ".json"

// UnmarshalJsonFile unmarshals the content of a json file into a given target.
func UnmarshalJsonFile(filePath string, target interface{}) error {

	tOf := reflect.ValueOf(target).Type()
	if tOf.Kind() != reflect.Ptr || tOf.Elem().Kind() != reflect.Struct {
		return errors.New(ErrNoPointerNorStruct)
	}

	if len(filePath) == 0 {
		return errors.New(ErrEmptyParameter)
	}

	if !strings.HasSuffix(filePath, jsonFileExtension) {
		return errors.New(ErrNoJsonExtensionFile)
	}

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		return fileErr
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err := decoder.Decode(target)
	if err != nil {
		return err
	}

	return nil

}
