package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
	"html/template"
)

type Page struct {
Title string
Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := "./index.html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func loadPosts(post string) (*Page, error) {
	var postt = "Awesome post#! " + post
	postData, _ := getPost(post)
	body := []byte(postData.Content)
	return &Page{Title: postt, Body: body}, nil
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.SplitAfter(r.URL.Path, "/blog/")[1]
	p, _ := loadPosts(title)
	t, _ := template.ParseFiles("blog.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("Home")
	fmt.Fprintf(w, "%s", p.Body)
}
