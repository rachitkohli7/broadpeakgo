package broadpeakgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type broadpeakgo struct {
	apiKey string
}

func (b broadpeakgo) getAllSources() {

	bearer := "Bearer " + b.apiKey

	url := "https://api.broadpeak.io/v1/sources?offset=0&limit=1000"

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
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}

}
