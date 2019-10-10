package Infrastructures

type IJsonReader interface {
	ReadDataFromJsonFile(path string, entity interface{}) error
}
