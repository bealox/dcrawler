package matcher

import (
	"encoding/json"
	"os"
)

const dataFile = "data/council.json"

type Feed struct {
	State string `json:state`
	Link  string `json:link`
	Site  string `json:site`
}

func RetrieveFeed() ([]*Feed, error) {
	//Open the file

	file, err := os.Open(dataFile)

	//Error then return back
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed

	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
