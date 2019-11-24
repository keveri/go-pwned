package api

import (
	"encoding/json"
	. "github.com/keveri/go-pwned/internal/types"
	"io/ioutil"
	"net/http"
	"os"
)

const API = "https://haveibeenpwned.com/api/v3/"

func addHeaders(request *http.Request) {
	apiKey := os.Getenv("GO_PWNED_API_KEY")
	request.Header.Add("hibp-api-key", apiKey)
	request.Header.Add("user-agent", "go-pwned")
}

func callAPI(path string) ([]byte, error) {
	client := &http.Client{}
	endpoint := API + path

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	addHeaders(request)

	resp, err := client.Do(request)
	// TODO: status code checking
	if err != nil {
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		return data, nil
	}
}

func GetAllBreachesForEmail(email string) ([]Breach, error) {
	data, err := callAPI("breachedaccount/" + email + "?truncateResponse=false")
	if err != nil {
		return nil, err
	} else {
		var breaches []Breach
		json.Unmarshal(data, &breaches)
		return breaches, nil
	}
}

func GetAllPastesForEmail(email string) ([]Paste, error) {
	data, err := callAPI("pasteaccount/" + email)
	if err != nil {
		return nil, err
	} else {
		var pastes []Paste
		json.Unmarshal(data, &pastes)
		return pastes, nil
	}
}
