package structs

func Perimeter(width float64, height float64) float64 {
	
	return 2 * (width + height)
}

func Area(wigth float64, package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
})