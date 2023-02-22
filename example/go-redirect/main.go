package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func redirect() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://www.baidu.com", http.StatusFound)
	})

	http.ListenAndServe(":8090", nil)
}

func forward() {
	targetUrl, _ := url.Parse("http://localhost:8090")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)
		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
