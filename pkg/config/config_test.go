package config

import (
	"github.com/stretchr/testify/assert"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

type TestUnmarshalJsonFileType struct {
	SomeParam1 string `json:"some_par_1"`
	SomeParam2 struct {
		SomeNestedParam string `json:"some_nested_par"`
	} `json:"some_par_2"`
}

func TestUnmarshalJsonFile(t *testing.T) {

	var out TestUnmarshalJsonFileType

	fileName := "../../configs/test-config.json"

	_, dirName, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirName), fileName)

	someVal := "some_val"
	someNestedVal := "some_nested_val"

	err := UnmarshalJsonFile(filePath, &out)

	assert.NoError(t, err)
	assert.Equal(t, someVal, out.SomeParam1)
	assert.Equal(t, someNestedVal, out.SomeParam2.SomeNestedParam)

	shouldBeEmptyParamErr := UnmarshalJsonFile("", &out)

	assert.Error(t, shouldBeEmptyParamErr)
	assert.Equal(t, shouldBeEmptyParamErr.Error(), ErrEmptyParameter)

	shouldBeNoJsonFileErr := UnmarshalJsonFile("somefilename", &out)

	assert.Error(t, shouldBeNoJsonFileErr)
	assert.Equal(t, shouldBeNoJsonFileErr.Error(), ErrNoJsonExtensionFile)

	shouldBeNoPtrOrStructErr1 := UnmarshalJsonFile(filePath, out)

	assert.Error(t, shouldBeNoPtrOrStructErr1)
	assert.Equal(t, shouldBeNoPtrOrStructErr1.Error(), ErrNoPointerNorStruct)

	shouldBeNoPtrOrStructErr2 := UnmarshalJsonFile(filePath, map[string]interface{}{})

	assert.Error(t, shouldBeNoPtrOrStructErr2)
	assert.Equal(t, shouldBeNoPtrOrStructErr2.Error(), ErrNoPointerNorStruct)

}
