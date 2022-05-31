package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CreateRepository(token string, organization string, repository *Repository) error {
	b, err := json.Marshal(repository)
	if err != nil {
		panic(err)
	}
	r, err := http.NewRequest("POST", "https://api.github.com/orgs/"+organization+"/repos", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	addHeaders(r, token)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return fmt.Errorf("failed to connect to github: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return fmt.Errorf("github error: %s", strconv.Itoa(resp.StatusCode))
	}

	return nil
}
