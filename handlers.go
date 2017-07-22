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

type Post struct {
	Title string
	Body template.HTML
	Author string
	PublishDate string
	Tags template.HTML
}

func loadPage(title string) (*Page, error) {
	filename := "./index.html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: template.HTML(string(body))}, nil
}

func loadPosts(post string) (*Post, error) {
	postData, _ := getPost(post)
	body, _ := readContents(postData.Content)
	var tags bytes.Buffer
	for i := 0; i < len(postData.Tags); i++ {
		tags.WriteString("<a href=\"" + postData.Tags[i] + "\">" + postData.Tags[i] +  "</a>")
	}
	return &Post{Title: postData.Title,
		Body: template.HTML(string(body)),
		Author:postData.Author,
		PublishDate: postData.Publishdate,
		Tags: template.HTML(tags.String()) }, nil
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
		t, _ := template.ParseFiles("blog-post.html")
		t.Execute(w, p)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("Home")
	fmt.Fprintf(w, "%s", template.HTML(string(p.Body)))
}
