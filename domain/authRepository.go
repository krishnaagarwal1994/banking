package domain

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type DefaultAuthRepository struct {
}

func (repository DefaultAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := repository.buildVerifyAPIUrl(token, routeName, vars)
	var response *http.Response
	var apiFailureError error
	if response, apiFailureError = http.Get(u); apiFailureError != nil {
		log.Print("Failed to verify the token")
		return false
	}
	m := map[string]bool{}
	if err := json.NewDecoder(response.Body).Decode(&m); err != nil {
		log.Print("Verify api response decoding failed")
		return false
	}
	isAuthorized := m["isAuthorized"]
	return isAuthorized
}

func (repository DefaultAuthRepository) buildVerifyAPIUrl(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8081", Path: "auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() AuthRepository {
	return DefaultAuthRepository{}
}
