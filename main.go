package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	//get Env var
	token, _ := os.LookupEnv("TEAMCITY_TOKEN")

	url := "http://build.exzeo.com/app/rest/buildTypes"

	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}

}
