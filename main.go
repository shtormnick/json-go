package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ContentMarketing struct {
	Type              string  `json:"type"`
	ID                string  `json:"harvesterId"`
	CommercialPartner string  `json:"commercialPartner"`
	LogoUrl           string  `string:"logoURL"`
	CerebroScore      float64 `jason:"cerebro-score"`
	URL               string  `json:"url"`
	Title             string  `json:"title"`
	CleanImage        string  `json:"cleanImage"`
}

type Items struct {
	Items []ContentMarketing `json:"items"`
}

type Response struct {
	Status   int   `json:"httpStatus"`
	Response Items `json:"response"`
}

// urls := []string {
// 	"https://storage.googleapis.com/aller-structure-task/articles.json",
// 	"https://storage.googleapis.com/aller-structure-task/contentmarketing.json"
// }

func GetArticles(url string) []ContentMarketing {

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	Data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var jsonData Response

	err = json.Unmarshal([]byte(Data), &jsonData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", jsonData.Response.Items)

	return jsonData.Response.Items
}

func GetMarketing(url string) []ContentMarketing {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	Data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var jsonData Response

	err = json.Unmarshal([]byte(Data), &jsonData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", jsonData.Response.Items)

	return jsonData.Response.Items
}

func ToSlice(article_result, content_result []ContentMarketing) []ContentMarketing {
	var j, am = 0, 0
	p := 5
	var f []ContentMarketing

	for i := 0; i < len(article_result); i++ {
		if p <= len(article_result) {
			f = append(f, article_result[j:p]...)
			p += 5
			j += 5
			if i == am && am < len(content_result) {
				f = append(f, content_result[am])
				am++
			} else {
				f = append(f, ContentMarketing{Type: "AD"})
			}
		} else {
			break
		}
	}

	return f
}

func main() {
	a := GetArticles("https://storage.googleapis.com/aller-structure-task/articles.json")
	b := GetMarketing("https://storage.googleapis.com/aller-structure-task/contentmarketing.json")
	twoJson := ToSlice(a, b)
	fmt.Println(twoJson)
}
