package arguments

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ArgURL() string {
	urlReader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe a URL Completa: \n")
	URL, _ := urlReader.ReadString('\n')
	URL = strings.Replace(URL, "\n", "", -1)

	return URL
}
func ArgRR() string {

	routeReader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe o INGRESS com PORTA: \n")
	RR, _ := routeReader.ReadString('\n')
	RR = strings.Replace(RR, "\n", "", -1)

	return RR
}
