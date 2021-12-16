package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	containers = []string{"core", "dsg"}
)
func main() {
	//获取命令行中输入的参数和对应的值
	//第一个参数，为参数名称，第二个参数为默认值，第三个参数是说明
	//containerlist := flag.String("list", "", "get all container")
	container := flag.String("get", "", "get container nglog")
	flag.Parse() //从os.Args[1:]中解析注册的flag


	//list(containerlist)
	//if *containerlist != "" {
		getLog(container)
	//}

	if *container != "" {
		test(container)
	}

}

func getLog(container *string) {
	fmt.Println("get nglog")
	command := exec.Command("./logcli", "get al2l")

	// set var to get the output
	var errOut bytes.Buffer
	var out bytes.Buffer
	// set the output to our variable
	command.Stderr = &errOut
	command.Stdout = &out
	err := command.Start()
	if err != nil {
		log.Println(err)
		fmt.Println(errOut.String())
	}

	fmt.Println(out.String())
}

func test(str *string) {
	f,err := os.Create(*str)
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		_,err=f.Write([]byte(*str))
	}

}
