package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type PostsData struct {
	Title string
	Content string
	Tags []string
	Post string
	Author string
	Publishdate string
	Seometa string
}

func readConfig() ([]byte, error){
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	return file, nil
}

func getPosts() ([]PostsData, error){
	var data []PostsData
	var file, err = readConfig()
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}

func getPost(post string) (PostsData, error){
	var data []PostsData
	var postData PostsData
	var file, err = readConfig()
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(data); i++ {
		if(strings.Contains(post, data[i].Post)){
			postData = data[i]
			break
		}else{
			for j := 0; j < len(data[i].Tags); j++ {
				if(strings.Contains(post, data[i].Tags[j])){
					postData = data[i]
					break
				}
			}
		}
	}

	return postData, nil
}