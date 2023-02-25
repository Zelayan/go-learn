package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Process func(conn net.Conn)

// 利用socket连接自己封装处理HTTP
func processHttp(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("client addr: %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	var buf [1024]byte
	n, err := reader.Read(buf[:])
	if err != nil {
		fmt.Printf("read from conn failed, err:%v\\n", err)
		return
	}
	rev := string(buf[:n])

	//fmt.Printf("recived: %v\n", rev)

	header := strings.Split(rev, "\r\n")
	path := strings.Split(header[0], " ")

	//fmt.Printf("path: %s", path[0])

	response := []byte("HTTP/1.0 200 OK\r\n\r\n")

	file, err := os.ReadFile(fmt.Sprintf("%s.go", path[1]))
	if err != nil {
		response = []byte("HTTP/1.1 404 Not Found\r\n\r\nfile not found")
	} else {
		response = append(response, file...)
	}
	conn.Write(response)

}

// TCP处理客户端请求
func process(conn net.Conn) {
	// 处理完关闭连接
	defer conn.Close()

	fmt.Printf("client addr: %s\n", conn.RemoteAddr())
	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		rev := string(buf[:n])
		fmt.Printf("recived：%v\n", rev)

		// 将接受到的数据返回给客户端
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func TCPServer(p Process) {
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go p(conn)
	}
}

func TCPClient() {

	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	// 2、使用 conn 连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed, err:%v\n", err)
			return
		}
		// 从服务端接收回复消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		fmt.Printf("收到服务端回复:%v\n", string(buf[:n]))
	}

}
