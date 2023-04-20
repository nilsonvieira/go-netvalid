package arguments

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Cleans the string removing special chars like "\r" (Carriage Return).
func CleanString(dirtyString string) string {
	newString := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, dirtyString)
	return newString
}

func ArgURL() string {
	urlReader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe a URL Completa: \n")
	URL, _ := urlReader.ReadString('\n')
	URL = CleanString(strings.Replace(URL, "\n", "", -1))

	return URL
}
func ArgRR() string {

	routeReader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe o INGRESS com PORTA: \n")
	RR, _ := routeReader.ReadString('\n')
	RR = CleanString(strings.Replace(RR, "\n", "", -1))

	return RR
}
