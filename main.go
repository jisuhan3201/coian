package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jisuhan3201/coian/explorer"
	"github.com/jisuhan3201/coian/rest"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 4000, "Choose port of server")
	mode := flag.String("mode", "rest", "Choose mode between 'rest' or 'html'")

	flag.Parse()

	switch *mode {
	case "html":
		explorer.Start(*port)
	case "rest":
		rest.Start(*port)
	default:
		usage()
	}
}

func usage() {
	fmt.Printf("Welcome to COIAN\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API (recommended)\n\n")
	os.Exit(0)
}
