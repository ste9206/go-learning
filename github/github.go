package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println(UserInfo("ardanlabs"))
}

func demo() {
	res, err := http.Get("https://api.github.com/users/ardanlabs")

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	if(res.StatusCode != http.StatusOK) {
		fmt.Printf("Error: bad status - %s\n", res.Status)
		return
	}

	cType := res.Header.Get("Content-Type")

	fmt.Println("content-type", cType)

	var reply struct {
		Name string
		Public_Repos int `json:"public_repos"`
	}

	dec := json.NewDecoder(res.Body)

	if err := dec.Decode(&reply); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(reply.Name, reply.Public_Repos);
	
}

func UserInfo(login string) (string, int, error){
	url := "https://api.github.io/users" + login

	res, err := http.Get(url)

	if err != nil {
		return "", 0, err
	}

	if res.StatusCode != http.StatusOK {
		return "", 0,fmt.Errorf("%q - bad status: %s", url, res.Status)
	}

	return parseResponse(res.Body)
}

func parseResponse(r io.Reader)(string, int,error) {
	var reply struct {
		Name string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(r)
	
	if err := dec.Decode(&reply); err != nil {
		return "", 0, err
	}

	return reply.Name, reply.NumRepos, nil
}

//io.Copy(os.Stdout, res.Body)

	/* JSON <-> Go 
	
	string <-> string
	true/false <-> bool
	number <-> float64, float42, int, int8 ...int64, uint, uint8 ...
	array <-> []T, []any
	object <-> map[string]any , struct
	

	encoding/json API

	JSON -> []byte -> Go: Unmarshal
	Go -> []byte -> JSON: Marshal
	JSON -> io.Reader -> Go: Decoder
	Go -> io.Writer -> JSON: Encoder
	*/