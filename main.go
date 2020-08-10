package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type body struct {
	UserName string `json:user_name`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Top Page.")
	})

	http.HandleFunc("/post_id", contributions)

	http.ListenAndServe(":8080", nil)
}

func contributions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Post only."))

		return
	}

	doc, err := goquery.NewDocument("https://github.com/users/t-kusakabe/contributions")
	if err != nil {
		fmt.Println("failed url")
	}

	var allFill uint64

	doc.Find("svg g g rect").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("fill")
		v, _ := strconv.ParseUint(url[1:6], 16, 0)
		allFill += v
	})

	json, _ := json.Marshal(allFill)

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(json)
}
