package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https:example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("type of url: %T\n", myURL)

	parsedURL, _ := url.Parse(myURL)
	fmt.Printf("type of url: %T\n", parsedURL)

	// getting components of parsed URL.
	fmt.Println("Scheme of URL:", parsedURL.Scheme)
	fmt.Println("Host of URL:", parsedURL.Host)
	fmt.Println("Path of URL:", parsedURL.Path)
	fmt.Println("Raw Query of URL:", parsedURL.RawQuery)
	fmt.Println("*********************************************")

	// modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=atul"
	fmt.Println("Scheme of URL:", parsedURL.Scheme)
	fmt.Println("Host of URL:", parsedURL.Host)
	fmt.Println("Path of URL:", parsedURL.Path)
	fmt.Println("Raw Query of URL:", parsedURL.RawQuery)

	// Constructing a URL string from URL object
	newURL := parsedURL.String()
	fmt.Println("new URL:", newURL)

}
