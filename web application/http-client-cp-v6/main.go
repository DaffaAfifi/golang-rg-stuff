package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := http.Client{}

	// Hit API https://animechan.xyz/api/quotes/anime?title=naruto with method GET:
	var animechan []Animechan

	req, err := http.NewRequest("GET", "https://animechan.xyz/api/quotes/anime?title=naruto", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &animechan)
	if err != nil {
		return nil, err
	}

	return animechan, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	// Hit API https://postman-echo.com/post with method POST:
	var postman Postman

	resp, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	if err != nil {
		return postman, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return postman, err
	}

	err = json.Unmarshal(body, &postman)
	if err != nil {
		return postman, err
	}

	return postman, nil
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
