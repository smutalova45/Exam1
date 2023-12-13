package store

import (
	"encoding/json"
	"fmt"
	"os"
)



func ReadJSon(fileName string) (Store, error) {
	var s Store
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error is while reading a json file", err.Error())
	}
	err = json.Unmarshal(file, &s.File)
	if err != nil {
		fmt.Println("error is while unmarshalling json file", err.Error())
	}
	return s, nil
}
