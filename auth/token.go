package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Key struct {
	Token string
}

func Token() string {

	filename, err := os.Open("helpers/token.json")
	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()

	data, err := ioutil.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	var result Key

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return result.Token

}
