package httpReusing

import (
	"io"
	"io/ioutil"
	"net/http"
)

func ReUsing() {
	count := 100
	for i := 0; i < count; i++ {
		resp, err := http.Get("https://www.oschina.net")
		if err != nil {
			panic(err)
		}

		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}

func NotReUsing() {
	count := 100
	for i := 0; i < count; i++ {
		resp, err := http.Get("https://www.oschina.net")
		if err != nil {
			panic(err)
		}

		//io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}