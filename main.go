package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

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

	userName := r.FormValue("username")
	doc, err := goquery.NewDocument("https://github.com/users/" + userName + "/contributions")
	if err != nil {
		fmt.Println("failed url")
	}

	var allFill uint64
	doc.Find("rect").Each(func(_ int, s *goquery.Selection) {
		hex, _ := s.Attr("fill")
		if hex[1:7] != "ebedf0" {
			v, _ := strconv.ParseUint(hex[1:7], 16, 0)
			allFill += 16777215 - v
		}
	})

	json, _ := json.Marshal(allFill)

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(json)
}
