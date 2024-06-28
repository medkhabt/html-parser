package htmlprs

import (
	"fmt"
	"io"
	"net/http"
)

// TODO handle the stream processing better [ look at streaming algorithms]
func Crawl() (string, error) {
	response, err := http.Get("https://www.medkha.com")
	if err != nil {
		return "", fmt.Errorf("Error with Get : [%s].\n", err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Error with Read of the stream : [%s].\n", err)
	}
	fmt.Printf("content is %s. \n", body)
	return "", nil

}
