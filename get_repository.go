package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func GetRepository(userID, token string, organization string, repository string) (*Repository, bool, error) {
	b, err := json.Marshal(repository)
	if err != nil {
		panic(err)
	}
	endpoint := url.URL{
		Scheme: "https",
		Host:   "api.github.com",
		Path:   "/repos/" + organization + "/" + repository,
	}
	r, err := http.NewRequest("GET", endpoint.String(), bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	addHeaders(r, userID, token)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, false, fmt.Errorf("failed to connect to github: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, false, nil
	}
	if resp.StatusCode != 200 {
		return nil, false, fmt.Errorf("github error: %s", strconv.Itoa(resp.StatusCode))
	}

	var repo Repository
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		return nil, false, fmt.Errorf("failed to decode github response: %s", err)
	}

	return &repo, true, nil
}
