**URL Shortener**

This Go program implements a simple URL shortener service that allows you to shorten long URLs and retrieve the original URLs using shortened keys.

Endpoints

1. Shorten URL
   Method: POST
   Endpoint: /v1/url/shorten
   Request Body: JSON object with a single field url containing the long URL to be shortened.
   Response: JSON object with a single field shortened_url containing the shortened URL.


2. Retrieve Original URL
   Method: GET
   Endpoint: /v1/url/:key
   URL Parameter: :key is the key representing the shortened URL.
   Response: Redirects to the original URL associated with the provided key.

**Usage**

1. Send a POST request to /v1/url/shorten with a JSON object containing the long URL.

```
 curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com/very/long/url"}' http://localhost:8080/v1/url/shorten
```

Response

```
{
   "url": "http://localhost:8080/jpv1zAbIyh"
}
```

2.  Use the shortened URL returned in the response or send a GET request directly to /v1/url/:key with the provided key to retrieve the original URL.

```
curl -L http://localhost:8080/v1/url/jpv1zAbIyh 
```

Response
```
{
   "url": "https://example.com/very/long/url"
}
```


**Installation and Setup**

1. Clone the repository: git clone https://github.com/fpsbr/shorten-url.git
2. Navigate to the project directory: cd shorten-url
3. Build and run the Go program: go run main.go
4. The server will start running at http://localhost:8080