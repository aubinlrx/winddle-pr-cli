package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PullRequest struct {
	Id      int64  `json:"id"`
	Url     string `json:"url"`
	Title   string `json:"title"`
	HtmlUrl string `json:"html_url"`
	Number  int32  `json:"number"`
	User    struct {
		Id    int64  `json:"id"`
		Login string `json:"login"`
	}
}

type PullRequests []PullRequest

var ApiToken = "65259272b07e6a452e31b096d95aadf432d59585"

func GetAll() (*PullRequests, error) {
	req, err := NewRequest("GET", "repos/winddle/winddle/pulls")
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	pullRequests := PullRequests{}

	err = json.Unmarshal([]byte(bodyBytes), &pullRequests)
	if err != nil {
		return nil, err
	}

	return &pullRequests, nil
}

func NewRequest(method string, url string) (*http.Request, error) {
	url = "https://api.github.com/" + url
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+ApiToken)

	return req, nil
}
