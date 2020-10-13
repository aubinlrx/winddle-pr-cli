package api

import (
	"io/ioutil"
	"net/http"
)

type PullRequest struct{}

var ApiToken = ""

func GetAll() (string, error) {
	req, err := NewRequest("GET", "repos/winddle/winddle/pulls")
	if err != nil {
		return "", err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func NewRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+ApiToken)

	return req, nil
}
