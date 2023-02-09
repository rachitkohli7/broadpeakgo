package broadpeakgo

import "fmt"

func (b Broadpeak) GetAllSources() string {
	url := "https://api.broadpeak.io/v1/sources?offset=0&limit=1000"
	resp := HttpGetRequest(b, url)
	return resp
}

func (b Broadpeak) GetSourceStatus(sourceType string, sourceUrl string) string {
	url := fmt.Sprintf("https://api.broadpeak.io/v1/sources/%s/check", sourceType)
	resp := HttpGetRequest(b, url)
	return resp
}
