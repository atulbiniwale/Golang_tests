package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get(" enter url endpoint here ")
	if err != nil {
		fmt.Println("error in getting API response", err)
	}
	defer res.Body.Close()                      // imp to close the API connection, but deferred.
	fmt.Printf("type of API response: %T", res) // get type of API response.

	// read API response body
	data, err := ioutil.ReadAll(res.Body)  
	if err != nil {
		fmt.Println("error in reading API response", err)
		return
	}

	fmt.Println("response:", string(data))

}
