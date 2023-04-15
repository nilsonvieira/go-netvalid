package main

import (
	"fmt"
	"netvalid/configs"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Execute in MacOs or Linux Operational System Only!")
	} else {
		configs.GetHostname()
		configs.StatusCode()
		configs.StatusK8sNode()
	}
}
