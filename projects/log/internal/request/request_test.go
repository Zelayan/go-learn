package request

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	http.HandleFunc("/list", List)
	http.HandleFunc("/get", GetService)
	http.ListenAndServe(":8080", nil)
}

func GetService(writer http.ResponseWriter, request *http.Request) {
	var log string
	query := request.URL.RawQuery
	split := strings.Split(query, "=")
	s := split[1]
	log = "1312312312312312313123123131231231231231231231231231233333333333" + s
	fmt.Fprintf(writer, log)
}


func List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "core audit monitor")
}