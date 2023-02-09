package broadpeakgo

import (
	"encoding/json"
	"fmt"
)

type category struct {
	Name string `json:"name"`
}

func (b Broadpeak) CreateCategory(name string) string {
	url := "https://api.broadpeak.io/v1/categories"
	newCategory := category{name}

	newCategoryJson, _ := json.Marshal(newCategory)

	resp := HttpPostRequest(b, url, string(newCategoryJson))
	return resp
}

func (b Broadpeak) GetAllCategories() string {
	url := "https://api.broadpeak.io/v1/categories"
	resp := HttpGetRequest(b, url)
	return resp
}

func (b Broadpeak) DeleteCategory(id int) string {
	url := fmt.Sprintf("https://api.broadpeak.io/v1/categories/%d", id)
	fmt.Println(url)
	resp := HttpDeleteRequest(b, url)
	return resp
}
