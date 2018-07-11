package config

import (
	"reflect"
	"errors"
	"runtime"
	"path"
	"path/filepath"
	"os"
	"encoding/json"
	"strings"
)

const jsonFileExtension = ".json"

// UnmarshalJsonFile unmarshals the content of a json file into a given target.
func UnmarshalJsonFile(fileName string, target interface{}) error {

	tOf := reflect.ValueOf(target).Type()
	if tOf.Kind() != reflect.Ptr || tOf.Elem().Kind() != reflect.Struct {
		return errors.New(ErrNoPointerNorStruct)
	}

	if len(fileName) == 0 {
		return errors.New(ErrEmptyParameter)
	}

	if !strings.HasSuffix(fileName, jsonFileExtension) {
		return errors.New(ErrNoJsonExtensionFile)
	}

	_, dirName, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirName), fileName)

	file, fileErr := os.Open(filePath); if fileErr != nil {
		return fileErr
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err := decoder.Decode(target); if err != nil {
		return err
	}

	return nil

}
