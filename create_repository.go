package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func CreateRepository(userID, token string, organization string, repository *Repository) error {
	b, err := json.Marshal(repository)
	if err != nil {
		panic(err)
	}
	url := "https://api.github.com/orgs/" + organization + "/repos"
	if userID == organization {
		url = "https://api.github.com/user/repos"
	}
	r, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	addHeaders(r, userID, token)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return fmt.Errorf("failed to connect to github: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 401 {
		return fmt.Errorf("unauthorized")
	}
	if resp.StatusCode == 404 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%s", string(body))
	}
	if resp.StatusCode != 201 {
		return fmt.Errorf("github error: %s", strconv.Itoa(resp.StatusCode))
	}

	return nil
}
