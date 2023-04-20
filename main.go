package main

import (
	"fmt"
	"netvalid/configs"
)

func main() {
	configs.GetHostname()
	fmt.Println("Got hostname")
	configs.StatusCode()
	fmt.Println("Got statuscode")
	configs.StatusK8sNode()
}
