package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Jeffail/gabs"
)

type ImageResult struct {
	Target string
	Class  string
	Type   string
}

func parseJSON(name string) {

	path := "result/" + name + ".json"
	jsonFile, _ := gabs.ParseJSONFile(path)
	// value = jsonFile.Path("Results.Target")

	result := &ImageResult{
		Target: jsonFile.Path("Results.Target").String(),
		Class:  jsonFile.Path("Results.Class").String(),
		Type:   jsonFile.Path("Results.Type").String(),
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
