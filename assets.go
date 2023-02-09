package broadpeakgo

import (
	"encoding/json"
)

type asset struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (b Broadpeak) CreateAsset(assetName string, sourceUrl string) string {
	url := "https://api.broadpeak.io/v1/sources/asset"
	newAsset := asset{assetName, sourceUrl}

	newAssetJson, _ := json.Marshal(newAsset)

	resp := HttpPostRequest(b, url, string(newAssetJson))
	return resp
}
