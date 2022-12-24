package utils

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
)

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetStringBetween(str string, start string, end string) string {
	startIndex := strings.Index(str, start)
	endIndex := strings.Index(str, end)
	if startIndex == -1 || endIndex == -1 {
		return ""
	}
	return str[startIndex+len(start) : endIndex]
}

func ReadFile2Str(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func IfElementExistInArray(element string, array []string) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}
