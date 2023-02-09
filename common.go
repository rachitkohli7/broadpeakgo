package broadpeakgo

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGetRequest(b Broadpeak, url string) string {

	bearer := "Bearer " + b.apiKey

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return ""
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		return string(data)
	}
}

func HttpPostRequest(b Broadpeak, url string, body string) string {

	bearer := "Bearer " + b.apiKey

	requestBody := []byte(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return ""
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(data))
		return string(data)
	}
}

func HttpDeleteRequest(b Broadpeak, url string) string {
	bearer := "Bearer " + b.apiKey

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return ""
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		return string(data)
	}
}
