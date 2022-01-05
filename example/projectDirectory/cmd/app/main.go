package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	flagvar = pflag.Int("flagname", 1234, "help message for flagname")
)

func main() {
	pflag.Parse()

	fmt.Printf("argument number is: %v\n", pflag.NArg())
	fmt.Printf("argument list is: %v\n", pflag.Args())
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))

	// preparing config file
	viper.AddConfigPath("./config/")
	viper.SetConfigName("test")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	host := viper.GetString("datastore.metric.host") // (返回 "127.0.0.1")
	fmt.Printf(host)
}
