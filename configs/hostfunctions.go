package configs

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"netvalid/arguments"
	"os/exec"
)

var URL string = arguments.ArgURL()
var RR string = arguments.ArgRR()

func GetHostname() {

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("+------- TESTE PELA URL ------+")
	fmt.Println("HOSTNAME: " + url.Host)
	fmt.Println("PROTOCOLO: " + url.Scheme)
	fmt.Println("PATH: " + url.Path)
	ips, _ := net.LookupIP(url.Host)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Println("IPV4 DO HOST: ", ipv4)
		}
	}
}

func StatusCode() {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("STATUS DO HOST:", resp.StatusCode, http.StatusText(resp.StatusCode))
}

func StatusK8sNode() {

	url, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}

	HostK8s := "Host: " + url.Host
	TS3RR01 := RR + url.Path

	var CURL = "curl -IH " + "'" + HostK8s + "' " + TS3RR01

	out, err := exec.Command("bash", "-c", CURL).Output()

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		output := string(out[:])
		fmt.Println(" ")
		fmt.Println("+------- TESTE PELO KUBERNETES ------+")
		fmt.Println("COMANDO: ", CURL)
		fmt.Println(" ")
		fmt.Println(output)
		fmt.Println("+------------------------------------+")
	}
}
