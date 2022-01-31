package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Msg string `json:"msg"`
}

func GetUrl(url string) (string, error) {
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
}

type val map[string][]string

/*func httpPostForm(url string) (string ,error) {
	var param val
	param = make(param)
	param = map[string][]string{
		lesson:[]string{"a", "b"},
	}
	resp, err := http.PostForm(url, )

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))

}*/

func main() {
	/*var url1 string = "http://127.0.0.1:8080/hello?lesson=从0学习golang语言"
	u1, err := GetUrl(url1)
	if err != nil {
		panic(err)
	}
	fmt.Println(GetUrl(u1))
	GetUrl(url1)
	fmt.Println("end....")*/

}
