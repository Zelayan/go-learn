package request

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get() {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/list")
	if err != nil {

	}
	defer func() {
		if err == nil {
			resp.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	log.Println(string(body))
}

func GetServiceLog(service string) []byte{
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/get?service=" + service)
	if err != nil {

	}
	defer func() {
		if err == nil {
			resp.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	return body
}