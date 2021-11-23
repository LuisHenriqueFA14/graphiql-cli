package internal

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
)

/**
* @params url: `url to call`, query: `the POST content`
* This function will execute and return the response body of the request
*/
func RunQuery(url string, query string) string {
	var jsonStr = []byte(`{"query": "` + query + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	
	fmt.Println("\nResponse Status:", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
