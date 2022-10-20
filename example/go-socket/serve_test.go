package go_socket

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	go TCPServer()
	go TCPClient()

	time.Sleep(time.Second * 10)
}

func Test_xxx(t *testing.T) {

}
