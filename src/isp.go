package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ISP(addr string) (string, error) {
	request, err := http.Get("http://ip-api.com/json/"+addr)
	if err != nil {
		return "", err
	}
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return "", err
	}
	if request.StatusCode != 200 {
		return "", fmt.Errorf("API returned %d", request.StatusCode)
	}

	var IP IpAPI

	err = json.Unmarshal(body, &IP)
	if err != nil {
		return "", err
	}

	return IP.Isp, nil
}