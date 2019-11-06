package main

import (
	"bytes"
	"encoding/json"
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

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	//create a custom error
	// var RedirectAttemptedError = errors.New("redirect")

	client := &http.Client{}
	// return the error, so client won't attempt redirects
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERR] -", err)
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var project Project
	err = json.Unmarshal(data, &project)
	if err != nil {
		fmt.Println("Error on response \n[ERR] -", err)
	}
	// ind, _ := json.Marshal(project)
	fmt.Println(project)
}
