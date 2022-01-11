package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/feeds"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	Author             = "Metin Emre Koral"
	Github             = "https://github.com/eiektro"
	Version            = "0.0.1"
	NumberOfTopArticle = 10
	BaseUrl            = "https://hacker-news.firebaseio.com/v0/"
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

	var new New

	var news []*feeds.Item

	for _, v := range getTopStories() {
		item, err := getItem(strconv.Itoa(v))
		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(item, &new)

		if new.Type == "story" {
			news = append(news, &feeds.Item{

				Title:       new.Title,
				Link:        &feeds.Link{Href: new.Url},
				Description: "No Description Defined",
				Created:     timeConverter(new.Time),
				Id:          uuid.New().String(),
				Author:      &feeds.Author{Name: new.By},
			})
		}
	}

	fmt.Println(news[9])

	/*
		feed := &feeds.Feed{}
		feed.Items = news

		rss, err := feed.ToRss()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(rss)
	*/
}

//function that converts unix time to time.Time
func timeConverter(unixTime int) time.Time {
	return time.Unix(int64(unixTime), 0)
}

//gets the top stories from the api
func getTopStories() []int {
	resp, err := http.Get(BaseUrl + "topstories.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var topStories []int
	json.Unmarshal(body, &topStories)
	return topStories[:NumberOfTopArticle]
}

// gets story item from hacker news by id
func getItem(itemId string) ([]byte, error) {
	url := BaseUrl + "item/" + itemId + ".json"
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
