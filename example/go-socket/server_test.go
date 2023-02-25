package main

import (
	"testing"
)

func Test_processHttp(t *testing.T) {
	TCPServer(processHttp)
}
