package Infrastructures

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JsonReader struct {
}

func (reader JsonReader) ReadDataFromJsonFile(path string, entity interface{}) error {
	pwd, _ := os.Getwd()
	file, _ := ioutil.ReadFile(pwd + path)
	err := json.Unmarshal([]byte(file), &entity)
	return err
}

func NewJsonReader() JsonReader {
	return JsonReader{}
}
