package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Result struct {
	Msg string `json:"msg"`
}

/*function GetUrl(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	var result Result
	err = json.Unmarshal(b, &result)
	if err != nil {
		return "", err
	}
	return result.Msg, nil
	//fmt.Println(result.Msg)
}*/

func httpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1:8080/sss",
		url.Values{"lesson": {"哈哈哈"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println("...end")

}
func main() {
	/*var url1 string = "http://127.0.0.1:8080/hello?lesson=从0学习golang语言"
	u1, err := GetUrl(url1)
	if err != nil {
		panic(err)
	}
	fmt.Println(GetUrl(u1))
	GetUrl(url1)
	fmt.Println("end....")*/

	httpPostForm()
}
