package main

import (
	"fmt"
	"net/http"
)

func main() {
	build := NewResponseBuilder().Code(http.StatusOK).Msg("success").Data("11111").Data("22222").Build()
	fmt.Print(build)
}
