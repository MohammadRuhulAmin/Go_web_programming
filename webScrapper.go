package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	baseUrl := "https://www.aiub.edu/"
	response, err := http.Get(baseUrl)
	checkErr(err)
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
	response.Body.Close()

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
