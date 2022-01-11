package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	Author  = "Metin Emre Koral"
	Github  = "https://github.com/eiektro"
	Version = "0.0.1"
	BaseUrl = "https://hacker-news.firebaseio.com/v0/"
)

type New struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

func main() {
	fmt.Println("Hacker News API")
	res, err := getItem("29894300")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

}

//get top stories from hacker news
func getTopStories() ([]int, error) {
	res, err := http.Get(BaseUrl + "topstories.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var stories []int
	err = json.Unmarshal(body, &stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// takes string and return http response
func getItem(itemId string) ([]byte, error) {
	url := BaseUrl + "item/" + itemId + ".json"
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
