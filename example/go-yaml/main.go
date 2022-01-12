package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	data, _ := ioutil.ReadFile("docker-compose.yml")

	list, _ := ioutil.ReadDir(".")

	fmt.Println(list)

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	services := m["services"]
	handler := HandlerInterfaceMapping(services)
	for s := range handler {
		i := handler[s]
		tmp := HandlerInterfaceMapping(i)
		fmt.Println(tmp["image"])
	}

}

// HandlerInterfaceMapping
func HandlerInterfaceMapping(res interface{}) (tmp map[string]interface{}) {
	// map 需要初始化一个出来
	tmp = make(map[string]interface{})
	switch res.(type) {
	case nil:
		return tmp
	case map[string]interface{}:
		return res.(map[string]interface{})
	case map[interface{}]interface{}:
		for k, v := range res.(map[interface{}]interface{}) {
			switch k.(type) {
			case string:
				switch v.(type) {
				case map[interface{}]interface{}:
					tmp[k.(string)] = HandlerInterfaceMapping(v)
					continue
				default:
					tmp[k.(string)] = v
				}

			default:
				continue
			}
		}
		return tmp
	default:
		log.Println("unknow data:", res)
	}
	return tmp
}
