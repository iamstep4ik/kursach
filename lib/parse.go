package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type RawData map[string][][]int
type Data map[string][]int

func Parse(file string) Data {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var rawData RawData
	err = json.Unmarshal(byteValue, &rawData)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	data := make(Data)
	for key, matrix := range rawData {
		flatArray := make([]int, 0, 25)
		for _, array := range matrix {
			flatArray = append(flatArray, array...)
		}
		data[key] = flatArray
	}

	return data
}
