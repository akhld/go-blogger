HTTP/2 is designed to address many of the failings of HTTP/1.x. Modern web pages use many resources: HTML, stylesheets, scripts, images, and so on. In HTTP/1.x, each of these resources must be requested explicitly. This can be a slow process. The browser starts by fetching the HTML, then learns of more resources incrementally as it parses and evaluates the page. Since the server must wait for the browser to make each request, the network is often idle and underutilized.
### Setting up the HTTP/2 server
You need to download the http2 package from the golang.org. If you have the $GOPATH configured correctly, then you can use the following command to get the package:
```
go get golang.org/x/net/http2
```
Once you have the http2 package, you can import it along with the net/http package:
```
import (
   "net/http"
   "golang.org/x/net/http2"
)
```
The http2 library integrates with the http package in the standard library. To enable HTTP/2 support, all you need to do is call http2.ConfigureServer()

```
var srv http.Server
srv.Addr = ":8080"
//Enable http2
http2.ConfigureServer(&srv, nil)
```
Although encryption is required, currently all the browser clients require HTTP/2 to be encrypted. You can call the ListenAndServeTLS pointing to the key and ssl certificate to start the server.

```
srv.ListenAndServeTLS("certs/localhost.cert", "certs/localhost.key")
```
### Generating the ssl key and certificate
You can create a key and certificate yourself using the following commands:

First, generate the key:
```
$ openssl genrsa -out localhost.key 2048
```
Now, generate the certificate:
```
$ openssl req -new -x509 -key localhost.key -out localhost.cert -days 3650 -subj /CN=localhost
```

### Start the server

```
go run main.go
```

Now visit https://localhost:8080 ignore the certificate error and you can see the webpage. To verify its using HTTP/2, you can open the developer tools and under the Network tab, enable Protocol and refresh the page.

![alt text](https://hacked.work/blog/wp-content/uploads/2017/04/go-http2.png)

You can check out the complete code from the following repository:

https://github.com/akhld/simple-go-http2-server

Credits:
- https://blog.golang.org/h2push
- https://golang.org


