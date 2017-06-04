package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"github.com/russross/blackfriday"
)

type Page struct {
Title string
Body  template.HTML
}

func loadPage(title string) (*Page, error) {
	filename := "./index.html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: template.HTML(string(body))}, nil
}

func loadPosts(post string) (*Page, error) {
	postData, _ := getPost(post)
	body, _ := readContents(postData.Content)
	return &Page{Title: postData.Title, Body: template.HTML(string(body))}, nil
}

func readContents(path string) ([]byte, error) {
	if(path == "" ){
		return []byte("Oops 404"), nil
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return []byte("Oops 404"), nil
	}
	output := blackfriday.MarkdownCommon(file)
	return output, nil
}

func loadPostsTitles() (*Page, error){
	allposts, _ := getPosts()
	var buffer bytes.Buffer
	buffer.WriteString("<ul>")
	for i:= len(allposts)-1; i>=0; i-- {
		buffer.WriteString("<a href=\"" + allposts[i].Post + "\"><li>" + allposts[i].Title + ", " + allposts[i].Publishdate + "</li></a>" )
	}
	buffer.WriteString("</ul>")

	return &Page{Title:"Available posts", Body: template.HTML(buffer.String())}, nil
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.SplitAfter(r.URL.Path, "/blog/")[1]
	if(title == "" ){
		p, _ := loadPostsTitles()
		t, _ := template.ParseFiles("blog.html")
		t.Execute(w, p)
	}else{
		p, _ := loadPosts(title)
		t, _ := template.ParseFiles("blog.html")
		t.Execute(w, p)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("Home")
	fmt.Fprintf(w, "%s", template.HTML(string(p.Body)))
}
