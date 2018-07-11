package config

const (

	// ErrNoPointerNorStruct is used when a parameter is not a pointer or not a struct
	ErrNoPointerNorStruct = "NoPointerNorStruct"

	// ErrEmptyParameter is used when a required parameter is empty
	ErrEmptyParameter = "ErrEmptyParameter"

	// ErrNoJsonExtensionFile is used when a string containing a file path
	// does not contain `.json` extension
	ErrNoJsonExtensionFile = "NoJsonExtensionFile"

)
