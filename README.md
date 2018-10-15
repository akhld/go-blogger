## GoBlogger
A tiny micro blogging platform written in Go!

### Setting up
```
$ git clone https://github.com/akhld/go-blogger
$ cd go-blogger
$ go build
$ ./go-blogger
```
Visit http://localhost:8080 after running.

### Adding / Editing blog posts
- Write the post in a markdown file
- Edit the config.json file
Example config for a post
```
{
    "title": "Hello from GoBlog!",
    "content" :"posts/hello-goblog.md",
    "tags":["go-blog", "go-blogging"],
    "post":"hello-goblog",
    "author":"AkhlD",
    "publishdate":"June 04, 2017"
  }
```

### Summary
[GoBlogger - A micro blogging platform wirrten in Go](https://hacked.work/blog/goblogger-a-micro-blogging-platform-written-in-go/)

### Contributing
Contributions / Suggestions are welcome - Please open an issue / PR in the repo.
