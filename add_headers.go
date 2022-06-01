package github

import "net/http"

func addHeaders(r *http.Request, userID, token string) {
	r.Header.Set("Authorization", "token "+token)
	r.Header.Set("Accept", "application/vnd.github.v3+json")
}
