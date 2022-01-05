package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"log"
)
func main() {
	cl, err := client.NewClientWithOpts()
	if err != nil {
		log.Println(err)
	}
	inspect, _ := cl.ContainerInspect(context.Background(), "prometheus")
	fmt.Println(inspect.NetworkSettings.Networks)
}