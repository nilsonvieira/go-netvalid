package main

import (
	"netvalid/configs"
)

func main() {
	configs.GetHostname()
	configs.StatusCode()
	configs.StatusK8sNode()
}
